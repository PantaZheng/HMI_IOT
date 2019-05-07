package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"strconv"
	"strings"
	"time"
)

type Project struct {
	gorm.Model
	Name			string			`gorm:"unique"`
	Type			string
	CreatorID		uint
	Creator			*User
	CreateTime		string
	StartTime		string
	EndTime			string
	Content			string
	Target			string
	LeaderID		uint
	Leader			*User
	Participants	[]*User		`gorm:"many2many:user_projects"`
	TagSet			string
	Tag				bool
}

type ProjectJson struct {
	ID				uint				`json:"id"`
	Name			string				`json:"name"`
	Type			string				`json:"type"`
	Creator			*UserBriefJson		`json:"creator"`
	CreateTime		string				`json:"create_time"`
	StartTime		string				`json:"start_time"`
	EndTime			string				`json:"end_time"`
	Content			string				`json:"content"`
	Targets			[]string			`json:"targets"`
	Leader			*UserBriefJson		`json:"leader"`
	Participants	[]*UserBriefJson	`json:"participants"`
	Tag				bool				`json:"tag"`		//create、update
	TagSet			[]*TagJson			`json:"tags"`
	Modules			[]*ModuleBriefJson	`json:"modules"`	//仅拉取更新
}

type ProjectBriefJson struct {
	ID				uint			`json:"id"`
	Name			string			`json:"name"`
	StartTime		string			`json:"startTime"`
	EndTime			string			`json:"endTime"`
	Leader			*UserBriefJson	`json:"leader"`
	Tag				bool			`json:"tag"`
	Content			string			`json:"content"`
}

type TagJson struct{
	ID	uint	`json:"id"`
	Tag	bool	`json:"tag"`
}

func projectTestData() {
	leader1:=&UserBriefJson{ID:2}
	leader2:=&UserBriefJson{ID:3}
	leader3:=&UserBriefJson{ID:4}
	leader4:=&UserBriefJson{ID:5}
	_,_=ProjectCreate(&ProjectJson{Name:"Project1",Targets:[]string{"t1"},Leader:leader1,Participants:[]*UserBriefJson{{ID:2}},TagSet:[]*TagJson{{ID:2,Tag:true},{ID:3,Tag:false}}})
	_,_=ProjectCreate(&ProjectJson{Name:"Project2",Targets:[]string{"t1","tt2"},Leader:leader2,Participants:[]*UserBriefJson{{ID:3}},TagSet:[]*TagJson{{ID:2,Tag:true},{ID:3,Tag:true}}})
	_,_=ProjectCreate(&ProjectJson{Name:"Project3",Targets:[]string{"t1","tt2","ttt3"},Leader:leader3,Participants:[]*UserBriefJson{{ID:2}},TagSet:[]*TagJson{{ID:2,Tag:true},{ID:3,Tag:false}}})
	_,_=ProjectCreate(&ProjectJson{Name:"Project4",Targets:[]string{"t1","tt2","ttt3"},Leader:leader4,Participants:[]*UserBriefJson{{ID:2}},TagSet:[]*TagJson{{ID:2,Tag:false},{ID:3,Tag:false}}})
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

func tagSet2TagsJson(tagSet string) (tags []*TagJson){
	temp:=strings.Split(tagSet,",")
	for _,v:=range temp{
		IdTag :=strings.Split(v,"+")
		if len(IdTag)==2 {
			id,_:=strconv.Atoi(IdTag[0])
			idU:=uint(id)
			t,_:=strconv.ParseBool(IdTag[1])
			tags=append(tags,&TagJson{ID: idU,Tag:t} )
		}
	}
	return
}

func tagsJson2TagSet(tags []*TagJson)(tag bool,tagSet string){
	/**
	@Author: PantaZheng
	@Description:TagJson转换为db中user表中的TagSet
	@Date: 2019/5/6 23:14
	*/
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
	project.CreatorID=projectJson.Creator.ID
	project.CreateTime=projectJson.CreateTime
	project.StartTime=projectJson.StartTime
	project.EndTime=projectJson.EndTime
	project.Content=projectJson.Content
	project.Target=targetsJson2Target(projectJson.Targets)
	project.LeaderID=projectJson.Leader.ID
	project.Tag,project.TagSet= tagsJson2TagSet(projectJson.TagSet)
}

func (projectJson *ProjectJson) project2ProjectJson(project *Project){
	projectJson.ID=project.ID
	projectJson.Name=project.Name
	projectJson.Type=project.Type
	projectJson.Creator.User2UserBriefJson(project.Creator)
	projectJson.CreateTime=project.CreateTime
	projectJson.StartTime=project.StartTime
	projectJson.EndTime=project.EndTime
	projectJson.Content=project.Content
	projectJson.Targets=target2TargetsJson(project.Target)
	projectJson.Leader.User2UserBriefJson(project.Leader)
	projectJson.Tag=project.Tag
	projectJson.TagSet=tagSet2TagsJson(project.TagSet)
	projectJson.Modules,_=ModulesFindByProject(project)
	participants :=make([]*User,len(project.Participants))
	database.DB.Model(&project).Related(&participants,"Participants")
	tempUser:=&UserBriefJson{}
	for _,v:=range participants {
		tempUser.User2UserBriefJson(v)
		projectJson.Participants=append(projectJson.Participants,tempUser)
	}
}

func (projectBriefJson *ProjectBriefJson)project2ProjectBriefJson(project *Project){
	projectBriefJson.ID=project.ID
	projectBriefJson.Name=project.Name
	projectBriefJson.StartTime=project.StartTime
	projectBriefJson.EndTime=project.EndTime
	projectBriefJson.Leader.User2UserBriefJson(project.Leader)
	projectBriefJson.Tag=project.Tag
	projectBriefJson.Content=project.Content
}

func ProjectCreate(projectJson *ProjectJson)(recordProjectJson ProjectJson,err error){
	newProject := new(Project)
	newProject.projectJson2Project(projectJson)
	newProject.CreateTime=time.Now().Format("2006-01-02 15:04:05")
	if err=database.DB.Create(&newProject).Error;err!=nil{
		return
	}
	if err=database.DB.Model(&newProject).First(&newProject).Error;err==nil{
		participants :=make([]User,len(projectJson.Participants))
		for i,v:=range projectJson.Participants{
			participants[i].ID=v.ID
		}
		err=database.DB.Model(&newProject).Association("Participants").Append(participants).Error
		recordProjectJson.project2ProjectJson(newProject)
	}
	return
}

func ProjectFind(project *Project)(recordProjectJson ProjectJson,err error){
	recordProject:=new(Project)
	if err=database.DB.First(&recordProject,&project).Error;err==nil{
		recordProjectJson.project2ProjectJson(recordProject)
	}
	return
}

func ProjectsFindByLeader(leader *User)(projectsBriefJson []ProjectBriefJson,err error){
	projects := make([]Project,1)
	if err=database.DB.Model(&leader).Related(&projects,"LeaderID").Error;err!=nil{
		return
	}
	if len(projects)==0{
		err=errors.New("ProjectsFindByLeader No Project Record")
	}else{
		for _,v :=range  projects{
			tempJson:=&ProjectBriefJson{}
			tempJson.project2ProjectBriefJson(&v)
			projectsBriefJson=append(projectsBriefJson,*tempJson)
		}
	}
	return
}

func ProjectsFindByParticipant(participant *User)(projectsBriefJson []ProjectBriefJson,err error){
	projects := make([]Project,1)
	if err=database.DB.Model(&participant).Related(&projects,"PProjects").Error;err!=nil{
		return
	}
	if len(projects)==0{
		err=errors.New("ProjectsFindByParticipants No Project Record")
	}else{
		for _,v :=range  projects{
			tempJson:=&ProjectBriefJson{}
			tempJson.project2ProjectBriefJson(&v)
			projectsBriefJson=append(projectsBriefJson,*tempJson)
		}
	}
	return
}

func ProjectUpdate(projectJson *ProjectJson)(recordProjectJson ProjectJson,err error){
	updateProject := new(Project)
	updateProject.projectJson2Project(projectJson)
	recordProject := new(Project)
	recordProject.ID=updateProject.ID
	if database.DB.First(&recordProject,&recordProject).RecordNotFound(){
		err = errors.New("ProjectUpdate No Module Record")
	}else{
		database.DB.Model(&recordProject).Updates(updateProject)
		if num:=len(projectJson.Participants);num!=0{
			users:=make([]User,num)
			for i,v:=range projectJson.Participants{
				users[i].ID=v.ID
			}
			err=database.DB.Model(&recordProject).Association("Participants").Replace(users).Error
		}
		recordProjectJson.project2ProjectJson(recordProject)
	}
	return
}

func ProjectDelete(project *Project)(recordProjectJson ProjectJson,err error){
	recordProject := new (Project)
	if database.DB.Find(&recordProject,&project).RecordNotFound(){
		err=errors.New("ProjectDelete No Module Record")
	}else{
		recordProjectJson.project2ProjectJson(recordProject)
		err=database.DB.Unscoped().Delete(&recordProject).Error
	}
	return
}