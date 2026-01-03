package db

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/entity/todo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)



func TestLoad(t *testing.T){
	
	tempDir:= t.TempDir()
	tempFilePath:= filepath.Join(tempDir,"tempFile.json")
	store:= New(tempFilePath)
	
	t.Run("empty list when file does not exist", func(t *testing.T) {
	//act
	got,err:= store.Load()

	//assert
	assert.NoError(t,err,"should not error when file does not exist")
	assert.Len(t,got,0,"should return an empty list when file does not exist")
})

t.Run("return non-empty list when file exists", func(t *testing.T) {
	//arrange
  data:= []*todo.Todo{ {Id:"123",Task:"old task",Done:false},}
	err:= store.Save(data)
	require.NoError(t,err,"saving test data should not fail")

	//act
	got,err:= store.Load()

	//assert
	assert.NoError(t,err,"should not error when file exist")
	assert.Len(t,got,1,"should return a one element list")
})


t.Run("invalid  JSON", func(t *testing.T) {
	//arrange
	_= os.WriteFile(store.FilePath,[]byte("{invalid-JSON"),0644)

	//act
	got,err:= store.Load()

	//assert
	assert.Error(t,err,"should fail on invalid JSON")
	assert.Nil(t,got,"should be nil on invalid JSON")
})


}

func TestSave(t *testing.T){
//arrange
tempDir:= t.TempDir()
tempFilePath:= filepath.Join(tempDir,"tempFile.json")
store:= New(tempFilePath)
wantData:= []*todo.Todo{{Id:"123",Task:"a task",Done:false},}

//act
gotErr:= store.Save(wantData)
byteData,err:= os.ReadFile(tempFilePath)
require.NoError(t,err,"should not fail test file read")

var gotData []*todo.Todo
err= json.Unmarshal(byteData,&gotData)
require.NoError(t,err,"should not fail test unmarshalling")

//assert
assert.NoError(t,gotErr,"should succeed on valid data")
assert.Equal(t,wantData,gotData,"saved data should match expected data")
}