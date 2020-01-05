package usecases_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-rest-chat/src/api/domain/check/repository/mocks"
	"go-rest-chat/src/api/domain/check/usecases"
	"testing"
)

func TestUseCases_Check_Success(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	repository := mocks.NewMock()
	repository.PatchCheck(true, nil)
	usecases := usecases.NewUseCasesMock(repository)
	check, err := usecases.Check()

	// Then
	ass.True(check)
	ass.Nil(err)
}

func TestUseCases_Check_ShouldFail(t *testing.T) {
	// Given
	ass := assert.New(t)

	// When
	repository := mocks.NewMock()
	repository.PatchCheck(false, errors.New("forced for test"))
	usecases := usecases.NewUseCasesMock(repository)
	check, err := usecases.Check()

	// Then
	ass.False(check)
	ass.NotNil(err)
	ass.EqualError(err, "forced for test")
}
