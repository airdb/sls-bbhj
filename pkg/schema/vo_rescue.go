package schema

type RescueListReq struct {
	PageNo   int `form:"pageNo"`
	PageSize int `form:"pageSize"`
}

type RescueListResp struct {
	Data    []*Rescue `json:"data"`
	Success bool      `json:"success"`
}

type RescueQueryResp struct {
	Data    *Rescue `json:"data"`
	Success bool    `json:"success"`
}
