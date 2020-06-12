package web

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/airdb/mina-api/model/vo"
	"github.com/airdb/sailor"
	"github.com/airdb/sailor/enum"
	"github.com/gin-gonic/gin"
)

func ListLost(c *gin.Context) {
	var req vo.ListLostReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &sailor.HTTPAirdbResponse{
			Code:    enum.AirdbUndefined,
			Success: false,
			Data:    nil,
			Error:   "不合法的请求参数",
		})

		return
	}

	log.Println(req.Category, req.Page, req.PageSize)

	var resp sailor.HTTPAirdbResponse

	resp.Data = vo.ListLost(req)
	resp.Code = enum.AirdbSuccess
	resp.Success = true

	c.JSON(http.StatusOK, resp)
}

func QueryLost(c *gin.Context) {
	// id := c.Param("id")
	uint64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		fmt.Print(err)
	}

	id := uint(uint64)

	var resp sailor.HTTPAirdbResponse
	resp.Data = vo.QueryLost(&id)
	resp.Code = enum.AirdbSuccess
	resp.Success = true

	c.JSON(http.StatusOK, resp)
}

func SearchLost(c *gin.Context) {
	var req vo.SearchLostReq

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, &sailor.HTTPAirdbResponse{
			Code:    enum.AirdbUndefined,
			Success: false,
			Data:    nil,
			Error:   "不合法的请求参数",
		})

		return
	}

	resp := vo.SearchLost(req.Keywords)

	c.JSON(http.StatusOK, &sailor.HTTPAirdbResponse{
		Code:    enum.AirdbSuccess,
		Success: true,
		Data:    resp,
	})
}

// QueryBBS godoc
// @Summary for QQ robot query article
// @Description Query article
// @Tags article
// @Accept  json
// @Produce  json
// @Param req body vo.QQRobotQueryReq true "Message"
// @Success 200 {string} string vo.QQRobotQueryResp
// @Router /robot/query [get].
func QueryBBS(c *gin.Context) {
	var req vo.QQRobotQueryReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.String(http.StatusBadRequest, "不合法的请求参数")
		return
	}

	var msg string
	if req.Message == "" {
		// help
		msg += "bbhj 机器人使用帮助\n"
		msg += "示例1：bbhj 4407\n"
		msg += "示例2：bbhj 山西 4407\n"
		msg += "示例3：bbhj 山西 运城 张\n"
		msg += "\n"
		msg += "说明：bbhj命令支持最多3个关键字的查询; 命令及各关键字只能以空格分隔。"
	} else {
		articles := vo.QueryBBSByKeyword(req.Message)
		if len(articles) == 0 {
			msg += "您的查询的信息，暂时无结果，可能是后台同步论坛数据失败。\n"
		} else {
			for _, article := range articles {
				msg += article.Subject + "\n" + article.DataFrom + "\n"
			}
		}
	}

	msg += "(出处: 宝贝回家论坛)\n"
	c.String(http.StatusOK, msg)
}
