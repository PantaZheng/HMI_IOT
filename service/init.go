package service

import (
	"github.com/pantazheng/bci/database"
	"github.com/pantazheng/bci/models"
)

/**
*@Author: PantaZheng
*@CreateAt: 2019/5/10 0:21
*@Title: init.go
*@Package: service
*@Description: service 初始化(用一句话描述该文件该做什么)
@Software: GoLand
*/

//func init() {
//	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&User{}, &Project{}, &Module{}, &Mission{}, &Gain{})
//	userTestData()
//	projectTestData()
//	moduleTestData()
//	missionTestData()
//	gainTestData()
//}

func init() {
	database.DB.DropTableIfExists("users").DropTableIfExists("gains")
	if database.DB.HasTable("users") {
		database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&models.User{}, &models.Gain{})
	} else {
		database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&models.User{}, &models.Gain{})
		userTestData()
		gainTestData()
	}
	//projectTestData()
	//moduleTestData()
	//missionTestData()
	//gainTestData()
}
