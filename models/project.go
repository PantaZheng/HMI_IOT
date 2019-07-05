package models

import (
	"github.com/pantazheng/bci/database"
	"log"
	"strconv"
	"time"
)

type ProjectCore struct {
	ID          uint   `gorm:"primary_key",json:"id"`
	Name        string `json:"name"`
	State       uint   `json:"state"`
	ManagerName string `gorm:"-",json:"managerName"`
}

type Project struct {
	ProjectCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `gorm:"-",json:"createTime"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-",json:"updateTime"`
	DeletedAt  *time.Time `sql:"index",json:"-"`
	StartTime  string     `json:"startTime"`
	EndTime    string     `json:"endTime"`
	Content    string     `json:"content"`
	Target     string     `json:"target"`

	ManagerID uint `json:"managerID"`
}

func projectTestData() {
	log.Println("projectTestData")
	l := 8
	projects := make([]Project, l)

	for i := 0; i < l; i++ {
		projects[i].Name = "Project" + strconv.Itoa(i)
		projects[i].ManagerID = uint(i/2 + 1)
	}

	for _, v := range projects {
		if err := v.Insert(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
}

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
		err = database.DB.Model(Project{}).Find(&project).Error
		return
	}
	if field == "member" {
		//project
		mission := Mission{}
		projectAmount := 0
		if err = database.DB.Model(Project{}).Count(&projectAmount).Error; err != nil {
			return
		}
		projectAmount++
		projectCount := make([]uint, projectAmount)
		//owner
		if missions, err := mission.Find("owner", id); err != nil {
			return
		} else {
			for _, v := range missions {
				projectCount[v.ProjectID]++
			}
		}
		//manager
		if leaderProjects, err := project.Find("manager", id); err != nil {
			return
		} else {
			for _, v := range leaderProjects {
				projectCount[v.ID]++
			}
		}
		//merge
		for i := 0; i < projectAmount; i++ {
			if projectCount[i] > 0 {
				project.ID = uint(i)
				_ = project.First()
				projects = append(projects, *project)
			}
		}
	}
	err = database.DB.Where(field+"_id=?", id).Find(&project).Error
	return
}

func (project *Project) FindBrief(field string, id uint) (projectsCore []ProjectCore, err error) {
	if projects, err := project.Find(field, id); err != nil {
		return
	} else {
		l := len(projects)
		projectsCore := make([]ProjectCore, l)
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

func (project *Project) Updates() (err error) {
	if err = database.DB.Model(Project{}).Where("id=?", project.ID).Updates(&project).Error; err != nil {
		return
	}
	err = project.First()
	return
}

func (project *Project) Delete() (err error) {
	if err = project.First(); err != nil {
		return
	}
	err = database.DB.Model(Module{}).Where("id=?", project.ID).Delete(&project).Error
	return
}
