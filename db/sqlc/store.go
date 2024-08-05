package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to exexute db Queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := New(tx)

	if err = fn(queries); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type Line struct {
	Ammount   int32 `json:"ammount"`
	ProductID int32 `json:"product_id"`
}

type OrderTxParams struct {
	AccountName string `json:"account_name"`
	UserID      int32  `json:"user_id"`
	PaymentID   int32  `json:"payment_id"`
	Lines       []Line `json:"order_lines"`
}

type OrderTxResult struct {
	SaleOrderID int32 `json:"order_id"`
}

func (store *Store) OrderTx(ctx context.Context, arg OrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Create sale order
		result.SaleOrderID, err = q.CreateSaleOrder(ctx, CreateSaleOrderParams{
			AccountName: arg.AccountName,
			UserID:      arg.UserID,
			PaymentID:   arg.PaymentID,
		})
		if err != nil {
			return err
		}

		// Create order lines
		for _, line := range arg.Lines {
			err = q.CreateOrderLine(context.Background(), CreateOrderLineParams{
				AccountName: arg.AccountName,
				SaleOrderID: result.SaleOrderID,
				Ammount:     line.Ammount,
				ProductID:   line.ProductID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}
