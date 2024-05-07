package main

import (
	"aiotools/proto"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var port = os.Getenv("PORT")
var dsn = os.Getenv("DSN")

type Application interface {
	Run(listener net.Listener)
}

type AppConfig struct {
	DSN  string
	PORT string
}
type ApplicationImpl struct {
	server         *grpc.Server
	shortenHandler *proto.ShortenerServiceServer
	config         AppConfig
}

func (app *ApplicationImpl) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", app.config.PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := app.server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewApplication(
	server *grpc.Server,
	shortenHandler proto.ShortenerServiceServer,
	config AppConfig,
) ApplicationImpl {
	proto.RegisterShortenerServiceServer(server, shortenHandler)
	return ApplicationImpl{
		server:         server,
		shortenHandler: &shortenHandler,
		config:         config,
	}
}

func main() {
	application, err := InitializeApp([]grpc.ServerOption{}, []gorm.Option{nil, nil},
		AppConfig{
			DSN:  dsn,
			PORT: port,
		},
	)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	application.Run()
}
