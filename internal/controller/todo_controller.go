package controller

import (
	"fmt"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/repository"
)

//TodoRepo defines the methods implemented by TodoRepo
type TodoRepo interface{
	GetList() ([]*todo.Todo,error)
	SaveTask(task *todo.Todo) error
	UpdateTask(taskId,task string) error
	ToggleDone(taskId string)error
	DeleteTask(taskId string)error
}


// Ensures *repository.TodoRepo implements the TodoRepo interface
var _ TodoRepo = (*repository.TodoRepo)(nil)


type TodoController struct{
	Repo TodoRepo
}


//NewController instantiate a new TodoController
func New(repo TodoRepo) *TodoController{
	return &TodoController{
		Repo: repo,
	}
}

//PrintList prints list of tasks
func(c *TodoController)PrintList()error{
	todos,err:= c.Repo.GetList()

	if err != nil{
		return err
	}

	for i,task:= range todos{
		taskNum := i+1
		todo:= task.Task
		completed := "no"
		if task.Done {
				completed = "yes"
		}
		
    
		line:= fmt.Sprintf("%v. %s  ID: %v  Completed: %v",taskNum,todo,task.Id,completed)

		fmt.Println(line)
	}


  return nil
}

func(c *TodoController)AddTask(taskReq string)error{
  task:= todo.NewTodo(taskReq)
	return c.Repo.SaveTask(task)
}
func(c *TodoController)UpdateTask(taskId,task string)error{
	return c.Repo.UpdateTask(taskId,task)
}
func(c *TodoController)ToggleDone(taskId string)error{
	return c.Repo.ToggleDone(taskId)
}
func(c *TodoController)HandleDelete(taskId string)error{

	return c.Repo.DeleteTask(taskId)
}