package handler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/controller"
)

type todoHandler struct{
	Reader *bufio.Reader
	Controller controller.TodoController
}

type TodoHandler interface{
	PromptInput()([]string,error)
	HandleCommand(fields []string) error
	HandleList() error
	HandleAdd(fields []string) error
	HandleUpdate(fields []string) error
	HandleDone(fields []string) error
	HandleDelete(fields []string) error
}

func NewTodoHandler(ctrl controller.TodoController) TodoHandler{
  return &todoHandler{
		Reader: bufio.NewReader(os.Stdin),
		Controller: ctrl,
	}
}

func (h *todoHandler)PromptInput()([]string,error){
	input,err:= h.Reader.ReadString('\n')

	if err != nil{
		return nil,err
	}

	fields:= strings.Fields(strings.TrimSpace(input))
	
	return fields,nil
}


func(h *todoHandler)HandleCommand(fields []string)error{

	if len(fields) == 0 {
		return errors.New("no command provided")
}

	command:= strings.ToLower(fields[0])
	
	
	switch command{
	case "list":
		return h.HandleList()
	case "add":
		return h.HandleAdd(fields[1:])
		
	case "update":
	
	 return h.HandleUpdate(fields[1:])

	case "delete":
		
		return h.HandleDelete(fields[1:])

	case "done":

		return h.HandleDone(fields[1:])

	case "end":
		fmt.Println("Program ended")
		os.Exit(0)
		return nil
	default:
		return errors.New("invalid command.  commands: list | add | update | done | delete")
	}

}


func(h *todoHandler)HandleList()error{
	return h.Controller.PrintList()
}

func(h *todoHandler)HandleAdd(fields []string)error{
	if len(fields) == 0{
		return errors.New("missing task | Usage: add <task>")
	}
	taskReq:= strings.Join(fields," ")

	return h.Controller.AddTask(taskReq)
	
}
func(h *todoHandler)HandleUpdate(fields []string)error{
	if len(fields) < 2{
		return errors.New("missing ID or task | Usage: update <id> <task>")
	}

 id:= fields[0]
 task:= strings.Join(fields[1:]," ")

 return h.Controller.HandleUpdate(id,task)

}

func(h *todoHandler)HandleDone(fields []string) error{
	if len(fields) == 0{
		return errors.New("missing ID | Usage: done <id>")
	}

	id:= fields[0]

	return h.Controller.ToggleDone(id)

}


func(h *todoHandler)HandleDelete(fields []string)error{
	if len(fields) < 1{
		return errors.New("missing ID | Usage: delete <id>")
	}

	id:= fields[0]

	return h.Controller.HandleDelete(id)
}