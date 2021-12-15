package main

import "fmt"

func main() {
	var i int
	ans := 0
	numbers := make([]int, 0)
	for {
		_, err := fmt.Scan(&i)
		if err != nil {
			break
		}
		numbers = append(numbers, i)
	}
	for i = 3; i < len(numbers); i++ {
		if numbers[i] > numbers[i-3] {
			ans++
		}
	}
	fmt.Println("ans=", ans)
}
