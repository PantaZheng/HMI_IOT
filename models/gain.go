package models

import (
    "github.com/jinzhu/gorm"
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
    ID      	uint	`json:"id"`
    Name		string	`json:"name"`
    Type		string	`json:"type"`
    File		string	`json:"file"`
    UpTime		string	`json:"up_time"`
    Remark		string	`json:"remark"`
    OwnerID		uint	`json:"owner_id"`
    MissionID	uint	`json:"mission_id"`
}

func (gain *Gain) gainJson2Gain(gainJson *GainJson){
    gain.ID=gainJson.ID
    gain.Name=gainJson.Name
    gain.Type=gainJson.Type
    gain.File=gainJson.File
    gain.UpTime=gainJson.UpTime
    gain.Remark=gainJson.Remark
    gain.OwnerID=gainJson.OwnerID
    gain.MissionID=gainJson.MissionID
    return
}

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

func gainTestData(){
    _,_=GainCreate(&GainJson{Name:"gain1",OwnerID:1,MissionID:1})
    _,_=GainCreate(&GainJson{Name:"gain2",OwnerID:2,MissionID:1})
    _,_=GainCreate(&GainJson{Name:"gain3",OwnerID:1,MissionID:2})
    _,_=GainCreate(&GainJson{Name:"gain4",OwnerID:2,MissionID:2})
    log.Println("gainTestData")
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
    if err=database.DB.Model(&newGain).First(&newGain).Error;err==nil{
        recordGainJson.gain2GainJson(newGain)
    }
    return
}

func GainFindByID(gain *Gain)(recordGainJson GainJson,err error){
    recordGain:=new(Gain)
    if err=database.DB.Find(&recordGain,&gain).Error;err==nil {
        recordGainJson.gain2GainJson(recordGain)
    }
    return
}

func GainsFindByOwner(owner *User)(gainsJson []GainJson,err error){
    gains:=make([]Gain,1)
    if err=database.DB.Model(&owner).Related(&gains,"OwnerID").Error;err==nil{
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
    if err=database.DB.Model(&mission).Related(&gains,"MissionID").Error;err==nil{
        for _,v :=range gains{
            tempJson:=&GainJson{}
            tempJson.gain2GainJson(&v)
            gainsJson=append(gainsJson,*tempJson)
        }
    }
    return
}

func GainUpdate(gainJson *GainJson) (recordGainJson GainJson,err error){
    recordGain:=new(Gain)
    recordGain.ID=gainJson.ID
    if err=database.DB.First(&recordGain).Error;err!=nil{
        return
    }
    newGain:=new(Gain)
    newGain.gainJson2Gain(gainJson)
    newGain.UpTime=time.Now().Format("2006-01-02 15:04:05")
    if err=database.DB.Model(&recordGain).Updates(newGain).Error;err!=nil{
        return
    }
    if err=database.DB.Model(&newGain).First(&newGain).Error;err==nil{
        recordGainJson.gain2GainJson(newGain)
    }
    return
}