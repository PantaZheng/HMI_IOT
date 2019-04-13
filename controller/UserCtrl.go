package controller

import (
	"../service"
	"github.com/kataras/iris"
)

func Check() {
	service.CheckTableUser()
}

func EnrollTeacher(ctx iris.Context) {
	service.UpdateTeacher(ctx)
}

func EnrollStudent(ctx iris.Context){
	service.UpdateStudent(ctx)
}

func Purify(ctx iris.Context){
	service.Purify(ctx)
}

func ListStudent(ctx iris.Context){
	service.GetStudents(ctx)
}

func ListTeacher(ctx iris.Context){
	service.GetTeachers(ctx)
}