package schema

type LostListRequest struct {
	Pagination
	Keyword string `form:"keyword"`
}

func (m *LostListRequest) Valadate() error {
	if err := m.Pagination.Valadate(); err != nil {
		return err
	}

	return nil
}

type LostListResponse struct {
	Data    []*Lost `json:"data"`
	Success bool    `json:"success"`
}

type LostGetResponse struct {
	Data    *Lost `json:"data"`
	Success bool  `json:"success"`
}
