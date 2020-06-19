package grpc

import (
	"context"
	"github.com/ciazhar/golang-grpc/app/server/pkg/user/usecase"
	"github.com/ciazhar/golang-grpc/common/rest"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
	"google.golang.org/grpc"
)

type userController struct {
	uc usecase.UserUseCase
}

func (u userController) AddRecipe(ctx context.Context, user *golang.User) (*golang.User, error) {
	if err := u.uc.Store(user); err != nil {
		return user, err
	}
	return user, nil
}

func (u userController) ListAllUser(request *golang.ListAllUserRequest, server golang.UserService_ListAllUserServer) error {
	user, err := u.uc.Fetch(rest.NewParam())
	if err != nil {
		return err
	}
	return server.Send(&golang.ListAllUserResponse{User: user})
}

func (u userController) UpdateService(ctx context.Context, user *golang.User) (*golang.User, error) {
	if err := u.uc.Update(user); err != nil {
		return user, err
	}
	return user, nil
}

func NewUserGRPCController(server *grpc.Server, useCase usecase.UserUseCase) {
	golang.RegisterUserServiceServer(server, &userController{
		uc: useCase,
	})
}
