package main

import (
	"aiotools/proto"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

var port = os.Getenv("PORT")
var dsn = os.Getenv("DSN")
var parseBool = func(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}
var enableReflection = os.Getenv("ENABLE_REFLECTION")

type Application interface {
	Run(listener net.Listener)
}

type AppConfig struct {
	DSN     string
	PORT    string
	REFLECT bool
}
type ApplicationImpl struct {
	server         *grpc.Server
	shortenHandler *proto.ShortenerServiceServer
	pubSubHandler  *proto.PubSubServiceServer
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
	pubSubHandler proto.PubSubServiceServer,
	config AppConfig,
) ApplicationImpl {
	proto.RegisterShortenerServiceServer(server, shortenHandler)
	proto.RegisterPubSubServiceServer(server, pubSubHandler)
	if config.REFLECT {
		reflection.Register(server)
	}
	return ApplicationImpl{
		server:         server,
		shortenHandler: &shortenHandler,
		pubSubHandler:  &pubSubHandler,
		config:         config,
	}
}

func main() {
	application, err := InitializeApp([]grpc.ServerOption{}, []gorm.Option{},
		AppConfig{
			DSN:     dsn,
			PORT:    port,
			REFLECT: parseBool(enableReflection),
		},
	)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}
	application.Run()
}
