package main

import "fmt"

func main() {
	err := hub()
	if err != nil {
		fmt.Println(err)
	}
}
