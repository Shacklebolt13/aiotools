package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"

	"aiotools/src/handlers"
)

var port = flag.Int("port", 50051, "The server port")
var dsn = flag.String("dsn", "", "The database connection string")

type Application interface {
	Run(listener net.Listener)
}

type ApplicationImpl struct {
	server         *grpc.Server
	shortenHandler *handlers.ShortenerServiceHandler
	config         map[string]string
	database       *gorm.DB
}

func NewApplication(server *grpc.Server, shortenHandler *handlers.ShortenerServiceHandler, dbConn *gorm.DB, config map[string]string) Application {
	return &ApplicationImpl{
		server:         server,
		database:       dbConn,
		shortenHandler: shortenHandler,
		config:         config,
	}
}

func (app *ApplicationImpl) Run(listener net.Listener) {
	log.Printf("server listening at %v", listener.Addr())
	if err := app.server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func main() {
	flag.Parse()

	NewApplication()

	// application := InitializeApp()

	// // create a listener
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// In

	// // create config
	// config := map[string]string{
	// 	"DSN": os.Getenv("DSN"),
	// }

	// conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("failed to connect to database: %v", err)
	// }

	// // register services
	// pb.RegisterShortenerServiceServer(s, handlers.NewShortenerServiceHandler(db))

}
