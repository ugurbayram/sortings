package main

import (
	"fmt"
	"testing"
)

func TestFibonacciLoop(t *testing.T) {
	fmt.Println(fibGen(8))
}
func TestFibonacciRec(t *testing.T) {
	fmt.Println(fibGenRec(8))
}
