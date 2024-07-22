package main

import "fmt"

func main() {
	x := 10
	y := 5000
	if x > 5 && y < 100 {
		fmt.Println("Greater than 5...")
	}
	fmt.Println("Hello Go...")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
