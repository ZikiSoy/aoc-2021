package main

import "fmt"
import "sort"

func dfs(vis *[][]bool, num *[]string, n int, m int, x int, y int) int {
	if x < 0 || y < 0 || x >= n || y >= m {
		return 0
	}
	if (*vis)[x][y] {
		return 0
	}
	if (*num)[x][y] == '9' {
		return 0
	}
	(*vis)[x][y] = true
	ret := 1
	ret += dfs(vis, num, n, m, x-1, y)
	ret += dfs(vis, num, n, m, x+1, y)
	ret += dfs(vis, num, n, m, x, y-1)
	ret += dfs(vis, num, n, m, x, y+1)
	return ret
}

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
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		vis[i] = make([]bool, m)
	}
	fmt.Println(len(vis), len(vis[0]))
	ans := 0
	anss := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i == 0 || num[i][j] < num[i-1][j]) &&
				(i == n-1 || num[i][j] < num[i+1][j]) &&
				(j == 0 || num[i][j] < num[i][j-1]) &&
				(j == m-1 || num[i][j] < num[i][j+1]) {
				if !vis[i][j] {
					anss = append(anss, dfs(&vis, &num, n, m, i, j))
				}
			}
		}
	}
	sort.Ints(anss)
	ans = anss[len(anss)-1] * anss[len(anss)-2] * anss[len(anss)-3]
	fmt.Println("ans=", ans)
}
