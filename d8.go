package main

import "fmt"
import "strings"

func main() {
	ans := 0
	for {
		var str string
		_, err := fmt.Scanf("%s", &str)
		if err != nil {
			break
		}
		a := strings.Split(str, ",")
		for i := 11; i < 15; i++ {
			if len(a[i]) == 2 || len(a[i]) == 3 || len(a[i]) == 4 || len(a[i]) == 7 {
				ans++
			}
		}
	}
	fmt.Println("ans=", ans)
}
