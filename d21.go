package main

import "fmt"

var d_num, rolled int

func dice() int {
	d_num += 1
	if d_num > 100 {
		d_num = 1
	}
	rolled++
	return d_num
}

func main() {
	p1, p2 := 7, 8
	s1, s2 := 0, 0
	for s1 < 1000 && s2 < 1000 {
		p1 = (dice()+dice()+dice()+p1-1)%10 + 1
		s1 += p1
		if s1 < 1000 && s2 < 1000 {
			p2 = (dice()+dice()+dice()+p2-1)%10 + 1
			s2 += p2
		}
		fmt.Println(p1, s1, p2, s2)

	}

	fmt.Println("ans=", rolled, s1, s2)
	fmt.Println("ans=", rolled*s1, rolled*s2)
}
