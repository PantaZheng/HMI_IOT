package main

import (
	"github.com/pantazheng/bci/models"
	"log"
	"testing"
)

func TestGain(t *testing.T) {
	gain:=new(models.GainJson)
	gain.ID=1
	mission:=new(models.Mission)
	mission.ID=2
	if gains,err:=models.GainsFindByMission(mission);err!=nil{
		log.Println(err.Error())
	}else{
		for _,v :=range gains {
			log.Println(v)
		}
	}
	mission.ID=4
	if gains,err:=models.GainsFindByMission(mission);err!=nil{
		log.Println(err.Error())
	}else{
		for _,v :=range gains {
			log.Println(v)
		}
	}
	gain.Name="gainTest1"
	gain.Type= gain.Name+".type"
	gain.File= gain.Name+".file"
	gain.Remark= gain.Name+".remark"
	gain.OwnerID=1
	gain.MissionID=2
	if gainJson,err:=models.GainUpdate(gain);err!=nil{
		log.Println(err.Error())
	}else{
		log.Println(gainJson)
	}
}
