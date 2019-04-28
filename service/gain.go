package service

import "github.com/pantazheng/bci/models"

func GainCreate(gain *models.GainJson)(gainJson models.GainJson,err error){
	return models.GainCreate(gain)
}

func GainFindByID(id uint)(gainJson models.GainJson,err error){
	gain:=new(models.Gain)
	gain.ID=id
	 return models.GainFindByID(gain)
}

//owner单一确定
func GainsFindByOwner(owner *models.User)(gainsJson []models.GainJson,err error){
	return models.GainsFindByOwner(owner)
}

//mission单一确定
func GainsFindByMission(mission *models.Mission)(gainsJson []models.GainJson,err error){
	return models.GainsFindByMission(mission)
}

