package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
)

type ModuleCore struct {
	gorm.Model
	Name      string
	StartTime string
	EndTime   string
	Content   string
	Target    string
	File      string
	State     uint
	LeaderID  uint
}

type Module struct {
	ModuleCore
	//const foreign
	CreatorID uint
	ProjectID uint
}

//Insert
func (module *Module) Insert() (err error) {
	module.State = 1
	project := Project{}
	project.ID = module.ProjectID
	if err = project.First(); err != nil {
		return
	}
	if err = database.DB.Create(&module).Error; err != nil {
		return
	}
	err = module.First()
	return
}

func (module *Module) First() (err error) {
	err = database.DB.Model(Module{}).Where("id = ?", module.ID).First(&module).Error
	return
}

//Find
func (module *Module) Find(field string) (modules []Module, err error) {
	if field == "creator_id" {
		err = database.DB.Where("creator_id = ?", module.CreatorID).Find(&modules).Error
	} else if field == "leader_id" {
		err = database.DB.Where("leader_id = ?", module.LeaderID).Find(&modules).Error
	} else if field == "project_id" {
		err = database.DB.Where("project_id = ?", module.ProjectID).Find(&modules).Error
	} else if field == "all" {
		err = database.DB.Model(Module{}).Find(&modules).Error
	} else {
		err = errors.New("no this field")
	}
	return
}

//Updates
func (module *Module) Update() (err error) {
	if err = database.DB.Where("id=?", module.ID).Updates(&module).Error; err != nil {
		return
	}
	err = module.First()
	return
}

// Delete
func (module *Module) Delete() (err error) {
	if err = module.First(); err != nil {
		return
	}
	err = database.DB.Model(Module{}).Where("id=?", module.ID).Delete(&module).Error
	return
}
