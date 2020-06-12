package po

import (
	"github.com/airdb/sailor/dbutils"
)

type PreForumPost struct {
	// gorm.Model
	Message  string `json:"message"`
	Subject  string `json:"subject"`
	Useip    string `json:"useip"`
	Pid      int64  `json:"pid"`
	Tid      uint   `json:"tid"`
	Authorid int64  `json:"authorid"`
}

func GetBBSArticles() (preForumPost []PreForumPost) {
	sqltext := ""
	sqltext = "select * from pre_forum_post where  subject != '' "
	// sqltext += " and message like '%登记信息%宝贝回家编号%' "
	// sqltext += " and message like '%登记信息%编号%' "
	sqltext += " and message like '%本网站及志愿者提供的寻人服务均是免费%' "
	//sqltext += " and subject like '%3313%' "
	// sqltext += " and tid = 193856 "
	sqltext += " order by pid desc"
	// sqltext += " limit 10 offset 0"
	dbutils.ReadDB(dbBbhjBBSRead).Raw(sqltext).Scan(&preForumPost)

	return
}

func GetBBSArticleByID(tid uint) *PreForumPost {
	var preForumPost PreForumPost

	dbutils.ReadDB(dbBbhjBBSRead).Table("pre_forum_post").Where("tid = ? and first is true", tid).First(&preForumPost)

	return &preForumPost
}
