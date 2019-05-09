package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"time"
)

type Module struct {
	gorm.Model
	Name         string
	CreatorID    uint
	CreateTime   string
	StartTime    string
	EndTime      string
	Content      string
	Tag          bool    //tag有Module负责人修改
	Participants []*User `gorm:"many2many:user_modules"`
	LeaderID     uint
	Leader       User
	ProjectID    uint
	Project      Project
}

type ModuleJson struct {
	ID           uint               `json:"id"`
	Name         string             `json:"name"`
	Creator      UserBriefJSON      `json:"creator"`
	CreateTime   string             `json:"createTime"` //创建时间
	StartTime    string             `json:"startTime"`  //开始时间
	EndTime      string             `json:"endTime"`    //结束时间
	Content      string             `json:"content"`
	Tag          bool               `json:"tag"`
	ProjectID    uint               `json:"projectID"`
	Leader       UserBriefJSON      `json:"leader"`
	Participants []UserBriefJSON    `json:"participants"` //参与人员
	Missions     []MissionBriefJson `json:"missions"`     //创建或更新不会修改该字段，仅拉取使用
}

type ModuleBriefJson struct {
	ID         uint          `json:"id"`
	Name       string        `json:"name"`
	CreateTime string        `json:"createTime"` //创建时间
	Content    string        `json:"content"`
	Tag        bool          `json:"tag"`
	Leader     UserBriefJSON `json:"leader"`
	ProjectID  uint          `json:"project"`
}

func moduleTestData() {
	u2 := &UserBriefJSON{ID: 2}
	u3 := &UserBriefJSON{ID: 3}
	u4 := &UserBriefJSON{ID: 4}
	u5 := &UserBriefJSON{ID: 5}
	u6 := &UserBriefJSON{ID: 6}
	_, _ = ModuleCreate(&ModuleJson{Name: "Module1", Creator: *u2, ProjectID: 1, Leader: *u2, Participants: []UserBriefJSON{*u2, *u3}})
	_, _ = ModuleCreate(&ModuleJson{Name: "Module1", Creator: *u2, ProjectID: 2, Leader: *u2, Participants: []UserBriefJSON{*u6, *u2}})
	_, _ = ModuleCreate(&ModuleJson{Name: "Module1", Creator: *u4, ProjectID: 1, Leader: *u3, Participants: []UserBriefJSON{*u2, *u3, *u4}})
	_, _ = ModuleCreate(&ModuleJson{Name: "Module1", Creator: *u4, ProjectID: 2, Leader: *u5, Participants: []UserBriefJSON{*u2, *u3, *u6}})
}

func (module *Module) moduleJson2Module(moduleJson *ModuleJson) {
	module.ID = moduleJson.ID
	module.Name = moduleJson.Name
	module.CreatorID = moduleJson.Creator.ID
	module.CreateTime = moduleJson.CreateTime
	module.StartTime = moduleJson.StartTime
	module.EndTime = moduleJson.EndTime
	module.Content = moduleJson.Content
	module.Tag = moduleJson.Tag
	module.ProjectID = moduleJson.ProjectID
	module.LeaderID = moduleJson.Leader.ID
}

func (moduleJson *ModuleJson) module2ModuleJson(module *Module) {
	moduleJson.ID = module.ID
	moduleJson.Name = module.Name
	creator := &User{}
	database.DB.Model(&module).Related(&creator, "CreatorID")
	moduleJson.Creator.User2UserBriefJSON(creator)
	moduleJson.CreateTime = module.CreateTime
	moduleJson.StartTime = module.StartTime
	moduleJson.EndTime = module.EndTime
	moduleJson.Content = module.Content
	moduleJson.Tag = module.Tag
	moduleJson.ProjectID = module.ProjectID
	leader := &User{}
	database.DB.Model(&module).Related(&leader, "LeaderID")
	moduleJson.Leader.User2UserBriefJSON(leader)
	participants := make([]*User, len(module.Participants))
	database.DB.Model(&module).Related(&participants, "Participants")
	tempUser := &UserBriefJSON{}
	for _, v := range participants {
		tempUser.User2UserBriefJSON(v)
		moduleJson.Participants = append(moduleJson.Participants, *tempUser)
	}
	moduleJson.Missions, _ = MissionsFindByModule(module)
}

func (moduleBriefJson *ModuleBriefJson) module2ModuleBriefJson(module *Module) {
	moduleBriefJson.ID = module.ID
	moduleBriefJson.Name = module.Name
	moduleBriefJson.CreateTime = module.CreateTime
	moduleBriefJson.Content = module.Content
	moduleBriefJson.Tag = module.Tag
	moduleBriefJson.ProjectID = module.ProjectID
	leader := &User{}
	database.DB.Model(&module).Related(&leader, "LeaderID")
	moduleBriefJson.Leader.User2UserBriefJSON(leader)
}

func ModuleCreate(moduleJson *ModuleJson) (recordModuleJson ModuleJson, err error) {
	newModule := new(Module)
	newModule.moduleJson2Module(moduleJson)
	newModule.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	if err = database.DB.Create(&newModule).Error; err != nil {
		return
	}
	if err = database.DB.Model(&newModule).First(&newModule).Error; err == nil {
		users := make([]User, len(moduleJson.Participants))
		for i, v := range moduleJson.Participants {
			users[i].ID = v.ID
		}
		err = database.DB.Model(&newModule).Association("Participants").Append(users).Error
		recordModuleJson.module2ModuleJson(newModule)
	}
	return
}

func ModuleFind(module *Module) (recordModuleJson ModuleJson, err error) {
	recordModule := new(Module)
	if err = database.DB.First(&recordModule, &module).Error; err == nil {
		recordModuleJson.module2ModuleJson(recordModule)
	}
	return
}

func ModulesFindByLeader(leader *User) (modulesBriefJson []ModuleBriefJson, err error) {
	modules := make([]Module, 1)
	if err = database.DB.Model(&leader).Related(&modules, "LeaderID").Error; err != nil {
		return
	}
	if len(modules) == 0 {
		err = errors.New("ModulesFindByLeader No Module Record")
	} else {
		for _, v := range modules {
			tempJson := &ModuleBriefJson{}
			tempJson.module2ModuleBriefJson(&v)
			modulesBriefJson = append(modulesBriefJson, *tempJson)
		}
	}
	return
}

func ModulesFindByProject(project *Project) (modulesBriefJson []ModuleBriefJson, err error) {
	modules := make([]Module, 1)
	if err = database.DB.Model(&project).Related(&modules, "ProjectID").Error; err != nil {
		return
	}
	if len(modules) == 0 {
		err = errors.New("ModulesFindByProject No Module Record")
	} else {
		for _, v := range modules {
			tempJson := &ModuleBriefJson{}
			tempJson.module2ModuleBriefJson(&v)
			modulesBriefJson = append(modulesBriefJson, *tempJson)
		}
	}
	return
}

func ModuleUpdate(moduleJson *ModuleJson) (recordModuleJson ModuleJson, err error) {
	updateModule := new(Module)
	updateModule.moduleJson2Module(moduleJson)
	recordModule := new(Module)
	recordModule.ID = updateModule.ID
	if database.DB.First(&recordModule, &recordModule).RecordNotFound() {
		err = errors.New("MissionUpdate No Module Record")
	} else {
		database.DB.Model(&recordModule).Updates(updateModule)
		if num := len(moduleJson.Participants); num != 0 {
			users := make([]User, num)
			for i, v := range moduleJson.Participants {
				users[i].ID = v.ID
			}
			err = database.DB.Model(&recordModule).Association("Participants").Replace(users).Error
		}
		recordModuleJson.module2ModuleJson(recordModule)
	}
	return
}

func ModuleDelete(module *Module) (recordModuleJson ModuleJson, err error) {
	recordModule := new(Module)
	if database.DB.First(&recordModule, &module).RecordNotFound() {
		err = errors.New("ModuleDelete No Module Record")
	} else {
		recordModuleJson.module2ModuleJson(recordModule)
		err = database.DB.Unscoped().Delete(&recordModule).Error
	}
	return
}
