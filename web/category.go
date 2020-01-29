package web

import (
	"github.com/airdb/mina-api/model/vo"
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	resp := vo.ListCategory()

	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		resp,
	)
}
