package db

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(
		"postgres",
		"postgres://salesregister:salesregister@localhost/salesregister?sslmode=disable",
	)
	if err != nil {
		log.Fatal("error connet to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

const alphebet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generate a random string of length n
func randomString(n int) string {
	var sb strings.Builder
	k := len(alphebet)

	for i := 0; i < n; i++ {
		c := alphebet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}
