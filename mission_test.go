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
		{ID:3}}
	res1,_:=models.MissionCreate(mission)
	log.Println(res1)
	mission.Name="mission1_test"
	mission.StartTime=mission.Name+"StartTime"
	mission.EndTime=mission.Name+"EndTime"
	mission.Content=mission.Name+"Content"
	mission.Participants =[]models.UserBriefJson{{ID: 1},
		{ID:2}}
	res2,_:=models.MissionCreate(mission)
	log.Println(res2)
}