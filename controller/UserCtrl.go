package controller

import (
	"../service"
	"github.com/kataras/iris"
)

func EnrollTeacher(ctx iris.Context) {
	service.UpdateTeacher(ctx)
	service.TestMenu()
}

func EnrollStudent(ctx iris.Context){
	service.UpdateStudent(ctx)
	service.TestMenu()
}

func Purify(ctx iris.Context){
	service.Purify(ctx)
	service.TestMenu()
}

func ListStudent(ctx iris.Context){
	service.GetStudents(ctx)
}

func ListTeacher(ctx iris.Context){
	service.GetTeachers(ctx)
}