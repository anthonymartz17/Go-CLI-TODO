package controller

import (
	"fmt"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/repository"
)
type todoController struct{
	Repo repository.TodoRepo
}

type TodoController interface{
	PrintList() error
	AddTask(task string) error
	HandleUpdate(taskId,task string) error
	ToggleDone(taskId string)error
	HandleDelete(taskId string) error

}
//NewController instantiate a new TodoController
func NewController(repo repository.TodoRepo) TodoController{
	return &todoController{
		Repo: repo,
	}
}

//PrintList prints list of tasks
func(c *todoController)PrintList()error{
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

func(c *todoController)AddTask(taskReq string)error{
  task:= todo.NewTodo(taskReq)
	return c.Repo.SaveTask(task)
}
func(c *todoController)HandleUpdate(taskId,task string)error{
	return c.Repo.UpdateTask(taskId,task)
}
func(c *todoController)ToggleDone(taskId string)error{
	return c.Repo.ToggleDone(taskId)
}
func(c *todoController)HandleDelete(taskId string)error{

	return c.Repo.DeleteTask(taskId)
}