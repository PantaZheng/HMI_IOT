package test

import (
	"github.com/pantazheng/bci/models"
	"log"
	"strconv"
	"testing"
)

func TestGain(t *testing.T) {
	gain := new(models.GainJson)
	gain.ID = 1
	g := &models.Gain{}
	g.ID = 1
	if a, err := models.GainFind(g); err != nil {
		log.Println(err)
	} else {
		log.Println(a)
	}
	mission := new(models.Mission)
	mission.ID = 2
	if gains, err := models.GainsFindByMission(mission); err != nil {
		log.Println(err.Error())
	} else {
		for _, v := range gains {
			log.Println(v)
		}
	}
	mission.ID = 4
	if gains, err := models.GainsFindByMission(mission); err != nil {
		log.Println(err.Error())
	} else {
		for _, v := range gains {
			log.Println(v)
		}
	}
	gain.Name = "gainTest" + strconv.Itoa(int(gain.ID))
	gain.Type = gain.Name + ".type"
	gain.File = gain.Name + ".file"
	gain.Remark = gain.Name + ".remark"
	gain.Owner.ID = 1
	gain.MissionID = 2
	if gainJson, err := models.GainUpdate(gain); err != nil {
		log.Println(err.Error())
	} else {
		log.Println(gainJson)
	}
	gain.ID = 4
	gain.Name = "gainTest" + strconv.Itoa(int(gain.ID))
	gain.Type = gain.Name + ".type"
	gain.File = gain.Name + ".file"
	gain.Remark = gain.Name + ".remark"
	gain.OwnerID = 1
	gain.MissionID = 2
	if gainJson, err := models.GainUpdate(gain); err != nil {
		log.Println(err.Error())
	} else {
		log.Println(gainJson)
	}
	if gainJson, err := models.GainDelete(gain); err != nil {
		log.Println(err.Error())
	} else {
		log.Println(gainJson)
	}
}
