package schema

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

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

// LostGetResponse 失踪信息 详情返回
type LostGetResponse struct {
	Data    *LostDetail `json:"data"`
	Success bool        `json:"success"`
}

type LostMpCode struct {
}

// LostGetMpCodeResponse 失踪信息 详情返回
type LostGetMpCodeResponse struct {
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
	// MissAt       string `json:"miss_at"`      // 丢失日期
	// MissAddr     string `json:"miss_addr"`    // 地点
	MissedAt      string `json:"missed_at"`      // 丢失日期
	MissedAddress string `json:"missed_address"` // 地点
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
	Details    string `json:"details"`     // 失踪详情

	// 寻亲信息
	Category string `json:"category"`  // 寻亲类型
	DataFrom string `json:"data_from"` // 信息来源
	Follower string `json:"follower"`  // 跟进志愿者

	// 第三方交互
	WxMore *WxMore `json:"wx_more"` // 微信相关信息
}

// LostCreateRequest 失踪信息 录入
type LostCreateRequest struct {
	// 基础信息
	Name      string    `json:"name"`       // 姓名
	Gender    uint      `json:"gender"`     // 性别: 1男 2女 0未知
	BirthedAt time.Time `json:"birthed_at"` // 出生日期
	Carousel  []string  `json:"carousel"`   // 寻亲目标轮播图

	// 失踪信息
	MissedAt       time.Time `json:"missed_at"`       // 失踪时间
	MissedCountry  string    `json:"missed_country"`  // 失踪国家
	MissedProvince string    `json:"missed_province"` // 失踪省
	MissedCity     string    `json:"missed_city"`     // 失踪市
	MissedAddr     string    `json:"missed_addr"`     // 详细地址
	MissedHeight   string    `json:"missed_height"`   // 失踪时身高
	Character      string    `json:"character"`       // 特征
	Details        string    `json:"details"`         // 失踪详情

	// 寻亲信息
	Category string `json:"category"`  // 寻亲类型
	DataFrom string `json:"data_from"` // 信息来源
	Follower string `json:"follower"`  // 跟进志愿者

	// 关联图片
	Images []string `json:"images"` // 图片列表
}

func (m *LostCreateRequest) Valadate() error {
	if len(m.Name) == 0 {
		return fmt.Errorf("请输入 姓名")
	}

	if m.BirthedAt.IsZero() {
		return fmt.Errorf("请输入 出生日期")
	}

	if m.MissedAt.IsZero() {
		return fmt.Errorf("请输入 失踪时间")
	}

	if len(m.MissedAddr) == 0 {
		return fmt.Errorf("请输入 失踪地点")
	}

	if len(m.MissedHeight) == 0 {
		return fmt.Errorf("请输入 失踪时身高")
	}

	if len(m.Character) == 0 {
		return fmt.Errorf("请输入 特征")
	}

	if len(m.Details) == 0 {
		return fmt.Errorf("请输入 失踪详情")
	}

	if len(m.Category) == 0 {
		return fmt.Errorf("请输入 寻亲类型")
	}

	if len(m.DataFrom) == 0 {
		return fmt.Errorf("请输入 信息来源")
	}

	if len(m.Follower) == 0 {
		return fmt.Errorf("请输入 跟进志愿者")
	}

	log.Println(m.MissedAt, m.MissedAt.IsZero())

	return nil
}

func (m LostCreateRequest) Bind(r *http.Request) error {
	return nil
}

// LostCreateResponse 失踪信息 录入
type LostCreateResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

// LostGetPresignedURLRequest 失踪信息 上传请求
type LostGetPresignedURLRequest struct {
	// 基础信息
	Filename string `json:"filename"` // 文件名
}

func (m *LostGetPresignedURLRequest) Valadate() error {
	if len(m.Filename) == 0 {
		return fmt.Errorf("请输入 文件名")
	}

	return nil
}

func (m LostGetPresignedURLRequest) Bind(r *http.Request) error {
	return nil
}

type LostGetPresignedURL struct {
	URL string `json:"url"`
}

// LostGetPresignedURLResponse 失踪信息 上传请求
type LostGetPresignedURLResponse struct {
	Message string              `json:"message"`
	Success bool                `json:"success"`
	Data    LostGetPresignedURL `json:"data"`
}
