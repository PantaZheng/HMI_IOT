package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"log"
	"strconv"
)

const titleGain = "service.gain."

type GainCore struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	OwnerName string `json:"ownerName"`
	State     uint   `json:"state"`
}

type GainJSON struct {
	GainCore

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Type      string `json:"type"`
	File      string `json:"file"`
	Remark    string `json:"remark"`

	//const inherit foreign
	MissionID   uint   `json:"missionID"`
	MissionName string `json:"missionName"`
	OwnerID     uint   `json:"ownerID"`

	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `json:"leaderName"`
	ModuleName  string `json:"moduleName"`
	ProjectName string `json:"projectName"`
}

func gainTestData() {
	log.Println("gainTestData")
	l := 64
	gains := make([]GainJSON, l)

	for i := 0; i < l; i++ {
		gains[i].Name = "gain" + strconv.Itoa(i)
		gains[i].OwnerID = uint(i/4 + 1)
		gains[i].MissionID = uint(i/2 + 1)
	}

	for _, v := range gains {
		if err := v.Insert(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}

	for i := 0; i < l-1; i++ {
		gains[i].ID = uint(i + 1)
		gains[i].State = 2
		if err := gains[i].Updates(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(gains[i])
		}
	}

	mission := MissionJSON{}
	if missionsJSON, err := mission.Find("all"); err != nil {
		log.Println(err)
	} else {
		log.Println(missionsJSON)
	}
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
	mission.ID = gainJSON.MissionID
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
	if gains, err := g.Find(field); err != nil {
		err = errors.New(titleGain + "Find:\t" + err.Error())
	} else {
		gainsJSON = make([]GainCore, len(gains))
		for i, v := range gains {
			gainsJSON[i].ID = v.ID
			gainsJSON[i].Name = v.Name
			owner := UserJSON{ID: v.OwnerID}
			_ = owner.First()
			gainsJSON[i].OwnerName = owner.Name
			gainsJSON[i].State = v.State
		}
	}
	return
}

//Updates 必须包含ID
func (gainJSON *GainJSON) Updates() (err error) {
	if gainJSON.ID == 0 {
		err = errors.New(titleGain + "Updates:\t id 不可缺")
		return
	}
	checkTag := false
	if gainJSON.State > 0 {
		checkTag = true
	}
	g := gainJSON.gainJSON2Gain()
	if err = g.Updates(); err == nil {
		gainJSON.gain2GainJSON(g)
	} else {
		err = errors.New(titleGain + "Updates:\t" + err.Error())
		return
	}
	if !checkTag {
		return
	}
	if gainsJSON, err := gainJSON.Find("mission_id"); err != nil {
		err = errors.New(titleGain + "Updates:\t" + err.Error())
	} else {
		count := len(gainsJSON) - 1
		tag := 0
		for i, v := range gainsJSON {
			if v.State != 2 {
				tag = i
				break
			}
		}
		if tag == count {
			mission := MissionJSON{}
			mission.ID = gainJSON.MissionID
			mission.State = 2
			err = mission.Update()
		}
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
