.PHONY: all

install:
	go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go

generate:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway \
		--go_out=plugins=grpc:./grpc \
		--grpc-gateway_out=logtostderr=true:./grpc \
		--swagger_out=allow_merge=true,merge_file_name=global:./grpc/generated/swagger \
		grpc/proto/**

client:
	go run app/client/main.go

server:
	go run app/server/main.go

gateway:
	go run app/gateway/main.go
