
genproto:
	protoc --go_out=proto --go-grpc_out=proto  --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative .\aiotools.proto
