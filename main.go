package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("The time now is " + time.Now().Format("2006-01-02 15:04:05"))
}