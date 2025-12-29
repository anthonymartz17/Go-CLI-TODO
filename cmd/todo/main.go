package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/anthonymartz17/Go-CLI-TODO.git/internal/router"
)




func main(){

if err:= Run(); err != nil{


	if errors.Is(err,router.ErrExit){
		fmt.Println("Program  has ended")
		os.Exit(0)
		 
	}else{
		
		fmt.Println(err)
		os.Exit(1)
	}
}



}