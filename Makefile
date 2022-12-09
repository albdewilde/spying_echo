.PHONY: grpc-go server server-run

grpc-go:
		protoc --go_out=./grpc --go_opt=paths=import \
		--go-grpc_out=./grpc --go-grpc_opt=paths=import \
		./proto/spying_echo.proto

server:
	@CGO_ENABLED=0 go build -o ./out/server ./server/

server-run: server
	@./out/server
