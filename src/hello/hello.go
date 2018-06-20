package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println(7/5)
	fmt.Print("Hello world!\n")


	for _, env := range os.Environ(){
		fmt.Println(env)
	}

}
