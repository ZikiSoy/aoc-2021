package main

import (
	"fmt"
	"sort"
)

type node struct {
	ins                    string
	x1, x2, y1, y2, z1, z2 int
}

func unique(b []int) []int {
	var ret []int
	for i := 0; i < len(b); i++ {
		if i == 0 || b[i] != b[i-1] {
			ret = append(ret, b[i])
		}
	}
	return ret
}

func idx(b []int, val int) int {
	for i := 0; i < len(b); i++ {
		if b[i] == val {
			return i
		}
	}
	return len(b)
}

type point struct {
	x, y, z int
}

func main() {
	var ins []node
	var allx, ally, allz []int
	mp := map[point]bool{}
	for cnt := 0; ; cnt++ {
		var (
			str string
		)
		_, err := fmt.Scan(&str)
		if err != nil {
			break
		}
		var x1, x2, y1, y2, z1, z2 int
		fmt.Scanf("x=%d..%d,y=%d..%d,z=%d..%d", &x1, &x2, &y1, &y2, &z1, &z2)
		fmt.Println(x1, x2, y1, y2, z1, z2)
		if x1 > 50 || x1 < -50 {
			// continue
		}
		x2++
		y2++
		z2++
		allx = append(allx, x1, x2)
		ally = append(ally, y1, y2)
		allz = append(allz, z1, z2)
		ins = append(ins, node{str, x1, x2, y1, y2, z1, z2})
	}
	sort.Ints(allx)
	allx = unique(allx)
	sort.Ints(ally)
	ally = unique(ally)
	sort.Ints(allz)
	allz = unique(allz)
	fmt.Println(allx)
	fmt.Println(len(allx), len(ally), len(allz))
	for i := 0; i < len(ins); i++ {
		ins[i] = node{ins[i].ins,
			idx(allx, ins[i].x1), idx(allx, ins[i].x2),
			idx(ally, ins[i].y1), idx(ally, ins[i].y2),
			idx(allz, ins[i].z1), idx(allz, ins[i].z2)}
	}
	ans := 0
	var mp2 [900][900][900]bool
	for i := 0; i < len(ins); i++ {
		for x := ins[i].x1; x < ins[i].x2; x++ {
			for y := ins[i].y1; y < ins[i].y2; y++ {
				for z := ins[i].z1; z < ins[i].z2; z++ {
					//ans++

					if ins[i].ins == "on" {
						// mp[point{x, y, z}] = true
						mp2[x][y][z] = true
					} else {
						// mp[point{x, y, z}] = false
						mp2[x][y][z] = false
					}

				}
			}
		}
	}
	for k, v := range mp {
		if v {
			ans += (allx[k.x+1] - allx[k.x]) * (ally[k.y+1] - ally[k.y]) * (allz[k.z+1] - allz[k.z])
		}
	}
	for x := 0; x < len(allx); x++ {
		for y := 0; y < len(ally); y++ {
			for z := 0; z < len(allz); z++ {
				if mp2[x][y][z] {
					ans += (allx[x+1] - allx[x]) * (ally[y+1] - ally[y]) * (allz[z+1] - allz[z])
				}
			}
		}
	}
	fmt.Println("ans=", ans)
}
