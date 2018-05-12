package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestIndexApi(t *testing.T) {
	// setup
	con := fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable", "postgres", "app", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	db, _ := gorm.Open("postgres", con)
	defer db.Close()
	// apiテーブルの全削除
	db.Delete(&Api{})
	// apiテーブルへテストデータ格納
	api := &Api{
		Path:       "/test",
		Method:     "POST",
		AccessPath: "http://google.com/",
	}
	for i := 1; i < 10; i++ {
		api.ID = 0
		api.Name = fmt.Sprintf("%s:%d", "test api", i)
		db.Create(&api)
	}
	// GETリクエスト呼び出し
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/api?offset=2&limit=3", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	if assert.NoError(t, Index(c)) {
		// status code check
		assert.Equal(t, http.StatusOK, rec.Code)
		// limitチェック
		json := rec.Body.String()
		r := gjson.Parse(json)
		assert.Equal(t, 3, len(r.Array()))
		// offsetチェック(DESC IDも共にチェック)
		assert.Equal(t, "test api:7", r.Array()[0].Get("name").String())
	}
}

func TestCreateApi(t *testing.T) {
	// Setup
	postJSON := `{"name":"path title","path":"/Images","method":"POST","access_path":"http://google.com/"}`
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	// Assertions
	if assert.NoError(t, Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// jsonチェック
		name := gjson.Get(rec.Body.String(), "name").String()
		assert.Equal(t, "path title", name)
		path := gjson.Get(rec.Body.String(), "path").String()
		assert.Equal(t, "/Images", path)
		method := gjson.Get(rec.Body.String(), "method").String()
		assert.Equal(t, "POST", method)
		access_path := gjson.Get(rec.Body.String(), "access_path").String()
		assert.Equal(t, "http://google.com/", access_path)
	}
}

func TestShowApi(t *testing.T) {
}

func TestUpdateApi(t *testing.T) {
}

func TestDestroyApi(t *testing.T) {
}
