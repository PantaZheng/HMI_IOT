package main

import (
	"github.com/pantazheng/bci/models"
	"log"
	"testing"
)

func TestProject(t *testing.T){
	tagsTransfer()
}

//func targetTransfer(){
//	log.Println("target")
//	t :=models.Target2TargetsJson("")
//	log.Println(t)
//	t1:=models.Target2TargetsJson("目标1,任务2,测试3")
//	log.Println(t1)
//	log.Println(models.TargetsJson2Target([]string{}))
//	log.Println(models.TargetsJson2Target([]string{"m1"}))
//	log.Println(models.TargetsJson2Target([]string{"m1","m2","k3"}))
//}

func tagsTransfer(){
	log.Println("target")
	log.Println(models.TagSet2Tags(""))
	log.Println(models.TagSet2Tags("1,2,3"))
	log.Println(models.TagSet2Tags("1+true,2+false,3+true"))
	log.Println(models.Tags2TagSet([]models.TagJson{}))
	log.Println(models.Tags2TagSet([]models.TagJson{{10,false}}))
	log.Println(models.Tags2TagSet([]models.TagJson{{11,true},{12,false},{13,true}}))
}