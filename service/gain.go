package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
)

const titleGain = "service.gain."

type GainCore struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Name      string `json:"type"`
	Type      string `json:"type"`
	File      string `json:"file"`
	Remark    string `json:"remark"`
	State     uint   `json:"state"`
}

type GainJSON struct {
	GainCore
	//const inherit foreign
	MissionID   uint   `json:"missionID"`
	MissionName string `json:"missionName"`
	OwnerID     uint   `json:"ownerID"`
	OwnerName   string `json:"ownerName"`
	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `json:"leaderName"`
	ModuleName  string `json:"moduleName"`
	ProjectName string `json:"projectName"`
}

//gain2GainJSON
func (gainJSON *GainJSON) gain2GainJSON(gain models.Gain) {
	gainJSON.ID = gain.ID
	gainJSON.CreatedAt = gain.CreatedAt.Format("2006-01-02")
	gainJSON.UpdatedAt = gain.UpdatedAt.Format("2006-01-02")
	gainJSON.Name = gain.Name
	gainJSON.Type = gain.Type
	gainJSON.File = gain.File
	gainJSON.Remark = gain.Remark
	gainJSON.State = gain.State
	gainJSON.LeaderID = gain.LeaderID
	gainJSON.OwnerID = gain.OwnerID

	gainJSON.MissionID = gain.MissionID
	mission := MissionJSON{}
	_ = mission.First()
	gainJSON.MissionName = mission.Name
	gainJSON.OwnerName = mission.OwnerName
	gainJSON.LeaderName = mission.LeaderName
	gainJSON.ModuleName = mission.ModuleName
	gainJSON.ProjectName = mission.ProjectName

	return
}

//gainJSON2Gain
func (gainJSON *GainJSON) gainJSON2Gain() (gain models.Gain) {
	gain.ID = gainJSON.ID
	gain.Name = gainJSON.Name
	gain.Type = gainJSON.Type
	gain.File = gainJSON.File
	gain.Remark = gainJSON.Remark
	gain.State = gainJSON.State

	gain.LeaderID = gainJSON.LeaderID
	gain.OwnerID = gainJSON.OwnerID
	gain.MissionID = gainJSON.MissionID
	return
}

//Insert
func (gainJSON *GainJSON) Insert() (err error) {
	g := gainJSON.gainJSON2Gain()
	if err = g.Insert(); err == nil {
		gainJSON.gain2GainJSON(g)
	} else {
		err = errors.New(titleGain + "Insert:\t" + err.Error())
	}
	return
}

//First 单Gain查找的原子方法.
func (gainJSON *GainJSON) First() (err error) {
	g := gainJSON.gainJSON2Gain()
	if err = g.First(); err == nil {
		gainJSON.gain2GainJSON(g)
	} else {
		err = errors.New(titleGain + "First:\t" + err.Error())
	}
	return
}

//Find
func (gainJSON *GainJSON) Find(field string) (gainsJSON []GainCore, err error) {
	g := gainJSON.gainJSON2Gain()
	if gains, err1 := g.Find(field); err1 == nil {
		gainsJSON = make([]GainCore, len(gains))
		gTemp := GainCore{}
		for i, v := range gains {
			gTemp = gainsJSON[i]
			gTemp.ID = v.ID
			gTemp.Name = v.Name
			gTemp.Type = v.Type
			gTemp.File = v.File
			gTemp.CreatedAt = v.CreatedAt.Format("2006-01-02")
			gTemp.UpdatedAt = v.UpdatedAt.Format("2006-01-02")
			gTemp.Remark = v.Remark
			gTemp.State = v.State
		}
	} else {
		err = errors.New(titleGain + "GainsFind:\t" + err1.Error())
	}
	return
}

//Update 必须包含ID
func (gainJSON *GainJSON) Update() (err error) {
	if gainJSON.ID == 0 {
		err = errors.New(titleGain + "Update:\t id 不可缺")
		return
	}
	g := gainJSON.gainJSON2Gain()
	if err = g.Update(); err == nil {
		gainJSON.gain2GainJSON(g)
	} else {
		err = errors.New(titleGain + "Update:\t" + err.Error())
	}
	return
}

//Delete 必须包含ID
func (gainJSON *GainJSON) Delete() (err error) {
	g := models.Gain{}
	g.ID = gainJSON.ID
	if err = g.Delete(); err == nil {
		gainJSON.gain2GainJSON(g)
	} else {
		err = errors.New(titleGain + "Delete:\t" + err.Error())
	}
	return
}
