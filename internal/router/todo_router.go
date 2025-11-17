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
	Route(input string) error
}



func NewRouter(handler handler.TodoHandler)Router{
	return &router{
		TodoHandler: handler,
	}
}


func(r *router)Route(input string)error{
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

	case "end":
		fmt.Println("Program ended")
		os.Exit(0)
		return nil
	default:
		return errors.New("invalid command.  commands: list | add | update | done | delete")
	}

}
