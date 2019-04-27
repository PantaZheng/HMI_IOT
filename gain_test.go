package main

import (
	"fmt"
	"github.com/pantazheng/bci/database"
	"github.com/pantazheng/bci/models"
	"testing"
)

func TestGain(t *testing.T) {
	gain:=new(models.GainJson)
	gain.Name="gainTest"
	gain.Type=gain.Name+".type"
	gain.File=gain.Name+".file"
	gain.UpTime=gain.Name+".up_time"
	gain.Remark=gain.Name+".remark"
	gain.OwnerID=1
	gain.MissionID=1
	res,err:=models.GainCreate(gain)
	fmt.Print(res)
	fmt.Println()
	fmt.Println(err)
	user:=new(models.User)
	database.DB.Model(&user).Related(&gain)
	fmt.Println(gain)
	fmt.Println(gain.MissionID)
	fmt.Println(gain.Remark)
}
