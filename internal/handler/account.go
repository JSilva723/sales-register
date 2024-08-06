package handler

import (
	"context"
	"net/http"
	db "sales-register/db/sqlc"
	"sales-register/internal/util"
)

func Account(ctx context.Context, q *db.Queries) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		type Req struct {
			ID       int32  `json:"id"`
			Name     string `json:"name"`
			Username string `json:"username"`
		}
		req, err := util.ReadJSON[Req](r.Body)
		if err != nil {
			util.WriteError(w, err, http.StatusBadRequest)
			return
		}

		name, err := q.CreateAccount(ctx, db.CreateAccountParams{
			ID:   req.ID,
			Name: req.Name,
		})
		if err != nil {
			util.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		user, err := q.CreateUser(ctx, db.CreateUserParams{
			Username:    req.Username,
			Password:    "admin",
			AccountName: name,
			Rol:         "ADMIN",
		})

		util.WriteJSON(w, user)
	}
}
