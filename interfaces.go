package main

import "fmt"

type Employee interface {
	Print(name string)
}

type emp int

func (e emp) Print(name string) {
	fmt.Println("My name is " + name)
}
