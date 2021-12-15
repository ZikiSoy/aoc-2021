package main

import "fmt"

func main() {
	var bin string
	var cnt0, cnt1 []int
	l := 0
	for {
		_, err := fmt.Scan(&bin)
		if err != nil {
			break
		}
		if l == 0 {
			l = len(bin)
			cnt0 = make([]int, l)
			cnt1 = make([]int, l)
		}
		for i := 0; i < l; i++ {
			if bin[i] == '1' {
				cnt1[i]++
			} else {
				cnt0[i]++
			}
		}
	}
	gamma, eps := 0, 0
	fmt.Println(cnt0, cnt1)
	for i := 0; i < l; i++ {
		if cnt1[l-1-i] > cnt0[l-1-i] {
			gamma += (1 << i)
		} else {
			eps += (1 << i)
		}
	}
	fmt.Println("ans=", gamma*eps)
}
