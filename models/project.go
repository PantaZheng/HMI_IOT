package models

import (
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"strconv"
	"strings"
)

type Project struct {
	gorm.Model
	Name			string
	Type			string
	Creator			string
	CreateTime		string
	StartTime		string
	EndTime			string
	Content			string
	Target			string
	LeaderID		uint
	Leader			User
	Teachers		[]*User		`gorm:"many2many:user_projects"`
	TagSet			string
	Tag				bool
}

type ProjectJson struct {
	ID				uint				`json:"id"`
	Name			string				`json:"name"`
	Type			string				`json:"type"`
	Creator			string				`json:"creator"`
	CreateTime		string				`json:"create_time"`
	StartTime		string				`json:"start_time"`
	EndTime			string				`json:"end_time"`
	Content			string				`json:"content"`
	Targets			[]string			`json:"targets"`
	LeaderID		uint				`json:"leader"`
	Teachers		[]UserBriefJson		`json:"teachers"`
	Tag				bool				`json:"tag"`
	TagSet			[]TagJson			`json:"tags"`
	Modules			[]ModuleBriefJson	`json:"modules"`
}

type BriefProject struct {
	ID				uint	`json:"id"`
	Name			string	`json:"name"`
	StartTime		string	`json:"startTime"`
	EndTime			string	`json:"endTime"`
	LeaderID		string	`json:"leader"`
	Tag				string	`json:"tag"`
	Content			string	`json:"content"`
}

type TagJson struct{
	ID	uint	`json:"id"`
	Tag	bool	`json:"tag"`
}

func target2TargetsJson (target string) []string{
		return strings.Split(target,",")
}

func targetsJson2Target(targets []string) (target string){
	l:=len(targets)
	if l>0 {
		for i, v := range targets {
			if i == 0 {
				target += v
			} else {
				target+=","+v
			}
		}
	}
	return
}

func tagSet2TagsJson(tagSet string) (tags []TagJson){
	temp:=strings.Split(tagSet,",")
	for _,v:=range temp{
		IdTag :=strings.Split(v,"+")
		if len(IdTag)==2 {
			id,_:=strconv.Atoi(IdTag[0])
			idU:=uint(id)
			t,_:=strconv.ParseBool(IdTag[1])
			tags=append(tags,*&TagJson{ID: idU,Tag:t} )
		}
	}
	return
}

func tagsJson2TagSet(tags []TagJson)(tag bool,tagSet string){
	l:=len(tags)
	count:=0
	if l>0 {
		for i, v := range tags {
			id:=strconv.FormatUint(uint64(v.ID),10)
			t:=strconv.FormatBool(v.Tag)
			if i == 0 {
				tagSet += id+"+"+t
			} else {
				tagSet += ","+id+"+"+t
			}
			if v.Tag==true{
				count++
			}
		}
	}
	if count==l{
		tag=true
	}
	return
}

func (project *Project) projectJson2Project(projectJson *ProjectJson){
	project.ID=projectJson.ID
	project.Name=projectJson.Name
	project.Type=projectJson.Type
	project.Creator=projectJson.Creator
	project.CreateTime=projectJson.CreateTime
	project.StartTime=projectJson.StartTime
	project.EndTime=projectJson.EndTime
	project.Content=projectJson.Content
	project.Target=targetsJson2Target(projectJson.Targets)
	project.LeaderID=projectJson.LeaderID
	project.Tag,project.TagSet= tagsJson2TagSet(projectJson.TagSet)
}

func (projectJson *ProjectJson) project2ProjectJson(project *Project){
	projectJson.ID=project.ID
	projectJson.Name=project.Name
	projectJson.Type=project.Type
	projectJson.Creator=project.Creator
	projectJson.CreateTime=project.CreateTime
	projectJson.StartTime=project.StartTime
	projectJson.EndTime=project.EndTime
	projectJson.Content=project.Content
	projectJson.Targets=target2TargetsJson(project.Target)
	projectJson.LeaderID=project.LeaderID
	projectJson.Tag=project.Tag
	projectJson.TagSet=tagSet2TagsJson(project.TagSet)
	projectJson.Modules,_=ModulesFindByProject(project)
	teachers:=make([]*User,len(project.Teachers))
	database.DB.Model(&project).Related(&teachers,"Teachers")
	tempUser:=&UserBriefJson{}
	for _,v:=range teachers{
		tempUser.user2UserBriefJson(v)
		projectJson.Teachers=append(projectJson.Teachers,*tempUser)
	}
}