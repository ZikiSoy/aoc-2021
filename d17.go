package main

import "fmt"

var ans int

func main() {
	for {
		var (
			x1, x2, y1, y2 int
		)
		_, err := fmt.Scanf("%d%d%d%d", &x1, &x2, &y1, &y2)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(err)
		fmt.Println(x1, x2, y1, y2)
		ans := 0
		ans2 := 0
		for dx := 0; dx < 500; dx++ {
			for dy := -500; dy < 500; dy++ {
				vx, vy := dx, dy
				mxY := 0
				for cx, cy := 0, 0; ; {
					cx += vx
					cy += vy
					if mxY < cy {
						mxY = cy
					}

					if cx >= x1 && cx <= x2 && cy >= y1 && cy <= y2 {
						fmt.Println(dx, dy, mxY)
						ans2++
						if ans < mxY {
							ans = mxY
						}
						break
					}
					if vx == 0 {
						if cx < x1 || cx > x2 {
							break
						}
					}
					if cx > x2 {
						break
					}
					if vy < 0 && cy < y1 {
						break
					}
					vy--
					if vx > 0 {
						vx--
					}
				}
			}
		}

		fmt.Println("ans=", ans, ans2)
	}
}
