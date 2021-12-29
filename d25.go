package main

import "fmt"

func main() {
	var cur, nxt, down [][]byte
	for {
		var str string
		_, err := fmt.Scan(&str)
		if err != nil {
			break
		}
		cur = append(cur, []byte(str))
		nxt = append(nxt, []byte(str))
		down = append(down, []byte(str))
	}
	step := 0
	n, m := len(cur), len(cur[0])
	for {
		for i := 0; i < n; i++ {
			copy(nxt[i], cur[i])
		}
		if step < 10 {
			fmt.Println(step)
			for i := 0; i < n; i++ {
				fmt.Println(string(cur[i]))
			}
		}
		// move right
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				x, y := i, (j+1)%m
				if cur[x][y] == '.' && cur[i][j] == '>' {
					nxt[x][y] = '>'
					nxt[i][j] = '.'
				}
			}
		}
		for i := 0; i < n; i++ {
			copy(down[i], nxt[i])
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				x, y := (i+1)%n, j
				if nxt[x][y] == '.' && nxt[i][j] == 'v' {
					down[x][y] = 'v'
					down[i][j] = '.'
				}
			}
		}
		moved := false
	MOVED:
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				if down[i][j] != cur[i][j] {
					moved = true
					break MOVED
				}
			}
		}
		step++
		if moved {
			cur, down = down, cur
		} else {
			break
		}
	}
	ans := step
	fmt.Println("ans=", ans, n, m)
}
