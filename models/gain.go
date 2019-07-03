package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
)

type Gain struct {
	gorm.Model
	Name   string
	Type   string
	File   string
	Remark string
	State  uint
	//foreign const
	MissionID uint
	LeaderID  uint
	OwnerID   uint
}

//Insert 必须包含MissionID
func (gain *Gain) Insert() (err error) {
	gain.State = 1
	mission := Mission{}
	mission.ID = gain.MissionID
	if err = mission.First(); err != nil {
		return
	}
	gain.LeaderID = mission.LeaderID
	gain.OwnerID = mission.OwnerID
	if err = database.DB.Create(&gain).Error; err != nil {
		return
	}
	err = gain.First()
	return
}

//First 根据id查找Gain.
func (gain *Gain) First() (err error) {
	if err = database.DB.Where("id = ? ", gain.ID).First(&gain).Error; err != nil {
		return
	}
	return
}

//FindGains
func (gain *Gain) Find(field string) (gains []Gain, err error) {
	if field == "leader_id" {
		err = database.DB.Where("leader_id = ?", gain.LeaderID).Find(&gains).Error
	} else if field == "owner_id" {
		err = database.DB.Where("owner_id = ?", gain.OwnerID).Find(&gains).Error
	} else if field == "mission_id" {
		err = database.DB.Where("mission_id = ?", gain.MissionID).Find(&gains).Error
	} else if field == "all" {
		err = database.DB.Model(Gain{}).Find(&gains).Error
	} else {
		err = errors.New("no this field")
	}
	return
}

//Updates 通用更新接口，ID必须，Uptime自动更新。
func (gain *Gain) Updates() (err error) {
	if err = database.DB.Model(Gain{}).Where("id=?", gain.ID).Updates(&gain).Error; err != nil {
		return
	}
	err = gain.First()
	return
}

//Delete
func (gain *Gain) Delete() (err error) {
	if err = gain.First(); err != nil {
		return
	}
	//硬删除
	err = database.DB.Model(Gain{}).Where("id=?", gain.ID).Delete(&gain).Error
	return
}
