package main

import (
	"fmt"
	"sync"
)

func g1() {
	fmt.Println("g1")
	wg.Done()
}

func g2() {
	fmt.Println("g2")
	wg.Done()

}

var wg sync.WaitGroup

func main() {
	//Khai báo 1 go routine: go func()
	// go g1()

	// ng := runtime.NumGoroutine()
	// fmt.Println(ng)

	// time.Sleep(time.Second)

	//Synchronized goroutines
	/*
		logic 1

		go g1()
		go g2()

		logic 2

	*/
	fmt.Println("Begin")
	wg.Add(2)

	go g1()
	go g2()

	wg.Wait()
	fmt.Println("End")

}
