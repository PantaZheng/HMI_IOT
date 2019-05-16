package test

import (
	"github.com/pantazheng/bci/models"
	"log"
	"testing"
)

func TestProject(t *testing.T) {
	project := projectCreate()
	project.Name += "Update"
	projectUpdate(&project)
	projectFind()
	projectsFindByLeader()
	projectsFindByParticipant()
	projectDelete()
	projectFind()
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

//func tagsTransfer(){
//	log.Println("tag")
//	log.Println(models.TagSet2Tags(""))
//	log.Println(models.TagSet2Tags("1,2,3"))
//	log.Println(models.TagSet2Tags("1+true,2+false,3+true"))
//	log.Println(models.Tags2TagSet([]models.TagJson{}))
//	log.Println(models.Tags2TagSet([]models.TagJson{{10,false}}))
//	log.Println(models.Tags2TagSet([]models.TagJson{{11,true},{12,false},{13,true}}))
//}

func projectCreate() (projectJson models.ProjectJson) {
	log.Println("projectCreate")
	project := new(models.ProjectJson)
	project.Name = "project_test"
	project.Participants = []models.UserBriefJSON{{ID: 5}}
	if tmp, err := models.ProjectCreate(project); err != nil {
		log.Println(err)
	} else {
		log.Println(tmp)
		projectJson = tmp
	}
	return
}

func projectUpdate(projectJson *models.ProjectJson) {
	log.Println("projectUpdate")
	if p, err := models.ProjectUpdate(projectJson); err != nil {
		log.Println(err)
	} else {
		log.Println(p)
	}
	return
}

func projectFind() {
	log.Println("projectFind")
	p := new(models.Project)
	for i := 1; i <= 6; i++ {
		p.ID = uint(i)
		if projectJson, err := models.ProjectFind(p); err != nil {
			log.Println(err)
		} else {
			log.Println(projectJson)
		}
	}
}

func projectsFindByLeader() {
	log.Println("projectsFindByLeader")
	l := new(models.User)
	for i := 1; i <= 3; i++ {
		l.ID = uint(i)
		if ps, err := models.ProjectsFindByLeader(l); err != nil {
			log.Println(err)
		} else {
			log.Println(ps)
		}
	}
}

func projectsFindByParticipant() {
	log.Println("projectsFindByParticipant")
	l := new(models.User)
	for i := 1; i <= 7; i++ {
		l.ID = uint(i)
		if ps, err := models.ProjectsFindByParticipant(l); err != nil {
			log.Println(err)
		} else {
			log.Println(ps)
		}
	}
}

func projectDelete() {
	log.Println("projectDelete")
	u := new(models.Project)
	for i := 1; i <= 3; i++ {
		u.ID = uint(i)
		if projectJson, err := models.ProjectDelete(u); err != nil {
			log.Println(err)
		} else {
			log.Println(projectJson)
		}
	}
}
