package schema

// LostListRequest 失踪信息 列表请求
type LostListRequest struct {
	Pagination
	Keyword  string `form:"keyword"`  // 查询关键字
	Category string `form:"category"` // 分类ID
}

func (m *LostListRequest) Valadate() error {
	if err := m.Pagination.Valadate(); err != nil {
		return err
	}

	return nil
}

// LostListRequest 失踪信息 列表返回
type LostListResponse struct {
	Data    []*LostItem `json:"data"`
	Success bool        `json:"success"`
}

// LostListRequest 失踪信息 详情返回
type LostGetResponse struct {
	Data    *LostDetail `json:"data"`
	Success bool        `json:"success"`
}

// LostItem 失踪信息 列表条目
type LostItem struct {
	ID           uint   `json:"id"`           // 丢失序号
	Title        string `json:"title"`        // 标题
	Category     string `json:"category"`     // 寻找类别
	Name         string `json:"name"`         // 姓名
	Babyid       string `json:"babyid"`       // 失踪人员登记编号
	Introduction string `json:"introduction"` // 简介内容
	MissAt       string `json:"miss_at"`      // 丢失日期
	MissAddr     string `json:"miss_addr"`    // 地点
}

// LostDetail 失踪信息 详情页条目
type LostDetail struct {
	ID    uint   `json:"id"`    // 丢失序号
	Title string `json:"title"` // 标题

	Name         string `json:"name"`         // 姓名
	Babyid       string `json:"babyid"`       // 失踪人员登记编号
	Introduction string `json:"introduction"` // 简介内容
	ShareCount   uint   `json:"share_count"`  // 累计转发助力
	ShowCount    uint   `json:"show_count"`   // 累计曝光助力

	// 基础信息
	NameMore  string         `json:"name_more"`  // 姓名
	Gender    string         `json:"gender"`     // 性别
	BirthedAt string         `json:"birthed_at"` // 出生日期
	Carousel  []CarouselItem `json:"carousel"`   // 寻亲目标轮播图

	// 失踪信息
	MissAt     string `json:"miss_at"`     // 失踪时间
	MissAddr   string `json:"miss_addr"`   // 失踪地点
	MissHeight string `json:"miss_height"` // 失踪时身高
	Character  string `json:"character"`   // 特征

	// 寻亲信息
	Category string `json:"category"`  // 寻亲类型
	DataFrom string `json:"data_from"` // 信息来源
	Follower string `json:"follower"`  // 跟进志愿者

	// 第三方交互
	WxMore *WxMore `json:"wx_more"` // 微信相关信息
}
