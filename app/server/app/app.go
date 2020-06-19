package app

import (
	pg2 "github.com/ciazhar/golang-grpc/common/db/pg"
	"github.com/ciazhar/golang-grpc/common/env"
	"github.com/ciazhar/golang-grpc/common/logger"
	"github.com/ciazhar/golang-grpc/common/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"os"
)

type Application struct {
	Env      *env.Environtment
	Postgres *pg.DB
}

func SetupApp() (*Application, error) {

	//env
	environment := env.InitEnv()

	//set default timezone
	if err := os.Setenv("TZ", "Asia/Jakarta"); err != nil {
		panic(err.Error())
	}

	//profile
	gin.SetMode(environment.Get("profile"))

	//logger
	logger.InitLogger()

	//postgres
	pgConn := pg2.InitPG(environment)

	//validator
	validator.Init()

	return &Application{
		Env:      environment,
		Postgres: pgConn,
	}, nil
}

func SetupAppWithPath(path string) (*Application, error) {

	//env
	environment := env.InitPath(path)

	//set default timezone
	if err := os.Setenv("TZ", "Asia/Jakarta"); err != nil {
		panic(err.Error())
	}

	//profile
	gin.SetMode(environment.Get("profile"))

	//logger
	logger.InitLogger()

	//postgres
	pgConn := pg2.InitPG(environment)

	//validator
	validator.Init()

	return &Application{
		Env:      environment,
		Postgres: pgConn,
	}, nil
}
