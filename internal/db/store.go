package db

import (
	"encoding/json"
	"os"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
)

type Store struct{
 FilePath string
}


func NewStore(filePath string) *Store{
 return &Store{
    FilePath:filePath,
	}
}

func(s *Store)Load()([]*todo.Todo,error){
  data,err:= os.ReadFile(s.FilePath)

	if err != nil{
		return nil,err
	 }

	var tasks []*todo.Todo

	err = json.Unmarshal(data,&tasks)

	if err != nil{
		return nil,err
	 }

	 return tasks,nil
}

func(s *Store)Save (payload []*todo.Todo) error{

 jsonBytes,err:=  json.Marshal(payload)

 if err != nil{
	return err
 }

 return os.WriteFile(s.FilePath,jsonBytes,0644)

}