package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/controller"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/db"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/handler"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/repository"
	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/router"
)

func main(){
	
filePath:="/Users/antoniomartinez/programacion/myProjects/go-projects/Go-CLI-TODO/internal/db/db.json"
store:= db.NewStore(filePath)
repo:= repository.NewRepo(store)
crtl:= controller.NewController(repo)
todoHandler:=	handler.NewTodoHandler(crtl)
router:= router.NewRouter(todoHandler)

reader:= bufio.NewReader(os.Stdin)

for{
	fmt.Print("> ")
	// waits for user input.
	input,err:= reader.ReadString('\n')
	 
	if err != nil{
		fmt.Println(err)
  continue	
	}
  
	err= router.Route(input)

	if err != nil{
		fmt.Println(err)
	}


}


}