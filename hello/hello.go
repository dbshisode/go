package main

import (
	"fmt"

	"example.com/godemo"
)

func main() {
	//p := godemo.Params{2, 3}
	a := 2
	b := 3
	fmt.Println(godemo.Calculate(&a, &b, "addition"))
	fmt.Println(godemo.Calculate(&a, &b, "multiplication"))

	in := make(chan int)
	out := make(chan int)
	// We want to run a goroutine to multiply n by 2
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)

	go func() {
		in <- 1
		in <- 2
		in <- 3
		in <- 4
	}()

	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

	fmt.Println("Done")
	// We pause the program so that the `multiplyByTwo` goroutine
	// can finish and print the output before the code exits
	//time.Sleep(time.Second)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("inititalizing GoRoutine..")
	for {
		num := <-in
		result := num * 2
		out <- result
	}
	//time.Sleep(time.Second)
}
