package main

import "fmt"

func g1() {
	fmt.Println("g1")
}

func main() {
	//Khai báo 1 go routine: go func()
	go g1()
}
