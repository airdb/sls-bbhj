package web

import (
	"net/http"

	"github.com/airdb/mina-api/model/vo"
	"github.com/airdb/sailor"
	"github.com/airdb/sailor/enum"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	var resp sailor.HTTPAirdbResponse
	resp.Data = vo.ListCategory()
	resp.Code = enum.AirdbSuccess
	resp.Success = true

	c.JSON(http.StatusOK, resp)
}
