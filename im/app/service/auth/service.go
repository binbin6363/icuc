package auth

import (
	"context"

	apppb "github.com/binbin6363/icuc-pb/protobuf/im/app"
)

type Service struct {
	apppb.UnimplementedAuthServiceServer
}

func (s *Service) Login(ctx context.Context, request *apppb.LoginRequest) (*apppb.LoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Logout(ctx context.Context, request *apppb.LogoutRequest) (*apppb.LogoutResponse, error) {
	//TODO implement me
	panic("implement me")
}

func New() *Service {
	return &Service{}
}
