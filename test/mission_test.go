package test

import (
	"github.com/pantazheng/bci/models"
	"log"
	"testing"
)

func TestMission(t *testing.T) {
	mission := new(models.MissionJson)
	mission.Name = "mission_test"
	mission.StartTime = mission.Name + "StartTime"
	mission.EndTime = mission.Name + "EndTime"
	mission.Content = mission.Name + "Content"
	mission.Participants = []models.UserBriefJSON{{ID: 1},
		{ID: 3}}
	res1, _ := models.MissionCreate(mission)
	log.Println(res1)
	m := new(models.Mission)
	m.ID = res1.ID
	res2, err := models.MissionDelete(m)
	log.Println(res2)
	log.Println(err)
	m.ID = 10
	res3, er3 := models.MissionDelete(m)
	log.Println(res3)
	log.Println(er3)
}
