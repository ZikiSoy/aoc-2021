package main

import "fmt"

func main() {
	var d int
	var direction string
	x := 0
	y := 0
	aim := 0
	for {
		_, err := fmt.Scan(&direction, &d)
		if err != nil {
			break
		}
		if direction == "forward" {
			x += d
			y += aim * d
		} else if direction == "up" {
			aim -= d
		} else {
			aim += d
		}
	}
	fmt.Println("ans=", x*y)
}
