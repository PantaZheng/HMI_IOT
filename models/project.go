package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"time"
)

const titleProject = "models.project."

type Project struct {
	gorm.Model
	Name         string `gorm:"unique"`
	Type         string
	CreatorID    uint
	Creator      User
	CreateTime   string
	StartTime    string
	EndTime      string
	Content      string
	Target       string
	LeaderID     uint
	Leader       User
	Participants []User `gorm:"many2many:user_projects"`
	TagSet       string
	Tag          bool
}

func (project *Project) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 21:34
	*/
	project.ID = 0
	project.CreateTime = time.Now().Format("2006-01-02")
	participants := project.Participants
	project.Participants = make([]User, 0)
	if err = database.DB.Create(&project).Error; err == nil {
		if participants != nil {
			err = database.DB.Model(&project).Association("Participants").Append(participants).Error
		}
		if err == nil {
			err = project.First()
		}
	}
	if err != nil {
		err = errors.New(titleProject + "Create:\t" + err.Error())
	}
	return
}

func (project *Project) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 21:34
	*/
	if project.ID > 0 {
		p := &Project{}
		p.ID = project.ID
		if err = database.DB.First(&p).Error; err == nil {
			*project = *p
			if err = database.DB.Model(&project).Association("Participants").Find(&project.Participants).Error; err == nil {
				project.Creator.ID = project.CreatorID
				project.Leader.ID = project.LeaderID
				if err = project.Creator.First(); err == nil {
					err = project.Leader.First()
				}
			}
		}
	} else {
		err = errors.New("ID必须为正数")
	}
	if err != nil {
		err = errors.New(titleProject + "First:\t" + err.Error())
	}
	return
}

func ProjectsFindAll() (projects []Project, err error) {
	projects = make([]Project, 1)
	if err = database.DB.Find(&projects).Error; err == nil {
		for i := 0; i < len(projects); i++ {
			if err = projects[i].First(); err != nil {
				break
			}
		}
	}
	if err != nil {
		err = errors.New(titleProject + "ProjectsFindAll:\t" + err.Error())
	}
	return
}

func ProjectsFindByCreatorID(id uint) (projects []Project, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 21:52
	*/
	creator := &User{}
	creator.ID = id
	if err = creator.First(); err == nil {
		if err = database.DB.Model(&creator).Related(&projects, "CreatorID").Error; err == nil {
			for i := 0; i < len(projects); i++ {
				if err = projects[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleProject + "ProjectsFindByCreatorID:\t" + err.Error())
	}
	return
}

func ProjectsFindByLeaderID(id uint) (projects []Project, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 21:52
	*/
	leader := &User{}
	leader.ID = id
	if err = leader.First(); err == nil {
		if err = database.DB.Model(&leader).Related(&projects, "LeaderID").Error; err == nil {
			for i := 0; i < len(projects); i++ {
				if err = projects[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleProject + "ProjectsFindByLeaderID:\t" + err.Error())
	}
	return
}

func ProjectsFindByParticipantID(id uint) (projects []Project, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 21:52
	*/
	participant := &User{}
	participant.ID = id
	if err = participant.First(); err == nil {
		if err = database.DB.Model(&participant).Related(&projects, "PProjects").Error; err == nil {
			for i := 0; i < len(projects); i++ {
				if err = projects[i].First(); err != nil {
					break
				}
			}
		}
	}
	if err != nil {
		err = errors.New(titleProject + "ProjectsFindByParticipantID:\t" + err.Error())
	}
	return
}

func (project *Project) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 21:34
	*/

	p := &Project{}
	p.ID = project.ID
	participants := project.Participants
	project.Participants = nil
	if err = database.DB.Model(&p).Updates(&project).Error; err == nil {
		if participants != nil {
			err = database.DB.Model(&p).Association("Participants").Replace(participants).Error
		}
		if err == nil {
			err = project.First()
		}
	}
	if err != nil {
		err = errors.New(titleProject + "Updates\t" + err.Error())
	}
	return
}

func (project *Project) Delete() (err error) {
	if err = project.First(); err == nil {
		p := Project{}
		p.ID = project.ID
		err = database.DB.Delete(&p).Error
	}
	if err != nil {
		err = errors.New(titleProject + "DeleteSoft\t" + err.Error())
	}
	return
}
