package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dowenliu-xyz/wild-workouts-go-ddd-walkthrough/internal/common/genproto/users"
)

type GrpcServer struct {
	users.UnimplementedUsersServiceServer

	db db
}

func (g GrpcServer) GetTrainingBalance(ctx context.Context, request *users.GetTrainingBalanceRequest) (*users.GetTrainingBalanceResponse, error) {
	user, err := g.db.GetUser(ctx, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &users.GetTrainingBalanceResponse{Amount: int64(user.Balance)}, nil
}

func (g GrpcServer) UpdateTrainingBalance(
	ctx context.Context,
	req *users.UpdateTrainingBalanceRequest,
) (*users.EmptyResponse, error) {
	err := g.db.UpdateBalance(ctx, req.UserId, int(req.AmountChange))
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to update balance: %s", err))
	}

	return &users.EmptyResponse{}, nil
}
