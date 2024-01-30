package config

import (
	"context"

	apipb "github.com/binbin6363/icuc-pb/protobuf/api"
	"github.com/binbin6363/icuc/common/log"
)

type Service struct {
	apipb.UnimplementedConfigServiceServer
}

func (s *Service) Get(ctx context.Context, request *apipb.ConfigRequest) (*apipb.ConfigResponse, error) {
	log.InfoContext(ctx, "recv req: %v", request)
	rsp := &apipb.ConfigResponse{}
	rsp.ServerInfoList = append(rsp.ServerInfoList, &apipb.ServerInfo{
		Type: 1,
		Endpoints: []*apipb.Endpoint{
			{Address: "127.0.0.1:443", Weight: 100},
			{Address: "127.0.0.1:8080", Weight: 90},
		},
	})
	log.InfoContextf(ctx, "done Get rsp: %v", rsp)
	return rsp, nil
}

func New() *Service {
	return &Service{}
}
