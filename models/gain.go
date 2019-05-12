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

//FindByOwner 通过OwnerID来查找多个用户.
func FindByOwnerID(ownerID uint) (gains []*Gain, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 0:29
	*/
	owner := &User{}
	owner.ID = ownerID
	if err = owner.First(); err == nil {
		if err = database.DB.Model(&owner).Related(&gains, "OwnerID").Error; err == nil {
			if len(gains) == 0 {
				err = errors.New("record not found")
			}
		}
	}
	if err != nil {
		err = errors.New(titleGain + "FindByOwner:\t" + err.Error())
	}
	return
}

//func GainsFindByMission(mission *Mission) (gainsJson []GainJson, err error) {
//	gains := make([]Gain, 1)
//	if err = database.DB.Model(&mission).Related(&gains, "MissionID").Error; err != nil {
//		return
//	}
//	if len(gains) == 0 {
//		err = errors.New("GainsFindByMission No Gain Record")
//	} else {
//		for _, v := range gains {
//			tempJson := &GainJson{}
//			tempJson.gain2GainJson(&v)
//			gainsJson = append(gainsJson, *tempJson)
//		}
//	}
//	return
//}

func (gain *Gain) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 1:09
	*/
	g := new(Gain)
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
