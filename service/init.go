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

func init() {
	database.DB.DropTableIfExists("users", "gains", "missions", "modules", "projects")
	if database.DB.HasTable("users") {
		database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&models.User{}, &models.Gain{}, models.Mission{}, &models.Module{}, &models.Project{})
	} else {
		database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 AUTO_INCREMENT=1;").AutoMigrate(&models.User{}, &models.Gain{}, models.Mission{}, models.Module{}, &models.Project{})
		userTestData()
		projectTestData()
		moduleTestData()
		missionTestData()
		gainTestData()
	}
}
