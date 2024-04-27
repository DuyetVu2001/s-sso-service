package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sso-service/api"
	db "sso-service/db/sqlc"

	"github.com/jackc/pgx/v5"
)

func main() {
	dbUrl := "postgres://default:lZsIkJCjEV97@ep-round-voice-37130748-pooler.ap-southeast-1.aws.neon.tech:5432/verceldb?sslmode=require"
	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect to db: %v\n", err)
	}
	defer conn.Close(context.Background())

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start("localhost:8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
