package repository

import (
	"errors"
	"fmt"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/db"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
)

// Store interface implemented by db.Store
type Store interface{
	Load()([]*todo.Todo,error)
	Save(payload []*todo.Todo) error
}

//Ensures db.Store implements Store interface
var _ Store = (*db.Store)(nil)


//TodoRepo provides access to the database store for todo tasks
type TodoRepo struct{
	Store Store
}


//NewRepo instantiate a new TodoRepo
func New(store Store) *TodoRepo{

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

//DeleteTask removes a task by id
func(r *TodoRepo)DeleteTask(taskId string)error{
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
