package main

import (
	"fmt"
	"github.com/pantazheng/bci/models"
	"testing"
)

func TestGain(t *testing.T) {
	gain:=new(models.GainJson)
	gain.Name="gainTest1"
	gain.Type= gain.Name+".type"
	gain.File= gain.Name+".file"
	gain.Remark= gain.Name+".remark"
	gain.OwnerID=1
	gain.MissionID=2


	_,_=models.GainCreate(gain)
	gain.Name="gainTest2"
	gain.OwnerID=1
	gain.MissionID=2
	_,_=models.GainCreate(gain)

	user:=new(models.User)
	user.ID=1
	gains,_:=models.GainsFindByMission(user)
	for _,v :=range gains {
		fmt.Println(v)
	}
}
