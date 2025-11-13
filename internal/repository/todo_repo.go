package repository

import (
	"errors"
	"fmt"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/db"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
)

//todoRepo provides access to the database store for todo tasks
type todoRepo struct{
	Store *db.Store
}

//TodoRepo defines the methods implemented by TodoRepo
type TodoRepo interface{
	GetList() ([]*todo.Todo,error)
	SaveTask(task *todo.Todo) error
	UpdateTask(taskId,task string) error
	ToggleDone(taskId string)error
	DeleteTask(taskId string)error
}

//NewRepo instantiate a new TodoRepo
func NewRepo(store *db.Store) TodoRepo{

	return &todoRepo{
		Store: store,
	}
}
//GetList loads the json database and returns the list of tasks.
func(r *todoRepo)GetList()([]*todo.Todo,error){

	list,err := r.Store.Load()

	if err != nil{
		return nil,err
	}

	return list,nil

}
//SaveTask loads json database, appends a new task and saves it back.
func(r *todoRepo)SaveTask(task *todo.Todo)error{
  
	data,err:= r.Store.Load()

	if err != nil{
		return err
	}

	data = append(data, task)

	return r.Store.Save(data)
	
}

func(r *todoRepo)UpdateTask(taskId,task string)error{
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



func(r *todoRepo)ToggleDone(taskId string)error{
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

//DeleteTask removes a task by id
func(r *todoRepo)DeleteTask(taskId string)error{
	data,err:= r.Store.Load()

	if err != nil{
		return err
	}
  
	_,err= findTaskById(data,taskId)
	if err != nil{
		return err
	}
 
	
	updatedData:= make([]*todo.Todo,0,len(data) - 1)

	for _,item:= range data{
		if item.Id != taskId{
			updatedData = append(updatedData, item)
		}
	}

	return r.Store.Save(updatedData)

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
