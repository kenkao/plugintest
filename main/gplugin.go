package main

import (
	"fmt"
	"mylib"
)

var V int

func F() {
	lib := mylib.New()
	fmt.Printf("Hello, number %d\n", lib.Add(V, V))
}

func main(){
	fmt.Printf("ss")
}
