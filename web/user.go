package web

import (
	"net/http"

	"github.com/airdb/mina-api/model/vo"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var resp vo.UserResp

	resp.Nickname = "xxx"
	resp.Token = "bb"
	c.JSON(http.StatusOK, resp)
}
