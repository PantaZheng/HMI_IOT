package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
)

const titleMission = "service.mission."

type MissionCore struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Content   string `json:"content"`
	Target    string `json:"target"`
	File      string `json:"file"`
	State     string `json:"state"`
	//foreign
	OwnerID   uint   `json:"ownerID"`
	OwnerName string `json:"ownerName"`
}

type MissionJSON struct {
	MissionCore
	//const inherit foreign
	ModuleID    uint   `json:"moduleID"`
	ModuleName  string `json:"moduleName"`
	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `json:"leaderName"`
	ProjectName string `json:"projectName"`
}

func (missionJSON *MissionJSON) mission2MissionJSON(mission models.Mission) {
	missionJSON.ID = mission.ID
	missionJSON.CreatedAt = mission.CreatedAt.Format("2006-01-02")
	missionJSON.UpdatedAt = mission.UpdatedAt.Format("2006-01-02")
	missionJSON.Name = mission.Name
	missionJSON.StartTime = mission.StartTime
	missionJSON.EndTime = mission.EndTime
	missionJSON.Content = mission.Content
	missionJSON.Target = mission.Target
	missionJSON.File = mission.File
	missionJSON.LeaderID = mission.LeaderID
	missionJSON.OwnerID = mission.OwnerID
	owner := UserJSON{ID: missionJSON.OwnerID}
	_ = owner.First()
	missionJSON.OwnerName = owner.Name

	missionJSON.ModuleID = mission.ModuleID
	module := ModuleJSON{}
	_ = module.First()
	missionJSON.LeaderName = module.LeaderName
	missionJSON.ModuleName = module.Name
	missionJSON.ProjectName = module.ProjectName

	return
}

func (missionJSON *MissionJSON) missionJSON2Mission() (mission models.Mission) {
	mission.ID = missionJSON.ID
	mission.Name = missionJSON.Name
	mission.StartTime = missionJSON.StartTime
	mission.EndTime = missionJSON.EndTime
	mission.Content = missionJSON.Content
	mission.Target = missionJSON.Target
	mission.File = missionJSON.File
	mission.ModuleID = missionJSON.ModuleID

	mission.LeaderID = missionJSON.LeaderID
	mission.OwnerID = missionJSON.OwnerID
	mission.ModuleID = missionJSON.ModuleID

	return
}

func (missionJSON *MissionJSON) Create() (err error) {
	m := missionJSON.missionJSON2Mission()
	if err = m.Insert(); err == nil {
		mission.miss
	}
	if err != nil {
		err = errors.New(titleMission + "Insert:\t" + err.Error())
	}
	return
}

func (missionJSON *MissionJSON) First() (err error) {
	m := missionJSON.missionJSON2Mission()
	if err = m.First(); err == nil {
		*missionJSON = mission2MissionJSON(&m)
	} else {
		err = errors.New(titleMission + "First:\t" + err.Error())
	}
	return
}

func MissionsFindAll() (missionsJSON []MissionJSON, err error) {
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
	m := missionJSON.missionJSON2Mission()
	if err = m.Update(); err == nil {
		*missionJSON = mission2MissionJSON(&m)
	} else {
		err = errors.New(titleMission + "Update:\t" + err.Error())
	}
	return
}

func (missionJSON *MissionJSON) Delete() (err error) {
	m := missionJSON.missionJSON2Mission()
	if err = m.Delete(); err == nil {
		*missionJSON = mission2MissionJSON(&m)
	} else {
		err = errors.New(titleMission + "Update:\t" + err.Error())
	}
	return
}

func MissionDeleteByID(id uint) (missionJSON MissionJSON, err error) {
	missionJSON = MissionJSON{ID: id}
	err = missionJSON.Delete()
	return
}
