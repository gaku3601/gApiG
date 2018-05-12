package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestCreateUser(t *testing.T) {
	// Setup
	postJSON := `{"name":"path title","path":"/Images","method":"POST","access_path":"http://google.com/"}`
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))

	// Assertions
	if assert.NoError(t, New(c)) {
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
