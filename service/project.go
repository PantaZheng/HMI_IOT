package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"strconv"
	"strings"
)

const titleProject = "service.project."

type ProjectJSON struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/16 15:07
	*/
	ID           uint         `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	CreatorID    uint         `json:"creatorID"`
	Creator      UserJSON     `json:"creator"`
	CreateTime   string       `json:"createTime"`
	StartTime    string       `json:"startTime"`
	EndTime      string       `json:"endTime"`
	Content      string       `json:"content"`
	Target       string       `json:"target"`
	LeaderID     uint         `json:"leaderID"`
	Leader       UserJSON     `json:"leader"`
	Participants []UserJSON   `json:"participants"`
	Tag          bool         `json:"tag"` //create、update
	TagSet       []TagJson    `json:"tagSet"`
	Modules      []ModuleJSON `json:"modules"` //仅拉取更新
}

type TagJson struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/16 15:53
	*/
	ID  uint `json:"id"`
	Tag bool `json:"tag"`
}

type FramePaceJSON struct {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/24 0:32
	*/
	ID        uint              `json:"id"`
	Name      string            `json:"name"`
	StartTime string            `json:"startTime"`
	EndTime   string            `json:"endTime"`
	Leader    UserJSON          `json:"leader"`
	Modules   []ModuleBriefJSON `json:"modules"` //仅拉取更新

}

func tagSet2TagsJson(tagSet string) (tags []TagJson) {
	temp := strings.Split(tagSet, ",")
	for _, v := range temp {
		IdTag := strings.Split(v, "+")
		if len(IdTag) == 2 {
			id, _ := strconv.Atoi(IdTag[0])
			idU := uint(id)
			t, _ := strconv.ParseBool(IdTag[1])
			tags = append(tags, TagJson{ID: idU, Tag: t})
		}
	}
	return
}

func tagsJson2TagSet(tags []TagJson) (tag bool, tagSet string) {
	/**
	@Author: PantaZheng
	@Description:TagJson转换为db中user表中的TagSet
	@Date: 2019/5/6 23:14
	*/
	l := len(tags)
	count := 0
	if l > 0 {
		for i, v := range tags {
			id := strconv.FormatUint(uint64(v.ID), 10)
			t := strconv.FormatBool(v.Tag)
			if i == 0 {
				tagSet += id + "+" + t
			} else {
				tagSet += "," + id + "+" + t
			}
			if v.Tag == true {
				count++
			}
		}
	}
	if count == l {
		tag = true
	}
	return
}

func project2ProjectJson(project *models.Project) (projectJSON ProjectJSON) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:53
	*/
	projectJSON.ID = project.ID
	projectJSON.Name = project.Name
	projectJSON.Type = project.Type
	projectJSON.CreatorID = project.CreatorID
	creator := user2UserJSON(project.Creator)
	projectJSON.Creator = userJSON2UserBriefJSON(creator)
	projectJSON.CreateTime = project.CreateTime
	projectJSON.StartTime = project.StartTime
	projectJSON.EndTime = project.EndTime
	projectJSON.Content = project.Content
	projectJSON.Target = project.Target
	projectJSON.LeaderID = project.LeaderID
	leader := user2UserJSON(project.Leader)
	projectJSON.Leader = userJSON2UserBriefJSON(leader)
	projectJSON.Participants = users2BriefUsersJSON(project.Participants)
	projectJSON.Tag = project.Tag
	projectJSON.TagSet = tagSet2TagsJson(project.TagSet)
	projectJSON.Modules, _ = ModulesFindByProjectID(project.ID)
	return
}

func project2ProjectBriefJson(projectJSON1 *ProjectJSON) (projectJSON2 ProjectJSON) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:53
	*/
	projectJSON2.ID = projectJSON1.ID
	projectJSON2.Name = projectJSON1.Name
	projectJSON2.StartTime = projectJSON1.StartTime
	projectJSON2.EndTime = projectJSON1.EndTime
	projectJSON2.Leader = projectJSON1.Leader
	projectJSON2.Tag = projectJSON1.Tag
	projectJSON2.Content = projectJSON1.Content
	return
}

func projects2ProjectsBriefJSON(projects []models.Project) (projectsJSON []ProjectJSON) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:53
	*/
	projectsJSON = make([]ProjectJSON, len(projects))
	for i, v := range projects {
		p := project2ProjectJson(&v)
		projectsJSON[i] = project2ProjectBriefJson(&p)
	}
	return
}

func (projectJSON *ProjectJSON) projectJSON2Project() (project models.Project) {
	project.ID = projectJSON.ID
	project.Name = projectJSON.Name
	project.Type = projectJSON.Type
	project.CreatorID = projectJSON.CreatorID
	project.CreateTime = projectJSON.CreateTime
	project.StartTime = projectJSON.StartTime
	project.EndTime = projectJSON.EndTime
	project.Content = projectJSON.Content
	project.Target = projectJSON.Target
	project.LeaderID = projectJSON.LeaderID
	project.Participants = usersJSON2Users(projectJSON.Participants)
	project.Tag, project.TagSet = tagsJson2TagSet(projectJSON.TagSet)
	return
}

func (projectJSON *ProjectJSON) Create() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 22:59
	*/
	creator := UserJSON{ID: projectJSON.CreatorID}
	if err = creator.First(); err == nil {
		l := len(projectJSON.Participants)
		projectJSON.TagSet = make([]TagJson, l)
		for i := 0; i < l; i++ {
			projectJSON.TagSet[i] = TagJson{ID: projectJSON.Participants[i].ID, Tag: false}
		}
		p := projectJSON.projectJSON2Project()
		if err = p.Create(); err == nil {
			*projectJSON = project2ProjectJson(&p)
		}
	}
	if err != nil {
		err = errors.New(titleProject + "Insert:\t" + err.Error())
	}
	return
}

func (projectJSON *ProjectJSON) First() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:03
	*/
	p := projectJSON.projectJSON2Project()
	if err = p.First(); err == nil {
		*projectJSON = project2ProjectJson(&p)
	} else {
		err = errors.New(titleProject + "First:\t" + err.Error())
	}
	return
}

func ProjectFindByID(id uint) (projectJSON ProjectJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:06
	*/
	projectJSON = ProjectJSON{ID: id}
	err = projectJSON.First()
	return
}

func ProjectFramePaceByID(id uint) (framePaceJSON FramePaceJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/17 10:51
	*/
	projectJSON := &ProjectJSON{ID: id}
	if err = projectJSON.First(); err == nil {
		framePaceJSON.ID = projectJSON.ID
		framePaceJSON.Name = projectJSON.Name
		framePaceJSON.StartTime = projectJSON.StartTime
		framePaceJSON.EndTime = projectJSON.EndTime
		framePaceJSON.Leader = projectJSON.Leader
		modules, _ := ModulesFindByProjectID(projectJSON.ID)
		l := len(modules)
		module := ModuleJSON{}
		framePaceJSON.Modules = make([]ModuleBriefJSON, l)
		for i := 0; i < l; i++ {
			module = modules[i]
			_ = module.First()
			framePaceJSON.Modules[i].ID = module.ID
			framePaceJSON.Modules[i].Name = module.Name
			framePaceJSON.Modules[i].StartTime = module.StartTime
			framePaceJSON.Modules[i].EndTime = module.EndTime
			framePaceJSON.Modules[i].Leader = module.Leader
			missions, _ := MissionsFindByModuleID(module.ID)
			m := len(missions)
			mission := MissionJSON{}
			framePaceJSON.Modules[i].Missions = make([]MissionBriefJSON, m)
			for j := 0; j < m; j++ {
				mission = missions[j]
				_ = mission.First()
				framePaceJSON.Modules[i].Missions[j].ID = mission.ID
				framePaceJSON.Modules[i].Missions[j].Name = mission.Name
				framePaceJSON.Modules[i].Missions[j].StartTime = mission.StartTime
				framePaceJSON.Modules[i].Missions[j].EndTime = mission.EndTime
				framePaceJSON.Modules[i].Missions[j].Participants = mission.Participants
			}
		}
	}
	if err != nil {
		err = errors.New(titleProject + "ProjectFrameByID:\t" + err.Error())
	}
	return
}

func ProjectsFindAll() (projectsJSON []ProjectJSON, err error) {
	if projects, err1 := models.ProjectsFindAll(); err1 == nil {
		projectsJSON = projects2ProjectsBriefJSON(projects)
	} else {
		err = errors.New(titleProject + "ProjectsFindAll:\t" + err1.Error())
	}
	return
}

func ProjectsFindByCreatorID(id uint) (
	projectsJSON []ProjectJSON, err error) {
	if projects, err1 := models.ProjectsFindByCreatorID(id); err1 == nil {
		projectsJSON = projects2ProjectsBriefJSON(projects)
	} else {
		err = errors.New(titleProject + "ProjectsFindByCreatorID:\t" + err1.Error())
	}
	return
}

func ProjectsFindByLeaderID(id uint) (
	projectsJSON []ProjectJSON, err error) {
	if projects, err1 := models.ProjectsFindByLeaderID(id); err1 == nil {
		projectsJSON = projects2ProjectsBriefJSON(projects)
	} else {
		err = errors.New(titleProject + "ProjectsFindByLeaderID:\t" + err1.Error())
	}
	return
}

func ProjectsFindByParticipantID(id uint) (
	projectsJSON []ProjectJSON, err error) {
	if projects, err1 := models.ProjectsFindByParticipantID(id); err1 == nil {
		projectsJSON = projects2ProjectsBriefJSON(projects)
	} else {
		err = errors.New(titleProject + " ProjectsFindByParticipantID:\t" + err1.Error())
	}
	return
}

func (projectJSON *ProjectJSON) Updates() (err error) {
	p := projectJSON.projectJSON2Project()
	l := len(projectJSON.TagSet)
	if l > 1 {
		err = errors.New("更新时，TagSet必须为空或者仅存在一个")
	} else if l == 0 {
		err = p.Updates()
	} else {
		set := projectJSON.TagSet[0]
		if err = p.First(); err == nil {
			tagSet := tagSet2TagsJson(p.TagSet)
			flag := false
			for i := 0; i < len(tagSet); i++ {
				if set.ID == tagSet[i].ID {
					set.Tag = tagSet[i].Tag
					flag = true
					break
				}
			}
			if !flag {
				err = errors.New("tagSet不存在该对象ID")
			}
		}
		err = p.Updates()
	}
	if err != nil {
		err = errors.New(titleProject + "Update\t" + err.Error())
	}
	return
}

func (projectJSON *ProjectJSON) Delete() (err error) {
	p := projectJSON.projectJSON2Project()
	if err = p.Delete(); err == nil {
		*projectJSON = project2ProjectJson(&p)
	} else {
		err = errors.New(titleProject + "Update:\t" + err.Error())
	}
	return
}

func ProjectDeleteByID(id uint) (projectJSON ProjectJSON, err error) {
	projectJSON = ProjectJSON{ID: id}
	err = projectJSON.Delete()
	return
}
