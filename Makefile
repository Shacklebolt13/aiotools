
genproto:
	protoc --go_out=. --go-grpc_out=. --proto_path=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative .\proto\src\main.proto
