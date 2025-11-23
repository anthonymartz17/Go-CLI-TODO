package handler

import (
	"errors"
	"strings"
	"testing"

	"github.com/anthonymartz17/Go-CLI-TODO.git/mocks"
	"github.com/stretchr/testify/suite"
)


type TodoHandlerSuite struct{
	suite.Suite
	mockCtrl *mocks.TodoController
	todoHandler *TodoHandler
}

func (s *TodoHandlerSuite)SetupTest(){
	s.mockCtrl= new(mocks.TodoController)
	s.todoHandler= New(s.mockCtrl)

}
func (s *TodoHandlerSuite) TestHandleAdd() {
	tt := []struct {
			name         string
			param        []string
			want         error
			msg          string
			callExpected bool
	}{
			{
					name: "empty input",
					param: []string{},
					want: errors.New("missing task | Usage: add <task>"),
					msg: "AddTask should return an error when input is empty",
					callExpected: false,
			},
			{
					name: "valid input",
					param: []string{"buy", "milk"},
					want: nil,
					msg: "AddTask should succeed when input is valid",
					callExpected: true,
			},
	}

	for _, tc := range tt {
			tc := tc // capture for closure
			s.T().Run(tc.name, func(t *testing.T) {
					task := strings.Join(tc.param, " ")

					if tc.callExpected {
							s.mockCtrl.On("AddTask", task).Return(tc.want)
					}

					err := s.todoHandler.HandleAdd(tc.param)

					if tc.want == nil {
							s.NoError( err, tc.msg)
					} else {
							s.EqualError( err, tc.want.Error(), tc.msg)
					}

					s.mockCtrl.AssertExpectations(s.T())
			})
	}
}

func(s *TodoHandlerSuite)TestHandleList(){
 s.T().Run("no error", func(t *testing.T) {
	  s.mockCtrl.On("PrintList").Return(nil).Once()

		err:= s.todoHandler.HandleList()
		
		s.NoError(err,"should not error on printing")
	})
	
	s.T().Run("handle error from controller",func (t *testing.T)  {
		s.mockCtrl.On("PrintList").Return(errors.New("something went wrong")).Once()

		err:= s.todoHandler.HandleList()
		wantErr:= errors.New("something went wrong")
		s.EqualError(err, wantErr.Error(),"should error on controller error")
 })

 s.mockCtrl.AssertExpectations(s.T())
}
func(s *TodoHandlerSuite)TestHandleUpdate(){

  tt:= []struct{
      name string
			param []string
			want error
			msg string
			callExpected bool
	}{
		{
			name: "empty input",
			param: []string{},
			want: errors.New("missing ID or task | Usage: update <id> <task>"),
			msg: "HandleUpdate should return an error when input is empty",
			callExpected: false,
	},
		{
			name: "input length is 1",
			param: []string{"test"},
			want: errors.New("missing ID or task | Usage: update <id> <task>"),
			msg: "HandleUpdate should return an error when input length is 1",
			callExpected: false,
	},
	{
			name: "valid input",
			param: []string{"123", "get", "milk"},
			want: nil,
			msg: "HandleUpdate should succeed when input is valid",
			callExpected: true,
	},
	{
			name: "valid input with longer task text",
			param: []string{"123", "get", "milk","at","super","market"},
			want: nil,
			msg: "HandleUpdate should succeed when input is valid",
			callExpected: true,
	},
	}

	for _,tc := range tt{
		s.T().Run(tc.name,func(t *testing.T) {

			if tc.callExpected{
				 id:= tc.param[0]
				 task:= strings.Join(tc.param[1:]," ")
				 s.mockCtrl.On("UpdateTask",id,task).Return(tc.want)
			 }
    
			 err:= s.todoHandler.HandleUpdate(tc.param)

			 if  tc.want == nil{
				  s.NoError(err,tc.msg)
			 }else{
				s.EqualError(err,tc.want.Error(),tc.msg)
			 }

		})
	}
    s.mockCtrl.AssertExpectations(s.T())
  
}

func(s *TodoHandlerSuite)TestHandleDone(){


	tt:= []struct{
		name string
		param []string
		wantErr error
		msg string
		callExpected bool
	}{
		{name:"empty fields/no id", param: []string{},wantErr: errors.New("missing ID | Usage: done <id>"),msg: "should error on empty fields",callExpected:false},
		{name:"invalid id", param: []string{"invalid"},wantErr: errors.New("invalid id"),msg: "should error on invalid id",callExpected:true},
		{name:"valid id", param: []string{"123"},wantErr: nil,msg: "should succeed on valid id",callExpected:true},
	}

 for _,tc:= range tt{
	s.T().Run(tc.name,func(t *testing.T) {

		 // reset mock expectations for each test
		 s.mockCtrl.ExpectedCalls = nil
		 s.mockCtrl.Calls = nil

		if tc.callExpected{
       id:= tc.param[0]
			s.mockCtrl.On("ToggleDone",id).Return(tc.wantErr)
		}

		wantErr:= s.todoHandler.HandleDone(tc.param)

		if tc.wantErr == nil{
			s.NoError(wantErr,tc.msg)
		}else{
			s.EqualError(wantErr,tc.wantErr.Error(),tc.msg)
		}
	})
 }

s.mockCtrl.AssertExpectations(s.T())

}


func(s *TodoHandlerSuite)TestHandleDelete(){

 tt:= []struct{
	  name string
		param []string
		wantErr error
		msg string
		callExpected bool

 }{
	{
		name: "empty args",
		param:[]string{},
		wantErr: errors.New("missing ID | Usage: delete <id>"),
		msg:"should error on empty args list",
		callExpected: false,
	},
	{
		name: "task not found",
		param:[]string{"123"},
		wantErr: errors.New("task with id 123, not found"),
		msg:"should error on task not found",
		callExpected: true,
	},
	{
		name: "valid id",
		param:[]string{"124"},
		wantErr: nil,
		msg:"should not fail on valid id",
		callExpected: true,
	},

 }


 for _,tc:= range tt{
	s.T().Run(tc.name,func(t *testing.T) {

		if tc.callExpected{
			id:= tc.param[0]
			s.mockCtrl.On("HandleDelete",id).Return(tc.wantErr)
		}

		gotErr:= s.todoHandler.HandleDelete(tc.param)

		 if tc.wantErr == nil{
			s.NoError(gotErr,tc.msg)
			}else{
			 s.EqualError(gotErr,tc.wantErr.Error(),tc.msg)
		 }


	})
 }

 s.mockCtrl.AssertExpectations(s.T())
}

func TestTodoHandlerSuite(t *testing.T){
	 suite.Run(t,new(TodoHandlerSuite))
}