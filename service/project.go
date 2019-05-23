package service

import (
	"errors"
	"github.com/pantazheng/bci/models"
	"log"
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
	Targets      []string     `json:"targets"`
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

func projectTestData() {
	log.Println("projectTestData")
	u1 := UserJSON{ID: 2}
	u2 := UserJSON{ID: 3}
	u3 := UserJSON{ID: 4}
	u4 := UserJSON{ID: 5}
	u5 := UserJSON{ID: 6}
	u6 := UserJSON{ID: 7}
	u7 := UserJSON{ID: 8}
	projects := make([]ProjectJSON, 4)
	projects[0] = ProjectJSON{Name: "复仇者联盟", CreatorID: 1, Targets: []string{"tag1", "tag2"}, LeaderID: 4, Participants: []UserJSON{u1, u2, u3, u4, u5, u6, u7}}
	projects[1] = ProjectJSON{Name: "复仇者联盟2：奥创纪元", CreatorID: 1, Targets: []string{"tag1", "tag3"}, LeaderID: 5, Participants: []UserJSON{u2, u3, u4, u5, u6, u7}}
	projects[2] = ProjectJSON{Name: "复仇者联盟3：无限战争", CreatorID: 1, Targets: []string{"tag3", "tag4"}, LeaderID: 6, Participants: []UserJSON{u1, u2, u3, u4, u5, u6, u7}}
	projects[3] = ProjectJSON{Name: "复仇者联盟4：终局之战", CreatorID: 2, Targets: []string{"tag8", "tag9"}, LeaderID: 7, Participants: []UserJSON{u1, u2, u3, u4, u5, u6, u7}}
	for _, v := range projects {
		if err := v.Create(); err != nil {
			log.Println(err.Error())
		} else {
			log.Println(v)
		}
	}
}

func target2TargetsJson(target string) []string {
	return strings.Split(target, ",")
}

func targetsJson2Target(targets []string) (target string) {
	l := len(targets)
	if l > 0 {
		for i, v := range targets {
			if i == 0 {
				target += v
			} else {
				target += "," + v
			}
		}
	}
	return
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
	projectJSON.Targets = target2TargetsJson(project.Target)
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
	project.Target = targetsJson2Target(projectJSON.Targets)
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
		err = errors.New(titleProject + "Create:\t" + err.Error())
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
		l := len(projectJSON.Modules)
		framePaceJSON.Modules = make([]ModuleBriefJSON, l)
		for i := 0; i < l; i++ {
			if err = projectJSON.Modules[i].First(); err != nil {
				framePaceJSON.Modules[i].ID = projectJSON.Modules[i].ID
				framePaceJSON.Modules[i].Name = projectJSON.Modules[i].Name
				framePaceJSON.Modules[i].StartTime = projectJSON.Modules[i].StartTime
				framePaceJSON.Modules[i].EndTime = projectJSON.Modules[i].EndTime
				framePaceJSON.Modules[i].Leader = projectJSON.Modules[i].Leader
				m := len(projectJSON.Modules[i].Missions)
				framePaceJSON.Modules[i].Missions = make([]MissionBriefJSON, m)
				for j := 0; j < m; j++ {
					if err = projectJSON.Modules[i].Missions[j].First(); err != nil {
						framePaceJSON.Modules[i].Missions[j].ID = projectJSON.Modules[i].Missions[j].ID
						framePaceJSON.Modules[i].Missions[j].Name = projectJSON.Modules[i].Missions[j].Name
						framePaceJSON.Modules[i].Missions[j].StartTime = projectJSON.Modules[i].Missions[j].StartTime
						framePaceJSON.Modules[i].Missions[j].EndTime = projectJSON.Modules[i].Missions[j].EndTime
						framePaceJSON.Modules[i].Missions[j].Participants = projectJSON.Modules[i].Missions[j].Participants
					} else {
						break
					}
				}
			} else {
				break
			}
		}
	}
	if err != nil {
		err = errors.New(titleProject + "ProjectFrameByID:\t" + err.Error())
	}
	return
}

func ProjectsFindAll() (projectsJSON []ProjectJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:49
	*/
	if projects, err1 := models.ProjectsFindAll(); err1 == nil {
		projectsJSON = projects2ProjectsBriefJSON(projects)
	} else {
		err = errors.New(titleProject + "ProjectsFindAll:\t" + err1.Error())
	}
	return
}

func ProjectsFindByCreatorID(id uint) (
	projectsJSON []ProjectJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:07
	*/
	if projects, err1 := models.ProjectsFindByCreatorID(id); err1 == nil {
		projectsJSON = projects2ProjectsBriefJSON(projects)
	} else {
		err = errors.New(titleProject + "ProjectsFindByCreatorID:\t" + err1.Error())
	}
	return
}

func ProjectsFindByLeaderID(id uint) (
	projectsJSON []ProjectJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:07
	*/
	if projects, err1 := models.ProjectsFindByLeaderID(id); err1 == nil {
		projectsJSON = projects2ProjectsBriefJSON(projects)
	} else {
		err = errors.New(titleProject + "ProjectsFindByLeaderID:\t" + err1.Error())
	}
	return
}

func ProjectsFindByParticipantID(id uint) (
	projectsJSON []ProjectJSON, err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:07
	*/
	if projects, err1 := models.ProjectsFindByParticipantID(id); err1 == nil {
		projectsJSON = projects2ProjectsBriefJSON(projects)
	} else {
		err = errors.New(titleProject + " ProjectsFindByParticipantID:\t" + err1.Error())
	}
	return
}

func (projectJSON *ProjectJSON) Updates() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:15
	*/
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
		err = errors.New(titleProject + "Updates\t" + err.Error())
	}
	return
}

func (projectJSON *ProjectJSON) Delete() (err error) {
	/**
	@Author: PantaZheng
	@Description:
	@Date: 2019/5/15 23:16
	*/
	p := projectJSON.projectJSON2Project()
	if err = p.Delete(); err == nil {
		*projectJSON = project2ProjectJson(&p)
	} else {
		err = errors.New(titleProject + "Updates:\t" + err.Error())
	}
	return
}

func ProjectDeleteByID(id uint) (projectJSON ProjectJSON, err error) {
	projectJSON = ProjectJSON{ID: id}
	err = projectJSON.Delete()
	return
}
