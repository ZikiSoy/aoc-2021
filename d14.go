package main

import "fmt"
import "sort"

func main() {
	mp := map[string]string{}
	var initial string
	fmt.Scanf("%s", &initial)
	cur := map[string]int{}
	for i := 0; i < len(initial)-1; i++ {
		cur[initial[i:i+2]]++
	}

	for {
		var (
			a string
			b string
		)
		_, err := fmt.Scanf("%s -> %s", &a, &b)
		if err != nil {
			fmt.Println(err)
			break
		}
		mp[a] = b
	}
	for step := 0; step < 40; step++ {
		nxt := map[string]int{}
		for key, val := range cur {
			if l, err := mp[key]; err {
				nxt[string(key[0])+l] += val
				nxt[l+string(key[1])] += val
			} else {
				nxt[key]++
			}
		}
		cur = nxt
		fmt.Println(cur)
	}

	cnt := map[byte]int{}
	for key, val := range cur {
		cnt[key[1]] += val
	}
	cnt[initial[0]]++

	arr := make([]int, 0)
	for _, val := range cnt {
		arr = append(arr, val)
	}
	sort.Ints(arr)
	ans := arr[len(arr)-1] - arr[0]
	fmt.Println("ans=", ans)
}
