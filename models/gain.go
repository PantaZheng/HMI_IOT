package models

import (
    "errors"
    "github.com/jinzhu/gorm"
    "github.com/pantazheng/bci/database"
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
    Owner		*User
    MissionID	uint
    Mission		*Mission
}

type GainJson struct {
    ID          uint	        `json:"id"`
    Name		string	        `json:"name"`
    Type		string	        `json:"type"`
    File		string	        `json:"file"`
    UpTime		string	        `json:"up_time"`
    Remark		string	        `json:"remark"`
    Owner		*UserBriefJson	`json:"owner"`
    MissionID	uint	        `json:"mission_id"`
}

func (gain *Gain) gainJson2Gain(gainJson *GainJson){
    gain.ID=gainJson.ID
    gain.Name=gainJson.Name
    gain.Type=gainJson.Type
    gain.File=gainJson.File
    gain.UpTime=gainJson.UpTime
    gain.Remark=gainJson.Remark
    gain.OwnerID=gainJson.Owner.ID
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
    gainJson.Owner.User2UserBriefJson(gain.Owner)
    gainJson.MissionID=gain.MissionID
}

func gainTestData(){
    _,_=GainCreate(&GainJson{Name:"gain1",Owner:&UserBriefJson{ID:1},MissionID:1})
    _,_=GainCreate(&GainJson{Name:"gain2",Owner:&UserBriefJson{ID:2},MissionID:1})
    _,_=GainCreate(&GainJson{Name:"gain3",Owner:&UserBriefJson{ID:1},MissionID:2})
    _,_=GainCreate(&GainJson{Name:"gain4",Owner:&UserBriefJson{ID:2},MissionID:2})
}

func GainCreate(gainJson *GainJson) (recordGainJson GainJson,err error){
    //构建Gain
    newGain:=new(Gain)
    newGain.gainJson2Gain(gainJson)
    newGain.UpTime=time.Now().Format("2006-01-02 15:04:05")
    if err=database.DB.Create(&newGain).Error;err!=nil{
        return
    }
    if err=database.DB.Model(&newGain).First(&newGain).Error;err==nil{
        recordGainJson.gain2GainJson(newGain)
    }
    return
}

func GainFind(gain *Gain)(recordGainJson GainJson,err error){
    recordGain:=new(Gain)
    if err=database.DB.Find(&recordGain,&gain).Error;err==nil {
        recordGainJson.gain2GainJson(recordGain)
    }
    return
}

func GainsFindByOwner(owner *User)(gainsJson []*GainJson,err error){
    gains:=make([]Gain,1)
    if err=database.DB.Model(&owner).Related(&gains,"OwnerID").Error;err!=nil{
        return
    }
    if len(gains)==0{
        err=errors.New("GainsFindByOwner No Gain Record")
    }else {
        for _,v :=range gains{
            tempJson:=&GainJson{}
            tempJson.gain2GainJson(&v)
            gainsJson=append(gainsJson,tempJson)
        }
    }
    return
}

func GainsFindByMission(mission *Mission)(gainsJson []*GainJson,err error){
    gains:=make([]Gain,1)
    if err=database.DB.Model(&mission).Related(&gains,"MissionID").Error;err!=nil{
        return
    }
    if len(gains)==0{
        err=errors.New("GainsFindByMission No Gain Record")
    }else{
        for _,v :=range gains{
            tempJson:=&GainJson{}
            tempJson.gain2GainJson(&v)
            gainsJson=append(gainsJson,tempJson)
        }
    }
    return
}

//TODO:更新逻辑添加
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

//TODO:删除逻辑添加
func GainDelete(gainJson *GainJson) (recordGainJson GainJson,err error) {
    /**
     * @field:
     * @filename: gain.go
     * @param: 
     * @return: GainJson{}, nil
     * @author: panta
     * @date: 2019/5/6 21:27
     */
    recordGain:=new(Gain)
    recordGain.ID=gainJson.ID
    if err=database.DB.First(&recordGain).Error;err==nil{
         recordGainJson.gain2GainJson(recordGain)
         err=database.DB.Unscoped().Delete(&recordGain).Error
    }
    return
}