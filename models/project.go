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
	ID 	uint	`json:"id"`
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

func TagSet2Tags(tagSet string) (tags []*TagJson){
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
	return tags
}

func Tags2TagSet(tags []TagJson)(tagSet string){
	l:=len(tags)
	if l>0 {
		for i, v := range tags {
			id:=strconv.FormatUint(uint64(v.ID),10)
			t:=strconv.FormatBool(v.Tag)
			if i == 0 {
				tagSet += id+"+"+t
			} else {
				tagSet += ","+id+"+"+t
			}
		}
	}
	return
}

func GetLeaders(id uint)(leaders []User){
	database.DB.Find(&leaders,id).Select("leaders")
	return
}

func GetInstructors(id uint)(instructors []User){
	database.DB.Find(&instructors,id).Select("instructors")
	return
}

func EnrollProject(project *Project){
	recordProject:=Project{}
	database.DB.FirstOrCreate(&recordProject,&Project{Name:recordProject.Name})

}
