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
	Name         string
	CreatorID    uint
	Creator      User
	CreateTime   string
	StartTime    string
	EndTime      string
	Content      string
	Tag          bool
	Participants []User `gorm:"many2many:user_modules"`
	LeaderID     uint
	Leader       User
	ProjectID    uint
	//Project      Project
}

//func (module *Module) checkForeignKey() (err error) {
//	/**
//	@Author: PantaZheng
//	@Description:
//	@Date: 2019/5/15 18:02
//	*/
//	m := &Module{}
//	m.ID = Module{}.Project
//	if err = m.First(); err == nil {
//		if mission.CreatorID > 0 {
//			err = errors.New("OwnerID not in mission's participants")
//			for _, v := range m.Participants {
//				if mission.CreatorID == v.ID {
//					err = nil
//				}
//			}
//		}
//	}
//	return
//}

func (module *Module) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 16:34
	*/
	module.ID = 0
	module.CreateTime = time.Now().Format("2006-01-02")
	participants := module.Participants
	module.Participants = make([]User, 0)
	if err = database.DB.Create(&module).Error; err == nil {
		if participants != nil {
			err = database.DB.Model(&module).Association("Participants").Append(participants).Error
		}
		if err == nil {
			err = module.First()
		}
	}
	if err != nil {
		err = errors.New(titleModule + "Create:\t" + err.Error())
	}
	return
}

func (module *Module) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 16:49
	*/
	if module.ID > 0 {
		m := &Module{}
		m.ID = module.ID
		if err = database.DB.First(&m).Error; err == nil {
			*module = *m
			if err = database.DB.Model(&module).Association("Participants").Find(&module.Participants).Error; err == nil {
				module.Creator.ID = module.CreatorID
				module.Leader.ID = module.LeaderID
				if err = module.Creator.First(); err == nil {
					err = module.Leader.First()
				}
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
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 17:47
	*/
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

func ModulesFindByParticipantsID(id uint) (modules []Module, err error) {
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
		err = errors.New(titleModule + "ModulesFindByParticipantsID:\t" + err.Error())
	}
	return
}

func (module *Module) Updates() (err error) {
	m := &Module{}
	m.ID = module.ID
	participants := module.Participants
	module.Participants = nil
	if err = database.DB.Model(&m).Updates(&module).Error; err == nil {
		if participants != nil {
			err = database.DB.Model(&m).Association("Participants").Replace(participants).Error
		}
		if err == nil {
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
