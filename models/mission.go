package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/pantazheng/bci/database"
	"log"
	"time"
)

type Mission struct{
	gorm.Model
	Name			string
	CreatorID		uint
	CreateTime		string
	StartTime		string
	EndTime			string
	Content			string
	File			string
	Tag				bool
	Participants	[]*User		`gorm:"many2many:user_missions"`
	ModuleID		uint
	Module			Module
}

type MissionJson struct{
	ID				uint				`json:"id"`
	Name			string				`json:"name"`
	CreatorID		uint				`json:"creator"`
	CreateTime		string				`json:"create_time"`
	StartTime		string				`json:"start_time"`
	EndTime			string				`json:"end_time"`
	Content			string				`json:"content"`
	File			string				`json:"file"`
	Tag				bool				`json:"tag"`
	Gains			[]GainJson			`json:"gains"`
	Participants	[]UserBriefJson		`json:"participants"`
	ModuleID		uint				`json:"module"`
}

type MissionBriefJson struct{
	ID			uint	`json:"id"`
	Name		string	`json:"name"`
	CreateTime	string	`json:"create_time"`
	Content		string	`json:"content"`
	Tag			bool	`json:"tag"`
	ModuleID	uint	`json:"module"`
}

func missionTestData(){
	_, _ =MissionCreate(&MissionJson{Name: "Mission1",ModuleID:1,Participants:[]UserBriefJson{{ID: 1},{ID:2}}})
	_, _ =MissionCreate(&MissionJson{Name: "Mission2",ModuleID:1,Participants:[]UserBriefJson{{ID: 1},{ID:3}}})
	_, _ =MissionCreate(&MissionJson{Name: "Mission3",ModuleID:2,Participants:[]UserBriefJson{{ID: 1},{ID:2},{ID:3}}})
	_, _ =MissionCreate(&MissionJson{Name: "Mission4",ModuleID:2,Participants:[]UserBriefJson{{ID: 1},{ID:2}}})
}

//缺失participants
func (mission *Mission) missionJson2Mission(missionJson *MissionJson){
	mission.ID=missionJson.ID
	mission.Name=missionJson.Name
	mission.CreatorID=missionJson.CreatorID
	mission.CreateTime=missionJson.CreateTime
	mission.StartTime=missionJson.StartTime
	mission.EndTime=missionJson.EndTime
	mission.Content=missionJson.Content
	mission.File=missionJson.File
	mission.Tag=missionJson.Tag
	mission.ModuleID=missionJson.ModuleID
}

func (missionJson *MissionJson) mission2MissionJSON(mission *Mission){
	missionJson.ID=mission.ID
	missionJson.Name= mission.Name
	missionJson.CreatorID= mission.CreatorID
	missionJson.CreateTime = mission.CreateTime
	missionJson.StartTime= mission.StartTime
	missionJson.EndTime= mission.EndTime
	missionJson.Content= mission.Content
	missionJson.File= mission.File
	missionJson.Tag= mission.Tag
	missionJson.ModuleID=mission.ModuleID
	log.Println(mission)
	participants:=make([]*User,len(mission.Participants))
	database.DB.Model(&mission).Related(&participants,"Participants")
	tempUser:=&UserBriefJson{}
	for _,v:=range participants{
		tempUser.user2UserBriefJson(v)
		missionJson.Participants=append(missionJson.Participants,*tempUser)
	}
	missionJson.Gains,_=GainsFindByMission(mission)
	log.Println(missionJson)
}

func(missionBriefJson *MissionBriefJson) mission2MissionBriefJson(mission *Mission){
	missionBriefJson.ID=mission.ID
	missionBriefJson.Name= mission.Name
	missionBriefJson.CreateTime = mission.CreateTime
	missionBriefJson.Content=mission.Content
	missionBriefJson.Tag=mission.Tag
	missionBriefJson.ModuleID=mission.ModuleID
}

func MissionCreate(missionJson *MissionJson) (recordMissionJson MissionJson,err error){
	newMission := new(Mission)
	newMission.missionJson2Mission(missionJson)
	newMission.CreateTime =time.Now().Format("2006-01-02 15:04:05")
	if err=database.DB.Create(&newMission).Error;err!=nil{
		return
	}
	if err=database.DB.Model(&newMission).First(&newMission).Error;err==nil{
		users:=make([]User,len(missionJson.Participants))
		for i,v:=range missionJson.Participants{
			users[i].ID=v.ID
		}
		err=database.DB.Model(&newMission).Association("Participants").Append(users).Error
		recordMissionJson.mission2MissionJSON(newMission)
	}
	return
}

func MissionFind(mission *Mission)(recordMissionJson MissionJson, err error){
	recordMission:=new(Mission)
	if err=database.DB.First(&recordMission,&mission).Error;err==nil{
		recordMissionJson.mission2MissionJSON(recordMission)
	}
	return
}

func MissionsFindByModule(module *Module)(missionsBriefJson []MissionBriefJson,err error){
	missions:=make([]Mission,1)
	if err=database.DB.Model(&module).Related(&missions,"ModuleID").Error;err!=nil{
		return
	}
	if len(missions)==0{
		err=errors.New("MissionsFindByModule No Mission Record")
	}else{
		for _,v:=range missions{
			tempJson:=&MissionBriefJson{}
			tempJson.mission2MissionBriefJson(&v)
			missionsBriefJson=append(missionsBriefJson,*tempJson)
		}
	}
	return
}

func MissionUpdate(missionJson *MissionJson)(recordMissionJson MissionJson,err error){
	updateMission:= new(Mission)
	updateMission.missionJson2Mission(missionJson)
	recordMission:=new(Mission)
	recordMission.ID=updateMission.ID
	if database.DB.First(&recordMission,&recordMission).RecordNotFound(){
		err = errors.New("MissionUpdate No Mission Record")
	}else{
		database.DB.Model(&recordMission).Updates(updateMission)
		if num:=len(missionJson.Participants);num!=0{
			users:=make([]User,num)
			for i,v:=range missionJson.Participants{
				users[i].ID=v.ID
			}
			err=database.DB.Model(&recordMission).Association("Participants").Replace(users).Error
		}
		recordMissionJson.mission2MissionJSON(recordMission)
	}
	return
}

func MissionDelete(mission *Mission)(recordMissionJson MissionJson,err error){
	recordMission:=new(Mission)
	if database.DB.Find(&recordMission,&mission).RecordNotFound(){
		err=errors.New("MissionDelete No Mission Record")
	}else{
		recordMissionJson.mission2MissionJSON(recordMission)
		err=database.DB.Delete(&recordMission).Error
	}
	return
}