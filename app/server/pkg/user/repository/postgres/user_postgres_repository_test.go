package postgres

import (
	"github.com/ciazhar/golang-grpc/app/server/app"
	"github.com/ciazhar/golang-grpc/common"
	"github.com/ciazhar/golang-grpc/common/env"
	"github.com/ciazhar/golang-grpc/common/rest"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
	"github.com/go-pg/pg/v9/orm"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var Application *app.Application

func init() {
	application, err := app.SetupAppWithPath(env.GetEnvPath() + "/config.json")
	if err != nil {
		panic(err)
	}
	Application = application
_:
	application.Postgres.DropTable((*golang.User)(nil), &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})
_:
	application.Postgres.CreateTable((*golang.User)(nil), nil)
}

var ID string

func NewActual() golang.User {
	var user golang.User
	common.ToStruct("app/server/testdata/user/actual.1.golden", &user)
	user.CreatedAt = ptypes.TimestampNow()
	user.UpdatedAt = ptypes.TimestampNow()
	return user
}

func NewActual2() golang.User {
	var user golang.User
	common.ToStruct("app/server/testdata/user/actual.2.golden", &user)
	user.CreatedAt = ptypes.TimestampNow()
	user.UpdatedAt = ptypes.TimestampNow()
	return user
}

func TestRepository_Store(t *testing.T) {
	actual := NewActual()
	actual2 := NewActual2()
	repo := NewUserPostgresRepository(Application)

	t.Run("default", func(t *testing.T) {
		err := repo.Store(&actual)
		assert.NoError(t, err)
	})
	t.Run("default2", func(t *testing.T) {
		err := repo.Store(&actual2)
		assert.NoError(t, err)
		ID = actual2.Id
	})
}

func TestRepository_Fetch(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		repo := NewUserPostgresRepository(Application)
		param := rest.NewParam()
		param.Offset = 0
		param.Limit = 10
		expected, err := repo.Fetch(param)

		assert.NotEmpty(t, expected)
		assert.NoError(t, err)
		assert.Len(t, expected, 2)
	})

	t.Run("error", func(t *testing.T) {
		repo := NewUserPostgresRepository(Application)
		param := rest.NewParam()
		param.Offset = 0
		param.Limit = -10
		expected, err := repo.Fetch(param)

		assert.Empty(t, expected)
		assert.Error(t, err)
	})
}

func TestRepository_GetByID(t *testing.T) {
	repo := NewUserPostgresRepository(Application)

	t.Run("default", func(t *testing.T) {
		expected, err := repo.GetByID(ID)

		assert.NotNil(t, expected)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		expected, err := repo.GetByID("100")

		assert.NotNil(t, expected)
		assert.Error(t, err)
	})
}

func TestRepository_Update(t *testing.T) {
	actual := NewActual()
	repo := NewUserPostgresRepository(Application)

	t.Run("default", func(t *testing.T) {
		actual.Id = ID
		actual.Name = "aye"
		err := repo.Update(&actual)
		assert.NoError(t, err)
	})
}
