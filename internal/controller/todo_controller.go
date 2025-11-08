package controller

import (
	"encoding/json"
	"fmt"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/repository"
)
type TodoController struct{
	Repo repository.TodoRepoInterface
}

type TodoControllerInterface interface{
	PrintList() error
	handleAdd(task string) error
	handleUpdate(taskId,task string) error
	handleDelete(taskId string) error

}

func NewController() TodoControllerInterface{
	return &TodoController{
		Repo: repository.NewRepo(),
	}
}


func(c *TodoController)PrintList()error{
	todos,err:= c.Repo.GetList()

	if err != nil{
		return err
	}

	jsonBytes, err := json.MarshalIndent(todos, "", "  ")
if err != nil {
    return err
}

  fmt.Println(string(jsonBytes))
  return nil
}

func(c *TodoController)handleAdd(task string)error{
	return nil
}
func(c *TodoController)handleUpdate(taskId,task string)error{
	return nil
}
func(c *TodoController)handleDelete(taskId string)error{
	return nil
}