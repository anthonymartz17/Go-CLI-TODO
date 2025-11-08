package main

import (
	"fmt"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/controller"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/db"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/handler"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/repository"
)

func main(){
	filePath:="/Users/antoniomartinez/programacion/myProjects/go-projects/Go-CLI-TODO/internal/db/db.json"
	store:= db.NewStore(filePath)
repo:= repository.NewRepo(store)
crtl:= controller.NewController(repo)
	TodoHandler:=	handler.NewTodoHandler(crtl)

for{
	fmt.Print("> ")
	// waits for user input.
	fields,err:= TodoHandler.PromptInput() 
	 
	if err != nil{
		fmt.Print(err)

	}
  
	err= TodoHandler.HandleCommand(fields)

	if err != nil{
		fmt.Println(err)
	}


	
}



}