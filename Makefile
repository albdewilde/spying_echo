.PHONY: grpc-go

grpc-go:
		protoc --go_out=./grpc --go_opt=paths=import \
		--go-grpc_out=./grpc --go-grpc_opt=paths=import \
		./proto/spying_echo.proto
