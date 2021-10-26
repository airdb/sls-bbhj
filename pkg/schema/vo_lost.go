package schema

type LostListReq struct {
	PageNo   int `form:"pageNo"`
	PageSize int `form:"pageSize"`
}

type LostListResp struct {
	Data    []*Lost `json:"data"`
	Success bool    `json:"success"`
}

type LostQueryResp struct {
	Data    *Lost `json:"data"`
	Success bool  `json:"success"`
}
