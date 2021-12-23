package main

import "fmt"

type key struct {
	x, y int
}

func main() {
	vis := map[key]int{}
	for {
		var a, b int
		_, err := fmt.Scanf("%d,%d", &a, &b)
		if err != nil {
			fmt.Println(err)
			break
		}
		vis[key{a, b}] = 1
	}
	fmt.Println(vis)
	for {
		var a string
		var b int
		_, err := fmt.Scanf("fold along %1s=%d", &a, &b)
		if err != nil {
			fmt.Println(err)
			break
		}
		nxt := map[key]int{}
		for k, _ := range vis {
			if a == "x" {
				tx := k.x
				if tx > b {
					tx = 2*b - tx
				}
				nxt[key{tx, k.y}]++
			} else {
				tx := k.y
				if tx > b {
					tx = 2*b - tx
				}
				nxt[key{k.x, tx}]++
			}
		}
		// fmt.Println(nxt)
		fmt.Println(a, b, len(nxt))
		vis = nxt
	}
	fmt.Println("start")
	for j := 0; j < 10; j++ {
		for i := 0; i < 40; i++ {
			if _, err := vis[key{i, j}]; err {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
	ans := 0
	fmt.Println(vis)
	fmt.Println("ans=", ans)
}
