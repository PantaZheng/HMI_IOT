package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"time"
)

const titleMission = "models.mission."

type Mission struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:43
	*/
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
	Participants []User `gorm:"many2many:user_missions"`
	ModuleID     uint
	//Module       Module
}

//Create 创建Mission, 不添加成员
func (mission *Mission) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 3:48
	*/
	if mission.ID != 0 {
		mission.ID = 0
	}
	mission.CreateTime = time.Now().Format("2006-01-02")
	participants := mission.Participants
	mission.Participants = nil
	if err = database.DB.Create(&mission).Error; err == nil {
		mission.Creator.ID = mission.CreatorID
		if participants != nil {
			if err = database.DB.Model(&mission).Association("Participants").Append(participants).Error; err == nil {
				err = mission.Creator.First()
			}
		} else {
			err = mission.Creator.First()
		}
	}
	if err != nil {
		err = errors.New(titleMission + "Create:\t" + err.Error())
	}
	return
}

//First 根据id查找Mission.
func (mission *Mission) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 15:49
	*/
	if mission.ID > 0 {
		m := &Mission{}
		m.ID = mission.ID
		if err = database.DB.First(&m).Error; err == nil {
			*mission = *m
			if err = database.DB.Model(&mission).Association("Participants").Find(&mission.Participants).Error; err == nil {
				mission.Creator.ID = mission.CreatorID
				err = mission.Creator.First()
			}
		}
	} else {
		err = errors.New("ID必须为正数")
	}
	if err != nil {
		err = errors.New(titleMission + "First:\t" + err.Error())
	}
	return
}

// MissionsFindByCID通过CreatorID查找Missions
func MissionsFindByCID(id uint) (missions []*Mission, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 16:06
	*/
	creator := &User{}
	creator.ID = id
	if err = creator.First(); err == nil {
		err = database.DB.Model(&creator).Related(&missions, "CreatorID").Error
	}
	if err != nil {
		err = errors.New(titleMission + "MissionsFindByCreatorID:\t" + err.Error())
	}
	return
}

func MissionsFindByPID(id uint) (missions []*Mission, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 16:06
	*/
	participant := &User{}
	participant.ID = id
	if err = participant.First(); err == nil {
		err = database.DB.Model(&participant).Related(&missions, "PMissions").Error
	}
	if err != nil {
		err = errors.New(titleMission + "MissionsFindByPID:\t" + err.Error())
	}
	return
}

func (mission *Mission) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 23:18
	*/
	m := &Mission{}
	m.ID = mission.ID
	participants := mission.Participants
	mission.Participants = nil
	if err = database.DB.Model(&m).Updates(&mission).Error; err == nil {
		*mission = *m
		mission.Creator.ID = mission.CreatorID
		if participants != nil {
			if err = database.DB.Model(&mission).Association("Participants").Replace(participants).Error; err == nil {
				err = mission.Creator.First()
			}
		} else {
			if err = database.DB.Model(&mission).Association("Participants").Find(&mission.Participants).Error; err == nil {
				err = mission.Creator.First()
			}
		}
	}
	if err != nil {
		err = errors.New(titleMission + "Updates\t" + err.Error())
	}
	return
}

func (mission *Mission) Delete() (err error) {
	m := &Mission{}
	m.ID = mission.ID
	participants := make([]User, 0)
	if err = database.DB.Model(&mission).Association("Participants").Find(&participants).Error; err == nil {
		m := &Mission{}
		m.ID = mission.ID
		if err = database.DB.Delete(&m).Error; err == nil {
			*mission = *m
			mission.Participants = participants
			mission.Creator.ID = mission.CreatorID
			err = mission.Creator.First()
		}
	}
	if err != nil {
		err = errors.New(titleMission + "Updates\t" + err.Error())
	}
	return
}
