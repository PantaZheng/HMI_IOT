package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"time"
)

const titleModule = "models.module."

type Module struct {
	gorm.Model
	Name       string
	CreatorID  uint
	Creator    User
	CreateTime string
	StartTime  string
	EndTime    string
	Content    string
	Target     string
	Tag        bool
	LeaderID   uint
	Leader     User
	ProjectID  uint
	Project    Project
}

func (module *Module) checkForeignKey() (err error) {
	p := &Project{}
	p.ID = module.ProjectID
	err = p.First()
	return
}

func (module *Module) Create() (err error) {
	module.ID = 0
	if err = module.checkForeignKey(); err == nil {
		module.CreateTime = time.Now().Format("2006-01-02")
		if err = database.DB.Create(&module).Error; err == nil {
			err = module.First()
		}
	}
	if err != nil {
		err = errors.New(titleModule + "Create:\t" + err.Error())
	}
	return
}

func (module *Module) First() (err error) {
	if module.ID > 0 {
		m := &Module{}
		m.ID = module.ID
		if err = database.DB.First(&m).Error; err == nil {
			*module = *m
			module.Creator.ID = module.CreatorID
			module.Leader.ID = module.LeaderID
			if err = module.Creator.First(); err == nil {
				err = module.Leader.First()
			}
		}
	} else {
		err = errors.New("ID必须为正数")
	}
	if err != nil {
		err = errors.New(titleModule + "First:\t" + err.Error())
	}
	return
}

func ModulesFindByCreatorID(id uint) (modules []Module, err error) {
	creator := &User{}
	creator.ID = id
	if err = creator.First(); err == nil {
		if err = database.DB.Model(&creator).Related(&modules, "CreatorID").Error; err == nil {
			for i := 0; i < len(modules); i++ {
				if err = modules[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleModule + "MissionsFindByCreatorID:\t" + err.Error())
	}
	return
}

func ModulesFindByLeaderID(id uint) (modules []Module, err error) {
	leader := &User{}
	leader.ID = id
	if err = leader.First(); err == nil {
		if err = database.DB.Model(&leader).Related(&modules, "LeaderID").Error; err == nil {
			for i := 0; i < len(modules); i++ {
				if err = modules[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleModule + "ModulesFindByLeaderID:\t" + err.Error())
	}
	return
}

func ModulesFindByParticipantID(id uint) (modules []Module, err error) {
	participant := &User{}
	participant.ID = id
	if err = participant.First(); err == nil {
		if err = database.DB.Model(&participant).Related(&modules, "PModules").Error; err == nil {
			for i := 0; i < len(modules); i++ {
				if err = modules[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleModule + "ModulesFindByParticipantID:\t" + err.Error())
	}
	return
}

func ModulesFindByProjectID(id uint) (modules []Module, err error) {
	project := &Project{}
	project.ID = id
	if err = project.First(); err == nil {
		if err = database.DB.Model(&project).Related(&modules, "ProjectID").Error; err == nil {
			for i := 0; i < len(modules); i++ {
				if err = modules[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleModule + "ModulesFindByProjectID:\t" + err.Error())
	}
	return
}

func (module *Module) Updates() (err error) {
	if err = module.checkForeignKey(); err == nil {
		m := &Module{}
		m.ID = module.ID
		if err = database.DB.Model(&m).Updates(&module).Error; err == nil {
			err = module.First()
		}
	}
	if err != nil {
		err = errors.New(titleModule + "Updates\t" + err.Error())
	}
	return
}

func (module *Module) Delete() (err error) {
	if err = module.First(); err == nil {
		m := Mission{}
		m.ID = module.ID
		err = database.DB.Delete(&m).Error
	}
	if err != nil {
		err = errors.New(titleModule + "Delete\t" + err.Error())
	}
	return
}
