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
	Creator			string
	CreateTime		string
	StartTime		string
	EndTime			string
	Content			string
	File			string
	Tag				bool
	Participants	[]*User `gorm:"many2many:user_missions"`
	ModuleID		uint
	Module			Module
}

type MissionJson struct{
	ID				uint				`json:"id"`
	Name			string				`json:"name"`
	Creator			string				`json:"creator"`
	CreateTime		string				`json:"create_time"`
	StartTime		string				`json:"start_time"`
	EndTime			string				`json:"end_time"`
	Content			string				`json:"content"`
	File			string				`json:"file"`
	Tag				bool				`json:"tag"`
	Participants	[]*UserBriefJson	`json:"participants"`
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
	_, _ =MissionCreate(&MissionJson{Name: "Mission1",ModuleID:1,Participants:[]*UserBriefJson{{ID: 1},
		{ID:2}}})
	_, _ =MissionCreate(&MissionJson{Name: "Mission2",ModuleID:2,Participants:[]*UserBriefJson{{ID: 1},
		{ID:2},{ID:3}}})
}

//缺失participants
func (mission *Mission) missionJson2Mission(missionJson *MissionJson){
	mission.ID=missionJson.ID
	mission.Name=missionJson.Name
	mission.Creator=missionJson.Creator
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
	missionJson.Creator= mission.Creator
	missionJson.CreateTime = mission.CreateTime
	missionJson.StartTime= mission.StartTime
	missionJson.EndTime= mission.EndTime
	missionJson.Content= mission.Content
	missionJson.File= mission.File
	missionJson.Tag= mission.Tag
	missionJson.ModuleID=mission.ModuleID
	participants:=make([]*User,len(mission.Participants))
	database.DB.Model(&mission).Related(&participants,"Participants")
	tempUser:=&UserBriefJson{}
	for _,v:=range participants{
		tempUser.user2UserBriefJson(v)
		missionJson.Participants=append(missionJson.Participants,tempUser)
	}
	return
}

func(missionBriefJson *MissionBriefJson) mission2MissionBriefJSON(mission *Mission){
	missionBriefJson.ID=mission.ID
	missionBriefJson.Name= mission.Name
	missionBriefJson.CreateTime = mission.CreateTime
	missionBriefJson.Content=mission.Content
	missionBriefJson.Tag=mission.Tag
	missionBriefJson.ModuleID=mission.ModuleID
}

func MissionCreate(missionJson *MissionJson) (missionBriefJson MissionBriefJson,err error){
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
		missionBriefJson.mission2MissionBriefJSON(newMission)
	}
	return
}

func MissionFind(mission *Mission)(recordMissionJSON MissionJson, err error){
	recordMission:=new(Mission)
	if err=database.DB.Find(&recordMission,&mission).Error;err==nil{
		log.Printf("recordMission:")
		log.Println(recordMission)
		database.DB.Model(&mission).Related(&mission.Participants,"Participants")
		log.Printf("mission")
		log.Println(mission)
		recordMissionJSON.mission2MissionJSON(mission)
		log.Printf("recordMission:")
		log.Println(recordMission)
	}
	return
}

func MissionsFindByModule(Module *Module)(missionsBriefJson []MissionBriefJson,err error){
	missions:=make([]Mission,1)
	if err=database.DB.Model(&Module).Related(&missions,"ModuleID").Error;err!=nil{
		return
	}
	if len(missions)==0{
		err=errors.New("MissionsFindByModule No Owner Record")
	}else{
		for _,v:=range missions{
			tempJson:=&MissionBriefJson{}
			tempJson.mission2MissionBriefJSON(&v)
			missionsBriefJson=append(missionsBriefJson,*tempJson)
		}
	}
	return
}

func MissionUpdate(missionJson *MissionJson)(missionBriefJson MissionBriefJson,err error){
	updateMission:= new(Mission)
	updateMission.missionJson2Mission(missionJson)
	recordMission:=new(Mission)
	if database.DB.First(&recordMission,&Mission{Name: updateMission.Name}).RecordNotFound(){
		err = errors.New("MISSION UPDATE NOT FOUND RECORD")
	}else{
		database.DB.Model(&recordMission).Updates(updateMission)
		missionBriefJson.mission2MissionBriefJSON(recordMission)
	}
	log.Printf("models.MissionUpdate:"+recordMission.Name)
	return
}

func MissionDelete(mission *Mission)(missionBriefJson MissionBriefJson,err error){
	recordMission:=new(Mission)
	if database.DB.Find(&recordMission,&mission).RecordNotFound(){
		err=errors.New("MISSION DELETE NOT FOUND RECORD")
	}else{
		missionBriefJson.mission2MissionBriefJSON(recordMission)
		err=database.DB.Delete(&recordMission).Error
	}
	log.Printf("models.MissionFind:"+recordMission.Name)
	return
}