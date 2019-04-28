package models

import (
	"github.com/pantazheng/bci/database"
)

func init(){
	database.DB.DropTable("projects")
	database.DB.DropTable("users")
	database.DB.DropTable("missions")
	database.DB.DropTable("gains")
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").AutoMigrate(&User{}).AutoMigrate(&Mission{}).AutoMigrate(&Project{}).AutoMigrate(&Gain{})
	userTestData()
	missionTestData()
	gainTestData()
}
