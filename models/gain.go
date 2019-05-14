package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"time"
)

const titleGain = "models.gain."

type Gain struct {
	gorm.Model
	Name      string
	Type      string
	File      string
	UpTime    string
	Remark    string
	OwnerID   uint
	Owner     User
	MissionID uint
	Mission   Mission
}

//Create Create()
func (gain *Gain) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 0:00
	*/
	gain.UpTime = time.Now().Format("2006-01-02")
	if err = database.DB.Create(&gain).Error; err != nil {
		err = errors.New(titleGain + "Create:\t" + err.Error())
	}
	return
}

//First 根据id查找Gain.
func (gain *Gain) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 0:57
	*/
	g := &Gain{}
	g.ID = gain.ID
	if err = database.DB.First(&g).Error; err != nil {
		err = errors.New(titleGain + "First:\t" + err.Error())
	} else {
		*gain = *g
	}
	return
}

//GainsFindByID 通过OwnerID来查找成果.
func GainsFindByID(id uint) (gains []*Gain, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 0:29
	*/
	owner := &User{}
	owner.ID = id
	if err = owner.First(); err == nil {
		if err = database.DB.Model(&owner).Related(&gains, "OwnerID").Error; err == nil {
			if len(gains) == 0 {
				err = errors.New("record not found")
			}
		}
	}
	if err != nil {
		err = errors.New(titleGain + "FindByOwnerID:\t" + err.Error())
	}
	return
}

//FindByMissionID 通过OwnerID来查找任务下的成果.
func GainsFindByMID(id uint) (gains []*Gain, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 12:15
	*/
	mission := &Mission{}
	mission.ID = id
	if err = mission.First(); err == nil {
		if err = database.DB.Model(&mission).Related(&gains, "MissionID").Error; err == nil {
			if len(gains) == 0 {
				err = errors.New("record not found")
			}
		}
	}
	if err != nil {
		err = errors.New(titleGain + "GainsFindByMID:\t" + err.Error())
	}
	return
}

func (gain *Gain) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 1:09
	*/
	g := &Gain{}
	g.ID = gain.ID
	gain.UpTime = time.Now().Format("2006-01-02")
	if err = database.DB.Model(&g).Updates(&gain).Error; err != nil {
		err = errors.New(titleGain + "Updates:\t" + err.Error())
	}
	return
}

func (gain *Gain) Delete() (err error) {
	g := &Gain{}
	g.ID = gain.ID
	if err = database.DB.Delete(&g).Error; err != nil {
		err = errors.New(titleGain + "Delete:\t" + err.Error())
	} else {
		*gain = *g
	}
	return
}
