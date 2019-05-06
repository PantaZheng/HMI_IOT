package service

import "github.com/pantazheng/bci/models"

func GainCreate(gain *models.GainJson)(gainJson models.GainJson,err error){
	return models.GainCreate(gain)
}

func GainFindByID(id uint)(gainJson models.GainJson,err error){
	gain:=new(models.Gain)
	gain.ID=id
	return models.GainFind(gain)
}

//owner单一确定
func GainsFindByOwnerID(id uint)(gainsJson []*models.GainJson,err error){
	owner:=new(models.User)
	owner.ID=id
	return models.GainsFindByOwner(owner)
}

//mission单一确定
func GainsFindByMissionID(id uint)(gainsJson []*models.GainJson,err error){
	mission:=new(models.Mission)
	mission.ID=id
	return models.GainsFindByMission(mission)
}

func GainUpdate(gainJson *models.GainJson)(recordGainJson models.GainJson,err error){
	return models.GainUpdate(gainJson)
}

func GainDeleteByID(id uint)(recordGainJson models.GainJson,err error) {
	gain:=new(models.GainJson)
	gain.ID=id
	return models.GainDelete(gain)
}