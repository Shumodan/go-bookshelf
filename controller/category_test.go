package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Shumodan/go-bookshelf/model"
	"github.com/Shumodan/go-bookshelf/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetCategoryList(t *testing.T) {
	router, context := test.Prepare()

	category := NewCategoryController(context)
	router.GET(APICategories, func(c echo.Context) error { return category.GetCategoryList(c) })

	req := httptest.NewRequest("GET", APICategories, nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	data := [...]*model.Category{
		{ID: 1, Name: "技術書"},
		{ID: 2, Name: "雑誌"},
		{ID: 3, Name: "小説"},
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
}
