package web

import (
	"net/http"

	"github.com/airdb/mina-api/model/vo"
	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
)

func ListLost(c *gin.Context) {
	resp := vo.ListLost()

	middlewares.SetResp(
		c,
		enum.AirdbSuccess,
		resp,
	)
}

/*
// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *LostController) GetAll() {
	// var state models.State
	var ret models.RetDataList
	ret.State.Code = 200
	ret.State.Status = "success"
	ret.State.Message = ""

	u.Data["json"] = ret

	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /list [get]
func (u *LostController) List() {
	category, _ := u.GetInt("category")
	page, _ := u.GetInt("page")
	pageSize, _ := u.GetInt("pageSize")

	beego.Error("=======", category, page, pageSize)
	ulist, _ := models.GetAllBabyinfo(category, page, pageSize)
	u.Data["json"] = ulist
	// var state models.State
	var ret models.RetDataList
	ret.State.Code = 200
	ret.State.Status = "success"
	ret.State.Message = ""

	// u.Data["json"] = map[string]string{"State": state}
	ret.Data = ulist
	u.Data["json"] = ret

	u.ServeJSON()
}

// @Title Get
// @Description get lost info by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /detail/:id [get]
func (u *LostController) Get() {

	id, _ := u.GetInt(":id")
	info, _ := models.GetBabyinfoById(id)

	var ret models.RetDataList
	ret.State.Code = 200
	ret.State.Status = "success"
	ret.State.Message = ""

	ret.Data = info
	u.Data["json"] = ret
	u.ServeJSON()
}

// @Title Search lost info
// @Description get lost info by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /search [post]
func (u *LostController) Search() {

	var condition models.SearchBabyinfo
	json.Unmarshal(u.Ctx.Input.RequestBody, &condition)

	beego.Error(condition)
	ulist, _ := models.GetAllBabyinfoByCondition(condition.Keywords)

	var ret models.RetDataList
	ret.State.Code = 200
	ret.State.Status = "success"
	ret.State.Message = ""

	ret.Data = ulist
	u.Data["json"] = ret
	u.ServeJSON()
}
*/

// QueryBBS godoc
// @Summary for QQ robot query article
// @Description Query article
// @Tags article
// @Accept  json
// @Produce  json
// @Param req body vo.QQRobotQueryReq true "Message"
// @Success 200 {string} string vo.QQRobotQueryResp
// @Router /robot/query [get]
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
