package schema

// CategoryItem 失踪信息 分类条目
type CategoryItem struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// CategoryItem 失踪信息 列表请求
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

// CategoryItem 失踪信息 列表返回
type CategoryListResponse struct {
	Data    []*CategoryItem `json:"data"`
	Success bool            `json:"success"`
}
