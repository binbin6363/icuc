package main

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	apipb "github.com/binbin6363/icuc-pb/protobuf/api"
	apppb "github.com/binbin6363/icuc-pb/protobuf/im/app"
	"github.com/binbin6363/icuc/common/log"
	cfg "github.com/binbin6363/icuc/im/app/config"
	"github.com/binbin6363/icuc/im/app/service/auth"
	"github.com/binbin6363/icuc/im/app/service/config"
	"github.com/binbin6363/icuc/im/app/service/message"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
)

var (
	serviceName  = os.Getenv("SERVICE_NAME")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("INSECURE_MODE")
)

func initTracer() func(context.Context) error {

	var secureOption otlptracegrpc.Option

	if strings.ToLower(insecure) == "false" || insecure == "0" || strings.ToLower(insecure) == "f" {
		secureOption = otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	} else {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorURL),
		),
	)

	if err != nil {
		log.Errorf("Failed to create exporter: %v", err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Fatalf("Could not set resources: %v", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}

func main() {
	confFile := flag.String("f", "../etc/conf.yaml", "配置文件路径")

	flag.Parse()
	cfg.Init(*confFile)
	gin.SetMode(cfg.AppConfig().ServerInfo.Mode)
	log.InitLogger(cfg.AppConfig().LogInfo.Path,
		cfg.AppConfig().LogInfo.MaxSize,
		cfg.AppConfig().LogInfo.MaxBackUps,
		cfg.AppConfig().LogInfo.MaxAge,
		cfg.AppConfig().LogInfo.Level,
		cfg.AppConfig().LogInfo.CallerSkip)

	cleanup := initTracer()
	defer cleanup(context.Background())

	// 初始化gin插件
	//r := plugins.Init(serviceName)
	r := gin.New()
	r.Use(gin.Recovery())
	//r.Use(gin.Logger())

	//service.Init()
	// 创建 gRPC 服务器
	// grpcServer := grpc.NewServer()
	// apipb.RegisterConfigServiceServer(grpcServer, config.New())

	// 创建 gRPC-Gateway 多路复用器
	mux := runtime.NewServeMux()
	// 创建一个 gRPC 连接
	//grpcConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("Failed to connect to gRPC server: %v", err)
	//}
	// 注册 gRPC-Gateway 处理程序
	err := apppb.RegisterAuthServiceHandlerServer(context.Background(), mux, auth.New())
	if err != nil {
		log.Fatalf("Failed to register gRPC-Gateway AuthService handler: %v", err)
	}
	err = apppb.RegisterMessageServiceHandlerServer(context.Background(), mux, message.New())
	if err != nil {
		log.Fatalf("Failed to register gRPC-Gateway MessageService handler: %v", err)
	}
	err = apipb.RegisterConfigServiceHandlerServer(context.Background(), mux, config.New())
	if err != nil {
		log.Fatalf("Failed to register gRPC-Gateway ConfigService handler: %v", err)
	}

	r.Use(gin.WrapH(mux))
	//r.Any("/api/", gin.WrapH(mux))

	if err := r.Run(cfg.AppConfig().ServerInfo.Listen); err != nil {
		log.Fatalf("startup service failed, err:%v", err)
	}
	log.Error("Error exit")
}
