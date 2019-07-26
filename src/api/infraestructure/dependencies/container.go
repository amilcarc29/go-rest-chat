package dependencies

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-resty/resty"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// Container defines a container for dependencies
type Container struct {
	db            *gorm.DB
	client        *resty.Client
	routerHandler *mux.Router
	mock          *sqlmock.Sqlmock
	clock         Clock
}

// NewContainer returns a container with the dependencies
func NewContainer() (*Container, error) {
	db, err := gorm.Open(
		"sqlite3",
		"challenge.db",
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
		clock:         NewClock(),
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

	gormDB, err := gorm.Open("sqlite3", db)
	if err != nil {
		return nil, err
	}

	routerHandler := mux.NewRouter()

	return &Container{
		db:            gormDB,
		routerHandler: routerHandler,
		mock:          &mock,
		clock:         NewClockMock(),
	}, nil
}

// SQLMock returns SQL mock
func (container *Container) SQLMock() *sqlmock.Sqlmock {
	return container.mock
}

type Clock interface {
	Now() time.Time
}

type clock struct {
	Clock
}

func NewClock() Clock {
	return &clock{}
}

func (clock *clock) Now() time.Time {
	return time.Now()
}

type mock struct {
	Clock
	nowDateTime time.Time
}

func NewClockMock() Clock {
	dateString := "2019-07-12T20:49:00Z"
	date, _ := time.Parse(time.RFC3339, dateString)
	return &mock{
		nowDateTime: date,
	}
}

func (clock *mock) Now() time.Time {
	return clock.nowDateTime
}

func (container *Container) Clock() Clock {
	return container.clock
}
