package models

import (
	"github.com/pantazheng/bci/database"
)

func init() {
	database.DB.DropTable("users", "projects", "modules", "missions", "gains")
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&User{}, &Project{}, &Module{}, &Mission{}, &Gain{})
	userTestData()
	projectTestData()
	moduleTestData()
	missionTestData()
	gainTestData()
}
