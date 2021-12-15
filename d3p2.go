package main

import "fmt"
import "strconv"

func main() {
	var bin string
	arr := make([]string, 0)
	l := 0
	for {
		_, err := fmt.Scan(&bin)
		if err != nil {
			break
		}
		arr = append(arr, bin)
	}
	l = len(bin)
	cur := arr
	for i := 0; i < l; i++ {
		cnt0, cnt1 := 0, 0
		for j := 0; j < len(cur); j++ {
			if cur[j][i] == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}
		nxt := make([]string, 0)
		for j := 0; j < len(cur); j++ {
			if cnt1 >= cnt0 {
				if cur[j][i] == '1' {
					nxt = append(nxt, cur[j])
				}
			} else {
				if cur[j][i] == '0' {
					nxt = append(nxt, cur[j])
				}
			}
		}
		cur = nxt
	}
	gamma := cur[0]

	cur = arr
	for i := 0; i < l; i++ {
		cnt0, cnt1 := 0, 0
		for j := 0; j < len(cur); j++ {
			if cur[j][i] == '0' {
				cnt0++
			} else {
				cnt1++
			}
		}
		nxt := make([]string, 0)
		for j := 0; j < len(cur); j++ {
			if cnt0 <= cnt1 {
				if cur[j][i] == '0' {
					nxt = append(nxt, cur[j])
				}
			} else {
				if cur[j][i] == '1' {
					nxt = append(nxt, cur[j])
				}
			}
		}
		cur = nxt
		if len(cur) == 1 {
			break
		}
		fmt.Println(cur)
	}
	eps := cur[0]
	a1, _ := strconv.ParseInt(gamma, 2, 64)
	a2, _ := strconv.ParseInt(eps, 2, 64)
	fmt.Println("ans=", gamma, eps, a1, a2, a1*a2)
}
