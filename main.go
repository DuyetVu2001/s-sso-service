package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sso-service/api"
	db "sso-service/db/sqlc"
	"sso-service/util"

	"github.com/jackc/pgx/v5"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := pgx.Connect(context.Background(), config.POSTGRES_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect to db: %v\n", err)
	}
	defer conn.Close(context.Background())

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
