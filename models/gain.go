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

func (gain *Gain) checkForeignKey() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/14 23:56
	*/
	m := &Mission{}
	m.ID = gain.MissionID
	if err = m.First(); err == nil {
		if gain.OwnerID != 0 {
			err = errors.New("OwnerID not in mission's participants")
			for _, v := range m.Participants {
				if gain.OwnerID == v.ID {
					err = nil
				}
			}
		}
	}
	return
}

//Create Create()
func (gain *Gain) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 0:00
	*/
	if err = gain.checkForeignKey(); err == nil {
		gain.UpTime = time.Now().Format("2006-01-02")
		if err = database.DB.Create(&gain).Error; err == nil {
			err = gain.First()
		}
	}
	if err != nil {
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
	g := Gain{}
	g.ID = gain.ID
	if err = database.DB.First(&g).Error; err == nil {
		*gain = g
		gain.Owner.ID = gain.OwnerID
		err = gain.Owner.First()
	}
	if err != nil {
		err = errors.New(titleGain + "First:\t" + err.Error())
	}
	return
}

//GainsFindByOwnerID 通过OwnerID来查找成果.
func GainsFindByOwnerID(id uint) (gains []Gain, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 0:29
	*/
	owner := User{}
	owner.ID = id
	if err = owner.First(); err == nil {
		if err = database.DB.Model(&owner).Related(&gains, "OwnerID").Error; err == nil {
			owner := User{}
			owner.ID = id
			err = owner.First()
			for _, v := range gains {
				v.Owner = owner
			}
		}
	}
	if err != nil {
		err = errors.New(titleGain + "FindByOwnerID:\t" + err.Error())
	}
	return
}

//FindByMissionID 通过OwnerID来查找任务下的成果.
func GainsFindByMissionID(id uint) (gains []Gain, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 12:15
	*/
	mission := &Mission{}
	mission.ID = id
	if err = mission.First(); err == nil {
		if err = database.DB.Model(&mission).Related(&gains, "MissionID").Error; err == nil {
			for _, v := range gains {
				v.Owner.ID = v.OwnerID
				err = v.Owner.First()
			}
		}
	}
	if err != nil {
		err = errors.New(titleGain + "GainsFindByMissionID:\t" + err.Error())
	}
	return
}

func (gain *Gain) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 1:09
	*/
	if err = gain.checkForeignKey(); err == nil {
		g := Gain{}
		g.ID = gain.ID
		gain.UpTime = time.Now().Format("2006-01-02")
		if err = database.DB.Model(&g).Updates(&gain).Error; err == nil {
			err = gain.First()
		}
	}
	if err != nil {
		err = errors.New(titleGain + "Updates:\t" + err.Error())
	}
	return
}

func (gain *Gain) Delete() (err error) {
	if err = gain.First(); err == nil {
		g := Gain{}
		g.ID = gain.ID
		err = database.DB.Delete(&g).Error
	}
	if err != nil {
		err = errors.New(titleGain + "Delete:\t" + err.Error())
	}
	return
}
