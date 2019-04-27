package models

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/core/errors"
	"github.com/pantazheng/bci/database"
	"log"
	"time"
)

type Gain struct{
	gorm.Model
	Name		string
	Type		string
	File		string
	UpTime		string
	Remark		string
	OwnerID		uint
	Owner		User
	MissionID	uint
	Mission		Mission
}

type GainJson struct {
	ID			uint	`json:"id"`
	Name		string	`json:"name"`		//新建必须
	Type		string	`json:"type"`		//新建类型
	File		string	`json:"file"`
	UpTime		string	`json:"up_time"`	//新建自动生成
	Remark		string	`json:"remark"`
	OwnerID		uint	`json:"owner_id"`
	MissionID	uint	`json:"mission_id"`
}

//name,type,file,remark,owner_id,mission_id
func (gain *Gain) gainJson2Gain(gainJson *GainJson){
	gain.Name=gainJson.Name
	gain.Type=gainJson.Type
	gain.File=gainJson.File
	gain.Remark=gainJson.Remark
	gain.OwnerID=gainJson.OwnerID
	gain.MissionID=gainJson.MissionID
	return
}

//id,name,type,file,up_time,remark,owner_id,mission_id
func (gainJson *GainJson) gain2GainJson(gain *Gain){
	gainJson.ID=gain.ID
	gainJson.Name=gain.Name
	gainJson.Type=gain.Type
	gainJson.File=gain.File
	gainJson.UpTime=gain.UpTime
	gainJson.Remark=gain.Remark
	gainJson.OwnerID=gain.OwnerID
	gainJson.MissionID=gain.MissionID
}


func GainCreate(gainJson *GainJson) (recordGainJson GainJson,err error){
	//构建Gain
	newGain:=new(Gain)
	newGain.gainJson2Gain(gainJson)
	newGain.UpTime=time.Now().Format("2006-01-02 15:04:05")
	log.Println(newGain.UpTime)
	//db新建
	if err=database.DB.Create(&newGain).Error;err!=nil{
		return
	}
	if err=database.DB.Model(&newGain).First(&newGain).Error;err!=nil{
	}else{
		recordGainJson.gain2GainJson(newGain)
	}
	return
}

func GainFindByID(gain *Gain)(recordGainJson GainJson,err error){
	recordGain:=new(Gain)
	if database.DB.Find(&recordGain,&gain).RecordNotFound() {
		err = errors.New("GAIN FIND BY ID NOT FOUND RECORD")
	}else{
		recordGainJson.gain2GainJson(recordGain)
	}
	return
}

func GainsFindByOwner(owner *User)(gainsJson []GainJson,err error){
	gains:=make([]Gain,1)
	if database.DB.Model(&owner).Related(&gains,"OwnerID").RecordNotFound(){
		err = errors.New("GAINS FIND BY OWNER NOT FOUND RECORD")
	}else{
		for _,v :=range gains{
			tempJson:=&GainJson{}
			tempJson.gain2GainJson(&v)
			gainsJson=append(gainsJson,*tempJson)
		}
	}
	return
}

func GainsFindByMission(mission *Mission)(gainsJson []GainJson,err error){
	gains:=make([]Gain,1)
	if database.DB.Model(&mission).Related(&gains,"MissionID").RecordNotFound(){
		err = errors.New("GAINS FIND BY OWNER NOT FOUND RECORD")
	}else{
		for _,v :=range gains{
			tempJson:=&GainJson{}
			tempJson.gain2GainJson(&v)
			gainsJson=append(gainsJson,*tempJson)
		}
	}
	return
}