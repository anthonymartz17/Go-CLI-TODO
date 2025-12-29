package repository

import (
	"errors"
	"testing"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
	"github.com/anthonymartz17/Go-CLI-TODO.git/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetList_Success(t *testing.T) {

tt:= []struct{
	name string
	want []*todo.Todo
	msg string
}{
	{
		name:"valid todo list",
	  want: []*todo.Todo{
		{Id:"123",Task:"buy milk",Done:false},{Id:"ff80acba-3d57-47dd-af34-ab03f7814594",Task:"talk to the girls",Done:false},},
		msg:"should succeed on valid data",
	},
	{
		name:"empty todo list",
	  want: []*todo.Todo{},
		msg:"should succeed on empty list",
	},
}


for _,tc:= range tt{
	t.Run(tc.name,func(t *testing.T) {
		mockStore:= new(mocks.Store)
		repo:= New(mockStore)
		mockStore.On("Load").Return(tc.want,nil).Once()

		got,err:= repo.GetList()

		assert.NoError(t,err,tc.msg)
		assert.NotNil(t,got,tc.msg)
		assert.Equal(t, got,tc.want,tc.msg)
		
		mockStore.AssertExpectations(t)
	})
}












	
}

func TestGetList_Failure(t *testing.T){
	//arrange
	mockStore:= new(mocks.Store)
	repo:= New(mockStore)
  err:= errors.New("failed to load list")
	mockStore.On("Load").Return(nil,err)

	//act
	todos,gotErr:= repo.GetList()
	
	//assert
	assert.Nil(t,todos,"should be nil on error")
	assert.Error(t,gotErr,"should fail on error from store")


	mockStore.AssertExpectations(t)
}

func TestSaveTask_Success(t *testing.T){
	//arrange
	mockStore:= new(mocks.Store)
	repo:= New(mockStore)

	data:=  []*todo.Todo{
		{Id:"ff80acba-3d57-47dd-af34-ab03f7814594",Task:"talk to the girls",Done:false},
	}
  task:= &todo.Todo{Id:"123",Task:"buy milk",Done:false}

	mockStore.On("Load").Return(data,nil)
	mockStore.On("Save",mock.Anything).Return(nil)

  
	//act
	gotErr:= repo.SaveTask(task)

	//assert
	assert.NoError(t,gotErr,"should succeed on valid todo")

	mockStore.AssertExpectations(t)

}



func TestUpdateTask(t *testing.T){
	


	
	tt:= []struct{
		name string
		task string
		taskId string
		wantErr error
		loadData []*todo.Todo
		loadErr error
		expectSave bool
		saveErr error
		msg string
		}{
			{
				name:"successful update",
				taskId:"123",
				task:"new task",
				wantErr: nil,
				loadData: []*todo.Todo{
					{Id:"123",Task:"old task",Done:false},
				},
				loadErr: nil,
				expectSave:true,
				saveErr:nil,
				msg:"should not error on successful case",
				
			},
			{
				name:"TASK does does not exist",
				taskId:"124",
				task:"new task",
				wantErr: errors.New("task with id: 124, does not exist"),
				loadData: []*todo.Todo{
					{Id:"123",Task:"old task",Done:false},
				},
				loadErr: nil,
				expectSave:false,
				saveErr:nil,
				msg:"should fail on task does not exist",
				
			},
			{
				name:"loading error",
				taskId:"123",
				task:"new task",
				wantErr: errors.New("error loading data"),
				loadData: []*todo.Todo{
					{Id:"123",Task:"old task",Done:false},
				},
				loadErr: errors.New("error loading data"),
				expectSave:false,
				saveErr:nil,
				msg:"should fail on loading data error",
				
			},
			{
				name:"saving error",
				taskId:"123",
				task:"new task",
				wantErr: errors.New("error saving task"),
				loadData: []*todo.Todo{
					{Id:"123",Task:"old task",Done:false},
				},
				loadErr: nil,
				expectSave:true,
				saveErr:errors.New("error saving task"),
				msg:"should fail on saving error",
				
			},
		}
		
		for _,tc:= range tt{
			t.Run(tc.name,func(t *testing.T) {
		//arrange
		mockStore:= new(mocks.Store)
		repo:= New(mockStore)
	 
		mockStore.On("Load").Return(tc.loadData,tc.loadErr).Once()
		
    
		if tc.expectSave{
			mockStore.On("Save",mock.Anything).Return(tc.saveErr).Once()
		}
	 
		//act
	  gotErr:= repo.UpdateTask(tc.taskId,tc.task)
	 
		//assert
		if tc.wantErr == nil{
			assert.NoError(t,gotErr,tc.msg)
		}else{

			assert.EqualError(t,gotErr,tc.wantErr.Error(),tc.msg)
		}

		if tc.expectSave{
			assert.Equal(t,tc.task,tc.loadData[0].Task)
		}

	 
		mockStore.AssertExpectations(t)
	})
 }

}



func TestToggleDone(t *testing.T){
	
	tt:= []struct{
		name string
		wantErr error
		wantDone bool
    loadedData []*todo.Todo
		loadErr error
		taskId string
		expectSave bool
		saveErr error
		msg string
		
		
		}{
			{
				name:"successful toggle",
				wantErr: nil,
				wantDone: true,
				loadedData: []*todo.Todo{
					{Id:"123",Task:"old task",Done:false},
				},
				loadErr:nil,
				taskId:"123",
				expectSave:true,
				saveErr:nil,
				msg:"should succeed on successful toggle",
			},
			{
				name:"error loading",
				wantErr: errors.New("error loading data"),
				wantDone: false,
				loadedData: []*todo.Todo{
					{Id:"123",Task:"old task",Done:false},
				},
				loadErr: errors.New("error loading data"),
				taskId:"123",
				expectSave:false,
	 saveErr:nil,
	 msg:"should should fail on error loading",
 },
			{
				name:"error saving",
				wantErr: errors.New("fail trying to save data"),
				wantDone: true,
				loadedData: []*todo.Todo{
					{Id:"123",Task:"old task",Done:false},
				},
				loadErr: nil,
				taskId:"123",
				expectSave:true,
	 saveErr: errors.New("fail trying to save data"),
	 msg:"should fail on error saving",
 },
			{
				
	name:"task does not exist",
				wantErr: errors.New("task with id: 124, does not exist"),
				wantDone: false,
				loadedData: []*todo.Todo{
					{Id:"123",Task:"old task",Done:false},
				},
				loadErr: nil,
				taskId:"124",
				expectSave:false,
	 saveErr: nil,
	 msg:"should fail on task does not exist",
 },
	}

	for _,tc:= range tt{
		t.Run(tc.name,func(t *testing.T) {
			//arrange
			mockStore:= new(mocks.Store)
			repo:= New(mockStore)


			mockStore.On("Load").Return(tc.loadedData,tc .loadErr).Once()

			if tc.expectSave{
				mockStore.On("Save",tc.loadedData).Return(tc.saveErr).Once()
			}


			//act
     gotErr:= repo.ToggleDone(tc.taskId)

			//assert
			if tc.wantErr == nil{
				assert.NoError(t,gotErr,tc.msg)
			}else{
				assert.EqualError(t, gotErr,tc.wantErr.Error(),tc.msg)
			}

			if tc.expectSave{
				assert.Equal(t,tc.loadedData[0].Done,tc.wantDone,"Done property should match wantDone after toggle")
			}

			mockStore.AssertExpectations(t)
		})
	}
}


func TestDeleteTask(t *testing.T){
	tt:= []struct{
		name string
		wantErr error
		loadedData  []*todo.Todo
		loadErr error
		saveErr error
		expectSave bool
		taskId string
		msg string
	}{
		{
			name:"successful deletion",
			wantErr: nil,
			loadedData:  []*todo.Todo{
				{Id:"123",Task:"old task",Done:false},
			},
			loadErr: nil,
			saveErr: nil,
			expectSave: true,
			taskId: "123",
			msg:"should succeed on valid data",
			
		},
		{
			name:"error loading",
			wantErr: errors.New("fail to load data"),
			loadedData:  nil,
			loadErr: errors.New("fail to load data"),
			saveErr: nil,
			expectSave: false,
			taskId: "123",
			msg:"should fail on error loading",
		
		},
		{
			name:"error saving",
			wantErr: errors.New("fail to save data"),
			loadedData:   []*todo.Todo{
				{Id:"123",Task:"old task",Done:false},
			},
			loadErr: nil,
			saveErr: errors.New("fail to save data"),
			expectSave: true,
			taskId: "123",
			msg:"should fail on error saving",
		},
		{
			name:"task not found",
			wantErr: errors.New("task with id: 124, does not exist"),
			loadedData:   []*todo.Todo{
				{Id:"123",Task:"old task",Done:false},
			},
			loadErr: nil,
			saveErr: nil,
			expectSave: false,
			taskId: "124",
			msg:"should fail on error saving",
		},
	}

	for _,tc:= range tt{
		t.Run(tc.name,func(t *testing.T) {
			//arrange
			mockStore:= new(mocks.Store)
			repo:= New(mockStore)

			mockStore.On("Load").Return(tc.loadedData,tc.loadErr).Once()

			if tc.expectSave{
				mockStore.On("Save",mock.Anything).Return(tc.saveErr).Once()
			}

			//act

			gotErr:= repo.DeleteTask(tc.taskId)

			//assert

			if tc.wantErr == nil{
				assert.NoError(t, gotErr,tc.msg)
			}else{
				assert.EqualError(t,gotErr,tc.wantErr.Error(),tc.msg)
			}

			if tc.expectSave{
				savedData:= mockStore.Calls[1].Arguments.Get(0).([]*todo.Todo)
        
				//assert only one task was removed
				assert.Len(t,savedData,0,tc.msg)
        
				//assert right task was removed
				for _,task:= range savedData{
					assert.NotEqual(t,task.Id,tc.taskId,tc.msg)
				}

			}


			mockStore.AssertExpectations(t)
		})

	}
	

}


func TestFindTaskById_Success(t *testing.T){
	
	data:=  []*todo.Todo{
		{Id:"123",Task:"buy milk",Done:false},
		{Id:"124",Task:"talk to the girls",Done:false},
		{Id:"125",Task:"do homework",Done:false},
	}

	tt:= []struct{
		name string
		taskId string
		want *todo.Todo
		msg string
		}{
			{
				name:"success id 123",
				taskId:"123",
				want: &todo.Todo{Id:"123",Task:"buy milk",Done:false},
				msg: "should find task with id 123",
			},
			{
				name:"success id 124",
				taskId:"124",
				want: &todo.Todo{Id:"124",Task:"talk to the girls",Done:false},
				msg: "should find task with id 124",
			},
			{
				name:"success id 125",
				taskId:"125",
				want: &todo.Todo{Id:"125",Task:"do homework",Done:false},
				msg: "should find task with id 125",
			},
	}

	for _,tc:= range tt{
		t.Run(tc.name,func(t *testing.T) {
			//arrange
			//act
			got,gotErr:= findTaskById(data,tc.taskId)

			//assert
			assert.NotNil(t,got,tc.msg)
			assert.NoError(t, gotErr, tc.msg)
			assert.Equal(t,got.Task,tc.want.Task,tc.msg)
			assert.Equal(t,got.Id,tc.want.Id,tc.msg)
			assert.Equal(t,got.Done,tc.want.Done,tc.msg)
		})
	}
 
}
func TestFindTaskById_FailToFind(t *testing.T){
	data:=  []*todo.Todo{
		{Id:"123",Task:"buy milk",Done:false},
		{Id:"124",Task:"talk to the girls",Done:false},
		{Id:"125",Task:"do homework",Done:false},
	}

	t.Run("Task not found", func(t *testing.T) {
		//arrange
		taskId:= "116"
		wantErr:= errors.New("task with id: 116, does not exist")

		//act

		got,gotErr:= findTaskById(data,taskId)

		//assert
		assert.Nil(t,got,"got data should be nil")
		assert.EqualError(t,gotErr,wantErr.Error(),"should fail on task not found")
	})
}