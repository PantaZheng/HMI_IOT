package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
)

type MissionCore struct {
	gorm.Model
	Name      string
	StartTime string
	EndTime   string
	Content   string
	Target    string
	File      string
	State     uint

	//foreign
	OwnerID uint
}

type Mission struct {
	MissionCore
	//foreign const
	ModuleID uint
	LeaderID uint
}

//Insert 创建Mission
func (mission *Mission) Insert() (err error) {
	mission.State = 1
	module := Module{}
	module.ID = mission.ModuleID
	if err = module.First(); err != nil {
		return
	}
	mission.LeaderID = module.LeaderID
	if err = database.DB.Create(&mission).Error; err != nil {
		return
	}
	err = mission.First()
	return
}

//First 根据id查找Mission.
func (mission *Mission) First() (err error) {
	err = database.DB.Where("id = ? ", mission.ID).First(&mission).Error
	return
}

//FindMissions
func (mission *Mission) Find(field string) (missions []Mission, err error) {
	if field == "leader_id" {
		err = database.DB.Where("leader_id = ?", mission.LeaderID).Find(&missions).Error
	} else if field == "owner_id" {
		err = database.DB.Where("owner_id = ?", mission.OwnerID).Find(&missions).Error
	} else if field == "module_id" {
		err = database.DB.Where("module_id = ?", mission.ModuleID).Find(&missions).Error
	} else if field == "all" {
		err = database.DB.Model(Mission{}).Find(&missions).Error
	} else {
		err = errors.New("no this field")
	}
	return
}

//Updates ID必须，Uptime自动更新
func (mission *Mission) Update() (err error) {
	if err = database.DB.Where("id=?", mission.ID).Updates(&mission).Error; err != nil {
		return
	}
	err = mission.First()
	return
}

func (mission *Mission) Delete() (err error) {
	if err = mission.First(); err != nil {
		return
	}
	//硬删除
	err = database.DB.Model(Mission{}).Where("id=?", mission.ID).Delete(&mission).Error
	return
}
