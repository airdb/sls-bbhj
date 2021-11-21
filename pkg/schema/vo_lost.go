package schema

type LostListReq struct {
	Keyword  string `form:"keyword"`
	PageNo   int    `form:"pageNo"`
	PageSize int    `form:"pageSize"`
}

func (m *LostListReq) Valadate() {
	if m.PageNo == 0 {
		m.PageNo = 1
	}
	if m.PageSize == 0 || m.PageSize > 100 {
		m.PageSize = 20
	}
}

type LostListResp struct {
	Data    []*Lost `json:"data"`
	Success bool    `json:"success"`
}

type LostGetResp struct {
	Data    *Lost `json:"data"`
	Success bool  `json:"success"`
}
