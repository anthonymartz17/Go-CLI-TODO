package router

import (
	"testing"

	"github.com/anthonymartz17/Go-CLI-TODO.git/mocks"
	"github.com/stretchr/testify/assert"
)


func TestRoute_CommandMatchMethod(t *testing.T){

 tt:=[]struct{
  name string
	command string
	args []string
	req string
	wantErr error
	msg string
 }{

	{
		name:"List command",
		req: "list",
		command: "list",
		wantErr: nil,
		msg: "should call HandleList",

	},

	{
		name:"add command",
		req: "add buy milk and butter",
		command: "add",
		args: []string{"buy","milk","and","butter"},
		wantErr: nil,
		msg: "should call HandleAdd",

	},
	{
		name:"update command",
		req: "update 123 buy milk and butter",
		command: "update",
		args: []string{"123","buy","milk","and","butter"},
		wantErr: nil,
		msg: "should call HandleUpdate",

	},
	{
		name:"delete command",
		req: "delete 123",
		command: "delete",
		args: []string{"123"},
		wantErr: nil,
		msg: "should call HandleDelete",

	},
	{
		name:"done command",
		req: "done 123",
		command: "done",
		args: []string{"123"},
		wantErr: nil,
		msg: "should call HandleDone",

	},


 }

 for _,tc:= range tt{
	t.Run(tc.name,func(t *testing.T) {
		//arange
		mockHandler:= new(mocks.TodoHandler)
		router:= New(mockHandler)
     
		switch tc.command{
		case "list":
			mockHandler.On("HandleList").Return(nil)
		case "add":
			mockHandler.On("HandleAdd",tc.args).Return(nil)
		case "update":
			mockHandler.On("HandleUpdate",tc.args).Return(nil)
			
		case "delete":
			mockHandler.On("HandleDelete",tc.args).Return(nil)
			
			
		case "done":
			mockHandler.On("HandleDone",tc.args).Return(nil)
		}
		
		//act

		gotErr := router.Route(tc.req)



		//assert
		assert.NoError(t,gotErr)

		   
		switch tc.command{
		case "list":
			mockHandler.AssertCalled(t,"HandleList")
		case "add":
			mockHandler.AssertCalled(t,"HandleAdd",tc.args)
		case "update":
			mockHandler.AssertCalled(t,"HandleUpdate",tc.args)
			
		case "delete":
			mockHandler.AssertCalled(t,"HandleDelete",tc.args)
		case "done":
			mockHandler.AssertCalled(t,"HandleDone",tc.args)
		}


		mockHandler.AssertExpectations(t)
	})
 }







}

func TestRoute_InvalidInputs(t *testing.T) {
	mockHandler := new(mocks.TodoHandler)
	router := New(mockHandler)

	tests := []struct {
		name    string
		req     string
		wantErr string
	}{
		{
			name:    "empty input",
			req:     "",
			wantErr: "no command provided",
		},
		{
			name:    "spaces only",
			req:     "    ",
			wantErr: "no command provided",
		},
		{
			name:    "invalid command",
			req:     "an invalid command",
			wantErr: "invalid command.  commands: list | add | update | done | delete",
		},
		{
			name:    "exit command",
			req:     "exit",
			wantErr: ErrExit.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantErr := router.Route(tt.req)
			if tt.wantErr == ErrExit.Error() {
				assert.ErrorIs(t, wantErr, ErrExit)
			} else {
				assert.EqualError(t, wantErr, tt.wantErr)
			}
		})
	}
}
