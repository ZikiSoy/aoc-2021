package main

import "fmt"

func main() {
	var num [9]int

	for {
		var x1 int
		_, err := fmt.Scanf("%d", &x1)
		if err != nil {
			break
		}
		num[x1]++
	}
	fmt.Println(num)
	ans := 0
	for i := 0; i < 256; i++ {
		var nxt [9]int
		nxt[8] = num[0]
		for j := 0; j < 8; j++ {
			nxt[j] = num[j+1]
		}
		nxt[6] += num[0]
		num = nxt
		ans = 0
		for j := 0; j < 9; j++ {
			ans += num[j]
		}
	}
	fmt.Println("ans=", ans)
}
