package service

import (
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/models"
)

func MissionCreate(mission *models.MissionJson)(missionBriefJson models.MissionBriefJson,err error){
	 return models.MissionCreate(mission)
}

func MissionFindByID(id uint)(missionJson models.MissionJson,err error){
	mission:=new(models.Mission)
	mission.ID=id
	return models.MissionFindOne(mission)
}

func MissionFindByName(name string)(missionJson models.MissionJson,err error){
	return models.MissionFindOne(&models.Mission{Name: name})
}

func MissionUpdate(missionJson *models.MissionJson)(missionBriefJson models.MissionBriefJson,err error){
	return models.MissionUpdate(missionJson)
}

func MissionDeleteByID(id uint)(missionBriefJson models.MissionBriefJson,err error){
	return models.MissionDelete(&models.Mission{ID:id})
}

func MissionDeleteByName(name string)(missionBriefJson models.MissionBriefJson,err error){
	return models.MissionDelete(&models.Mission{Name:name})
}