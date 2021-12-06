package schema

type Pagination struct {
	PageNo   int `form:"pageNo"`
	PageSize int `form:"pageSize"`
}

func (m *Pagination) Valadate() error {
	if m.PageNo == 0 {
		m.PageNo = 1
	}
	if m.PageSize == 0 || m.PageSize > 100 {
		m.PageSize = 20
	}

	return nil
}

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}
