package web

import (
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"

	"net/http"
)

func Status(c *gin.Context) {
	c.String(http.StatusOK, "hello\n")
	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		"",
	)
}
