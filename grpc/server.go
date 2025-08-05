package grpc

import (
	"context"
	"user/cmd/user/usecase"
	"user/proto/userpb"
)

type GRPCServer struct {
	// Add fields and methods as needed for your gRPC server implementation.
	userpb.UnimplementedUserServiceServer
	UserUsecase usecase.UserUsecase
}

func (s *GRPCServer) GetUserInfoByUserID(ctx context.Context, req *userpb.GetUserInfoRequest) (*userpb.GetUserInfoResult, error) {
	// Implement the logic to handle the request and return the response.
	userInfo, err := s.UserUsecase.GetUserInfoByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &userpb.GetUserInfoResult{
		Id:    userInfo.ID,
		Name:  userInfo.Name,
		Email: userInfo.Email,
		Role:  userInfo.Role,
	}, nil
}
