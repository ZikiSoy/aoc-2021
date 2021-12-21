package main

import "fmt"
import "sort"

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func main() {
	num := make([]int, 0)

	for {
		var x1 int
		_, err := fmt.Scanf("%d", &x1)
		if err != nil {
			break
		}
		num = append(num, x1)
	}
	sort.Ints(num)
	fmt.Println(num)
	mid := len(num) / 2
	ans := 0
	for i := 0; i < len(num); i++ {
		ans += abs(num[i], num[mid])
	}
	fmt.Println("ans=", ans)
}
