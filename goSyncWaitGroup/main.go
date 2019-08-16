package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunc(wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println(("finished"))
	wg.Done()
}

func main() {
	fmt.Println(("go wait"))

	var wg sync.WaitGroup
	wg.Add(2)
	go myFunc(&wg)
	go myFunc(&wg)
	wg.Wait()

	fmt.Println("finished exec my go")

}
