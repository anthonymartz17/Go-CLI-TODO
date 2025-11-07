package repository

import (
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/todo"
)

type TodoRepo struct{}

type TodoRepoInterface interface{
	GetList() ([]todo.Todo,error)
	SaveTask(task string) error
	UpdateTask(taskId,task string) error
	DeleteTask(taskId string)error
}


func NewRepo() TodoRepoInterface{
	return &TodoRepo{}
}

func(r *TodoRepo)GetList()([]todo.Todo,error){

  return nil,nil
}

func(r *TodoRepo)SaveTask(task string)error{
	return nil
}
func(r *TodoRepo)UpdateTask(taskId,task string)error{
	return nil
}
func(r *TodoRepo)DeleteTask(taskId string)error{
	return nil
}