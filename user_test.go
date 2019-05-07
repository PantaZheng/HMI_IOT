package main

import (
	"github.com/pantazheng/bci/models"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	user := userCreate()
	user.Name += "test"
	userUpdate(&user)
	userFind()
	usersFindByLevel()
	userDelete()
	userFind()
}

func userCreate() (userJson models.UserJSON) {
	log.Println("userCreate")
	user := new(models.UserJSON)
	user.Name = "project_test"
	if tmp, err := models.UserCreate(user); err != nil {
		log.Println(err)
	} else {
		log.Println(tmp)
		userJson = tmp
	}
	return
}

func userUpdate(userJson *models.UserJSON) {
	log.Println("userUpdate")
	if p, err := models.UserUpdate(userJson); err != nil {
		log.Println(err)
	} else {
		log.Println(p)
	}
	return
}

func userFind() {
	log.Println("userFind")
	p := new(models.User)
	for i := 1; i <= 10; i++ {
		p.ID = uint(i)
		if userJson, err := models.UserFind(p); err != nil {
			log.Println(err)
		} else {
			log.Println(userJson)
		}
	}
}

func usersFindByLevel() {
	log.Println("usersFindByLevel")
	for i := 1; i <= 10; i++ {
		if ps, err := models.UsersFindByLevel(i); err != nil {
			log.Println(err)
		} else {
			log.Println(ps)
		}
	}
}

func userDelete() {
	log.Println("userDelete")
	u := new(models.User)
	for i := 1; i <= 5; i++ {
		u.ID = uint(i)
		if userJson, err := models.UserDelete(u); err != nil {
			log.Println(err)
		} else {
			log.Println(userJson)
		}
	}
}
