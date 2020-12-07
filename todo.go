package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (a *App) FindTodos(c echo.Context) error {
	ts := []Todo{}

	err := a.DB.Find(&ts).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ts)
}

func (a *App) CreateTodo(c echo.Context) error {
	t := &Todo{}
	if err := c.Bind(t); err != nil {
		return err
	}

	err := a.DB.Save(t).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, t)
}

func (a *App) UpdateTodo(c echo.Context) error {
	t := &Todo{}
	if err := c.Bind(t); err != nil {
		return err
	}

	err := a.DB.Model(t).Updates(map[string]interface{}{
		"id":    c.Param("id"),
		"done":  t.Done,
		"title": t.Title,
	}).Error
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, t)
}

func (a *App) DeleteTodo(c echo.Context) error {
	err := a.DB.Delete(&Todo{}, c.Param("id")).Error
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("todo id %s deleted", c.Param("id")))
}
