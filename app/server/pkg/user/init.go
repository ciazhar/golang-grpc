package user

import (
	"github.com/ciazhar/golang-grpc/app/server/app"
	"github.com/ciazhar/golang-grpc/app/server/pkg/user/controller/grpc"
	"github.com/ciazhar/golang-grpc/app/server/pkg/user/repository/postgres"
	"github.com/ciazhar/golang-grpc/app/server/pkg/user/usecase"
	grpc2 "google.golang.org/grpc"
)

func InitGRPC(server *grpc2.Server, app *app.Application) {
	repo := postgres.NewUserPostgresRepository(app)
	uc := usecase.NewUserUseCase(repo)
	grpc.NewUserGRPCController(server, uc)
}
