package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"log"
)

const titleGain = "service.gain."

type GainJSON struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 1:17
	*/
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	File      string   `json:"file"`
	UpTime    string   `json:"upTime"`
	Remark    string   `json:"remark"`
	OwnerID   uint     `json:"ownerID"`
	Owner     UserJSON `json:"owner"`
	MissionID uint     `json:"missionID"`
}

func gainTestData() {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 2:39
	*/
	log.Println("gainTestData")
	gains := make([]GainJSON, 6)
	gains[0] = GainJSON{Name: "gain1", OwnerID: 2, MissionID: 1}
	gains[1] = GainJSON{Name: "gain2", OwnerID: 4, MissionID: 1}
	gains[2] = GainJSON{Name: "gain3", OwnerID: 5, MissionID: 2}
	gains[3] = GainJSON{Name: "gain4", OwnerID: 6, MissionID: 2}
	gains[4] = GainJSON{Name: "gain5", OwnerID: 7, MissionID: 3}
	gains[5] = GainJSON{Name: "gain6", OwnerID: 3, MissionID: 3}
	for _, v := range gains {
		if err := v.Create(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
}
func gain2GainJSON(gain *models.Gain) (gainJSON GainJSON) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 2:05
	*/
	gainJSON.ID = gain.ID
	gainJSON.Name = gain.Name
	gainJSON.Type = gain.Type
	gainJSON.File = gain.File
	gainJSON.UpTime = gain.UpTime
	gainJSON.Remark = gain.Remark
	gainJSON.OwnerID = gain.OwnerID
	gainJSON.Owner = user2UserJSON(&gain.Owner)
	gainJSON.MissionID = gain.MissionID
	return
}

func gainJSON2GainBriefJSON(gainJSON1 *GainJSON) (gainJSON2 GainJSON) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 3:14
	*/
	gainJSON2.ID = gainJSON1.ID
	gainJSON2.Name = gainJSON1.Name
	gainJSON2.UpTime = gainJSON1.UpTime
	gainJSON2.OwnerID = gainJSON1.OwnerID
	gainJSON2.Owner = userJSON2UserBriefJSON(&gainJSON1.Owner)
	gainJSON2.MissionID = gainJSON1.MissionID
	return
}

func gains2BriefGainsJSON(gains []models.Gain) (gainsJSON []GainJSON) {
	gainsJSON = make([]GainJSON, len(gains))
	for i, v := range gains {
		g := gain2GainJSON(&v)
		gainsJSON[i] = gainJSON2GainBriefJSON(&g)
	}
	return
}

//gainJSON2Gain GainJSON转换到Gain.
func (gainJSON *GainJSON) gainJSON2Gain() (gain models.Gain) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 2:39
	*/
	gain.ID = gainJSON.ID
	gain.Name = gainJSON.Name
	gain.Type = gainJSON.Type
	gain.File = gainJSON.File
	gain.UpTime = gainJSON.UpTime
	gain.Remark = gainJSON.Remark
	gain.OwnerID = gainJSON.OwnerID
	gain.MissionID = gainJSON.MissionID
	return
}

func (gainJSON *GainJSON) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 2:39
	*/
	//TODO:检查成果归属者是否在Mission的参与者中,前端选择
	g := gainJSON.gainJSON2Gain()
	if err = g.Create(); err == nil {
		*gainJSON = gain2GainJSON(&g)
	}
	if err != nil {
		err = errors.New(titleGain + "Create:\t" + err.Error())
	}
	return
}

//First 单Gain查找的原子方法.
func (gainJSON *GainJSON) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 2:39
	*/
	g := gainJSON.gainJSON2Gain()
	if err = g.First(); err == nil {
		*gainJSON = gain2GainJSON(&g)
	} else {
		err = errors.New(titleGain + "First:\t" + err.Error())
	}
	return
}

//GainFindByID 通过数据库ID查找单Gain.
func GainFindByID(id uint) (gainJSON GainJSON, err error) {
	gainJSON = GainJSON{ID: id}
	err = gainJSON.First()
	return
}

//owner单一确定
func GainsFindByOID(id uint) (gainsJson []GainJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 2:44
	*/
	if gains, err := models.GainsFindByOwnerID(id); err == nil {
		gainsJson = gains2BriefGainsJSON(gains)
	} else {
		err = errors.New(titleGain + "GainsFindByOID:\t" + err.Error())
	}
	return
}

//mission单一确定
func GainsFindByMID(id uint) (gainsJson []GainJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 12:25
	*/
	if gains, err := models.GainsFindByMissionID(id); err == nil {
		gainsJson = gains2BriefGainsJSON(gains)
	} else {
		err = errors.New(titleGain + "GainsFindByMissionID:\t" + err.Error())
	}
	return
}

//Updates 更新成果数据，id定位成果记录.
func (gainJSON *GainJSON) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 3:25
	*/
	g := gainJSON.gainJSON2Gain()
	if err = g.Updates(); err == nil {
		*gainJSON = gain2GainJSON(&g)
	} else {
		err = errors.New(titleGain + "Updates:\t" + err.Error())
	}
	return
}

func (gainJSON *GainJSON) Delete() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/13 3:30
	*/
	g := gainJSON.gainJSON2Gain()
	if err = g.Delete(); err == nil {
		*gainJSON = gain2GainJSON(&g)
	} else {
		err = errors.New(titleGain + "Delete:\t" + err.Error())
	}
	return
}

//GainDeleteByID 通过数据库ID删除Gain.
func GainDeleteByID(id uint) (gainJSON GainJSON, err error) {
	gainJSON = GainJSON{ID: id}
	err = gainJSON.Delete()
	return
}
