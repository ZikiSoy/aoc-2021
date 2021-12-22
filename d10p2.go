package main

import "fmt"
import "sort"

func main() {
	score := map[byte]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	cls := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	anss := make([]int, 0)
Line:
	for {
		var str string
		_, err := fmt.Scanf("%s", &str)
		if err != nil {
			break
		}
		stk := make([]byte, 0)
		for i := 0; i < len(str); i++ {
			switch str[i] {
			case '(', '[', '{', '<':
				stk = append(stk, str[i])
			default:
				{
					if len(stk) > 0 && stk[len(stk)-1] == cls[str[i]] {
						stk = stk[:len(stk)-1]
					} else {
						continue Line
					}
				}
			}
		}
		var tmp int
		for i := len(stk) - 1; i >= 0; i-- {
			tmp = tmp*5 + score[stk[i]]
		}
		anss = append(anss, tmp)
		fmt.Println("cap=", cap(anss))
	}
	sort.Slice(anss, func(i, j int) bool { return anss[i] < anss[j] })
	ans := anss[len(anss)/2]
	fmt.Println("ans=", ans)
}
