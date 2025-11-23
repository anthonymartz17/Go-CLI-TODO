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

type Router interface{
	Route(input string) error
}

// Ensures *router.Router implements the Router interface
var _ Router = (*router.Router)(nil)

type App struct{
	router Router
}

//InitApp sets up and inject dependencies
func InitApp()(*App,error){

	dbPath:= os.Getenv("TODO_DB_PATH")
	
	if dbPath == "" {
		dbPath = "internal/db/db.json"
	}
	
	store:= db.New(dbPath)
	repo:= repository.New(store)
	crtl:= controller.New(repo)
	todoHandler:=	handler.New(crtl)
	r:= router.New(todoHandler)

	return &App{router:r},nil
}


//Run initializes and runs the app
func Run() error{
app,err:= InitApp()

if err != nil{
	return err
}

reader:= bufio.NewReader(os.Stdin)

fmt.Println("TODO CLI App")
	fmt.Println("Type 'exit' to quit")
	fmt.Println("----------------------------------")

for{
	fmt.Print("> ")
	// waits for user input.
	input,err:= reader.ReadString('\n')
	 
	if err != nil{
		fmt.Println(err)
  continue	
	}

	

	if err := app.router.Route(input); err != nil {
		fmt.Println(err)
	}


}

}