package schema

// WxMore 微信相关信息
type WxMore struct {
	// Refer: https://developers.weixin.qq.com/miniprogram/dev/reference/api/Page.html#onShareTimeline
	ShareAppMessage *WxShareAppMessage `json:"share_app_message"` // 分享到对话
	ShareTimeline   *WxShareTimeline   `json:"share_timeline"`    // 分享到朋友圈
	CodeUnlimit     *WxCodeUnlimit     `json:"code_unlimit"`      // 小程序二维码
}

// WxShareAppMessage 分享到对话
type WxShareAppMessage struct {
	ShareKey string `json:"share_key"` // 用于分享后的通知回传
	Title    string `json:"title"`     // 转发标题
	ImageURL string `json:"image_url"` // 自定义图片路径
}

// WxShareTimeline 分享到朋友圈
type WxShareTimeline struct {
	ShareKey string `json:"share_key"` // 用于分享后的通知回传
	Title    string `json:"title"`     // 自定义标题
	Query    string `json:"query"`     // 自定义页面路径中携带的参数
	ImageURL string `json:"image_url"` // 自定义图片路径
}

// WxCodeUnlimit 小程序二维码图片链接
type WxCodeUnlimit struct {
	URL string `json:"url"` // 图片链接
}
