package main

import (
	"fmt"
	"time"
)

// create two functions
//  1 function is to print numbers till 10
//  2 function is to print numbes till e

func printNumbers() {
	var num int = 1
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", num)
		num++
		fmt.Println()
		time.Sleep(time.Millisecond * 500) // Sleep for a short duration
	}
}
func printLetters() {
	for char := 'a'; char <= 'e'; char++ {
		fmt.Println("Consistency :(")
		fmt.Printf("%c", char)
		fmt.Println()
		time.Sleep(time.Microsecond * 700)
	}
}
func main() {
	go printNumbers()
	go printLetters()
	// Keep the main goroutine running for a while
	time.Sleep(time.Second * 3)
	fmt.Println("Hello! Welcome to go routines")
}
