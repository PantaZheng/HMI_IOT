package models

import "github.com/pantazheng/bci/database"

/**
*@Author: PantaZheng
*@CreateAt: 2019/7/5 0:25
*@Title: common.go
*@Package: models
*@Description: (用一句话描述该文件该做什么)
@Software: GoLand
*/

func init() {
	database.DB.DropTable(&User{}, &Gain{}, Mission{}, Module{}, Project{})
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=1;").AutoMigrate(User{}, Gain{}, Mission{}, Module{}, Project{})
	userTestData()
	projectTestData()
	moduleTestData()
	missionTestData()
	gainTestData()
}
