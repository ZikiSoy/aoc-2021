package main

import "fmt"

func main() {
	cur := make([][]byte, 0)
	nxt := make([][]byte, 0)
	tmp := make([][]byte, 0)
	n, m := 0, 0
	for {
		var str string
		_, err := fmt.Scanf("%s", &str)
		if err != nil {
			break
		}
		n++
		m = len(str)
		stk := make([]byte, len(str))

		for i := 0; i < len(str); i++ {
			stk[i] = str[i] - '0'
		}
		cur = append(cur, stk)
		nxt = append(nxt, stk)
		tmp = append(tmp, stk)
	}
	ans := 0
	fmt.Println(cur)
	for i := 0; i < 100; i++ {
		flash := false
		for x := 0; x < n; x++ {
			for y := 0; y < m; y++ {
				nxt[x][y] = cur[x][y] + 1
				if nxt[x][y] == 10 {
					flash = true
				}
			}
		}
		for flash {
			flash = false
			for x := 0; x < n; x++ {
				for y := 0; y < m; y++ {
					if nxt[x][y] == 10 {
						nxt[x][y] = 11
						ans++
						for u := -1; u <= 1; u++ {
							for v := -1; v <= 1; v++ {
								tx, ty := x+u, y+v
								if tx >= 0 && tx < n && ty >= 0 && ty < m && nxt[tx][ty] < 10 {
									nxt[tx][ty]++
								}
							}
						}
					}
				}
			}
			for x := 0; x < n; x++ {
				for y := 0; y < m; y++ {
					if nxt[x][y] == 10 {
						flash = true
					}
				}
			}
		}
		for x := 0; x < n; x++ {
			for y := 0; y < m; y++ {
				if nxt[x][y] == 11 {
					nxt[x][y] = 0
				}
				cur[x][y] = nxt[x][y]
			}
		}
	}
	fmt.Println("ans=", ans)
}
