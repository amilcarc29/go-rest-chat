package dependencies

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-resty/resty"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Container defines a container for dependencies
type Container struct {
	db            *gorm.DB
	client        *resty.Client
	routerHandler *mux.Router
	mock          *sqlmock.Sqlmock
}

// export DBUsername=root
// export DBPassword=
// export DBHost=127.0.0.1:3306
// export DBName=db_test

// NewContainer returns a container with the dependencies
func NewContainer() (*Container, error) {
	db, err := gorm.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", os.Getenv("DBUsername"),
			os.Getenv("DBPassword"), os.Getenv("DBHost"), os.Getenv("DBName")),
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("DB connected successfully.")
	db.LogMode(true)

	routerHandler := mux.NewRouter()
	client := resty.New()

	return &Container{
		db:            db,
		client:        client,
		routerHandler: routerHandler,
	}, nil
}

// Database returns database
func (container *Container) Database() *gorm.DB {
	return container.db
}

// RouterHandler returns router handler
func (container *Container) RouterHandler() *mux.Router {
	return container.routerHandler
}

// HTTPClient returns http client
func (container *Container) HTTPClient() *resty.Client {
	return container.client
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

	gormDB.LogMode(true)
	routerHandler := mux.NewRouter()

	return &Container{
		db:            gormDB,
		routerHandler: routerHandler,
		mock:          &mock,
	}, nil
}

// SQLMock returns SQL mock
func (container *Container) SQLMock() *sqlmock.Sqlmock {
	return container.mock
}
