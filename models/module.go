package models

import (
	"github.com/pantazheng/bci/database"
	"log"
	"strconv"
	"time"
)

type ModuleCore struct {
	ID         uint   `gorm:"primary_key"`
	Name       string `json:"name"`
	State      uint   `json:"state"`
	LeaderName string `gorm:"-" json:"leaderName"`
}

type Module struct {
	ModuleCore
	CreatedAt  time.Time  `json:"-"`
	CreateTime string     `gorm:"-" json:"createTime"`
	UpdatedAt  time.Time  `json:"-"`
	UpdateTime string     `gorm:"-" json:"updateTime"`
	DeletedAt  *time.Time `sql:"index" json:"-"`
	StartTime  string     `json:"startTime"`
	EndTime    string     `json:"endTime"`
	Content    string     `json:"content"`
	Target     string     `json:"target"`

	LeaderID    uint   `json:"leaderID"`
	ProjectID   uint   `json:"projectID"`
	ProjectName string `gorm:"-" json:"projectName"`
	ManagerID   uint   `json:"managerID"`
	ManagerName string `gorm:"-" json:"managerName"`
}

func moduleTestData() {
	log.Println("moduleTestData")
	l := 8
	modules := make([]Module, l)
	for i := 0; i < l; i++ {
		modules[i].Name = "module" + strconv.Itoa(i)
		modules[i].LeaderID = uint(i/2 + 1)
		modules[i].ProjectID = uint(i/2 + 1)
	}

	for _, v := range modules {
		if err := v.Insert(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
}

//Insert
func (module *Module) Insert() (err error) {
	module.State = 1
	project := Project{}
	project.ID = module.ProjectID
	if err = project.First(); err != nil {
		return
	}
	module.ManagerID = project.ManagerID
	if err = database.DB.Create(&module).Error; err != nil {
		return
	}
	err = module.First()
	return
}

func (module *Module) First() (err error) {
	if err = database.DB.Model(Module{}).Where("id = ?", module.ID).First(&module).Error; err != nil {
		return
	}
	project := Project{}
	project.ID = module.ProjectID
	if err = project.First(); err != nil {
		return
	}
	leader := UserCore{ID: module.LeaderID}
	if err = leader.First(); err != nil {
		return
	}
	module.CreateTime = module.CreatedAt.Format("2006-01-02")
	module.UpdateTime = module.UpdatedAt.Format("2006-01-02")
	module.LeaderName = leader.Name
	module.ProjectName = project.Name
	module.ManagerName = project.ManagerName
	return
}

//Find
func (module *Module) Find(field string, id uint) (modules []Module, err error) {
	if field == "all" {
		err = database.DB.Model(Mission{}).Find(&modules).Error
		return
	}
	if field == "member" {
		//participant
		mission := Mission{}
		moduleAmount := 0
		if err = database.DB.Model(Module{}).Count(&moduleAmount).Error; err != nil {
			return
		}
		moduleAmount++
		moduleCount := make([]uint, moduleAmount)
		//owner
		if missions, e := mission.Find("owner", id); e != nil {
			err = e
			return
		} else {
			for _, v := range missions {
				moduleCount[v.ModuleID]++
			}
		}
		//leader
		if leaderModules, e := module.Find("leader", id); e != nil {
			err = e
			return
		} else {
			for _, v := range leaderModules {
				moduleCount[v.ID]++
			}
		}
		//merge
		for i := 0; i < moduleAmount; i++ {
			if moduleCount[i] > 0 {
				module.ID = uint(i)
				_ = module.First()
				modules = append(modules, *module)
			}
		}
		//leader
	}
	err = database.DB.Where(field+"_id=?", id).Find(&modules).Error
	return
}

func (module *Module) FindBrief(field string, id uint) (modulesCore []ModuleCore, err error) {
	if modules, e := module.Find(field, id); e != nil {
		err = e
		return
	} else {
		l := len(modules)
		modulesCore = make([]ModuleCore, l)
		for i, v := range modules {
			modulesCore[i] = v.ModuleCore
			leader := UserCore{ID: v.LeaderID}
			if err = leader.First(); err != nil {
				break
			}
			modulesCore[i].LeaderName = leader.Name
		}
	}
	return
}

//Updates
func (module *Module) Updates() (err error) {
	if err = database.DB.Model(Module{}).Where("id=?", module.ID).Updates(&module).Error; err != nil {
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
