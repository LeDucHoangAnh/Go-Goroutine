package main

import (
	"fmt"
	"runtime"
	"time"
)

func g1() {
	fmt.Println("g1")
}

func main() {
	//Khai báo 1 go routine: go func()
	go g1()

	ng := runtime.NumGoroutine()
	fmt.Println(ng)

	time.Sleep(time.Second)
}
