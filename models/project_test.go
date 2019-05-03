package models

import (
	"github.com/pantazheng/bci/models"
	"log"
	"testing"
)

func TestProject(t *testing.T){
	println("target")
}

func targetTransfer(){
	println(models.Target2TargetsJson(""))
	println(models.Target2TargetsJson("目标1,任务2，测试3"))
	println(models.TargetsJson2Target([]string{}))
	println(models.TargetsJson2Target([]string{"m1"}))
	println(models.TargetsJson2Target([]string{"m1","m2","k3"}))
}