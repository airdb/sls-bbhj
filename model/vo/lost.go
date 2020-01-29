package vo

type LostListReq struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

type LostListResp struct {
}

type LostQueryReq struct {
}

type LostQueryResp struct {
}

type LostSearchReq struct {
}

type LostSearchResp struct {
}

func ListLost() *CategoryListResp {
	return nil
}
