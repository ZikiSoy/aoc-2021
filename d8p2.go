package main

import "fmt"
import "strings"
import "sort"

func main() {
	ans := 0
	mapping := [10]string{
		"abcefg",
		"cf",
		"acdeg",
		"acdfg",
		"bcdf",
		"abdfg",
		"abdefg",
		"acf",
		"abcdefg",
		"abcdfg",
	}
	for {
		var str string
		_, err := fmt.Scanf("%s", &str)
		if err != nil {
			break
		}
		a := strings.Split(str, ",")
		mp := make(map[byte]byte)
		for i := 0; i < 10; i++ {
			if len(a[i]) == 2 {
				for j := 0; j < len(a[i]); j++ {
					cnt := 0
					for k := 0; k < 10; k++ {
						if strings.Contains(a[k], string(a[i][j])) {
							cnt++
						}
					}
					if cnt == 8 {
						mp[a[i][j]] = 'c'
					} else if cnt == 9 {
						mp[a[i][j]] = 'f'
					} else {
						fmt.Println("e1")
					}
				}
			}
		}
		for i := 0; i < 10; i++ {
			if len(a[i]) == 3 {
				for j := 0; j < len(a[i]); j++ {
					_, ok := mp[a[i][j]]
					if !ok {
						mp[a[i][j]] = 'a'
					}
				}
			}
		}
		for i := 0; i < 10; i++ {
			if len(a[i]) == 4 {
				for j := 0; j < len(a[i]); j++ {
					_, ok := mp[a[i][j]]
					if ok {
						continue
					}
					cnt := 0
					for k := 0; k < 10; k++ {
						if strings.Contains(a[k], string(a[i][j])) {
							cnt++
						}
					}
					if cnt == 6 {
						mp[a[i][j]] = 'b'
					} else if cnt == 7 {
						mp[a[i][j]] = 'd'
					} else {
						fmt.Println("e2")
					}
				}
			}
		}
		for i := 0; i < 10; i++ {
			for j := 0; j < len(a[i]); j++ {
				_, ok := mp[a[i][j]]
				if ok {
					continue
				}
				cnt := 0
				for k := 0; k < 10; k++ {
					if strings.Contains(a[k], string(a[i][j])) {
						cnt++
					}
				}
				if cnt == 7 {
					mp[a[i][j]] = 'g'
				} else if cnt == 4 {
					mp[a[i][j]] = 'e'
				} else {
					fmt.Println("e3")
				}
			}
		}
		temp := 0
		for i := 11; i < 15; i++ {
			s := strings.Split(a[i], "")
			for j := 0; j < len(s); j++ {
				s[j] = string(mp[s[j][0]])
			}
			sort.Strings(s)
			str := strings.Join(s, "")
			digit := -1
			for k := 0; k < 10; k++ {
				if str == mapping[k] {
					digit = k
				}
			}
			if digit != -1 {
				temp = temp*10 + digit
			} else {
				fmt.Println("not found ", digit)
			}
		}
		fmt.Println(temp)
		ans += temp
	}
	fmt.Println("ans=", ans)
}
