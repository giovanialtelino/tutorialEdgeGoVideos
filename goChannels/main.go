package main

import (
	"fmt"
	"time"
)

func SendValue(c chan string) {
	fmt.Println("exec go")
	time.Sleep(1 * time.Second)
	c <- "Hello World"
	fmt.Println("finished the routine")
}

func main() {
	fmt.Println("channels")

	values := make(chan string, 2)
	defer close(values)

	//calling channels
	go SendValue(values)
	go SendValue(values)

	value := <-values

	fmt.Println(value)

}
