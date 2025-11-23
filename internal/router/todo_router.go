package router

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/handler"
)

type TodoHandler interface{

	HandleList() error
	HandleAdd(fields []string) error
	HandleUpdate(fields []string) error
	HandleDone(fields []string) error
	HandleDelete(fields []string) error
}

//Ensures *handler.TodoHandler implements TodoHandler interface
var _ TodoHandler = (*handler.TodoHandler)(nil)

type Router struct{
	TodoHandler TodoHandler
}


func New(handler TodoHandler) *Router{
	return &Router{
		TodoHandler: handler,
	}
}


func(r *Router)Route(input string)error{
	fields:= strings.Fields(strings.TrimSpace(input))

	if len(fields) == 0 {
		return errors.New("no command provided")
}

	cmd:= strings.ToLower(fields[0])
	args:= fields[1:]
	
	switch cmd{
	case "list":
		return r.TodoHandler.HandleList()
	case "add":
		return r.TodoHandler.HandleAdd(args)
		
	case "update":
	
	 return r.TodoHandler.HandleUpdate(args)

	case "delete":
		
		return r.TodoHandler.HandleDelete(args)

	case "done":

		return r.TodoHandler.HandleDone(args)

	case "exit":
		fmt.Println("Program ended")
		os.Exit(0)
		return nil
	default:
		return errors.New("invalid command.  commands: list | add | update | done | delete")
	}

}
