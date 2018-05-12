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

type User struct {
	Name string `json:"name" form:"name"`
}

func New(c echo.Context) error {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(c.Request().Body())
	body := bufbody.String()
	name := gjson.Get(body, "name")
	path := gjson.Get(body, "path")
	method := gjson.Get(body, "method")
	accessPath := gjson.Get(body, "access_path")
	fmt.Println(name)
	fmt.Println(path)
	fmt.Println(method)
	fmt.Println(accessPath)
	u := new(User)
	u.Name = "gaku"

	return c.JSON(http.StatusOK, u)
}
