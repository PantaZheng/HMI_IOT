package service

import (
	"github.com/pantazheng/bci/database"
	"github.com/pantazheng/bci/models"
)

func init() {
	database.DB.DropTable(&models.User{}, &models.Gain{}, models.Mission{}, models.Module{}, &models.Project{})
	database.DB.Set("gorm:table_options", "DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci AUTO_INCREMENT=1;").AutoMigrate(&models.User{}, &models.Gain{}, models.Mission{}, models.Module{}, &models.Project{})
	userTestData()
	projectTestData()
	moduleTestData()
	missionTestData()
	gainTestData()
}
