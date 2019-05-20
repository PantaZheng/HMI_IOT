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
	CreateTime   string
	StartTime    string
	EndTime      string
	Content      string
	Target       string
	File         string
	Tag          bool
	Participants []User `gorm:"many2many:user_missions"`
	ModuleID     uint
	Module       Module
}

func (mission *Mission) checkForeignKey() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 18:02
	*/
	m := &Module{}
	m.ID = mission.ModuleID
	err = m.First()
	return
}

//Create 创建Mission, 不添加成员
func (mission *Mission) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 3:48
	*/
	mission.ID = 0
	if err = mission.checkForeignKey(); err == nil {
		mission.CreateTime = time.Now().Format("2006-01-02")
		participants := mission.Participants
		mission.Participants = make([]User, 0)
		if err = database.DB.Create(&mission).Error; err == nil {
			if participants != nil {
				err = database.DB.Model(&mission).Association("Participants").Append(participants).Error
			}
			if err == nil {
				err = mission.First()
			}
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
			err = database.DB.Model(&mission).Association("Participants").Find(&mission.Participants).Error
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
func MissionsFindByCreatorID(id uint) (missions []Mission, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 16:06
	*/
	creator := &User{}
	creator.ID = id
	if err = creator.First(); err == nil {
		if err = database.DB.Model(&creator).Related(&missions, "CreatorID").Error; err == nil {
			for i := 0; i < len(missions); i++ {
				if err = missions[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleMission + "MissionsFindByCreatorID:\t" + err.Error())
	}
	return
}

func MissionsFindByParticipantID(id uint) (missions []Mission, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 16:06
	*/
	participant := &User{}
	participant.ID = id
	if err = participant.First(); err == nil {
		if err = database.DB.Model(&participant).Related(&missions, "PMissions").Error; err == nil {
			for i := 0; i < len(missions); i++ {
				if err = missions[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleMission + "MissionsFindByParticipantID:\t" + err.Error())
	}
	return
}

func MissionsFindByModuleID(id uint) (missions []Mission, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 12:15
	*/
	module := &Module{}
	module.ID = id
	if err = module.First(); err == nil {
		if err = database.DB.Model(&module).Related(&missions, "ModuleID").Error; err == nil {
			for i := 0; i < len(missions); i++ {
				if err = missions[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleGain + "GainsFindByMissionID:\t" + err.Error())
	}
	return
}

func (mission *Mission) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 23:18
	*/
	if err = mission.checkForeignKey(); err == nil {
		m := &Mission{}
		m.ID = mission.ID
		participants := mission.Participants
		mission.Participants = nil
		if err = database.DB.Model(&m).Updates(&mission).Error; err == nil {
			if participants != nil {
				err = database.DB.Model(&m).Association("Participants").Replace(participants).Error
			}
			if err == nil {
				err = mission.First()
			}
		}
	}
	if err != nil {
		err = errors.New(titleMission + "Updates\t" + err.Error())
	}
	return
}

func (mission *Mission) Delete() (err error) {
	if err = mission.First(); err == nil {
		m := Mission{}
		m.ID = mission.ID
		if err = database.DB.Model(&m).Association("Participants").Clear().Error; err == nil {
			err = database.DB.Delete(&m).Error
		}
	}
	if err != nil {
		err = errors.New(titleMission + "Delete\t" + err.Error())
	}
	return
}
