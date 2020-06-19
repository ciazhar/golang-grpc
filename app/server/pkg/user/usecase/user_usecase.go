package usecase

import (
	"github.com/ciazhar/golang-grpc/app/server/pkg/user/repository/postgres"
	"github.com/ciazhar/golang-grpc/common/logger"
	"github.com/ciazhar/golang-grpc/common/rest"
	"github.com/ciazhar/golang-grpc/common/validator"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
	"github.com/golang/protobuf/ptypes"
	"github.com/imdario/mergo"
)

type UserUseCase interface {
	Fetch(param rest.Param) ([]*golang.User, error)
	GetByID(id string) (golang.User, error)
	Store(req *golang.User) error
	Update(req *golang.User) error
	Delete(id string) error
}

type userUseCase struct {
	UserRepository postgres.UserPostgresRepository
}

func (c userUseCase) GetByID(id string) (golang.User, error) {
	return c.UserRepository.GetByID(id)
}

func (c userUseCase) Update(req *golang.User) error {
	oldReq, err := c.UserRepository.GetByID(req.Id)
	if err != nil {
		return logger.WithError(err)
	}

	if err := mergo.Merge(req, oldReq); err != nil {
		return logger.WithError(err)
	}
	if err := validator.Struct(req); err != nil {
		return logger.WithError(err)
	}

	req.CreatedAt = oldReq.CreatedAt
	req.UpdatedAt = ptypes.TimestampNow()
	req.DeletedAt = oldReq.DeletedAt

	return c.UserRepository.Update(req)
}

func (c userUseCase) Delete(id string) error {
	payload, err := c.GetByID(id)
	if err != nil {
		return logger.WithError(err)
	}
	//TODO
	//if !payload.DeletedAt.IsZero() {
	//	return logger.WithError(errors.New("not found"))
	//}
	payload.DeletedAt = ptypes.TimestampNow()
	return c.UserRepository.Update(&payload)
}

func (c userUseCase) Fetch(param rest.Param) ([]*golang.User, error) {
	return c.UserRepository.Fetch(param)
}

func (c userUseCase) Store(req *golang.User) error {
	if err := validator.Struct(req); err != nil {
		return logger.WithError(err)
	}
	req.CreatedAt = ptypes.TimestampNow()
	req.UpdatedAt = ptypes.TimestampNow()
	return c.UserRepository.Store(req)
}

func NewUserUseCase(UserRepository postgres.UserPostgresRepository) UserUseCase {
	return userUseCase{UserRepository: UserRepository}
}
