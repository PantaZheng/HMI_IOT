package models

import (
	"github.com/pantazheng/bci/database"
	"time"
)

type MissionCore struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	State     uint   `json:"state"`
	OwnerName string `gorm:"-" json:"ownerName"`
}

type Mission struct {
	MissionCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `json:"createTime" gorm:"-"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-" json:"updateTime"`
	DeletedAt  *time.Time `sql:"index" json:"-"`
	StartTime  string     `json:"startTime"`
	EndTime    string     `json:"endTime"`
	Content    string     `json:"content"`
	Target     string     `json:"target"`

	OwnerID     uint   `json:"ownerID"`
	ModuleID    uint   `json:"moduleID"`
	ModuleName  string `gorm:"-" json:"moduleName"`
	LeaderID    uint   `json:"leaderID"`
	LeaderName  string `gorm:"-" json:"leaderName"`
	ProjectID   uint   `json:"projectID"`
	ProjectName string `gorm:"-" json:"projectName"`
	ManagerID   uint   `json:"managerID"`
	ManagerName string `gorm:"-" json:"managerName"`
}

//func missionTestData() {
//	log.Println("missionTestData")
//	l := 16
//	missions := make([]Mission, l)
//	for i := 0; i < l; i++ {
//		missions[i].Name = "mission" + strconv.Itoa(i)
//		missions[i].OwnerID = uint(i/2 + 1)
//		missions[i].ModuleID = uint(i/2 + 1)
//	}
//	for _, v := range missions {
//		if err := v.Insert(); err != nil {
//			log.Println(err.Error())
//		} else {
//			log.Println(v)
//		}
//	}
//}

//Insert 创建Mission
func (mission *Mission) Insert() (err error) {
	module := Module{}
	module.ID = mission.ModuleID
	if err = module.First(); err != nil {
		return
	}
	mission.LeaderID = module.LeaderID
	mission.ProjectID = module.ProjectID
	mission.ManagerID = module.ManagerID

	project := Project{}
	project.ID = module.ProjectID
	if err = project.First(); err != nil {
		return
	}
	mission.State = project.State

	if err = database.DB.Create(&mission).Error; err != nil {
		return
	}
	err = mission.First()
	return
}

//First 根据id查找Mission.
func (mission *Mission) First() (err error) {
	if err = database.DB.Where("id = ? ", mission.ID).First(&mission).Error; err != nil {
		return
	}
	module := Module{}
	module.ID = mission.ModuleID
	if err = module.First(); err != nil {
		return
	}
	owner := UserCore{ID: mission.OwnerID}
	if err = owner.First(); err != nil {
		return
	}
	mission.CreateTime = mission.CreatedAt.Format("2006-01-02")
	mission.UpdateTime = mission.UpdatedAt.Format("2006-01-02")
	mission.OwnerName = owner.Name
	mission.ModuleName = module.Name
	mission.LeaderName = module.LeaderName
	mission.ProjectName = module.ProjectName
	mission.ManagerName = module.ManagerName
	return
}

//FindMissions
func (mission *Mission) Find(field string, id uint) (missions []Mission, err error) {
	if field == "all" {
		err = database.DB.Model(Mission{}).Find(&missions).Error
	} else {
		err = database.DB.Where(field+"_id=?", id).Find(&missions).Error
	}
	return
}

func (mission *Mission) FindBrief(field string, id uint) (missionsCore []MissionCore, err error) {
	if missions, e := mission.Find(field, id); e != nil {
		err = e
		return
	} else {
		l := len(missions)
		missionsCore = make([]MissionCore, l)
		for i, v := range missions {
			missionsCore[i] = v.MissionCore
			owner := UserCore{ID: v.OwnerID}
			if err = owner.First(); err != nil {
				break
			}
			missionsCore[i].OwnerName = owner.Name
		}
	}
	return
}

//Updates ID必须，Uptime自动更新
func (mission *Mission) Updates() (err error) {
	if err = database.DB.Model(Mission{}).Where("id=?", mission.ID).Updates(&mission).Error; err != nil {
		return
	}
	err = mission.First()
	return
}

func (mission *Mission) Delete() (err error) {
	if err = mission.First(); err != nil {
		return
	}
	//硬删除
	if err = database.DB.Model(Mission{}).Where("id=?", mission.ID).Delete(&mission).Error; err != nil {
		return
	}
	gain := Gain{}
	_, err = gain.DeleteByField("mission", mission.ID)
	return
}

func (mission *Mission) DeleteByField(field string, id uint) (missions []Mission, err error) {
	err = database.DB.Model(Mission{}).Where(field+"_id=?", id).Delete(&missions).Error
	return
}
