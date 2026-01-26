package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'A'; i < 'A' + 10; i++ {
		fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers()
	printLetters()
}

func goPrint1() {
	go printNumbers()
	go printLetters()
}

func printNumbers2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%d ", i)
	}
}

func printLetters2() {
	for i := 'A'; i < 'A' + 10; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%c ", i)
	}
}

func goPrint2() {
	go printNumbers2()
	go printLetters2()
}

func print2() {
	printNumbers2()
	printLetters2()
}

func main() {
	print1()
	// fmt.Println("\n------------------------||------------------------\n")
	goPrint1()

	time.Sleep(1 * time.Second)
}