package main

import (
	"github.com/pantazheng/bci/models"
	"log"
	"testing"
)

func TestMission(t *testing.T) {
	mission:=new(models.MissionJson)
	mission.Name="mission_test"
	mission.StartTime=mission.Name+"StartTime"
	mission.EndTime=mission.Name+"EndTime"
	mission.Content=mission.Name+"Content"
	mission.Participants =[]models.UserBriefJson{{ID: 1},
		{ID:3},}
	_,_=models.MissionCreate(mission)
	module:=new(models.Module)
	module.ID=1
	r1,_:=models.MissionsFindByModule(module)
	log.Println(r1)
	module.ID=2
	r2,_:=models.MissionsFindByModule(module)
	log.Println(r2)
	module.ID=3
	r3,_:=models.MissionsFindByModule(module)
	log.Println(r3)
}