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
	m := &Mission{}
	m.ID = gain.MissionID
	if err = m.First(); err == nil {
		if gain.OwnerID > 0 {
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
	gain.ID = 0
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
	if gain.ID > 0 {
		g := Gain{}
		g.ID = gain.ID
		if err = database.DB.First(&g).Error; err == nil {
			*gain = g
			gain.Owner.ID = gain.OwnerID
			err = gain.Owner.First()
		}
	} else {
		err = errors.New("ID必须为正数")
	}
	if err != nil {
		err = errors.New(titleGain + "First:\t" + err.Error())
	}
	return
}

//GainsFindByOwnerID 通过OwnerID来查找成果.
func GainsFindByOwnerID(id uint) (gains []Gain, err error) {
	owner := User{}
	owner.ID = id
	if err = owner.First(); err == nil {
		if err = database.DB.Model(&owner).Related(&gains, "OwnerID").Error; err == nil {
			for i := 0; i < len(gains); i++ {
				gains[i].Owner = owner
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
	mission := &Mission{}
	mission.ID = id
	if err = mission.First(); err == nil {
		if err = database.DB.Model(&mission).Related(&gains, "MissionID").Error; err == nil {
			for i := 0; i < len(gains); i++ {
				if err = gains[i].First(); err != nil {
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

func (gain *Gain) Updates() (err error) {
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
		err = errors.New(titleGain + "DeleteSoft:\t" + err.Error())
	}
	return
}
