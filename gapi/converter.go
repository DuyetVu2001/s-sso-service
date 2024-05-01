package gapi

import (
	db "sso-service/db/sqlc"
	"sso-service/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertAccount(account db.CreateAccountRow) *pb.Account {
	return &pb.Account{
		Id:        account.ID,
		Username:  account.Username,
		CreatedAt: timestamppb.New(account.CreatedAt),
	}
}
