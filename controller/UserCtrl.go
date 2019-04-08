package controller

import (
	"../service"
	"github.com/kataras/iris"
)

func EnrollTeacher(ctx iris.Context) {
	service.UpdateTeacher(ctx)
}

func EnrollStudent(ctx iris.Context){
	service.UpdateStudent(ctx)
}

func ListStudent(ctx iris.Context){
	service.GetStudents(ctx)
}

func ListTeacher(ctx iris.Context){
	service.GetTeachers(ctx)
}