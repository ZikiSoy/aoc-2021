package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
		dis = append(dis, make([]int, len(a)))
		vis = append(vis, make([]bool, len(a)))
	}
	n, m := len(mp), len(mp[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dis[i][j] = 200000000000
		}
	}
	dis[0][0] = 0
	for dis[n-1][m-1] >= 200000000000 {
		cur := 200000000001
		var tx, ty int
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
		if tx > 0 {
			dis[tx-1][ty] = min(dis[tx-1][ty], dis[tx][ty]+int(mp[tx-1][ty]-'0'))
		}
		if tx+1 < n {
			dis[tx+1][ty] = min(dis[tx+1][ty], dis[tx][ty]+int(mp[tx+1][ty]-'0'))
		}
		if ty > 0 {
			dis[tx][ty-1] = min(dis[tx][ty-1], dis[tx][ty]+int(mp[tx][ty-1]-'0'))
		}
		if ty+1 < m {
			dis[tx][ty+1] = min(dis[tx][ty+1], dis[tx][ty]+int(mp[tx][ty+1]-'0'))
		}
	}
	fmt.Println(n, m)

	ans := dis[n-1][m-1]
	fmt.Println("ans=", ans)
}
