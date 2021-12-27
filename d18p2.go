package main

import "fmt"

type Node struct {
	val              int
	lhs, rhs, parent *Node
}

func parse(str string) (*Node, int) {
	ret := Node{-1, nil, nil, nil}
	tot := 0
	if str[0] == '[' {
		lhs, used := parse(str[1:])
		rhs, used2 := parse(str[1+used+1:])
		ret.lhs, ret.rhs = lhs, rhs
		tot = used + used2 + 3
		lhs.parent, rhs.parent = &ret, &ret
	} else {
		ret.val = int(str[0] - '0')
		tot = 1
	}
	return &ret, tot
}

func add(a, b *Node) *Node {
	ret := Node{-1, a, b, nil}
	a.parent, b.parent = &ret, &ret
	return &ret
}

func calc(a *Node) int {
	if a.val >= 0 {
		return a.val
	}
	return calc(a.lhs)*3 + calc(a.rhs)*2
}

func explode(a *Node, dep int) bool {
	if a.val >= 0 {
		return false
	}
	if dep == 4 {
		lv, rv := a.lhs.val, a.rhs.val
		a.val = 0
		a.lhs, a.rhs = nil, nil
		p := a.parent
		cur := a
		for p != nil && p.lhs == cur {
			p = p.parent
			cur = cur.parent
		}
		if p != nil {
			cur = p.lhs
			for cur.val == -1 {
				cur = cur.rhs
			}
			cur.val += lv
		}

		p = a.parent
		cur = a
		for p != nil && p.rhs == cur {
			p = p.parent
			cur = cur.parent
		}
		if p != nil {
			cur = p.rhs
			for cur.val == -1 {
				cur = cur.lhs
			}
			cur.val += rv
		}

		return true
	}
	return explode(a.lhs, dep+1) || explode(a.rhs, dep+1)
}

func split(a *Node) bool {
	if a.val >= 0 {
		if a.val > 9 {
			a.lhs = &Node{a.val / 2, nil, nil, a}
			a.rhs = &Node{a.val - a.val/2, nil, nil, a}
			a.val = -1
			return true
		}
		return false
	}
	return split(a.lhs) || split(a.rhs)
}

func work(a *Node) {
	for {
		if explode(a, 0) {
			continue
		}
		if split(a) {
			continue
		}
		return
	}
}
func show(a *Node) string {
	if a.val >= 0 {
		return fmt.Sprintf("%d", a.val)
	}
	return "[" + show(a.lhs) + "," + show(a.rhs) + "]"
}

func main() {
	var strs []string
	for cnt := 0; ; cnt++ {
		var (
			str string
		)
		_, err := fmt.Scanf("%s", &str)
		if err != nil {
			fmt.Println(err)
			break
		}
		strs = append(strs, str)
	}
	ans := 0
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(strs); j++ {
			if i != j {
				a, _ := parse(strs[i])
				b, _ := parse(strs[j])
				root := add(a, b)
				work(root)
				tmp := calc(root)
				if ans < tmp {
					ans = tmp
				}
			}
		}
	}
	fmt.Println("ans=", ans)
}
