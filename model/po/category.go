package po

import (
	"log"

	"github.com/airdb/sailor/dbutils"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ListCateories() []*Category {
	var c []*Category

	dbutils.DefaultDB().Debug().Find(&c)
	log.Println(c[0])

	return c
}

func QueryCategory() *Category {
	var c Category

	dbutils.DefaultDB().First(&c)

	return &c
}

/*
func AddCategory(c Category) (err error) {
	conn.FirstOrCreate(&c, Category{Name: c.Name})
	return
}
func GetCategory(id int) (c Category, err error) {
	dbutils.DefaultDB().First(&c, id)
	if conn.Error != nil && conn.RowsAffected > 0 {
		err = nil
	}
	return
}
*/
