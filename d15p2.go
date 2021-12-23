package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func getval(x, y int, mp []string) int {
	n, m := len(mp), len(mp[0])
	add := x/n + y/m
	dx, dy := x%n, y%m
	ret := int(mp[dx][dy]-'0') + add
	if ret > 9 {
		ret = ret + 1 - 10
	}
	return ret
}
func main() {
	mp := []string{}
	dis := make([][]int, 0)
	vis := make([][]bool, 0)

	for {
		var (
			a string
		)
		_, err := fmt.Scanf("%s", &a)
		if err != nil {
			fmt.Println(err)
			break
		}
		mp = append(mp, a)
		for i := 0; i < 5; i++ {
			dis = append(dis, make([]int, len(a)*5))
			vis = append(vis, make([]bool, len(a)*5))
		}
	}
	n, m := len(dis), len(dis[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dis[i][j] = 200000000000
		}
	}
	dis[0][0] = 0
	step := 0
	for dis[n-1][m-1] >= 200000000000 {
		cur := 200000000001
		var tx, ty int
		// use priority_queue
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if dis[i][j] < cur && !vis[i][j] {
					cur = dis[i][j]
					tx = i
					ty = j
				}
			}
		}
		vis[tx][ty] = true
		step++
		fmt.Println(step, tx, ty)
		if tx > 0 {
			dis[tx-1][ty] = min(dis[tx-1][ty], dis[tx][ty]+getval(tx-1, ty, mp))
		}
		if tx+1 < n {
			dis[tx+1][ty] = min(dis[tx+1][ty], dis[tx][ty]+getval(tx+1, ty, mp))
		}
		if ty > 0 {
			dis[tx][ty-1] = min(dis[tx][ty-1], dis[tx][ty]+getval(tx, ty-1, mp))
		}
		if ty+1 < m {
			dis[tx][ty+1] = min(dis[tx][ty+1], dis[tx][ty]+getval(tx, ty+1, mp))
		}
	}
	fmt.Println(n, m)

	ans := dis[n-1][m-1]
	fmt.Println("ans=", ans)
}
