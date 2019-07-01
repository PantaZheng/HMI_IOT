package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
)

const titleModule = "service.module."

type ModuleJSON struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	CreatorID   uint          `json:"creatorID"`
	Creator     UserJSON      `json:"creator"`
	CreateTime  string        `json:"createTime"` //创建时间
	StartTime   string        `json:"startTime"`  //开始时间
	EndTime     string        `json:"endTime"`    //结束时间
	Content     string        `json:"content"`
	Target      string        `json:"target"`
	Tag         bool          `json:"tag"`
	LeaderID    uint          `json:"leaderID"`
	LeaderName  string        `json:"leaderName"`
	Leader      UserJSON      `json:"leader"`
	Missions    []MissionJSON `json:"missions"` //创建或更新不会修改该字段，仅拉取使用
	ProjectName string        `json:"projectName"`
}

type ModuleBriefJSON struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	StartTime string             `json:"startTime"`
	EndTime   string             `json:"endTime"`
	Leader    UserJSON           `json:"leader"`
	Missions  []MissionBriefJSON `json:"missions"`
}

func module2ModuleJson(module *models.Module) (moduleJSON ModuleJSON) {
	moduleJSON.ID = module.ID
	moduleJSON.Name = module.Name
	moduleJSON.CreatorID = module.CreatorID
	creator := user2UserJSON(module.Creator)
	moduleJSON.Creator = userJSON2UserBriefJSON(creator)
	moduleJSON.CreateTime = module.CreateTime
	moduleJSON.StartTime = module.StartTime
	moduleJSON.EndTime = module.EndTime
	moduleJSON.Content = module.Content
	moduleJSON.Target = module.Target
	moduleJSON.Tag = module.Tag
	moduleJSON.LeaderID = module.LeaderID
	leader := user2UserJSON(module.Leader)
	moduleJSON.Leader = userJSON2UserBriefJSON(leader)
	moduleJSON.Missions, _ = MissionsFindByModuleID(module.ID)
	moduleJSON.ProjectID = module.ProjectID
	return
}

func moduleJSON2ModuleBriefJson(moduleJSON1 *ModuleJSON) (moduleJSON2 ModuleJSON) {
	moduleJSON2.ID = moduleJSON1.ID
	moduleJSON2.Name = moduleJSON1.Name
	moduleJSON2.CreateTime = moduleJSON1.CreateTime
	moduleJSON2.Content = moduleJSON1.Content
	moduleJSON2.Tag = moduleJSON1.Tag
	moduleJSON2.Leader = moduleJSON1.Leader
	moduleJSON2.ProjectID = moduleJSON1.ProjectID
	return
}

func modules2ModulesBriefJSON(modules []models.Module) (modulesJSON []ModuleJSON) {
	modulesJSON = make([]ModuleJSON, len(modules))
	for i, v := range modules {
		m := module2ModuleJson(&v)
		modulesJSON[i] = moduleJSON2ModuleBriefJson(&m)
	}
	return
}

func (moduleJSON *ModuleJSON) moduleJSON2Module() (module models.Module) {
	module.ID = moduleJSON.ID
	module.Name = moduleJSON.Name
	module.CreatorID = moduleJSON.CreatorID
	module.CreateTime = moduleJSON.CreateTime
	module.StartTime = moduleJSON.StartTime
	module.EndTime = moduleJSON.EndTime
	module.Content = moduleJSON.Content
	module.Target = moduleJSON.Target
	module.Tag = moduleJSON.Tag
	module.LeaderID = moduleJSON.LeaderID
	module.ProjectID = moduleJSON.ProjectID
	return
}

func (moduleJSON *ModuleJSON) Create() (err error) {
	creator := UserJSON{ID: moduleJSON.CreatorID}
	if err = creator.First(); err == nil {
		m := moduleJSON.moduleJSON2Module()
		if err = m.Create(); err == nil {
			*moduleJSON = module2ModuleJson(&m)
		}
	}
	if err != nil {
		err = errors.New(titleModule + "Insert:\t" + err.Error())
	}
	return
}

func (moduleJSON *ModuleJSON) First() (err error) {
	m := moduleJSON.moduleJSON2Module()
	if err = m.First(); err == nil {
		*moduleJSON = module2ModuleJson(&m)
	} else {
		err = errors.New(titleModule + "First:\t" + err.Error())
	}
	return
}

func ModuleFindByID(id uint) (moduleJSON ModuleJSON, err error) {
	moduleJSON = ModuleJSON{ID: id}
	err = moduleJSON.First()
	return
}

func ModulesFindByCreatorID(id uint) (modulesJSON []ModuleJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:15
	*/
	if modules, err1 := models.ModulesFindByCreatorID(id); err1 == nil {
		modulesJSON = modules2ModulesBriefJSON(modules)
	} else {
		err = errors.New(titleModule + "MissionsFindByCreatorID:\t" + err1.Error())
	}
	return
}

func ModulesFindByLeaderID(id uint) (
	modulesJSON []ModuleJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:15
	*/
	if modules, err1 := models.ModulesFindByLeaderID(id); err1 == nil {
		modulesJSON = modules2ModulesBriefJSON(modules)
	} else {
		err = errors.New(titleModule + "ModulesFindByLeaderID:\t" + err1.Error())
	}
	return
}

func ModulesFindByParticipantID(id uint) (modulesJSON []ModuleJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:14
	*/
	if modules, err1 := models.ModulesFindByParticipantID(id); err1 == nil {
		modulesJSON = modules2ModulesBriefJSON(modules)
	} else {
		err = errors.New(titleModule + "ModulesFindByLeaderID:\t" + err1.Error())
	}
	return
}

func ModulesFindByProjectID(id uint) (modulesJSON []ModuleJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:14
	*/
	if modules, err1 := models.ModulesFindByProjectID(id); err1 == nil {
		modulesJSON = modules2ModulesBriefJSON(modules)
	} else {
		err = errors.New(titleModule + "ModulesFindByLeaderID:\t" + err1.Error())
	}
	return
}

func (moduleJSON *ModuleJSON) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:16
	*/
	m := moduleJSON.moduleJSON2Module()
	if err = m.Updates(); err == nil {
		*moduleJSON = module2ModuleJson(&m)
	} else {
		err = errors.New(titleMission + "Update:\t" + err.Error())
	}
	return
}

func (moduleJSON *ModuleJSON) Delete() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:18
	*/
	m := moduleJSON.moduleJSON2Module()
	if err = m.Delete(); err == nil {
		*moduleJSON = module2ModuleJson(&m)
	} else {
		err = errors.New(titleMission + "Update:\t" + err.Error())
	}
	return
}

func ModuleDeleteByID(id uint) (moduleJSON ModuleJSON, err error) {
	moduleJSON = ModuleJSON{ID: id}
	err = moduleJSON.Delete()
	return
}
