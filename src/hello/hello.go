package main

import (
	"fmt"
)

type Error struct {
	errCode uint8
}

func main() {
	fmt.Println(7 / 5)
	fmt.Print("Hello world!\n")
	var err Error

	fmt.Println(err)
	println(err)
	println(7)

	//for _, env := range os.Environ(){
	//	fmt.Println(env)
	//}

	//var e interface{}
	//fmt.Println(&e)
	//if e != nil {
	//	fmt.Println("why")
	//
	//}

	fmt.Println(f())


}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}