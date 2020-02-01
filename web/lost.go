package web

import (
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
