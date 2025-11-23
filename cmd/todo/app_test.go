package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestInitApp_defaultPath(t *testing.T){
  os.Unsetenv("TODO_DB_PATH")

	got,err:= InitApp()

  assert.NoError(t,err,"Failed to initialize app")
	assert.NotNil(t,got,"App can not be nil")

}