package controller

import (
	"fmt"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/repository"
)
type TodoController struct{
	Repo repository.TodoRepoInterface
}

type TodoControllerInterface interface{
	PrintList() error
	AddTask(task string) error
	HandleUpdate(taskId,task string) error
	HandleDelete(taskId string) error

}
//NewController instantiate a new TodoController
func NewController(repo repository.TodoRepoInterface) TodoControllerInterface{
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
func(c *TodoController)HandleUpdate(taskId,task string)error{
	return c.Repo.UpdateTask(taskId,task)
}
func(c *TodoController)HandleDelete(taskId string)error{
	return nil
}