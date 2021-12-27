package main

import "fmt"

func main() {
	var dp [2][22][11][22][11]int
	start1, start2 := 7, 8
	dp[0][0][start1][0][start2] = 1
	var win [2]int
	for s1 := 0; s1 < 21; s1++ {
		for s2 := 0; s2 < 21; s2++ {
			for p1 := 1; p1 <= 10; p1++ {
				for p2 := 1; p2 <= 10; p2++ {
					for pl := 0; pl < 2; pl++ {
						c := dp[pl][s1][p1][s2][p2]
						if c == 0 {
							continue
						}
						for r := 0; r < 27; r++ {
							r1, r2, r3 := r%3, r/3%3, r/9
							if pl == 0 {
								nxtp := (p1+r1+r2+r3+3-1)%10 + 1
								nxts := s1 + nxtp
								if nxts >= 21 {
									win[0] += c
								} else {
									dp[1][nxts][nxtp][s2][p2] += c
								}
							} else {
								nxtp := (p2+r1+r2+r3+3-1)%10 + 1
								nxts := s2 + nxtp
								if nxts >= 21 {
									win[1] += c
								} else {
									dp[0][s1][p1][nxts][nxtp] += c
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println("ans=", win)
}
