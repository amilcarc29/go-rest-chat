package mock

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/suite"

	"github.com/DATA-DOG/go-sqlmock"
)

// Container defines a container for dependencies
type Container struct {
	db            *gorm.DB
	routerHandler *mux.Router
	suite.Suite
	mock sqlmock.Sqlmock
}

// NewMockContainer returns a container mocked
func NewMockContainer() (*Container, error) {

	var db *sql.DB
	var err error
	var mock sqlmock.Sqlmock

	db, mock, err = sqlmock.New()
	if err != nil {
		return nil, err
	}

	gormDB, err := gorm.Open("mysql", db)
	if err != nil {
		return nil, err
	}

	routerHandler := mux.NewRouter()

	return &Container{
		db:            gormDB,
		routerHandler: routerHandler,
		mock:          mock,
	}, nil
}

// Database returns database
func (mockContainer *Container) Database() *gorm.DB {
	return mockContainer.db
}

// RouterHandler returns router handler
func (mockContainer *Container) RouterHandler() *mux.Router {
	return mockContainer.routerHandler
}

// SQLMock returns SQL mock
func (mockContainer *Container) SQLMock() sqlmock.Sqlmock {
	return mockContainer.mock
}
