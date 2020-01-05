package database_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-rest-chat/src/api/domain/check/repository/database"
	"go-rest-chat/src/api/infraestructure/dependencies"
	"testing"
)

func TestCheckDatabaseRepository_Check(t *testing.T) {
	// Given
	ass := assert.New(t)

	tests := []struct {
		name       string
		repository func() *database.CheckDatabaseRepository
		asserts    func(check bool, err error)
	}{
		{
			name: "Success",
			repository: func() *database.CheckDatabaseRepository {
				container, _ := dependencies.NewMockContainer()
				row := []map[string]interface{}{
					{
						"": 1,
					},
				}
				container.Catcher().Reset().NewMock().WithQuery(`SELECT 1`).WithReply(row)
				return database.NewRepository(container)
			},
			asserts: func(check bool, err error) {
				ass.True(check)
				ass.Nil(err)
			},
		},
		{
			name: "Should Fail - Response different than 1",
			repository: func() *database.CheckDatabaseRepository {
				container, _ := dependencies.NewMockContainer()
				row := []map[string]interface{}{
					{
						"": 2,
					},
				}
				container.Catcher().Reset().NewMock().WithQuery(`SELECT 1`).WithReply(row)
				return database.NewRepository(container)
			},
			asserts: func(check bool, err error) {
				ass.False(check)
				ass.NotNil(err)
				ass.EqualError(err, "Unexpected query result")
			},
		},
		{
			name: "Should Fail - Response different than 1",
			repository: func() *database.CheckDatabaseRepository {
				container, _ := dependencies.NewMockContainer()

				container.Catcher().Reset().NewMock().WithQuery(`SELECT 1`).WithError(errors.New("forced for test"))
				return database.NewRepository(container)
			},
			asserts: func(check bool, err error) {
				ass.False(check)
				ass.NotNil(err)
				ass.EqualError(err, "DB connection error")
			},
		},
	}

	for _, test := range tests {
		// Given
		repository := test.repository()

		// Given
		check, err := repository.Check()

		// Then
		test.asserts(check, err)
	}
}
