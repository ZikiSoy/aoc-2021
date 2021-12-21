package main

import "fmt"

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func main() {
	loc := make([]int, 2001)

	for {
		var x1 int
		_, err := fmt.Scanf("%d", &x1)
		if err != nil {
			break
		}
		loc[x1]++
	}
	var costl, costr [2001]int
	num := 0
	deltaCost := 0
	currentCost := 0
	for i := 0; i <= 2000; i++ {
		deltaCost += num
		currentCost += deltaCost
		costl[i] += currentCost
		num += loc[i]
	}
	deltaCost = 0
	currentCost = 0
	num = 0
	for i := 2000; i >= 0; i-- {
		deltaCost += num
		currentCost += deltaCost
		costr[i] += currentCost
		num += loc[i]
	}
	fmt.Println(costl)
	fmt.Println(costr)
	ans := 0
	for i := 0; i <= 2000; i++ {
		if ans == 0 || ans > costl[i]+costr[i] {
			ans = costl[i] + costr[i]
		}
	}
	fmt.Println("ans=", ans)
}
