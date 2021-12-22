package main

import "fmt"

func main() {
	score := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	cls := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}

	ans := 0
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
						ans += score[str[i]]
						continue Line
					}
				}
			}
		}
	}
	fmt.Println("ans=", ans)
}
