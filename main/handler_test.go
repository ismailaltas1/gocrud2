package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
func TestGetUsers(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := getUsers(c)

	assert.NoError(t, err)
	assert.Equal(t, 200, rec.Code)

}
*/
type MockBooksRepository struct {
	mock.Mock
}

func (m *MockBooksRepository) GetBooks(ctx context.Context) ([]Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]Book), args.Error(1)
}

func (m *MockBooksRepository) InsertBooks(ctx context.Context, book Book) error {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func TestGetBooks(t *testing.T) {
	e := echo.New()

	mockBooksRepo := new(MockBooksRepository)
	mockBooks := []Book{
		{ID: "1", Title: "Book1", Author: "Author1"},
		{ID: "2", Title: "Book2", Author: "Author2"},
	}

	mockBooksRepo.On("GetBooks", mock.Anything).Return(mockBooks, nil)

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := BookHandler(mockBooksRepo)
	err := h.getBooks(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `[{"_id":"1","title":"Book1","author":"Author1"},{"_id":"2","title":"Book2","author":"Author2"}]`, rec.Body.String())
	mockBooksRepo.AssertExpectations(t)
}

func TestCreateBook(t *testing.T) {
	e := echo.New()

	mockBooksRepo := new(MockBooksRepository)
	mockBook := Book{ID: "1", Title: "Book1", Author: "Author1"}

	mockBooksRepo.On("InsertBooks", mock.Anything, mockBook).Return(nil)

	bookJson, _ := json.Marshal(mockBook)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(bookJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := BookHandler(mockBooksRepo)
	err := h.createBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"_id":"1","title":"Book1","author":"Author1"}`, rec.Body.String())
	mockBooksRepo.AssertExpectations(t)
}
