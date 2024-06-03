package main

import (
	"fmt"
	"time"
)

func sayHello() int{
    fmt.Println("Hello, goroutine!")
		return 9
}

func main() {
    // Create a goroutine
    go sayHello()

    // Print from main function
    fmt.Println("Hello from main!")

    // Sleep for a while to allow goroutine to execute
    time.Sleep(1 * time.Millisecond)
}
