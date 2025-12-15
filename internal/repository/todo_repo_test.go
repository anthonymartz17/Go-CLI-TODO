package repository

import (
	"errors"
	"testing"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
	"github.com/anthonymartz17/Go-CLI-TODO.git/mocks"
	"github.com/stretchr/testify/assert"
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