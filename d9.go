package main

import "fmt"

func main() {
	num := make([]string, 0)

	for {
		var str string
		_, err := fmt.Scanf("%s", &str)
		if err != nil {
			break
		}
		num = append(num, str)
	}
	n, m := len(num), len(num[0])
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i == 0 || num[i][j] < num[i-1][j]) &&
				(i == n-1 || num[i][j] < num[i+1][j]) &&
				(j == 0 || num[i][j] < num[i][j-1]) &&
				(j == m-1 || num[i][j] < num[i][j+1]) {
				ans += int(num[i][j] + 1 - '0')
			}
		}
	}
	fmt.Println("ans=", ans)
}
