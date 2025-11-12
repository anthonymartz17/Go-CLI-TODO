package repository

import (
	"errors"
	"fmt"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/db"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
)

type TodoRepo struct{
	Store *db.Store
}

type TodoRepoInterface interface{
	GetList() ([]*todo.Todo,error)
	SaveTask(task *todo.Todo) error
	UpdateTask(taskId,task string) error
	ToggleDone(taskId string)error
	DeleteTask(taskId string)error
}

//NewRepo instantiate a new TodoRepo
func NewRepo(store *db.Store) TodoRepoInterface{

	return &TodoRepo{
		Store: store,
	}
}
//GetList loads the json database and returns the list of tasks.
func(r *TodoRepo)GetList()([]*todo.Todo,error){

	list,err := r.Store.Load()

	if err != nil{
		return nil,err
	}

	return list,nil

}
//SaveTask loads json database, appends a new task and saves it back.
func(r *TodoRepo)SaveTask(task *todo.Todo)error{
  
	data,err:= r.Store.Load()

	if err != nil{
		return err
	}

	data = append(data, task)

	return r.Store.Save(data)
	
}

func(r *TodoRepo)UpdateTask(taskId,task string)error{
	 data,err := r.Store.Load()

	 if err != nil{
		return err
	 }

	 taskExist:= false

	 for _,item:= range data{
		if item.Id == taskId{
			item.Task = task
			taskExist = true
		}
	}

	if !taskExist{
		errorMsg:= fmt.Sprintf("task with id: %v, does not exist",taskId)
		return errors.New(errorMsg)
	}

	return r.Store.Save(data)
   
}



func(r *TodoRepo)ToggleDone(taskId string)error{
	 data,err:= r.Store.Load()

	 if err != nil{
		return err
	 }

   t,err:= findTaskById(data,taskId)
    
	 if err != nil{
		return err
	 }
	  
	 t.Done = !t.Done
	return r.Store.Save(data)
}

func(r *TodoRepo)DeleteTask(taskId string)error{
	return nil
}

func findTaskById(tasks []*todo.Todo, id string)(*todo.Todo,error){

  for _,item:= range tasks{
		if item.Id == id{
			return item,nil
		}
	 
	}
		errorMsg:= fmt.Sprintf("task with id: %v, does not exist",id)
		return nil,errors.New(errorMsg)
	 
}
