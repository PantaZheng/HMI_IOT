package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"log"
)

const titleMission = "service.mission."

type MissionJson struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 23:50
	*/
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
	//ModuleID     uint       `json:"module"`
}

func missionTestData() {
	u1 := UserJSON{ID: 2}
	u2 := UserJSON{ID: 3}
	u3 := UserJSON{ID: 4}
	u4 := UserJSON{ID: 5}
	u5 := UserJSON{ID: 6}
	u6 := UserJSON{ID: 7}
	missions := make([]MissionJson, 3)
	missions[0] = MissionJson{Name: "钢铁侠", CreatorID: 5, StartTime: "2008-1-1", EndTime: "2017-1-2", Content: "你不是世界上唯一的超级英雄。", File: "朝花夕拾", Participants: []UserJSON{u1, u3, u4, u5}}
	missions[1] = MissionJson{Name: "无敌浩克", CreatorID: 4, StartTime: "2008-3-1", EndTime: "2017-1-2", Content: "浩克应该加入他（复仇者联盟）", File: "", Participants: []UserJSON{u4, u5}}
	missions[2] = MissionJson{Name: "海王", CreatorID: 7, StartTime: "2008-3-1", EndTime: "2017-1-2", Content: "浩克应该加入他（复仇者联盟）", File: "", Participants: []UserJSON{u2, u6}}
	for _, v := range missions {
		if err := v.Create(); err != nil {
			log.Println(err.Error())
		}
	}
}

func mission2MissionJSON(mission *models.Mission) (missionJSON MissionJson) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/14 9:59
	*/
	missionJSON.ID = mission.ID
	missionJSON.Name = mission.Name
	missionJSON.CreatorID = mission.CreatorID
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
	//missionJSON.ModuleID = mission.ModuleID
	//user创建接口
	if participants, err := mission.FindParticipants(); err != nil {
		log.Println(err.Error())
	} else {
		missionJSON.Participants = users2BriefUsersJSON(participants)
	}
	missionJSON.Gains, _ = GainsFindByMID(mission.ID)
	return
}

func missionJSON2MissionBriefJSON(missionJSON1 *MissionJson) (missionJSON2 MissionJson) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/14 9:59
	*/
	missionJSON2.ID = missionJSON1.ID
	missionJSON2.Name = missionJSON1.Name
	missionJSON2.CreateTime = missionJSON1.CreateTime
	missionJSON2.Content = missionJSON1.Content
	missionJSON2.Tag = missionJSON1.Tag
	//missionJSON2.ModuleID=missionJSON1.ModuleID
	return
}

func missions2MissionsBriefJSON(missions []*models.Mission) (missionsJSON []MissionJson) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/14 10:00
	*/
	missionsJSON = make([]MissionJson, len(missions))
	for i, v := range missions {
		m := mission2MissionJSON(v)
		missionsJSON[i] = missionJSON2MissionBriefJSON(&m)
	}
	return
}

//缺失participants
func (missionJSON *MissionJson) missionJSON2Mission() (mission models.Mission) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/14 10:17
	*/
	mission.ID = missionJSON.ID
	mission.Name = missionJSON.Name
	mission.CreatorID = missionJSON.CreatorID
	mission.CreateTime = missionJSON.CreateTime
	mission.StartTime = missionJSON.StartTime
	mission.EndTime = missionJSON.EndTime
	mission.Content = missionJSON.Content
	mission.File = missionJSON.File
	mission.Tag = missionJSON.Tag
	//mission.ModuleID = missionJSON.ModuleID
	return
}

func (missionJSON *MissionJson) Create() (err error) {
	//TODO:检查creator是否归属module
	creator := UserJSON{ID: missionJSON.CreatorID}
	if err = creator.First(); err == nil {
		m := missionJSON.missionJSON2Mission()
		if err = m.Create(); err == nil {
			users := usersJSON2Users(missionJSON.Participants)
			if err = m.AppendParticipants(users); err == nil {
				*missionJSON = mission2MissionJSON(&m)
			}
		}
	}
	if err != nil {
		err = errors.New(titleMission + "Create:\t" + err.Error())
	}
	return
}

func (missionJSON *MissionJson) IfParticipants(id uint) (err error) {
	err = errors.New("成员不在任务的参与者中")
	m := missionJSON.missionJSON2Mission()
	if participants, err1 := m.FindParticipants(); err1 == nil {
		for _, v := range participants {
			if v.ID == id {
				err = nil
			}
		}
	}
	if err != nil {
		err = errors.New(titleMission + "IfParticipants:\t" + err.Error())
	}
	return
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
