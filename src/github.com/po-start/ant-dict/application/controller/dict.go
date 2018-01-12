package controller

import (
	//"fmt"
//	"encoding/json"
	"net/http"
//	"strconv"

	"github.com/labstack/echo"

	"github.com/po-start/ant-dict/application/service"
)

type DictController struct {
	BaseController
}

var dictService = &service.DictService{}

func (d *DictController) Search (c echo.Context) error {
	key := c.FormValue("key")
	if key != "" {
		data := dictService.Search(key).([]map[string]interface{})
		if len(data) > 0 {
			return c.JSON(http.StatusOK, data)
		}
	}
	return nil
}
