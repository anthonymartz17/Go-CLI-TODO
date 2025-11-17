package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

for{
	fmt.Print("> ")
	// waits for user input.
  reader:= bufio.NewReader(os.Stdin)
	input,err:= reader.ReadString('\n')

if err != nil{
	fmt.Println(err)
	os.Exit(1)
}

fields:= strings.Fields(strings.TrimSpace(input))
	
	err= router.HandleCommand(fields)

	if err != nil{
		fmt.Println(err)
	}


}


}