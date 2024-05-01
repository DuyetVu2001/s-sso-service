package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sso-service/api"
	db "sso-service/db/sqlc"
	"sso-service/gapi"
	"sso-service/pb"
	"sso-service/util"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// runGinServer(store, config)
	runGrpcServer(store, config)
}

func runGinServer(store *db.Store, config util.Config) {
	server := api.NewServer(store)

	err := server.Start(config.HTTP_SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func runGrpcServer(store *db.Store, config util.Config) {
	server := gapi.NewServer(store)

	// err := server.Start(config.HTTP_SERVER_ADDRESS)
	// if err != nil {
	// 	log.Fatal("cannot start gRPC server:", err)
	// }

	grpcServer := grpc.NewServer()
	pb.RegisterAppServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPC_SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
