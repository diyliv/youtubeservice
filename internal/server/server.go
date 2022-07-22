package server

import (
	"net"
	"os"
	"os/signal"

	"github.com/diyliv/youtubeservice/configs"
	services "github.com/diyliv/youtubeservice/internal/yt/delivery/grpc"
	ytservicepb "github.com/diyliv/youtubeservice/proto/ytservice"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type server struct {
	logger *zap.Logger
	config *configs.Config
	YToken string
}

func NewServer(logger *zap.Logger, config *configs.Config, YToken string) *server {
	return &server{logger: logger, config: config, YToken: YToken}
}

func (s *server) RunGRPC() {
	s.logger.Info("Starting gRPC server")
	lis, err := net.Listen("tcp", s.config.Server.Host+s.config.Server.Port)
	if err != nil {
		s.logger.Error("Error while listening: " + err.Error())
	}

	service := services.NewYTService(s.logger, s.YToken)

	serv := grpc.NewServer()
	ytservicepb.RegisterYTserviceServer(serv, service)

	go func() {
		if err := serv.Serve(lis); err != nil {
			s.logger.Error("Error while serving: " + err.Error())
		}
	}()

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt)
	<-done

	serv.GracefulStop()
	s.logger.Info("Exiting was successful")
}
