package main

import "fmt"

type node struct {
	x, y, z int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func manh(a, b node) int {
	return abs(a.x-b.x) + abs(a.y-b.y) + abs(a.z-b.z)
}

func main() {
	var all [][]node
	var bec []node
	for cnt := 0; ; cnt++ {
		var (
			str string
		)
		_, err := fmt.Scan(&str)
		if err != nil || str[:3] == "---" {
			if len(bec) > 0 {
				all = append(all, bec)
				bec = make([]node, 0)
			}
			if err != nil {
				fmt.Println(err)
				break
			} else {
				continue
			}
		}
		var x, y, z int
		fmt.Sscanf(str, "%d,%d,%d", &x, &y, &z)
		bec = append(bec, node{x, y, z})
	}
	var dis []node
	c := 1
	dis = append(dis, node{0, 0, 0})
	fullmap := map[node]int{}
	fmt.Println(len(all), all)
	for i := 0; i < c; i++ {
		for j := c; j < len(all); j++ {
			fmt.Println(i, j)
			for t := 0; t < 8*6; t++ {
				tmp := make([]node, len(all[j]))
				copy(tmp, all[j])
				// fmt.Println("tmp=", tmp)
				t1 := t / 6
				t2 := t % 6
				for k := 0; k < len(tmp); k++ {
					if (t1 & 4) > 0 {
						tmp[k].x *= -1
					}
					if (t1 & 2) > 0 {
						tmp[k].y *= -1
					}
					if (t1 & 1) > 0 {
						tmp[k].z *= -1
					}
					switch t2 % 3 {
					case 1:
						tmp[k].x, tmp[k].y = tmp[k].y, tmp[k].x
					case 2:
						tmp[k].x, tmp[k].z = tmp[k].z, tmp[k].x
					}
					if t2/3 > 0 {
						tmp[k].y, tmp[k].z = tmp[k].z, tmp[k].y
					}
				}
				cmp := all[i]
				match := false
				var delta node
			MATCHED:
				for u := 0; u < len(cmp); u++ {
					for v := 0; v < len(tmp); v++ {
						cnt := map[node]int{}
						for s := 0; s < len(cmp); s++ {
							cnt[cmp[s]]++

						}
						for s := 0; s < len(tmp); s++ {
							// cnt[node{cmp[s].x - cmp[u].x, cmp[s].y - cmp[u].y, cmp[s].z - cmp[u].z}]++
							cnt[node{tmp[s].x - tmp[v].x + cmp[u].x, tmp[s].y - tmp[v].y + cmp[u].y, tmp[s].z - tmp[v].z + cmp[u].z}]++
						}
						if i == 1 && j == 4 {
							// fmt.Println("dbg: ", len(cmp), len(tmp), len(cnt))
						}
						if len(cmp)+len(tmp)-len(cnt) >= 12 {
							match = true
							newtmp := make([]node, len(tmp))
							for s := 0; s < len(tmp); s++ {
								// tmp[s].x += -tmp[v].x + cmp[u].x
								// tmp[s].y += -tmp[v].y + cmp[u].y
								// tmp[s].z += -tmp[v].z + cmp[u].z
								newtmp[s] = node{tmp[s].x - tmp[v].x + cmp[u].x, tmp[s].y - tmp[v].y + cmp[u].y, tmp[s].z - tmp[v].z + cmp[u].z}
							}
							delta = node{-tmp[v].x + cmp[u].x, -tmp[v].y + cmp[u].y, -tmp[v].z + cmp[u].z}
							tmp = newtmp
							fmt.Println(cnt)
							fmt.Println(cmp)
							fmt.Println(tmp)
							break MATCHED
						}
					}
				}
				if match {
					fmt.Println("matched", i, j, c, len(all[j]), len(tmp))
					// fmt.Println(tmp)
					all[j] = tmp
					if j != c {
						all[c], all[j] = all[j], all[c]
					}
					dis = append(dis, delta)
					c++
					break
				}
			}
		}
	}
	fmt.Println(all)
	for i := 0; i < len(all); i++ {
		for j := 0; j < len(all[i]); j++ {
			fullmap[all[i][j]]++
		}
	}
	ans := len(fullmap)
	fmt.Println(fullmap)
	manhatten := 0
	for i := 0; i < len(dis); i++ {
		for j := i + 1; j < len(dis); j++ {
			if manhatten < manh(dis[i], dis[j]) {
				manhatten = manh(dis[i], dis[j])
			}
		}
	}
	fmt.Println("ans=", ans, manhatten)
}
