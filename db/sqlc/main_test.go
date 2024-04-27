package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sso-service/util"
	"testing"

	"github.com/jackc/pgx/v5"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := pgx.Connect(context.Background(), config.POSTGRES_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	testQueries = New(conn)

	os.Exit(m.Run())
}
