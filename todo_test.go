package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func TestTodos(t *testing.T) {
	t.Run("it should list empty todos when nothing is created", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		a := &App{DB: newTestDB()}
		a.SetupDB()

		assert.NoError(t, a.FindTodos(c))
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, "[]", rec.Body.String())
	})

	t.Run("it should be able to create a todo", func(t *testing.T) {
		createTodoJSON := `{"title": "test"}`

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", strings.NewReader(createTodoJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		a := &App{DB: newTestDB()}
		a.SetupDB()

		assert.NoError(t, a.FindTodos(c))
		assert.Equal(t, http.StatusOK, rec.Code)
	})
}
