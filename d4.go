package main

import "fmt"
import "strings"
import "strconv"

func main() {
	var seq string
	fmt.Scan(&seq)
	seqs := strings.Split(seq, ",")
	nums := make([]int, len(seqs))
	for i := 0; i < len(seqs); i++ {
		nums[i], _ = strconv.Atoi(seqs[i])
	}

	times := 0
	sc := 0
	for {
		var mat, yes [5][5]int
		_, err := fmt.Scan(&mat[0][0])
		if err != nil {
			break
		}
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if i != 0 || j != 0 {
					fmt.Scan(&mat[i][j])
				}
			}
		}
		for x := 0; x < len(nums); x++ {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if mat[i][j] == nums[x] {
						yes[i][j] = 1
					}
				}
			}

			ok := false
			for i := 0; i < 5; i++ {
				full := true
				for j := 0; j < 5; j++ {
					if yes[i][j] != 1 {
						full = false
					}
				}
				if full {
					ok = true
				}
				full = true
				for j := 0; j < 5; j++ {
					if yes[j][i] != 1 {
						full = false
					}
				}
				if full {
					ok = true
				}
			}
			if ok {
				score := 0
				for i := 0; i < 5; i++ {
					for j := 0; j < 5; j++ {
						if yes[i][j] != 1 {
							score += mat[i][j]
						}
					}
				}
				score *= nums[x]
				if x > times {
					times = x
					sc = score
				}
				fmt.Println(ok, yes, x, score)
				break
			}
		}
	}
	fmt.Println("ans=", times, sc)
}
