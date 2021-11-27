package schema

type CategoryItem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type CategoryListRequest struct {
	Pagination
	Keyword string `form:"keyword"`
}

func (m *CategoryListRequest) Valadate() error {
	if err := m.Pagination.Valadate(); err != nil {
		return err
	}

	return nil
}

type CategoryListResponse struct {
	Data    []*CategoryItem `json:"data"`
	Success bool            `json:"success"`
}
