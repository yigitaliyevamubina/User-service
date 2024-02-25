package main

import (
	"google.golang.org/grpc"
	"net"
	"user-service/config"
	pb "user-service/genproto/user-service"
	"user-service/pkg/db"
	"user-service/pkg/logger"
	"user-service/service"
	grpcclient "user-service/service/grpc_client"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "user-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, _, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sql connection to postgres error", logger.Error(err))
	}

	client, err := grpcclient.New(cfg)
	if err != nil {
		log.Fatal("error while creating a new client")
	}

	userService := service.NewUserService(connDB, log, client)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)

	log.Info("main: server is running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("error while listening: %v", logger.Error(err))
	}
}
