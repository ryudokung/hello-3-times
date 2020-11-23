package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += i
		fmt.Println("Helloworld!")
		time.Sleep(1 * time.Second)
	}
}