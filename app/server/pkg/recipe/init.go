package recipe

import (
	"github.com/ciazhar/golang-grpc/app/server/pkg/recipe/controller/grpc"
	grpc2 "google.golang.org/grpc"
)

func InitGRPC(server *grpc2.Server) {
	grpc.NewSocialGRPCController(server)
}
