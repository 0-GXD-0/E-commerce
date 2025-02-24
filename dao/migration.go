package dao

import (
	"E-commerce/model"
	"fmt"
)

func Migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Address{},
			&model.Admin{},
			&model.BasePage{},
			&model.Carousel{},
			&model.Cart{},
			&model.Category{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.ProductImg{},
			&model.Product{},
			&model.User{},
		)
	if err != nil {
		fmt.Println("Migration error:", err)
	}
	return
}
