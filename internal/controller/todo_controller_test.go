package controller

import (
	"errors"
	"testing"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
	"github.com/anthonymartz17/Go-CLI-TODO.git/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)




func TestPrintList(t *testing.T){

		tt:=[]struct{
			name string
			expectedGetListError error
			expectedGetListResponse []*todo.Todo
			wantErr error
			msg string
		}{
			{
				name:"mock service call returns error",
				expectedGetListResponse: nil,
				expectedGetListError: errors.New("something went wrong"),
				wantErr: errors.New("something went wrong") ,
				msg: "should fail on mock service error",
			},
			{
				name:"Empty list",
				expectedGetListResponse: []*todo.Todo{},
				expectedGetListError: nil,
				wantErr: errors.New("list is empty") ,
				msg: "should fail on empty list",
			},
			{
				name:"valid list",
				expectedGetListResponse: []*todo.Todo{
					{Id:"cb0ad4ef-e64d-4d8e-8049-12c3054f97aa",Task:"buy milk",Done:false},{Id:"ff80acba-3d57-47dd-af34-ab03f7814594",Task:"talk to the girls",Done:false},
				},
				expectedGetListError: nil,
				wantErr: nil ,
				msg: "should suceed on valid list",
			},
		}

		for _,tc:= range tt{
			t.Run(tc.name,func(t *testing.T) {

				mockRepo:= new(mocks.TodoRepo)
				ctrl:= New(mockRepo)
				
				require.NotNil(t,ctrl)

				mockRepo.On("GetList").Return(tc.expectedGetListResponse, tc.expectedGetListError).Once()
				

				err:= ctrl.PrintList()
        
          if(tc.wantErr != nil){
						assert.EqualError(t,err,tc.wantErr.Error(),tc.msg)
					}else{
						assert.NoError(t,err,tc.msg)
					}

					mockRepo.AssertExpectations(t)
			})
		}

		
}


func TestAddTask_Success(t *testing.T){
  mockRepo:= new(mocks.TodoRepo)
	crtl:= New(mockRepo)

  task:= "buy milk"
	mockRepo.On("SaveTask",task).Return(nil).Once()
 
  gotErr:= crtl.AddTask(task)

	assert.NoError(t,gotErr,"should succeed on valid data")
	mockRepo.AssertExpectations(t)
}
func TestAddTask_RepoFails(t *testing.T){
	mockRepo:= new(mocks.TodoRepo)
	crtl:= New(mockRepo)

  task:= "buy milk"
	err:= errors.New("internal error")
	mockRepo.On("SaveTask",task).Return(err).Once()
	gotErr:= crtl.AddTask(task)


	assert.ErrorIs(t,gotErr,err,"should fail on internal error")
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask_Success(t *testing.T){
  mockRepo:= new(mocks.TodoRepo)
	ctrl:= New(mockRepo)
  taskId:= "123"
	task:= "buy yogurt"

  mockRepo.On("UpdateTask",taskId,task).Return(nil).Once()

	gotErr:= ctrl.UpdateTask(taskId,task)

	assert.NoError(t,gotErr,"should succeed on valid data")

	mockRepo.AssertExpectations(t)
}
func TestUpdateTask_RepoFails(t *testing.T){
  mockRepo:= new(mocks.TodoRepo)
	ctrl:= New(mockRepo)
  taskId:= "123"
	task:= "buy yogurt"
	err:= errors.New("repository internal error")

  mockRepo.On("UpdateTask",taskId,task).Return(err).Once()

	gotErr:= ctrl.UpdateTask(taskId,task)

	assert.Error(t,gotErr,"should fail on repo internal error")

	mockRepo.AssertExpectations(t)
}

func TestToggle_Success(t *testing.T){
 mockRepo:= new(mocks.TodoRepo)
 ctrl:= New(mockRepo)
 taskId:= "123"
 mockRepo.On("ToggleDone",taskId).Return(nil).Once()

 gotErr:= ctrl.ToggleDone(taskId)
 assert.NoError(t,gotErr,"should succeed on valid id")

 mockRepo.AssertExpectations(t)
}

func TestToggle_RepoFails(t *testing.T){
	mockRepo:= new(mocks.TodoRepo)
	ctrl:= New(mockRepo)
	taskId:= "124"
	err:= errors.New("repository internal error")

	mockRepo.On("ToggleDone",taskId).Return(err).Once()

	gotErr:= ctrl.ToggleDone(taskId)

	assert.Error(t,gotErr,"should faile on repository internal error")

	mockRepo.AssertExpectations(t)

}
func TestHandleDelete_Success(t *testing.T){
 mockRepo:= new(mocks.TodoRepo)
 ctrl:= New(mockRepo)
 taskId:= "123"
 mockRepo.On("DeleteTask",taskId).Return(nil).Once()

 gotErr:= ctrl.HandleDelete(taskId)
 assert.NoError(t,gotErr,"should succeed on valid id")

 mockRepo.AssertExpectations(t)
}

func TestHandleDelete_RepoFails(t *testing.T){
	mockRepo:= new(mocks.TodoRepo)
	ctrl:= New(mockRepo)
	taskId:= "124"
	err:= errors.New("repository internal error")

	mockRepo.On("DeleteTask",taskId).Return(err).Once()

	gotErr:= ctrl.HandleDelete(taskId)

	assert.Error(t,gotErr,"should fail on repository internal error")

	mockRepo.AssertExpectations(t)

}