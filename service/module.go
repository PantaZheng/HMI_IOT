package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"log"
)

const titleModule = "service.module."

type ModuleJSON struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	CreatorID    uint          `json:"creatorID"`
	Creator      UserJSON      `json:"creator"`
	CreateTime   string        `json:"createTime"` //创建时间
	StartTime    string        `json:"startTime"`  //开始时间
	EndTime      string        `json:"endTime"`    //结束时间
	Content      string        `json:"content"`
	Tag          bool          `json:"tag"`
	LeaderID     uint          `json:"leaderID"`
	Leader       UserJSON      `json:"leader"`
	Participants []UserJSON    `json:"participants"` //参与人员
	Missions     []MissionJSON `json:"missions"`     //创建或更新不会修改该字段，仅拉取使用
	ProjectID    uint          `json:"projectID"`
}

func moduleTestData() {
	log.Println("moduleTestData")
	u1 := UserJSON{ID: 2}
	u2 := UserJSON{ID: 3}
	u3 := UserJSON{ID: 4}
	u4 := UserJSON{ID: 5}
	u5 := UserJSON{ID: 6}
	u6 := UserJSON{ID: 7}
	u7 := UserJSON{ID: 8}
	modules := make([]ModuleJSON, 5)
	modules[0] = ModuleJSON{Name: "钢铁侠与浩克", CreatorID: 1, StartTime: "2001-1-1", EndTime: "11111-1-1", Content: "不得不说的秘密", LeaderID: 2, Participants: []UserJSON{u1, u2, u3, u4, u5}}
	modules[1] = ModuleJSON{Name: "海王", CreatorID: 2, StartTime: "2001-1-1", EndTime: "11111-1-1", Content: "弟弟被绿", LeaderID: 5, Participants: []UserJSON{u2, u3, u4, u5, u6}}
	modules[2] = ModuleJSON{Name: "雷神1", CreatorID: 2, StartTime: "2001-1-1", EndTime: "11111-1-1", Content: "徐", LeaderID: 7, Participants: []UserJSON{u5, u6, u7}}
	modules[3] = ModuleJSON{Name: "雷神2", CreatorID: 2, StartTime: "2001-1-1", EndTime: "11111-1-1", Content: "锦", LeaderID: 7, Participants: []UserJSON{u5}}
	modules[3] = ModuleJSON{Name: "雷神3", CreatorID: 2, StartTime: "2001-1-1", EndTime: "11111-1-1", Content: "江", LeaderID: 7, Participants: []UserJSON{u5}}
	for _, v := range modules {
		if err := v.Create(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
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
	moduleJSON.Tag = module.Tag
	moduleJSON.LeaderID = module.LeaderID
	leader := user2UserJSON(module.Leader)
	moduleJSON.Creator = userJSON2UserBriefJSON(leader)
	moduleJSON.Participants = users2BriefUsersJSON(module.Participants)
	moduleJSON.Missions, _ = MissionsFindByModuleID(module.ID)
	return
}

func moduleJSON2ModuleBriefJson(moduleJSON1 *ModuleJSON) (moduleJSON2 ModuleJSON) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 19:29
	*/
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
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 19:34
	*/
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
	module.Tag = moduleJSON.Tag
	module.LeaderID = moduleJSON.LeaderID
	module.Participants = usersJSON2Users(moduleJSON.Participants)
	module.ProjectID = moduleJSON.ProjectID
	return
}

func (moduleJSON *ModuleJSON) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:15
	*/
	creator := UserJSON{ID: moduleJSON.CreatorID}
	if err = creator.First(); err == nil {
		m := moduleJSON.moduleJSON2Module()
		if err = m.Create(); err == nil {
			*moduleJSON = module2ModuleJson(&m)
		}
	}
	if err != nil {
		err = errors.New(titleModule + "Create:\t" + err.Error())
	}
	return
}

func (moduleJSON *ModuleJSON) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 19:41
	*/
	m := moduleJSON.moduleJSON2Module()
	if err = m.First(); err == nil {
		*moduleJSON = module2ModuleJson(&m)
	} else {
		err = errors.New(titleModule + "First:\t" + err.Error())
	}
	return
}

func ModuleFindByID(id uint) (moduleJSON ModuleJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:15
	*/
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
		err = errors.New(titleMission + "Updates:\t" + err.Error())
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
		err = errors.New(titleMission + "Updates:\t" + err.Error())
	}
	return
}

func ModuleDeleteByID(id uint) (moduleJSON ModuleJSON, err error) {
	moduleJSON = ModuleJSON{ID: id}
	err = moduleJSON.Delete()
	return
}
