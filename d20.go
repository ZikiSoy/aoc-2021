package main

import "fmt"
import "strconv"

func main() {
	var code string
	fmt.Scan(&code)
	padding_c := 55
	mp := make([]string, padding_c)
	padding := ""
	for i := 0; i < padding_c; i++ {
		padding += "."
	}

	for cnt := 0; ; cnt++ {
		var (
			str string
		)
		_, err := fmt.Scan(&str)
		if err != nil {
			break
		}
		str = padding + str + padding
		mp = append(mp, str)
	}
	empty := ""
	for i := 0; i < len(mp[padding_c]); i++ {
		empty = empty + "."
	}
	for i := 0; i < padding_c; i++ {
		mp[i] = empty
		mp = append(mp, empty)
	}
	n, m := len(mp), len(mp[0])
	for rep := 0; rep < 50; rep++ {
		nxt := make([]string, n)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				var ch byte
				ch = '.'
				if rep%2 == 0 {
					ch = code[0]
				}
				if i > 0 && i+1 < n && j > 0 && j+1 < m {
					bin := ""
					for x := i - 1; x <= i+1; x++ {
						for y := j - 1; y <= j+1; y++ {
							if mp[x][y] == '.' {
								bin += "0"
							} else {
								bin += "1"
							}
						}
					}
					loc, _ := strconv.ParseInt(bin, 2, 64)
					ch = code[loc]
				}
				nxt[i] += string(ch)
			}
		}
		mp = nxt
		fmt.Println(rep)
		for i := 0; i < n; i++ {
			// fmt.Println(mp[i])
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mp[i][j] == '#' {
				ans++
			}
		}
	}
	fmt.Println("ans=", ans)
}
