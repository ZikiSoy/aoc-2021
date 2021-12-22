package main

import "fmt"

func dfs(node string, mp map[string]([]string), vis map[string]bool, extra int) int {
	if node == "end" {
		fmt.Println(node)
		return 1
	}
	used := false
	if node[0] >= 'a' && node[0] <= 'z' {
		if vis[node] {
			if extra <= 0 || node == "start" {
				return 0
			}
			extra = 0
			used = true
		}
		vis[node] = true
	}
	fmt.Println(node)
	ret := 0
	for _, nxt := range mp[node] {
		ret += dfs(nxt, mp, vis, extra)
	}
	if node[0] >= 'a' && node[0] <= 'z' {
		if !used {
			vis[node] = false
		}
	}
	return ret
}
func main() {
	mp := map[string]([]string){}
	vis := map[string]bool{}
	for {
		var a, b string
		_, err := fmt.Scanf("%s %s", &a, &b)
		if err != nil {
			break
		}
		fmt.Println(a, b)
		mp[a] = append(mp[a], b)
		mp[b] = append(mp[b], a)
		vis[a] = false
		vis[b] = false
	}
	ans := dfs("start", mp, vis, 1)
	fmt.Println(mp)
	fmt.Println("ans=", ans)
}
