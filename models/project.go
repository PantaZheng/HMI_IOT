package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
)

type ProjectCore struct {
	gorm.Model
	Name      string
	StartTime string
	EndTime   string
	Type      string
	Content   string
	Target    string
	File      string
	State     uint
	CreatorID uint
	LeaderID  uint
}

type Project struct {
	ProjectCore
	Participants []User `gorm:"many2many:user_projects"`
}

func (project *Project) Insert() (err error) {
	project.State = 1
	if err = database.DB.Create(&project).Error; err != nil {
		return
	}
	if len(project.Participants) != 0 {
		err = project.Update()
	} else {
		err = project.First()
	}
	return
}

func (project *Project) First() (err error) {
	if err = database.DB.Where("id = ?", project.ID).First(&project).Error; err != nil {
		return
	}
	database.DB.Model(&project).Association("Participants").Find(&project.Participants)
	return
}

func (project *Project) Find(field string) (projects []Project, err error) {
	if field == "creator_id" {
		err = database.DB.Where("creator_id=?", project.CreatorID).Find(&projects).Error
	} else if field == "leader_id" {
		err = database.DB.Where("leader_id=?", project.LeaderID).Find(&projects).Error
	} else if field == "participant_id" {
		err = database.DB.Model(User{}).Where("id?=", project.LeaderID).Related(&projects, "PProjects").Error
	} else if field == "member_id" {
		err = database.DB.Model(User{}).Where("id?=", project.LeaderID).Related(&projects, "PProjects").Error
		projectsTemp := make([]Project, 0)
		err = database.DB.Where("leader_id=?", project.LeaderID).Find(projectsTemp).Error
		for i := 0; i < len(projectsTemp); i++ {
			notExist := true
			for j := 0; j < len(projects); j++ {
				if projects[i].ID == projectsTemp[j].ID {
					notExist = false
					break
				}
			}
			if notExist {
				projects = append(projects, projectsTemp[i])
			}
		}
	} else if field == "all " {
		err = database.DB.Model(Project{}).Find(projects).Error
	} else {
		err = errors.New("no this field")
	}
	if err == nil {
		for i := 0; i < len(projects); i++ {
			if err = projects[i].First(); err != nil {
				break
			}
		}
	}
	return
}

func (project *Project) Update() (err error) {
	participants := project.Participants
	if err = database.DB.Where("id=?", project.ID).Updates(&project).Error; err != nil {
		return
	}
	if participants != nil {
		if err = database.DB.Model(&project).Association("Participants").Replace(participants).Error; err != nil {
		}
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
