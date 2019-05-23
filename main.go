package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello pity!")

		time.Sleep(time.Second * 20)
	}
}
