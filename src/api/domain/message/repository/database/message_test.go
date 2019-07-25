package database_test

// import (
// 	"errors"
// 	"go-rest-chat/src/api/domain/message/repository"
// 	"go-rest-chat/src/api/infraestructure/dependencies"
// 	"testing"
// 	"time"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/stretchr/testify/assert"
// )

// /*
// GetResourceByID Stmt
// ^SELECT (.*) FROM `resources`  WHERE `resources`.`deleted_at` IS NULL AND ((`resources`.`id` = 7)) ORDER BY `resources`.`id` AS
// C LIMIT 1

// GetResources Stmt
// ^SELECT (.*) FROM `resources`  WHERE `resources`.`deleted_at` IS NULL
// */

// func Test_GetResources_OneResult_Success(t *testing.T) {
// 	// Given
// 	assert, container, repository := getDependenciesMock(t)
// 	defer container.Database().Close()

// 	// When
// 	dateString := "2019-06-17T23:35:59Z"
// 	layout := "2014-09-12T11:45:26.371Z"
// 	date, err := time.Parse(layout, dateString)
// 	rows := sqlmock.NewRows([]string{"link", "name", "description", "author", "tags", "created_at", "updated_at", "deleted_at", "id"}).
// 		AddRow("link_test", "name_test", "description_test", "author_test", "tags_test", date, date, date, 1)
// 	(*container.SQLMock()).ExpectQuery("^SELECT (.*) FROM `resources`  WHERE `resources`.`deleted_at` IS NULL").WillReturnRows(rows)

// 	// Then
// 	resources, err := repository.GetResources()

// 	assert.Len(resources, 1)
// 	assert.Equal(uint(1), resources[0].ID)
// 	assert.Nil(err)
// }

// func Test_GetResources_NoResult_Success(t *testing.T) {
// 	// Given
// 	assert, container, repository := getDependenciesMock(t)
// 	defer container.Database().Close()

// 	// When
// 	rows := sqlmock.NewRows([]string{"link", "name", "description", "author", "tags", "created_at", "updated_at", "deleted_at", "id"})
// 	(*container.SQLMock()).ExpectQuery("^SELECT (.*) FROM `resources`  WHERE `resources`.`deleted_at` IS NULL").WillReturnRows(rows)

// 	// Then
// 	resources, err := repository.GetResources()

// 	assert.Len(resources, 0)
// 	assert.Nil(err)
// }

// func Test_GetResources_NoResult_Fail(t *testing.T) {
// 	// Given
// 	assert, container, repository := getDependenciesMock(t)
// 	defer container.Database().Close()

// 	// When
// 	(*container.SQLMock()).ExpectQuery("^SELECT (.*) FROM `resources`  WHERE `resources`.`deleted_at` IS NULL").WillReturnError(errors.New("forced for test"))

// 	// Then
// 	resources, err := repository.GetResources()

// 	assert.Len(resources, 0)
// 	assert.NotNil(err)
// 	assert.EqualError(err, "forced for test")
// }

// func Test_GetResourceByID_Success(t *testing.T) {
// 	// Given
// 	assert, container, repository := getDependenciesMock(t)
// 	defer container.Database().Close()

// 	// When
// 	dateString := "2019-06-17T23:35:59Z"
// 	layout := "2014-09-12T11:45:26.371Z"
// 	date, err := time.Parse(layout, dateString)
// 	rows := sqlmock.NewRows([]string{"link", "name", "description", "author", "tags", "created_at", "updated_at", "id"}).
// 		AddRow("link_test", "name_test", "description_test", "author_test", "tags_test", date, date, 1)
// 	(*container.SQLMock()).ExpectQuery("^SELECT (.*) FROM `resources`  WHERE `resources`.`deleted_at` IS NULL AND ((`resources`.`id` = (.+))) ORDER BY `resources`.`id` ASC LIMIT 1").WillReturnRows(rows)

// 	// Then
// 	resource, err := repository.GetResource("1")

// 	assert.NotNil(resource)
// 	assert.Equal(uint(1), resource.ID)
// 	assert.Nil(err)
// }

// func getDependenciesMock(t *testing.T) (*assert.Assertions, *dependencies.Container, repository.ResourceRepository) {
// 	assert := assert.New(t)
// 	mockContainer, err := dependencies.NewMockContainer()
// 	assert.Nil(err)
// 	repository := repository.NewResourceRepository(mockContainer)
// 	return assert, mockContainer, repository
// }
