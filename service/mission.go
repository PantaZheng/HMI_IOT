package service

import (
	"github.com/pantazheng/bci/models"
)

func MissionCreate(mission *models.MissionJson)(missionJson models.MissionJson,err error){
	 return models.MissionCreate(mission)
}

func MissionFindByID(id uint)(missionJson models.MissionJson,err error){
	mission:=new(models.Mission)
	mission.ID=id
	return models.MissionFind(mission)
}

func MissionFindByName(name string)(missionJson models.MissionJson,err error){
	mission:=new(models.Mission)
	mission.Name=name
	return models.MissionFind(mission)
}

func MissionsFindByModuleID(id uint)(missionsBriefJson []*models.MissionBriefJson,err error){
	module:=new(models.Module)
	module.ID=id
	return models.MissionsFindByModule(module)
}

func MissionUpdate(missionJson *models.MissionJson)(recordMissionJson models.MissionJson,err error){
	return models.MissionUpdate(missionJson)
}

func MissionDeleteByID(id uint)(recordMissionJson models.MissionJson,err error){
	mission:=new(models.Mission)
	mission.ID=id
	return models.MissionDelete(mission)
}

func MissionDeleteByName(name string)(recordMissionJson models.MissionJson,err error){
	mission:=new(models.Mission)
	mission.Name=name
	return models.MissionDelete(mission)
}