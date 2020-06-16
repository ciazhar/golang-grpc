package app

import (
	"github.com/ciazhar/golang-grpc/common/env"
	"github.com/ciazhar/golang-grpc/common/logger"
	"github.com/ciazhar/golang-grpc/common/validator"
	"github.com/gin-gonic/gin"
	"os"
)

type Application struct {
	Env *env.Environtment
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

	//validator
	validator.Init()

	return &Application{
		Env: environment,
	}, nil
}

func SetupAppWithPath(path string) (*Application, error) {
	environment := env.InitPath(path)
	logger.InitLogger()

	validator.Init()

	return &Application{
		Env: environment,
	}, nil
}
