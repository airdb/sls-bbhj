package schema

import "net/http"

// PassportLoginRequest 登录信息 列表请求
type PassportLoginRequest struct {
	Code string `form:"bode"` // 查询关键字
}

func (m PassportLoginRequest) Bind(r *http.Request) error {
	return nil
}

func (m *PassportLoginRequest) Valadate() error {
	return nil
}

// PassportLoginRequest 登录信息 列表返回
type PassportLoginResponse struct {
	Token   string `json:"token"`
	Success bool   `json:"success"`
}
