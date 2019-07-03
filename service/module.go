package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"log"
	"strconv"
)

const titleModule = "service.module."

type ModuleCore struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Content   string `json:"content"`
	Target    string `json:"target"`
	File      string `json:"file"`
	State     uint   `json:"state"`

	LeaderID   uint   `json:"leaderID"`
	LeaderName string `json:"leaderName"`

	CreatorID   uint   `json:"creatorID"`
	CreatorName string `json:"creatorName"`
}

type ModuleJSON struct {
	ModuleCore

	Missions    []MissionJSON `json:"missions"`
	ProjectID   uint          `json:"projectID"`
	ProjectName string        `json:"projectName"`
}

func moduleTestData() {
	log.Println("moduleTestData")
	l := 16
	modules := make([]ModuleJSON, l)
	for i := 0; i < l; i++ {
		modules[i].Name = "module" + strconv.Itoa(i)
		modules[i].LeaderID = uint(i/2 + 1)
		modules[i].ProjectID = uint(i/4 + 1)
	}

	for _, v := range modules {
		if err := v.Insert(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
}

func (moduleJSON *ModuleJSON) module2ModuleJson(module models.Module) {
	moduleJSON.ID = module.ID
	moduleJSON.CreatedAt = module.CreatedAt.Format("2006-01-02")
	moduleJSON.UpdatedAt = module.UpdatedAt.Format("2006-01-02")
	moduleJSON.Name = module.Name
	moduleJSON.StartTime = module.StartTime
	moduleJSON.EndTime = module.EndTime
	moduleJSON.Content = module.Content
	moduleJSON.Target = module.Target
	moduleJSON.File = module.File
	moduleJSON.State = module.State

	moduleJSON.LeaderID = module.LeaderID
	leader := UserJSON{ID: moduleJSON.LeaderID}
	_ = leader.First()
	moduleJSON.LeaderName = leader.Name

	moduleJSON.CreatorID = module.CreatorID
	creator := UserJSON{ID: moduleJSON.CreatorID}
	_ = creator.First()
	moduleJSON.CreatorName = creator.Name

	mission := MissionJSON{}
	moduleJSON.Missions, _ = mission.Find("module_id")

	moduleJSON.ProjectID = module.ProjectID
	project := ProjectJSON{}
	project.ID = moduleJSON.ProjectID
	_ = project.First()
	moduleJSON.ProjectName = project.Name

	return
}

func (moduleJSON *ModuleJSON) moduleJSON2Module() (module models.Module) {
	module.ID = moduleJSON.ID
	module.Name = moduleJSON.Name
	module.StartTime = moduleJSON.StartTime
	module.EndTime = moduleJSON.EndTime
	module.Content = moduleJSON.Content
	module.Target = moduleJSON.Target
	module.File = moduleJSON.File
	module.State = moduleJSON.State

	module.CreatorID = moduleJSON.CreatorID
	module.LeaderID = moduleJSON.LeaderID
	module.ProjectID = moduleJSON.ProjectID
	return
}

//Insert
func (moduleJSON *ModuleJSON) Insert() (err error) {
	m := moduleJSON.moduleJSON2Module()
	if err = m.Insert(); err == nil {
		moduleJSON.module2ModuleJson(m)
	} else {
		err = errors.New(titleModule + "Insert:\t" + err.Error())
	}
	return
}

//First
func (moduleJSON *ModuleJSON) First() (err error) {
	m := moduleJSON.moduleJSON2Module()
	if err = m.First(); err == nil {
		moduleJSON.module2ModuleJson(m)
	} else {
		err = errors.New(titleModule + "First:\t" + err.Error())
	}
	return
}

//Find
func (moduleJSON *ModuleJSON) Find(field string) (modulesJSON []ModuleJSON, err error) {
	m := moduleJSON.moduleJSON2Module()
	if modules, err := m.Find(field); err != nil {
		err = errors.New(titleModule + "Find:\t" + err.Error())
	} else {
		modulesJSON = make([]ModuleJSON, len(modules))
		for i, v := range modules {
			modulesJSON[i].module2ModuleJson(v)
		}
	}
	return
}

func (moduleJSON *ModuleJSON) Update() (err error) {
	if moduleJSON.ID == 0 {
		err = errors.New(titleModule + "Updates:\t id 不可缺")
		return
	}
	m := moduleJSON.moduleJSON2Module()
	if err = m.Update(); err == nil {
		moduleJSON.module2ModuleJson(m)
	} else {
		err = errors.New(titleMission + "Updates:\t" + err.Error())
	}
	return
}

func (moduleJSON *ModuleJSON) Delete() (err error) {
	m := models.Module{}
	if err = m.Delete(); err == nil {
		moduleJSON.module2ModuleJson(m)
	} else {
		err = errors.New(titleMission + "Delete:\t" + err.Error())
	}
	return
}
