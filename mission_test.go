package main

import (
	"github.com/pantazheng/bci/models"
	"testing"
)

func TestMission(t *testing.T) {
	mission:=new(models.MissionJson)
	mission.Name="mission_test"
	mission.StartTime=mission.Name+"StartTime"
	mission.EndTime=mission.Name+"EndTime"
	mission.Content=mission.Name+"Content"
	mission.Participants =[]*models.UserBriefJson{{ID: 1},
		{ID:3},}
	_,_=models.MissionCreate(mission)
	m1:=new(models.Mission)
	m1.ID=1
	_,_=models.MissionFind(m1)
	m1.ID=2
	_,_=models.MissionFind(m1)
	m1.ID=3
	_,_=models.MissionFind(m1)

}