package vo

import (
	"time"

	"github.com/airdb/mina-api/model/po"
)

type CategoryListReq struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

type CategoryListResp struct {
	Categories []*Category
}

type Category struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}

func FromPoCategory(category *po.Category) *Category {
	if category == nil {
		return nil
	}

	return &Category{
		ID:          category.ID,
		CreatedAt:   category.CreatedAt,
		UpdatedAt:   category.UpdatedAt,
		Name:        category.Name,
		Description: category.Description,
	}
}

func FromPoCategories(categories []*po.Category) []*Category {
	_categories := make([]*Category, 0, len(categories))

	for _, category := range categories {
		_categories = append(_categories, FromPoCategory(category))
	}

	return _categories
}

func ListCategory() []*Category {
	categories := []*Category{}
	categories = append(categories, FromPoCategories(po.ListCateories())...)

	return categories
}
