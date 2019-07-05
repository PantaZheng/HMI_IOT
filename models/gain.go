package models

import (
	"github.com/pantazheng/bci/database"
	"log"
	"strconv"
	"time"
)

type GainCore struct {
	ID        uint   `gorm:"primary_key",json:"id"`
	Name      string `json:"name"`
	State     uint   `json:"state"`
	OwnerName string `gorm:"-",json:"ownerName"`
}

type Gain struct {
	GainCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `gorm:"-",json:"createTime"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-",json:"updateTime"`
	DeletedAt  *time.Time `sql:"index",json:"-"`
	File       string     `json:"file"`
	Remark     string     `json:"remark"`

	MissionID   uint   `json:"missionID"`
	MissionName string `gorm:"-",json:"missionName"`
	OwnerID     uint   `json:"ownerID"`
	ModuleID    uint   `json:"moduleID"`
	ModuleName  string `gorm:"-",json:"moduleName"`
	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `gorm:"-",json:"leaderName"`
	ProjectID   uint   `json:"projectID"`
	ProjectName string `gorm:"-",json:"projectName"`
	ManagerID   uint   `json:"managerID"`
	ManagerName string `gorm:"-",json:"managerName"`
}

func gainTestData() {
	log.Println("gainTestData")
	l := 32
	gains := make([]Gain, l)
	for i := 0; i < l; i++ {
		gains[i].Name = "gain" + strconv.Itoa(i)
	}

	for _, v := range gains {
		if err := v.Insert(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
}

//Insert 必须包含MissionID
func (gain *Gain) Insert() (err error) {
	gain.State = 1
	mission := Mission{}
	mission.ID = gain.MissionID
	if err = mission.First(); err != nil {
		return
	}
	gain.OwnerID = mission.OwnerID
	gain.ModuleID = mission.ModuleID
	gain.LeaderID = mission.LeaderID
	gain.ProjectID = mission.ProjectID
	gain.ManagerID = mission.ManagerID
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
	mission := Mission{}
	mission.ID = gain.MissionID
	if err = mission.First(); err != nil {
		return
	}
	gain.CreateTime = gain.CreatedAt.Format("2006-01-02")
	gain.UpdateTime = gain.UpdatedAt.Format("2006-01-02")
	gain.MissionName = mission.Name
	gain.OwnerName = mission.OwnerName
	gain.ModuleName = mission.ModuleName
	gain.LeaderName = mission.LeaderName
	gain.ProjectName = mission.ProjectName
	gain.ManagerName = mission.ManagerName
	return
}

//FindGains
func (gain *Gain) Find(field string, id uint) (gains []Gain, err error) {
	if field == "all" {
		err = database.DB.Model(Gain{}).Find(&gains).Error
	} else {
		err = database.DB.Where(field+"_id = ?", id).Find(&gains).Error
	}
	return
}

func (gain *Gain) FindBrief(field string, id uint) (gainsCore []GainCore, err error) {
	if gains, e := gain.Find(field, id); e != nil {
		err = e
		return
	} else {
		l := len(gains)
		gainsCore = make([]GainCore, l)
		for i := 0; i < l; i++ {
			gainsCore[i] = gains[i].GainCore
			owner := UserCore{ID: gains[i].OwnerID}
			if err = owner.First(); err != nil {
				break
			}
			gainsCore[i].OwnerName = owner.Name
		}
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
