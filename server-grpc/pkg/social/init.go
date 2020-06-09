package social

import (
	"github.com/ciazhar/golang-grpc/server-grpc/pkg/social/repository/grpc"
	grpc2 "google.golang.org/grpc"
)

func InitGRPC(server *grpc2.Server) {
	grpc.NewSocialGRPCRepository(server)
}
