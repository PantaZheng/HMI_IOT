package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"log"
	"strconv"
)

const titleMission = "service.mission."

type MissionCore struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	OwnerName string `json:"ownerName"`
	State     uint   `json:"state"`
}

type MissionJSON struct {
	MissionCore
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Content   string `json:"content"`
	Target    string `json:"target"`
	File      string `json:"file"`

	//foreign
	OwnerID uint `json:"ownerID"`
	//const inherit foreign
	ModuleID    uint   `json:"moduleID"`
	ModuleName  string `json:"moduleName"`
	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `json:"leaderName"`
	ProjectName string `json:"projectName"`
}

func missionTestData() {
	log.Println("missionTestData")
	l := 32
	missions := make([]MissionJSON, l)
	for i := 0; i < l; i++ {
		missions[i].Name = "mission" + strconv.Itoa(i)
		missions[i].OwnerID = uint(i/2 + 1)
		missions[i].ModuleID = uint(i/2 + 1)
	}
	for _, v := range missions {
		if err := v.Insert(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
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
	missionJSON.State = mission.State
	missionJSON.LeaderID = mission.LeaderID
	missionJSON.OwnerID = mission.OwnerID
	owner := UserJSON{ID: missionJSON.OwnerID}
	_ = owner.First()
	missionJSON.OwnerName = owner.Name

	missionJSON.ModuleID = mission.ModuleID
	module := ModuleJSON{}
	module.ID = mission.ModuleID
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
	mission.State = missionJSON.State

	mission.LeaderID = missionJSON.LeaderID
	mission.OwnerID = missionJSON.OwnerID
	mission.ModuleID = missionJSON.ModuleID
	return
}

//Insert
func (missionJSON *MissionJSON) Insert() (err error) {
	m := missionJSON.missionJSON2Mission()
	if err = m.Insert(); err == nil {
		missionJSON.mission2MissionJSON(m)
	} else {
		err = errors.New(titleMission + "Insert:\t" + err.Error())
	}
	return
}

//First
func (missionJSON *MissionJSON) First() (err error) {
	m := missionJSON.missionJSON2Mission()
	if err = m.First(); err == nil {
		missionJSON.mission2MissionJSON(m)
	} else {
		err = errors.New(titleMission + "First:\t" + err.Error())
	}
	return
}

//Find
func (missionJSON *MissionJSON) Find(field string) (missionsJSON []MissionCore, err error) {
	m := missionJSON.missionJSON2Mission()
	if missions, err := m.Find(field); err != nil {
		err = errors.New(titleMission + "Find:\t" + err.Error())
	} else {
		missionsJSON = make([]MissionCore, len(missions))
		for i, v := range missions {
			missionsJSON[i].ID = v.ID
			missionsJSON[i].Name = v.Name
			owner := UserJSON{ID: v.OwnerID}
			_ = owner.First()
			missionsJSON[i].OwnerName = owner.Name
			missionsJSON[i].State = v.State
		}
	}
	return
}

func (missionJSON *MissionJSON) Updates() (err error) {
	if missionJSON.ID == 0 {
		err = errors.New(titleMission + "Updates:\t id 不可缺")
		return
	}
	m := missionJSON.missionJSON2Mission()
	if err = m.Updates(); err == nil {
		missionJSON.mission2MissionJSON(m)
	} else {
		err = errors.New(titleMission + "Updates:\t" + err.Error())
	}
	return
}

func (missionJSON *MissionJSON) Delete() (err error) {
	m := models.Mission{}
	m.ID = missionJSON.ID
	if err = m.Delete(); err == nil {
		missionJSON.mission2MissionJSON(m)
	} else {
		err = errors.New(titleMission + "Delete:\t" + err.Error())
	}
	return
}
