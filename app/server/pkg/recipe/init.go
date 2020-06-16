package recipe

import (
	"github.com/ciazhar/golang-grpc/app/server/pkg/recipe/repository/grpc"
	grpc2 "google.golang.org/grpc"
)

func InitGRPC(server *grpc2.Server) {
	grpc.NewSocialGRPCRepository(server)
}
