package service

import (
	"github.com/pantazheng/bci/models"
	"log"
)

type MissionJson struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	CreatorID    uint
	Creator      UserJSON   `json:"creator"`
	CreateTime   string     `json:"createTime"`
	StartTime    string     `json:"startTime"`
	EndTime      string     `json:"endTime"`
	Content      string     `json:"content"`
	File         string     `json:"file"`
	Tag          bool       `json:"tag"` //tag由module负责人决定
	Gains        []GainJSON `json:"gains"`
	Participants []UserJSON `json:"participants"`
	ModuleID     uint       `json:"module"`
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

func mission2MissionJSON(mission *models.Mission) (missionJSON MissionJson) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 4:09
	*/
	missionJSON.ID = mission.ID
	missionJSON.Name = mission.Name
	creator := &UserJSON{ID: missionJSON.CreatorID}
	if err := creator.First(); err != nil {
		log.Println(err.Error())
	}
	missionJSON.Creator = userJSON2UserBriefJSON(creator)
	missionJSON.CreateTime = mission.CreateTime
	missionJSON.StartTime = mission.StartTime
	missionJSON.EndTime = mission.EndTime
	missionJSON.Content = mission.Content
	missionJSON.File = mission.File
	missionJSON.Tag = mission.Tag
	missionJSON.ModuleID = mission.ModuleID
	//user创建接口
	participants := make([]*User, len(mission.Participants))
	database.DB.Model(&mission).Related(&participants, "Participants")
	tempUser := &UserBriefJSON{}
	for _, v := range participants {
		tempUser.User2UserBriefJSON(v)
		missionJson.Participants = append(missionJson.Participants, *tempUser)
	}
	missionJson.Gains, _ = GainsFindByMission(mission)
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

}

func (missionBriefJson *MissionBriefJson) mission2MissionBriefJson(mission *Mission) {
	missionBriefJson.ID = mission.ID
	missionBriefJson.Name = mission.Name
	missionBriefJson.CreateTime = mission.CreateTime
	missionBriefJson.Content = mission.Content
	missionBriefJson.Tag = mission.Tag
	missionBriefJson.ModuleID = mission.ModuleID
}
func MissionCreate(mission *models.MissionJson) (missionJson models.MissionJson, err error) {
	return models.MissionCreate(mission)
}

//func MissionFindByID(id uint) (missionJson models.MissionJson, err error) {
//	mission := new(models.Mission)
//	mission.ID = id
//	return models.MissionFind(mission)
//}
//
//func MissionFindByName(name string) (missionJson models.MissionJson, err error) {
//	mission := new(models.Mission)
//	mission.Name = name
//	return models.MissionFind(mission)
//}
//
//func MissionsFindByModuleID(id uint) (missionsBriefJson []models.MissionBriefJson, err error) {
//	module := new(models.Module)
//	module.ID = id
//	return models.MissionsFindByModule(module)
//}
//
//func MissionUpdate(missionJson *models.MissionJson) (recordMissionJson models.MissionJson, err error) {
//	return models.MissionUpdate(missionJson)
//}
//
//func MissionDeleteByID(id uint) (recordMissionJson models.MissionJson, err error) {
//	mission := new(models.Mission)
//	mission.ID = id
//	return models.MissionDelete(mission)
//}
//
//func MissionDeleteByName(name string) (recordMissionJson models.MissionJson, err error) {
//	mission := new(models.Mission)
//	mission.Name = name
//	return models.MissionDelete(mission)
//}
