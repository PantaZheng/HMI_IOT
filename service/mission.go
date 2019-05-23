package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"log"
)

const titleMission = "service.mission."

type MissionJSON struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 23:50
	*/
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	CreateTime   string     `json:"createTime"`
	StartTime    string     `json:"startTime"`
	EndTime      string     `json:"endTime"`
	Content      string     `json:"content"`
	Target       string     `json:"target"`
	File         string     `json:"file"`
	Tag          bool       `json:"tag"` //tag由module负责人决定
	Gains        []GainJSON `json:"gains"`
	Participants []UserJSON `json:"participants"`
	ModuleID     uint       `json:"moduleID"`
}

type MissionBriefJSON struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/24 0:32
	*/
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	StartTime    string     `json:"startTime"`
	EndTime      string     `json:"endTime"`
	Participants []UserJSON `json:"participants"`
}

func missionTestData() {
	log.Println("missionTestData")
	u1 := UserJSON{ID: 2}
	u2 := UserJSON{ID: 3}
	u3 := UserJSON{ID: 4}
	u4 := UserJSON{ID: 5}
	u5 := UserJSON{ID: 6}
	u6 := UserJSON{ID: 7}
	u7 := UserJSON{ID: 8}
	missions := make([]MissionJSON, 4)
	missions[0] = MissionJSON{Name: "钢铁侠1", StartTime: "2008-1-1", EndTime: "2017-1-2", Content: "你不是世界上唯一的超级英雄。", File: "朝花夕拾", Participants: []UserJSON{u1, u3, u4, u5}, ModuleID: 1}
	missions[1] = MissionJSON{Name: "无敌浩克", StartTime: "2008-3-1", EndTime: "2017-1-2", Content: "复仇者联盟", File: "", Participants: []UserJSON{u4, u5}, ModuleID: 1}
	missions[2] = MissionJSON{Name: "海王1", StartTime: "2008-3-1", EndTime: "2017-1-2", Content: "你永远与我同在也永远是我的兄弟", File: "", Participants: []UserJSON{u2, u6}, ModuleID: 2}
	missions[3] = MissionJSON{Name: "雷神1", StartTime: "2008-3-1", EndTime: "2017-1-2", Content: "那些你笑就跟着你笑的人，如果不是傻子，就是爱你的人。", File: "", Participants: []UserJSON{u5, u7}, ModuleID: 3}
	for _, v := range missions {
		if err := v.Create(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
}

func mission2MissionJSON(mission *models.Mission) (missionJSON MissionJSON) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/14 9:59
	*/
	missionJSON.ID = mission.ID
	missionJSON.Name = mission.Name
	missionJSON.CreateTime = mission.CreateTime
	missionJSON.StartTime = mission.StartTime
	missionJSON.EndTime = mission.EndTime
	missionJSON.Content = mission.Content
	missionJSON.Target = mission.Target
	missionJSON.File = mission.File
	missionJSON.Tag = mission.Tag
	missionJSON.ModuleID = mission.ModuleID
	missionJSON.Participants = users2BriefUsersJSON(mission.Participants)
	missionJSON.Gains, _ = GainsFindByMissionID(mission.ID)
	return
}

func missionJSON2MissionBriefJSON(missionJSON1 *MissionJSON) (missionJSON2 MissionJSON) {
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
	missionJSON2.ModuleID = missionJSON1.ModuleID
	return
}

func missions2MissionsBriefJSON(missions []models.Mission) (missionsJSON []MissionJSON) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/14 10:00
	*/
	missionsJSON = make([]MissionJSON, len(missions))
	for i, v := range missions {
		m := mission2MissionJSON(&v)
		missionsJSON[i] = missionJSON2MissionBriefJSON(&m)
	}
	return
}

func (missionJSON *MissionJSON) missionJSON2Mission() (mission models.Mission) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/14 10:17
	*/
	mission.ID = missionJSON.ID
	mission.Name = missionJSON.Name
	mission.CreateTime = missionJSON.CreateTime
	mission.StartTime = missionJSON.StartTime
	mission.EndTime = missionJSON.EndTime
	mission.Content = missionJSON.Content
	mission.Target = missionJSON.Target
	mission.File = missionJSON.File
	mission.Tag = missionJSON.Tag
	mission.Participants = usersJSON2Users(missionJSON.Participants)
	mission.ModuleID = missionJSON.ModuleID
	return
}

func (missionJSON *MissionJSON) Create() (err error) {
	m := missionJSON.missionJSON2Mission()
	if err = m.Create(); err == nil {
		*missionJSON = mission2MissionJSON(&m)
	}
	if err != nil {
		err = errors.New(titleMission + "Create:\t" + err.Error())
	}
	return
}

func (missionJSON *MissionJSON) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 18:42
	*/
	m := missionJSON.missionJSON2Mission()
	if err = m.First(); err == nil {
		*missionJSON = mission2MissionJSON(&m)
	} else {
		err = errors.New(titleMission + "First:\t" + err.Error())
	}
	return
}

func MissionsFindAll() (missionsJSON []MissionJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/21 12:32
	*/
	if missions, err1 := models.MissionsFindAll(); err1 == nil {
		missionsJSON = missions2MissionsBriefJSON(missions)
	} else {
		err = errors.New(titleMission + "MissionsFindAll:\t" + err1.Error())
	}
	return
}

//MissionFindByID 通过数据库ID查找单Mission.
func MissionFindByID(id uint) (missionJSON MissionJSON, err error) {
	missionJSON = MissionJSON{ID: id}
	err = missionJSON.First()
	return
}

func MissionsFindByParticipantID(id uint) (missionsJSON []MissionJSON, err error) {
	if missions, err1 := models.MissionsFindByParticipantID(id); err1 == nil {
		missionsJSON = missions2MissionsBriefJSON(missions)
	} else {
		err = errors.New(titleMission + "MissionsFindByParticipantID:\t" + err1.Error())
	}
	return
}

func MissionsFindByModuleID(id uint) (missionsJSON []MissionJSON, err error) {
	if missions, err1 := models.MissionsFindByModuleID(id); err1 == nil {
		missionsJSON = missions2MissionsBriefJSON(missions)
	} else {
		err = errors.New(titleMission + "MissionsFindByModuleID:\t" + err1.Error())
	}
	return
}

func (missionJSON *MissionJSON) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 18:53
	*/
	m := missionJSON.missionJSON2Mission()
	if err = m.Updates(); err == nil {
		*missionJSON = mission2MissionJSON(&m)
	} else {
		err = errors.New(titleMission + "Updates:\t" + err.Error())
	}
	return
}

func (missionJSON *MissionJSON) Delete() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 18:53
	*/
	m := missionJSON.missionJSON2Mission()
	if err = m.Delete(); err == nil {
		*missionJSON = mission2MissionJSON(&m)
	} else {
		err = errors.New(titleMission + "Updates:\t" + err.Error())
	}
	return
}

func MissionDeleteByID(id uint) (missionJSON MissionJSON, err error) {
	missionJSON = MissionJSON{ID: id}
	err = missionJSON.Delete()
	return
}
