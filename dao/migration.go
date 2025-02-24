package dao

import "fmt"

func Migration() {
	err := _db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate()
	if err != nil {
		fmt.Println("Migration error:", err)
	}
	return
}
