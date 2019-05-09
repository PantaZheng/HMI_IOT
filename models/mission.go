package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"time"
)

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

type MissionJson struct {
	ID           uint            `json:"id"`
	Name         string          `json:"name"`
	Creator      UserBriefJSON   `json:"creator"`
	CreateTime   string          `json:"createTime"`
	StartTime    string          `json:"startTime"`
	EndTime      string          `json:"endTime"`
	Content      string          `json:"content"`
	File         string          `json:"file"`
	Tag          bool            `json:"tag"` //tag由module负责人决定
	Gains        []GainJson      `json:"gains"`
	Participants []UserBriefJSON `json:"participants"`
	ModuleID     uint            `json:"module"`
}

type MissionBriefJson struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
	Content    string `json:"content"`
	Tag        bool   `json:"tag"`
	ModuleID   uint   `json:"module"`
}

func missionTestData() {
	u2 := &UserBriefJSON{ID: 2}
	u3 := &UserBriefJSON{ID: 3}
	u4 := &UserBriefJSON{ID: 4}
	u5 := &UserBriefJSON{ID: 5}
	u6 := &UserBriefJSON{ID: 6}
	_, _ = MissionCreate(&MissionJson{Name: "Mission1", Creator: *u2, ModuleID: 1, Participants: []UserBriefJSON{*u2, *u3}})
	_, _ = MissionCreate(&MissionJson{Name: "Mission2", Creator: *u3, ModuleID: 1, Participants: []UserBriefJSON{*u2, *u4}})
	_, _ = MissionCreate(&MissionJson{Name: "Mission3", Creator: *u4, ModuleID: 2, Participants: []UserBriefJSON{*u2, *u3, *u4}})
	_, _ = MissionCreate(&MissionJson{Name: "Mission4", Creator: *u5, ModuleID: 2, Participants: []UserBriefJSON{*u2, *u3, *u6}})
}

//缺失participants
func (mission *Mission) missionJson2Mission(missionJson *MissionJson) {
	mission.ID = missionJson.ID
	mission.Name = missionJson.Name
	mission.CreatorID = missionJson.Creator.ID
	mission.CreateTime = missionJson.CreateTime
	mission.StartTime = missionJson.StartTime
	mission.EndTime = missionJson.EndTime
	mission.Content = missionJson.Content
	mission.File = missionJson.File
	mission.Tag = missionJson.Tag
	mission.ModuleID = missionJson.ModuleID
}

func (missionJson *MissionJson) mission2MissionJSON(mission *Mission) {
	missionJson.ID = mission.ID
	missionJson.Name = mission.Name
	creator := &User{}
	database.DB.Model(&mission).Related(&creator, "CreatorID")
	missionJson.Creator.User2UserBriefJSON(creator)
	missionJson.CreateTime = mission.CreateTime
	missionJson.StartTime = mission.StartTime
	missionJson.EndTime = mission.EndTime
	missionJson.Content = mission.Content
	missionJson.File = mission.File
	missionJson.Tag = mission.Tag
	missionJson.ModuleID = mission.ModuleID
	participants := make([]*User, len(mission.Participants))
	database.DB.Model(&mission).Related(&participants, "Participants")
	tempUser := &UserBriefJSON{}
	for _, v := range participants {
		tempUser.User2UserBriefJSON(v)
		missionJson.Participants = append(missionJson.Participants, *tempUser)
	}
	missionJson.Gains, _ = GainsFindByMission(mission)
}

func (missionBriefJson *MissionBriefJson) mission2MissionBriefJson(mission *Mission) {
	missionBriefJson.ID = mission.ID
	missionBriefJson.Name = mission.Name
	missionBriefJson.CreateTime = mission.CreateTime
	missionBriefJson.Content = mission.Content
	missionBriefJson.Tag = mission.Tag
	missionBriefJson.ModuleID = mission.ModuleID
}

func MissionCreate(missionJson *MissionJson) (recordMissionJson MissionJson, err error) {
	newMission := new(Mission)
	newMission.missionJson2Mission(missionJson)
	newMission.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	if err = database.DB.Create(&newMission).Error; err != nil {
		return
	}
	if err = database.DB.Model(&newMission).First(&newMission).Error; err == nil {
		users := make([]User, len(missionJson.Participants))
		for i, v := range missionJson.Participants {
			users[i].ID = v.ID
		}
		err = database.DB.Model(&newMission).Association("Participants").Append(users).Error
		recordMissionJson.mission2MissionJSON(newMission)
	}
	return
}

func MissionFind(mission *Mission) (recordMissionJson MissionJson, err error) {
	recordMission := new(Mission)
	if err = database.DB.First(&recordMission, &mission).Error; err == nil {
		recordMissionJson.mission2MissionJSON(recordMission)
	}
	return
}

func MissionsFindByModule(module *Module) (missionsBriefJson []MissionBriefJson, err error) {
	missions := make([]Mission, 1)
	if err = database.DB.Model(&module).Related(&missions, "ModuleID").Error; err != nil {
		return
	}
	if len(missions) == 0 {
		err = errors.New("MissionsFindByModule No Mission Record")
	} else {
		for _, v := range missions {
			tempJson := &MissionBriefJson{}
			tempJson.mission2MissionBriefJson(&v)
			missionsBriefJson = append(missionsBriefJson, *tempJson)
		}
	}
	return
}

func MissionUpdate(missionJson *MissionJson) (recordMissionJson MissionJson, err error) {
	updateMission := new(Mission)
	updateMission.missionJson2Mission(missionJson)
	recordMission := new(Mission)
	recordMission.ID = updateMission.ID
	if database.DB.First(&recordMission, &recordMission).RecordNotFound() {
		err = errors.New("MissionUpdate No Mission Record")
	} else {
		database.DB.Model(&recordMission).Updates(updateMission)
		if num := len(missionJson.Participants); num != 0 {
			users := make([]User, num)
			for i, v := range missionJson.Participants {
				users[i].ID = v.ID
			}
			err = database.DB.Model(&recordMission).Association("Participants").Replace(users).Error
		}
		recordMissionJson.mission2MissionJSON(recordMission)
	}
	return
}

func MissionDelete(mission *Mission) (recordMissionJson MissionJson, err error) {
	recordMission := new(Mission)
	if database.DB.Find(&recordMission, &mission).RecordNotFound() {
		err = errors.New("MissionDelete No Mission Record")
	} else {
		recordMissionJson.mission2MissionJSON(recordMission)
		err = database.DB.Unscoped().Delete(&recordMission).Error
	}
	return
}
