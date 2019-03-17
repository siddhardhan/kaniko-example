package main

import (
	"fmt"
	"time"
)

func main() {
	for {

		fmt.Println(time.Now().Format(time.RFC850) + " Hello Skaffold - this is Sid!")

		time.Sleep(time.Second * 5)
	}
}
