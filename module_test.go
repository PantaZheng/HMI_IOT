package main

import (
	"github.com/pantazheng/bci/models"
	"log"
	"testing"
)

func TestModule(t *testing.T){
	log.Println("moduleCreate")
	module:=moduleCreate()
	module.Name=module.Name+"Update测试"
	log.Println("moduleUpdate")
	moduleUpdate(&module)
	mod:=new(models.Module)
	log.Println("moduleFind")
	for i:=1;i<6;i++{
		mod.ID=uint(i)
		moduleFind(mod)
	}
	log.Println("ModulesFindByLeader")
	modulesFindByLeader()
	log.Println("modulesFindByProject")
	modulesFindByProject()
	log.Println("moduleDelete")
	for i:=1;i<3;i++{
		mod.ID=uint(i)
		moduleDelete(mod)
	}
	log.Println("moduleFind")
	for i:=1;i<6;i++{
		mod.ID=uint(i)
		moduleFind(mod)
	}
}


func moduleCreate()(moduleJson models.ModuleJson){
	module:=new(models.ModuleJson)
	module.Name="module_test5"
	module.Participants=[]models.UserBriefJson{{ID:2},{ID:6}}
	tmp,err:=models.ModuleCreate(module)
	log.Println(moduleJson)
	log.Println(err)
	moduleJson=tmp
	return
}

func moduleUpdate(module *models.ModuleJson){
	moduleJson,err:=models.ModuleUpdate(module)
	log.Println(moduleJson)
	log.Println(err)
	return
}

func moduleFind(module *models.Module){
	moduleJson,err:=models.ModuleFind(module)
	log.Println(moduleJson)
	log.Println(err)
}

func moduleDelete(module *models.Module){
	moduleJson,err:=models.ModuleDelete(module)
	log.Println(moduleJson)
	log.Println(err)
}

func modulesFindByLeader(){
	leader:=new(models.User)
	for i:=1;i<=3;i++{
		leader.ID=uint(i)
		modules,err:=models.ModulesFindByLeader(leader)
		log.Println(modules)
		log.Println(err)
	}
}

func modulesFindByProject(){
	p:=new(models.Project)
	for i:=1;i<=3;i++{
		p.ID=uint(i)
		modules,err:=models.ModulesFindByProject(p)
		log.Println(modules)
		log.Println(err)
	}
}