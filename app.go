package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type App struct {
	// Config
	Host          string
	Port          string
	DevSqlitePath string

	// Dependencies
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewAppWithDefaults() *App {
	a := &App{
		Host:          "0.0.0.0",
		Port:          "8080",
		DevSqlitePath: "./dev.sqlite",
	}
	a.SetupDB()
	a.SetupHTTP()
	return a
}

func (a *App) SetupDB() {
	if a.DB == nil {
		os.Remove(a.DevSqlitePath)
		db, err := gorm.Open(sqlite.Open(a.DevSqlitePath), &gorm.Config{})
		if err != nil {
			logrus.WithField("err", err).Errorf("failed to connect to db:%s", a.DevSqlitePath)
			panic(err)
		}
		a.DB = db
	}

	a.DB.AutoMigrate(
		Todo{},
	)
}

func (a *App) SetupHTTP() {
	if a.Echo != nil {
		return
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// todos
	e.GET("/api/v1/todos", a.FindTodos)
	e.POST("/api/v1/todos", a.CreateTodo)
	e.PUT("/api/v1/todos/:id", a.UpdateTodo)
	e.DELETE("/api/v1/todos/:id", a.DeleteTodo)

	a.Echo = e
}

func (a *App) Start() error {
	return a.Echo.Start(
		fmt.Sprintf("%s:%s", a.Host, a.Port),
	)
}
