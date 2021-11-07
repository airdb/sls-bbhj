package schema

type LostListReq struct {
	Keyword  string `form:"keyword"`
	PageNo   int    `form:"pageNo"`
	PageSize int    `form:"pageSize"`
}

type LostListResp struct {
	Data    []*Lost `json:"data"`
	Success bool    `json:"success"`
}

type LostGetResp struct {
	Data    *Lost `json:"data"`
	Success bool  `json:"success"`
}
