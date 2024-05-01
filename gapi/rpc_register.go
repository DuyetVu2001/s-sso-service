package gapi

import (
	"context"
	db "sso-service/db/sqlc"
	"sso-service/pb"
	"sso-service/util"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password %s", err)
	}

	args := db.CreateAccountParams{
		Username:     req.GetUsername(),
		PasswordHash: &hashedPassword,
	}

	account, err := server.store.CreateAccount(ctx, args)
	if err != nil {
		if db.ErrorCode(err) == db.ErrUniqueViolation.Code {
			return nil, status.Errorf(codes.AlreadyExists, "username already exists")

		}

		return nil, status.Errorf(codes.Internal, "failed to create user %s", err)
	}

	res := &pb.RegisterResponse{
		Account: convertAccount(account),
	}

	return res, nil
}
