package main

import (
	"fmt"
)

type state struct {
	hall [11]byte
	row  [4][4]byte
}

var vis map[state]int

var ans int

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func printState(s state) {
	fmt.Print(string(s.hall[:]), " ")
	for i := 0; i < 4; i++ {
		fmt.Print(string(s.row[i][:]), " ")
	}
	fmt.Println(vis[s])
}

func bfs(q []state) int {
	cost_per_step := []int{1, 10, 100, 1000}
	row_num := []int{2, 4, 6, 8}
	ans := -1
	run := 0
	for len(q) > 0 {
		now := q[0]
		run++
		cost, _ := vis[now]
		q = q[1:]
		end_state := true
		for i := 0; i < len(now.row); i++ {
			for j := 0; j < len(now.row[i]); j++ {
				if now.row[i][j] != byte('A'+i) {
					end_state = false
				}
			}
		}
		if end_state {
			if ans < 0 || ans > cost {
				ans = cost
				fmt.Println("ans:=", ans)
			}
			continue
		}

		for i := 0; i < len(now.hall); i++ {
			ch := now.hall[i]
			if ch != '.' {
				row := int(ch - 'A')
				var l []byte
				l = now.row[row][:]
				all_same := true
				pos := -1
				for j := len(l) - 1; j >= 0; j-- {
					if l[j] != '.' && l[j] != ch {
						all_same = false
					}
					if l[j] == '.' && pos == -1 {
						pos = j
					}
				}
				if all_same {
					delta := 1
					if row_num[row] < i {
						delta = -1
					}
					no_blocker := true
					for t := i + delta; t != row_num[row]; t += delta {
						if now.hall[t] != '.' {
							no_blocker = false
							break
						}
					}
					if no_blocker {
						nxt := now
						nxt.hall[i], nxt.row[row][pos] = nxt.row[row][pos], nxt.hall[i]
						pre_cost, lookup := vis[nxt]
						new_cost := cost + cost_per_step[row]*(abs(row_num[row]-i)+pos+1)
						if !lookup {
							vis[nxt] = new_cost
							q = append(q, nxt)
						} else if pre_cost > new_cost {
							vis[nxt] = new_cost
						}
					}
				}
			}
		}
		for i := 0; i < 4; i++ {
			end := 3
			for end >= 0 && now.row[i][end] == byte('A'+i) {
				end--
			}
			for j := 0; j <= end; j++ {
				if now.row[i][j] != '.' {
					tp := int(now.row[i][j] - 'A')
					for k, step := row_num[i], j+1; k >= 0; k-- {
						if k%2 == 0 && k >= 2 && k <= 8 {
							step++
							continue
						}
						if now.hall[k] != '.' {
							break
						}
						nxt := now
						nxt.row[i][j], nxt.hall[k] = nxt.hall[k], nxt.row[i][j]
						new_cost := cost + cost_per_step[tp]*step
						pre_cost, lookup := vis[nxt]
						if !lookup {
							vis[nxt] = new_cost
							q = append(q, nxt)
						} else if pre_cost > new_cost {
							vis[nxt] = new_cost
						}
						step++
					}
					for k, step := row_num[i], j+1; k < len(now.hall); k++ {
						if k%2 == 0 && k >= 2 && k <= 8 {
							step++
							continue
						}
						if now.hall[k] != '.' {
							break
						}
						nxt := now
						nxt.row[i][j], nxt.hall[k] = nxt.hall[k], nxt.row[i][j]
						new_cost := cost + cost_per_step[tp]*step
						pre_cost, lookup := vis[nxt]
						if !lookup {
							vis[nxt] = new_cost
							q = append(q, nxt)
						} else if pre_cost > new_cost {
							vis[nxt] = new_cost
						}
						step++
					}
					break
				}
			}
		}
	}
	return ans
}

func main() {
	var init state
	for i := 0; i < len(init.hall); i++ {
		init.hall[i] = '.'
	}
	// s := []string{"BDDA", "CCBD", "BBAC", "DACA"}
	// #############
	// #...........#
	// ###B#C#C#B###
	//   #D#C#B#A#
	//   #D#B#A#C#
	//   #D#D#A#A#
	//   #########
	s := []string{"BDDD", "CCBD", "CBAA", "BACA"}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			init.row[i][j] = s[i][j]
		}
	}
	q := []state{init}
	vis = map[state]int{}
	vis[init] = 0

	ans := bfs(q)
	fmt.Println("ans=", ans)
}
