package service

import (
	"github.com/pantazheng/bci/models"
)

func ModuleCreate(module *models.ModuleJson)(moduleJson models.ModuleJson,err error){
	return models.ModuleCreate(module)
}

func ModuleFindByID(id uint) (recordModuleJson models.ModuleJson,err error){
	module:=new(models.Module)
	module.ID=id
	return models.ModuleFind(module)
}

func ModulesFindByLeaderID(id uint)(modulesBriefJson []models.ModuleBriefJson,err error){
	leader:=new(models.User)
	leader.ID=id
	return models.ModulesFindByLeader(leader)
}

func ModulesFindByProjectID(id uint)(modulesBriefJson []models.ModuleBriefJson,err error){
	project:=new(models.Project)
	project.ID=id
	return models.ModulesFindByProject(project)
}

func ModuleUpdate(moduleJson *models.ModuleJson)(recordModuleJson models.ModuleJson,err error){
	return models.ModuleUpdate(moduleJson)
}

func ModuleDeleteByID(id uint)(recordModuleJson models.ModuleJson, err error) {
	module:=new(models.Module)
	module.ID=id
	return models.ModuleDelete(module)
}