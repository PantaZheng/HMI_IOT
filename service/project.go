package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"log"
	"strconv"
)

const titleProject = "service.project."

type ProjectCore struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Content   string `json:"content"`
	Target    string `json:"target"`
	State     uint   `json:"state"`
}

type ProjectJSON struct {
	ProjectCore
	LeaderID     uint         `json:"leaderID"`
	LeaderName   string       `json:"leaderName"`
	CreatorID    uint         `json:"creatorID"`
	CreatorName  string       `json:"creatorName"`
	Participants []UserJSON   `json:"participants"`
	Modules      []ModuleJSON `json:"modules"` //仅拉取更新
}

func projectTestData() {
	log.Println("projectTestData")
	l := 8
	projects := make([]ProjectJSON, l)

	for i := 1; i <= l; i++ {
		projects[i].Name = "Project" + strconv.Itoa(i)
		projects[i].LeaderID = uint(i / 2)
	}

	for _, v := range projects {
		if err := v.Insert(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
}

func (projectJSON *ProjectJSON) project2ProjectJson(project models.Project) {
	projectJSON.ID = project.ID
	projectJSON.CreatedAt = project.CreatedAt.Format("2006-01-02")
	projectJSON.UpdatedAt = project.UpdatedAt.Format("2006-01-02")
	projectJSON.Name = project.Name
	projectJSON.Type = project.Type
	projectJSON.StartTime = project.StartTime
	projectJSON.EndTime = project.EndTime
	projectJSON.Content = project.Content
	projectJSON.Target = project.Target
	projectJSON.State = project.State

	projectJSON.LeaderID = project.LeaderID
	leader := UserJSON{ID: projectJSON.LeaderID}
	_ = leader.First()
	projectJSON.LeaderName = leader.Name

	projectJSON.CreatorID = project.CreatorID
	creator := UserJSON{ID: projectJSON.CreatorID}
	_ = creator.First()
	projectJSON.CreatorName = creator.Name

	projectJSON.Participants = users2BriefUsersJSON(project.Participants)

	module := ModuleJSON{}
	projectJSON.Modules, _ = module.Find("project_id")

	return
}

func (projectJSON *ProjectJSON) projectJSON2Project() (project models.Project) {
	project.ID = projectJSON.ID
	project.Name = projectJSON.Name
	project.Type = projectJSON.Type
	project.StartTime = projectJSON.StartTime
	project.EndTime = projectJSON.EndTime
	project.Content = projectJSON.Content
	project.Target = projectJSON.Target

	project.CreatorID = projectJSON.CreatorID
	project.LeaderID = projectJSON.LeaderID
	project.Participants = usersJSON2Users(projectJSON.Participants)
	return
}

//Insert
func (projectJSON *ProjectJSON) Insert() (err error) {
	p := projectJSON.projectJSON2Project()
	if err = p.Insert(); err == nil {
		projectJSON.project2ProjectJson(p)
	} else {
		err = errors.New(titleProject + "Insert:\t" + err.Error())
	}
	return
}

func (projectJSON *ProjectJSON) First() (err error) {
	p := projectJSON.projectJSON2Project()
	if err = p.First(); err == nil {
		projectJSON.project2ProjectJson(p)
	} else {
		err = errors.New(titleProject + "First:\t" + err.Error())
	}
	return
}

func (projectJSON *ProjectJSON) Find(field string) (projectsJSON []ProjectJSON, err error) {
	p := projectJSON.projectJSON2Project()
	if projects, err := p.Find(field); err != nil {
		err = errors.New(titleProject + "Find:\t" + err.Error())
	} else {
		projectsJSON = make([]ProjectJSON, len(projects))
		for i, v := range projects {
			projectsJSON[i].project2ProjectJson(v)
		}
	}
	return
}

func (projectJSON *ProjectJSON) Update() (err error) {
	if projectJSON.ID == 0 {
		err = errors.New(titleProject + "Updates:\t id 不可缺")
		return
	}
	p := projectJSON.projectJSON2Project()
	if err = p.Update(); err == nil {
		projectJSON.project2ProjectJson(p)
	} else {
		err = errors.New(titleProject + "Updates:\t" + err.Error())
	}
	return
}

func (projectJSON *ProjectJSON) Delete() (err error) {
	p := projectJSON.projectJSON2Project()
	if err = p.Delete(); err == nil {
		projectJSON.project2ProjectJson(p)
	} else {
		err = errors.New(titleProject + "Delete:\t" + err.Error())
	}
	return
}
