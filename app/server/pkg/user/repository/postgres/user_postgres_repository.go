package postgres

import (
	"github.com/ciazhar/golang-grpc/app/server/app"
	"github.com/ciazhar/golang-grpc/common/logger"
	"github.com/ciazhar/golang-grpc/common/rest"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
	"github.com/go-pg/pg/v9/orm"
	uuid "github.com/satori/go.uuid"
)

type UserPostgresRepository interface {
	Fetch(param rest.Param) ([]*golang.User, error)
	GetByID(id string) (golang.User, error)
	Store(req *golang.User) error
	Update(req *golang.User) error
}

type repository struct {
	app *app.Application
}

func (r repository) Fetch(param rest.Param) ([]*golang.User, error) {
	users := make([]*golang.User, 0)
	query := r.app.Postgres.Model(&users).
		Where("deleted_at is null").
		Order("created_at DESC").
		Offset(param.Offset).
		Limit(param.Limit)
	if err := query.Select(); err != nil {
		return users, logger.WithError(err)
	}
	return users, nil
}

func (r repository) GetByID(id string) (golang.User, error) {
	user := golang.User{Id: id}
	if err := r.app.Postgres.Select(&user); err != nil {
		return user, logger.WithError(err)
	}
	return user, nil
}

func (r repository) Store(req *golang.User) error {
	id := uuid.Must(uuid.NewV4(), nil)
	req.Id = id.String()
	return r.app.Postgres.Insert(req)
}

func (r repository) Update(req *golang.User) error {
	return r.app.Postgres.Update(req)
}

func NewUserPostgresRepository(app *app.Application) UserPostgresRepository {
	r := repository{
		app: app,
	}

	if err := r.app.Postgres.CreateTable((*golang.User)(nil), &orm.CreateTableOptions{
		IfNotExists: true,
		Temp:        false,
	}); err != nil {
		panic(err)
	}

	return r
}
