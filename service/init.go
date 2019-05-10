package service

import "github.com/pantazheng/bci/database"

/**
*@Author: PantaZheng
*@CreateAt: 2019/5/10 0:21
*@Title: init.go
*@Package: service
*@Description: service 初始化(用一句话描述该文件该做什么)
@Software: GoLand
*/

func init() {
	if !database.DB.HasTable("users") {
		userTestData()
	}

	//projectTestData()
	//moduleTestData()
	//missionTestData()
	//gainTestData()
}
