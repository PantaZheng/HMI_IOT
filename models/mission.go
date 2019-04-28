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
	Participants	[]*User `gorm:"many2many:user_missions"` //多对多
	ModuleID		uint                                     //归属模块
	Module			Module
}

type MissionJson struct{
	ID				uint             `json:"id"`
	Name			string           `json:"name"`         //名称
	Creator			string           `json:"creator"`      //创建者
	CreateTime		string           `json:"create_time"`  //创建时间
	StartTime		string           `json:"start_time"`   //开始时间
	EndTime			string           `json:"end_time"`     //结束时间
	Content			string           `json:"content"`      //任务详细内容
	File			string           `json:"file"`         //附件
	Tag				bool             `json:"tag"`          //标记
	Participants	[]*UserBriefJson `json:"participants"` //参与人员
	ModuleID     uint             `json:"module"`
}

type MissionBriefJson struct{
	ID         uint   `json:"id"`
	Name       string `json:"name"`       //名称
	CreateTime string `json:"create_time"` //创建时间
	Content    string `json:"content"`    //任务详细内容
	Tag		   string `json:"tag"`
}

func missionTestData(){
	_, _ =MissionCreate(&MissionJson{Name: "Mission1",Content:"Mission1"})
	_, _ =MissionCreate(&MissionJson{Name: "Mission2",Content:"Mission2"})
	log.Println("missionTestData")
}

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
	for _,v:=range mission.Participants {
		missionJson.Participants =append(missionJson.Participants, &UserBriefJson{ID: v.ID,Name:v.Name})
	}
}

func(missionBriefJson *MissionBriefJson) mission2MissionBriefJSON(mission *Mission){
	missionBriefJson.ID=mission.ID
	missionBriefJson.Name= mission.Name
	missionBriefJson.CreateTime = mission.CreateTime
	missionBriefJson.Content=mission.Content
}



func MissionCreate(missionJson *MissionJson) (missionBriefJson MissionBriefJson,err error){
	newMission := new(Mission)
	newMission.Name=missionJson.Name
	//createdTime
	newMission.CreateTime =time.Now().Format("2006-01-02 15:04:05")
	recordMission:=new(Mission)
	if database.DB.First(&recordMission,&Mission{Name: newMission.Name}).RecordNotFound(){
		newMission.CreateTime =time.Now().Format("2006-01-02 15:04:05")
		newMission.Creator=missionJson.Creator
		newMission.Content=missionJson.Content
		newMission.StartTime=missionJson.StartTime
		newMission.EndTime=missionJson.EndTime
		newMission.Content=missionJson.Content
		newMission.File=missionJson.File
		newMission.Tag=missionJson.Tag
		database.DB.Create(&newMission)
		var users []User
		for _,v:=range missionJson.Participants {
			recordUser:=User{}
			recordUser.ID=v.ID
			users=append(users,recordUser)
		}
		database.DB.Model(&newMission).Association("Participants").Append(users)
	}else{
		err=errors.New("MISSION CREATE FOUND RECORD")
	}
	log.Printf("models.MissionCreate:"+missionJson.Name)
	return
}

func MissionFindOne(mission *Mission)(recordMissionJSON MissionJson, err error){
	recordMission:=new(Mission)
	if database.DB.Find(&recordMission,&mission).RecordNotFound(){
		err=errors.New("MISSION FIND NOT FOUND RECORD")
	}else{
		database.DB.Model(&mission).Related(&mission.Participants)
		recordMissionJSON.mission2MissionJSON(recordMission)
	}
	log.Printf("models.MissionFindOne:"+mission.Name)
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
	log.Printf("models.MissionFindOne:"+recordMission.Name)
	return
}