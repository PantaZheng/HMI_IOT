package main

import(
	"fmt"
	"github.com/pantazheng/bci/models"
)

func main(){
	mission:=new(models.MissionJson)
	mission.Name="mission_test"
	mission.StartTime=mission.Name+"StartTime"
	mission.EndTime=mission.Name+"EndTime"
	mission.Content=mission.Name+"Content"
	mission.Users=[]*models.UserBriefJson{{ID:1},
		{ID:2},}
	res,err:=models.MissionCreate(mission)
	fmt.Println(res)
	fmt.Println(err)
}