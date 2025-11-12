package handler

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/controller"
)

type TodoHandler struct{
	Reader *bufio.Reader
	Controller controller.TodoControllerInterface
}

func NewTodoHandler(ctrl controller.TodoControllerInterface) *TodoHandler{
  return &TodoHandler{
		Reader: bufio.NewReader(os.Stdin),
		Controller: ctrl,
	}
}

func (h *TodoHandler)PromptInput()([]string,error){
	input,err:= h.Reader.ReadString('\n')

	if err != nil{
		return nil,err
	}

	fields:= strings.Fields(strings.TrimSpace(input))
	
	return fields,nil
}


func(h *TodoHandler)HandleCommand(fields []string)error{

	if len(fields) == 0 {
		return errors.New("no command provided")
}

	command:= strings.ToLower(fields[0])
	
	
	switch command{
	case "list":
		return h.handleList()
	case "add":
		return h.handleAdd(fields[1:])
		
	case "update":
	
	 return h.handleUpdate(fields[1:])

	case "delete":
		if len(fields) < 2{
			return errors.New("missing ID | Usage: delete <id> <task>")
		}
		id:= fields[1]

		return h.handleDelete(id)

	case "end":
		fmt.Println("Program ended")
		os.Exit(0)
		return nil
	default:
		return errors.New("invalid command.  commands: list | add | update | done | delete")
	}

}


func(h *TodoHandler)handleList()error{
	return h.Controller.PrintList()
}

func(h *TodoHandler)handleAdd(fields []string)error{
	if len(fields) == 0{
		return errors.New("missing task | Usage: add <task>")
	}
	taskReq:= strings.Join(fields," ")

	return h.Controller.AddTask(taskReq)
	
}
func(h *TodoHandler)handleUpdate(fields []string)error{
    fmt.Println(fields)
	if len(fields) < 2{
		return errors.New("missing ID or task | Usage: update <id> <task>")
	}

 id:= fields[0]
 task:= strings.Join(fields[1:]," ")

 return h.Controller.HandleUpdate(id,task)

}
func(h *TodoHandler)handleDelete(taskId string)error{
	return nil
}