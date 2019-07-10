package models

import (
	"github.com/pantazheng/bci/database"
	"log"
	"time"
)

type ProjectCore struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `json:"name"`
	State       uint   `json:"state"`
	ManagerName string `gorm:"-" json:"managerName"`
	StartTime   string `json:"startTime"`
	EndTime     string `json:"endTime"`
}

type Project struct {
	ProjectCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `gorm:"-" json:"createTime"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-" json:"updateTime"`
	DeletedAt  *time.Time `sql:"index" json:"-"`

	Content string `json:"content"`
	Target  string `json:"target"`

	ManagerID uint `json:"managerID"`
}

type ModuleFrame struct {
	ModuleCore
	Missions []MissionCore `json:"missions"`
}

type ProjectFrame struct {
	ProjectCore
	Modules []ModuleFrame `json:"modules"`
}

//func projectTestData() {
//	log.Println("projectTestData")
//	l := 4
//	projects := make([]Project, l)
//
//	for i := 0; i < l; i++ {
//		projects[i].Name = "Project" + strconv.Itoa(i)
//		projects[i].ManagerID = uint(i/2 + 1)
//	}
//
//	for _, v := range projects {
//		if err := v.Insert(); err != nil {
//			log.Println(err.Error())
//		} else {
//			log.Println(v)
//		}
//	}
//}

func (project *Project) Insert() (err error) {
	project.State = 1
	if err = database.DB.Create(&project).Error; err != nil {
		return
	}
	err = project.First()
	return
}

func (project *Project) First() (err error) {
	if err = database.DB.Model(Project{}).Where("id = ?", project.ID).First(&project).Error; err != nil {
		return
	}
	manager := UserCore{ID: project.ManagerID}
	if err = manager.First(); err != nil {
		return
	}
	project.ManagerName = manager.Name
	return
}

func (project *Project) Find(field string, id uint) (projects []Project, err error) {
	if field == "all" {
		err = database.DB.Model(Project{}).Find(&projects).Error
		return
	}
	if field == "member" {
		//p
		mission := Mission{}
		p := Project{}
		if err = database.DB.Model(Project{}).Last(p).Error; err != nil {
			return
		}
		projectAmount := int(p.ID)
		projectCount := make([]uint, projectAmount)
		//owner
		missions, e := mission.Find("owner", id)
		if e != nil {
			err = e
			return
		}
		for _, v := range missions {
			projectCount[v.ProjectID]++
		}

		//manager
		if leaderProjects, e := p.Find("manager", id); e != nil {
			err = e
			return
		} else {
			for _, v := range leaderProjects {
				log.Println(v.ID)
				projectCount[v.ID]++
			}
		}
		//merge
		for i := 0; i < projectAmount; i++ {
			if projectCount[i] > 0 {
				p.ID = uint(i)
				_ = p.First()
				projects = append(projects, p)
			}
		}
		return
	}
	err = database.DB.Where(field+"_id=?", id).Find(&projects).Error
	return
}

func (project *Project) FindBrief(field string, id uint) (projectsCore []ProjectCore, err error) {
	if projects, e := project.Find(field, id); e != nil {
		err = e
		return
	} else {
		l := len(projects)
		projectsCore = make([]ProjectCore, l)
		for i, v := range projects {
			projectsCore[i] = v.ProjectCore
			manager := UserCore{ID: v.ManagerID}
			if err = manager.First(); err != nil {
				break
			}
			projectsCore[i].ManagerName = manager.Name
		}
	}
	return
}

func (project *Project) FindFrame() (projectFrame ProjectFrame, err error) {
	if err = project.First(); err != nil {
		return
	}
	projectFrame.ProjectCore = project.ProjectCore
	module := Module{}
	modules, e := module.FindBrief("project", project.ID)
	if e != nil {
		err = e
		return
	}
	modulesLength := len(modules)
	projectFrame.Modules = make([]ModuleFrame, modulesLength)
	mission := Mission{}
	for i := 0; i < len(modules); i++ {
		projectFrame.Modules[i].ModuleCore = modules[i]
		projectFrame.Modules[i].Missions, err = mission.FindBrief("module", projectFrame.Modules[i].ID)
		if err != nil {
			return
		}
	}
	return
}

func (project *Project) Updates() (err error) {
	nullState := uint(0)
	moduleKeyState := uint(3)
	if err = database.DB.Model(Project{}).Where("id=?", project.ID).Updates(&project).Error; err != nil {
		return
	}
	if project.State != nullState {
		mission := Mission{}
		missions, _ := mission.Find("project", project.ID)
		for _, v := range missions {
			v.State = project.State
			_ = v.Updates()
		}
	}
	if project.State == moduleKeyState {
		module := Module{}
		modules, e := module.FindByField("all", project.ID)
		if e != nil {
			err = e
			return
		}
		for _, v := range modules {
			if v.State != moduleKeyState {
				v.State = moduleKeyState
				if err = v.Updates(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		return
	}
	err = project.First()
	return
}

func (project *Project) Delete() (err error) {
	if err = project.First(); err != nil {
		return
	}
	if err = database.DB.Model(Module{}).Where("id=?", project.ID).Delete(&project).Error; err != nil {
		return
	}
	module := Module{}
	if _, err = module.DeleteByField("project", project.ID); err != nil {
		return
	}
	mission := Mission{}
	if _, err = mission.DeleteByField("project", project.ID); err != nil {
		return
	}
	gain := Gain{}
	_, err = gain.DeleteByField("project", project.ID)
	return
}
