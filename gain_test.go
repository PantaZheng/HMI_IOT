package main

import (
	"fmt"
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
	fmt.Println(res)
	fmt.Println(err)
}
