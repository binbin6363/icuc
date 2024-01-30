package message

import (
	"context"

	apppb "github.com/binbin6363/icuc/protobuf/im/app"
)

type Service struct {
	apppb.UnimplementedMessageServiceServer
}

func (s *Service) SingleMessage(ctx context.Context, request *apppb.SingleMessageRequest) (*apppb.SingleMessageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) GroupMessage(ctx context.Context, request *apppb.GroupMessageRequest) (*apppb.GroupMessageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) ImageMessage(ctx context.Context, request *apppb.ImageMessageRequest) (*apppb.ImageMessageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) FileMessage(ctx context.Context, request *apppb.FileMessageRequest) (*apppb.FileMessageResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) mustEmbedUnimplementedMessageServiceServer() {
	//TODO implement me
	panic("implement me")
}

func New() *Service {
	return &Service{}
}
