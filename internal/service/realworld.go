package service

import (
	"context"
	v1 "realworld/api/realworld/v1"
	"realworld/internal/biz"
)

// RealworldService is a greeter service.
type RealworldService struct {
	v1.UnimplementedRealworldServer

	uc *biz.GreeterUsecase
}

// NewRealworldService new a realworld service.
func NewRealworldService(uc *biz.GreeterUsecase) *RealworldService {
	return &RealworldService{uc: uc}
}

func (r *RealworldService) Login(context.Context, *v1.LoginRequest) (*v1.UsersReply, error) {
	return &v1.UsersReply{
		User: &v1.UsersReply_User{
			Username: "test",
		},
	}, nil
}
