package main

import "fmt"

func main() {
	var mat [1000][1000]int

	for {
		var x1, y1, x2, y2 int
		n, err := fmt.Scanf("%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		fmt.Println(n, err)
		if err != nil {
			break
		}
		fmt.Println(x1, y1, x2, y2)
		dx, dy := 0, 0
		if x1 == x2 {
			dy = 1
			if y2 < y1 {
				y1, y2 = y2, y1
			}
		} else if y1 == y2 {
			dx = 1
			if x2 < x1 {
				x1, x2 = x2, x1
			}
		} else {
			if x2 > x1 {
				dx = 1
			} else {
				dx = -1
			}
			if y2 > y1 {
				dy = 1
			} else {
				dy = -1
			}
		}
		for i, j := x1, y1; ; {
			fmt.Println(i, j)
			mat[i][j]++
			if i == x2 && j == y2 {
				break
			}
			i += dx
			j += dy
		}
	}
	ans := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if mat[i][j] > 1 {
				ans++
			}
		}
	}
	fmt.Println("ans=", ans)
}
