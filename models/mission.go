package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"time"
)

const titleMission = "models.mission."

type Mission struct {
	gorm.Model
	Name         string
	CreatorID    uint
	Creator      User
	CreateTime   string
	StartTime    string
	EndTime      string
	Content      string
	File         string
	Tag          bool
	Participants []*User `gorm:"many2many:user_missions"`
	ModuleID     uint
	Module       Module
}

//Create 创建Mission, 不添加成员
func (mission *Mission) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 3:48
	*/
	mission.CreateTime = time.Now().Format("2006-01-02")
	if err = database.DB.Create(&mission).Error; err != nil {
		err = errors.New(titleMission + "Create:\t" + err.Error())
	}
	return
}

func (mission *Mission) AppendParticipants(participants []*User) (err error) {
	m := &Mission{}
	m.ID = mission.ID
	if err = database.DB.Model(&m).Association("Participants").Append(participants).Error; err != nil {
		err = errors.New(titleMission + "AppendParticipants:\t" + err.Error())
	} else {
		*mission = *m
	}
	return
}

//func MissionFind(mission *Mission) (recordMissionJson MissionJson, err error) {
//	recordMission := new(Mission)
//	if err = database.DB.First(&recordMission, &mission).Error; err == nil {
//		recordMissionJson.mission2MissionJSON(recordMission)
//	}
//	return
//}
//
//func MissionsFindByModule(module *Module) (missionsBriefJson []MissionBriefJson, err error) {
//	missions := make([]Mission, 1)
//	if err = database.DB.Model(&module).Related(&missions, "ModuleID").Error; err != nil {
//		return
//	}
//	if len(missions) == 0 {
//		err = errors.New("MissionsFindByModule No Mission Record")
//	} else {
//		for _, v := range missions {
//			tempJson := &MissionBriefJson{}
//			tempJson.mission2MissionBriefJson(&v)
//			missionsBriefJson = append(missionsBriefJson, *tempJson)
//		}
//	}
//	return
//}
//
//func MissionUpdate(missionJson *MissionJson) (recordMissionJson MissionJson, err error) {
//	updateMission := new(Mission)
//	updateMission.missionJson2Mission(missionJson)
//	recordMission := new(Mission)
//	recordMission.ID = updateMission.ID
//	if database.DB.First(&recordMission, &recordMission).RecordNotFound() {
//		err = errors.New("MissionUpdate No Mission Record")
//	} else {
//		database.DB.Model(&recordMission).Updates(updateMission)
//		if num := len(missionJson.Participants); num != 0 {
//			users := make([]User, num)
//			for i, v := range missionJson.Participants {
//				users[i].ID = v.ID
//			}
//			err = database.DB.Model(&recordMission).Association("Participants").Replace(users).Error
//		}
//		recordMissionJson.mission2MissionJSON(recordMission)
//	}
//	return
//}
//
//func MissionDelete(mission *Mission) (recordMissionJson MissionJson, err error) {
//	recordMission := new(Mission)
//	if database.DB.Find(&recordMission, &mission).RecordNotFound() {
//		err = errors.New("MissionDelete No Mission Record")
//	} else {
//		recordMissionJson.mission2MissionJSON(recordMission)
//		err = database.DB.Unscoped().Delete(&recordMission).Error
//	}
//	return
//}
