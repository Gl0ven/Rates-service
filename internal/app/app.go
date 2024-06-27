package app

import (
	"Gl0ven/kata_projects/rates/config"
	db "Gl0ven/kata_projects/rates/internal/database"
	grpc "Gl0ven/kata_projects/rates/internal/grpc/gen"
	"Gl0ven/kata_projects/rates/internal/handlers"
	"Gl0ven/kata_projects/rates/internal/provider/garantex"
	"Gl0ven/kata_projects/rates/internal/service"
	"Gl0ven/kata_projects/rates/internal/storage"
	"Gl0ven/kata_projects/rates/pkg/logs"
	"net"

	"go.uber.org/zap"
	grpc_init "google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type App struct {
	Host string
	Port string
	Server *grpc_init.Server
	Listener net.Listener
	Logger      *zap.Logger
}

func Bootstrap() *App {
	logger, err := logs.NewLogger()
	if err != nil {
		panic(err)
	}

	appConf := config.NewAppConf()

	db, err := db.NewDB(config.NewDBConf())
	if err != nil {
		logger.Fatal("error init db", zap.Error(err))
	}
	
	storage := storage.NewStorage(db)

	provider := garantex.NewGarantexProvider("usdtrub")

	service := service.NewService(storage, provider, logger)

	handlers := handlers.NewHandler(service)

	server := grpc_init.NewServer()
	grpc.RegisterRatesServiceServer(server, handlers)

    healthServer := health.NewServer()
    grpc_health_v1.RegisterHealthServer(server, healthServer)
    healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	listener, err := net.Listen("tcp", appConf.Host + ":" + appConf.Port)
	if err != nil {
		logger.Fatal("init listener error", zap.Error(err))
	}

	return &App{
		Host: appConf.Host,
		Port: appConf.Port,
		Server: server,
		Listener: listener,
		Logger: logger,
	}
}
