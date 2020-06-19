package usecase

import (
	"bou.ke/monkey"
	"errors"
	"github.com/ciazhar/golang-grpc/app/server/internal/mocks"
	"github.com/ciazhar/golang-grpc/common"
	"github.com/ciazhar/golang-grpc/common/rest"
	"github.com/ciazhar/golang-grpc/grpc/generated/golang"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func NewActual() golang.User {
	var user golang.User
	common.ToStruct("app/server/testdata/user/actual.1.golden", &user)
	return user
}

func NewActual2() golang.User {
	var user golang.User
	common.ToStruct("app/server/testdata/user/actual.2.golden", &user)
	return user
}

func TestUserUseCase_Store(t *testing.T) {
	repo := new(mocks.UserPostgresRepository)
	uc := NewUserUseCase(repo)
	testCases := []struct {
		name        string
		user        golang.User
		returnError error
	}{
		{"default", NewActual(), nil},
		{"default2", NewActual2(), nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("Store", &testCase.user).Return(testCase.returnError)

			err := uc.Store(&testCase.user)

			assert.NoError(t, err)
			repo.AssertExpectations(t)
		})
	}
}

func TestUserUseCase_Fetch(t *testing.T) {
	repo := new(mocks.UserPostgresRepository)
	uc := NewUserUseCase(repo)
	testCases := []struct {
		name        string
		offset      int
		limit       int
		returnUser  []golang.User
		returnError error
	}{
		{"default", 0, 10, []golang.User{NewActual(), NewActual2()}, nil},
		{"default2", 0, 5, []golang.User{NewActual(), NewActual2()}, nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			param := rest.NewParam()
			param.Offset = 1
			param.Limit = 10

			repo.On("Fetch", param).Return(testCase.returnUser, testCase.returnError)

			expected, err := uc.Fetch(param)

			assert.NotEmpty(t, expected)
			assert.NoError(t, err)
			assert.Len(t, expected, len(testCase.returnUser))
			repo.AssertExpectations(t)
		})
	}
}

func TestUserUseCase_GetByID(t *testing.T) {
	repo := new(mocks.UserPostgresRepository)
	uc := NewUserUseCase(repo)
	testCases := []struct {
		name        string
		id          string
		returnUser  golang.User
		returnError error
	}{
		{"default", "1", NewActual(), nil},
		{"default2", "2", NewActual(), nil},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("GetByID", testCase.id).Return(testCase.returnUser, testCase.returnError)

			expected, err := uc.GetByID(testCase.id)

			assert.NoError(t, err)
			assert.NotNil(t, expected)
			repo.AssertExpectations(t)
		})
	}
}

func TestUserUseCase_Update(t *testing.T) {
	repo := new(mocks.UserPostgresRepository)
	uc := NewUserUseCase(repo)
	actual := NewActual()
	actual.Id = "100"
	testCases := []struct {
		name        string
		user        golang.User
		returnError error
	}{
		{"default", NewActual(), nil},
		{"default2", actual, errors.New("not found")},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repo.On("GetByID", testCase.user.Id).Return(testCase.user, testCase.returnError)
			repo.On("Update", &testCase.user).Return(testCase.returnError)

			err := uc.Update(&testCase.user)

			assert.Equal(t, err, testCase.returnError)
			repo.AssertExpectations(t)
		})
	}
}

func TestUserUseCase_Delete(t *testing.T) {
	repo := new(mocks.UserPostgresRepository)
	uc := NewUserUseCase(repo)
	actual := NewActual()
	actual.Id = "100"
	testCases := []struct {
		name        string
		user        golang.User
		returnError error
	}{
		{"default", NewActual(), nil},
		{"error", actual, errors.New("not found")},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			wayback := time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC)
			patch := monkey.Patch(time.Now, func() time.Time { return wayback })
			defer patch.Unpatch()

			repo.On("GetByID", testCase.user.Id).Return(testCase.user, testCase.returnError)

			testCase.user.DeletedAt = ptypes.TimestampNow()
			repo.On("Update", &testCase.user).Return(testCase.returnError)

			err := uc.Delete(testCase.user.Id)

			assert.Equal(t, err, testCase.returnError)
			//repo.AssertExpectations(t)
		})
	}
}
