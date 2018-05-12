package api

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/tidwall/gjson"
)

func Index(c echo.Context) error {
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	list := indexApi(offset, limit)
	return c.JSON(http.StatusOK, list)
}

func Create(c echo.Context) error {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(c.Request().Body())
	body := bufbody.String()
	name := gjson.Get(body, "name").String()
	path := gjson.Get(body, "path").String()
	method := gjson.Get(body, "method").String()
	accessPath := gjson.Get(body, "access_path").String()
	obj, err := createApi(name, path, method, accessPath)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, obj)
}
