package router

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/handler"
)

type router struct{
	TodoHandler handler.TodoHandler
}

type Router interface{
	HandleCommand(fields []string) error
}



func NewRouter(handler handler.TodoHandler)Router{
	return &router{
		TodoHandler: handler,
	}
}


func(r *router)HandleCommand(fields []string)error{

	if len(fields) == 0 {
		return errors.New("no command provided")
}

	command:= strings.ToLower(fields[0])
	
	
	switch command{
	case "list":
		return r.TodoHandler.HandleList()
	case "add":
		return r.TodoHandler.HandleAdd(fields[1:])
		
	case "update":
	
	 return r.TodoHandler.HandleUpdate(fields[1:])

	case "delete":
		
		return r.TodoHandler.HandleDelete(fields[1:])

	case "done":

		return r.TodoHandler.HandleDone(fields[1:])

	case "end":
		fmt.Println("Program ended")
		os.Exit(0)
		return nil
	default:
		return errors.New("invalid command.  commands: list | add | update | done | delete")
	}

}
