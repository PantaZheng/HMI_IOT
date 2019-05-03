package main

import (
	"github.com/pantazheng/bci/models"
	"testing"
)

func TestProject(t *testing.T){
	println("target")
	targetTransfer()
}

func targetTransfer(){
	t :=models.Target2TargetsJson("")
	println(t)
	t1:=models.Target2TargetsJson("目标1,任务2，测试3")
	println(t1)
	println(models.TargetsJson2Target([]string{}))
	println(models.TargetsJson2Target([]string{"m1"}))
	println(models.TargetsJson2Target([]string{"m1","m2","k3"}))
}