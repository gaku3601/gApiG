package api

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/tidwall/gjson"
)

func Index(c echo.Context) error {
	offset := c.QueryParam("offset")
	limit := c.QueryParam("limit")
	return c.String(http.StatusOK, fmt.Sprintf("Hello, World!\noffset:%s\nlimit:%s\n", offset, limit))
}

func New(c echo.Context) error {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(c.Request().Body())
	body := bufbody.String()
	name := gjson.Get(body, "name")
	email := gjson.Get(body, "email")

	return c.JSON(http.StatusOK, "{\"name\":\"gaku\"}")
}
