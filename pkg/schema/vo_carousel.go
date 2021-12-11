package schema

// CarouselItem 轮播图条目
type CarouselItem struct {
	ID    uint   `json:"id"`    // 序号
	Title string `json:"title"` // 图片名称
	URL   string `json:"url"`   // 图片链接
}
