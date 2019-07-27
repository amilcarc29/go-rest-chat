package dependencies

import (
	"fmt"
	"time"

	"github.com/go-resty/resty"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	mocket "github.com/selvatico/go-mocket"
)

// Container defines a container for dependencies
type Container struct {
	db            *gorm.DB
	client        *resty.Client
	catcher       *mocket.MockCatcher
	routerHandler *mux.Router
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

// Catcher returns mocket catcher
func (container *Container) Catcher() *mocket.MockCatcher {
	return container.catcher
}

// NewMockContainer returns a container mocked
func NewMockContainer() (*Container, error) {

	catcher := mocket.Catcher
	catcher.Register()
	catcher.Logging = true

	gormDB, _ := gorm.Open(mocket.DriverName, "connection_string")

	routerHandler := mux.NewRouter()
	client := resty.New()

	return &Container{
		db:            gormDB,
		routerHandler: routerHandler,
		clock:         NewClockMock(),
		client:        client,
		catcher:       catcher,
	}, nil
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
