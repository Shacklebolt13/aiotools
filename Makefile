WIRE_GEN_FILE = server/wire_gen.go
PROTO_GEN_FOLDER = proto
PROTO_SRC = aiotools.proto
SERVER_FOLDER = ./server

genproto:
	@$(shell mkdir $(PROTO_GEN_FOLDER))
	protoc --go_out=$(PROTO_GEN_FOLDER) --go-grpc_out=$(PROTO_GEN_FOLDER) --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative $(PROTO_SRC)

genwire:
	ifeq ("$(wildcard $(file))","")
		wire $(SERVER_FOLDER)
	endif
	go generate ./server

all: clean genproto genwire

server: all
	go run .\server

client: 
	go run ./client
	
clean:
	$(shell rm $(WIRE_GEN_FILE))
	$(shell rm $(PROTO_GEN_FOLDER)/*.go)
	$(shell rmdir -rf $(PROTO_GEN_FOLDER))

.PHONY: all clean genproto genwire server client