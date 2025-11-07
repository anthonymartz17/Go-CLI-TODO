package main

import (
	"fmt"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/handler"
)

func main(){
	TodoHandler:=	handler.NewTodoHandler()

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