package main

import (
	"fmt"
	"strconv"
)

type node struct {
	value int
	// 0, value
	// 1, read
	// 2, op
	tp       int
	lhs, rhs *node
	op       string
	min, max int
}

var digit []int

func fmtNode(a *node) string {
	if a.tp == 0 {
		return fmt.Sprint(a.value)
	}
	if a.tp == 1 {
		return fmt.Sprintf("read(%d)", a.value)
	}
	return "(" + fmtNode(a.lhs) + a.op + fmtNode(a.rhs) + ")"
}

func reduce(p *node) *node {
	if p.tp == 0 {
		return p
	}
	a := *p
	if a.tp == 1 {
		if digit[a.value] != 0 {
			a.value = digit[a.value]
			a.tp = 0
		}
	}
	if a.tp == 2 {
		a.lhs, a.rhs = reduce(a.lhs), reduce(a.rhs)
		if a.lhs.tp == 0 && a.rhs.tp == 0 {
			a.tp = 0
			switch a.op {
			case "*":
				a.value = a.lhs.value * a.rhs.value
			case "%":
				a.value = a.lhs.value % a.rhs.value
			case "+":
				a.value = a.lhs.value + a.rhs.value
			case "/":
				a.value = a.lhs.value / a.rhs.value
			case "==":
				a.value = 0
				if a.lhs.value == a.rhs.value {
					a.value = 1
				}
			}
		} else if a.op == "==" && (a.lhs.tp == 0 && a.rhs.tp == 1 ||
			a.lhs.tp == 1 && a.rhs.tp == 0) {
			if (a.lhs.tp == 0 && (a.lhs.value > 9 || a.lhs.value <= 0)) ||
				(a.rhs.tp == 0 && (a.rhs.value > 9 || a.rhs.value <= 0)) {
				a.value = 0
				a.tp = 0
			}
		} else if a.op == "*" && a.rhs.tp == 0 && a.rhs.value == 0 {
			a.tp = 0
			a.value = 0
		} else if a.op == "*" && a.lhs.tp == 0 && a.lhs.value == 0 {
			a.tp = 0
			a.value = 0
		} else if a.op == "%" && a.rhs.tp == 0 && a.rhs.value == 1 {
			a.tp = 0
			a.value = 0
		} else if a.op == "/" && a.rhs.tp == 0 && a.rhs.value == 1 {
			a = *a.lhs
		} else if a.op == "+" && a.rhs.tp == 0 && a.rhs.value == 0 {
			a = *a.lhs
		} else if a.op == "+" && a.lhs.tp == 0 && a.lhs.value == 0 {
			a = *a.rhs
		} else if a.op == "*" && a.rhs.tp == 0 && a.rhs.value == 1 {
			a = *a.lhs
		} else if a.op == "*" && a.lhs.tp == 0 && a.lhs.value == 1 {
			a = *a.rhs
		} else {
			s := fmtNode(&a)
			if s == "((((read(0)+8)%26)+13)==read(1))" {
				a.tp = 0
				a.value = 0
			} else if s == "((((((read(0)+8)*26)+(read(1)+8))%26)+13)==read(2))" {
				a.tp = 0
				a.value = 0
			} else if s == "((((((((read(0)+8)*26)+(read(1)+8))*26)+(read(2)+3))%26)+12)==read(3))" {
				a.tp = 0
				a.value = 0
			}
		}
	}
	return &a
}

func eval(a *node) (int, bool) {
	if a.tp == 0 {
		return a.value, false
	}
	if a.tp == 1 {
		return digit[a.value], false
	}
	if a.tp == 2 {
		lhs, e1 := eval(a.lhs)
		if e1 {
			return 0, e1
		}
		rhs, e2 := eval(a.rhs)
		if e2 {
			return 0, e2
		}
		switch a.op {
		case "*":
			return lhs * rhs, false
		case "%":
			if lhs < 0 || rhs <= 0 {
				return 0, true
			}
			return lhs % rhs, false
		case "+":
			return lhs + rhs, false
		case "/":
			if rhs == 0 {
				return 0, true
			}
			return lhs / rhs, false
		case "==":
			if lhs == rhs {
				return 1, false
			}
			return 0, false
		}
	}
	return -1, true
}

var found [20]int

func find(a *node) {
	if a.tp == 0 {
		return
	}
	if a.tp == 1 {
		// fmt.Printf("read(%d)", a.value)
		found[a.value] = a.value
		return
	}
	find(a.lhs)
	find(a.rhs)
}

func printNode(a *node) {
	if a.tp == 0 {
		fmt.Print(a.value)
		return
	}
	if a.tp == 1 {
		fmt.Printf("read(%d)", a.value)
		return
	}
	fmt.Print("(")
	printNode(a.lhs)
	fmt.Print(a.op)
	printNode(a.rhs)
	fmt.Print(")")
}

func find2(a *node, pos int) int {
	if a.tp == 0 {
		return -1
	}
	if a.tp == 1 {
		return -1
	}
	if a.lhs.tp == 1 && a.lhs.value == pos {
		if a.rhs.tp != 0 {
			fmt.Println("error")
		}
		return a.rhs.value
	}
	if a.rhs.tp == 1 && a.rhs.value == pos {
		if a.lhs.tp != 0 {
			fmt.Println("error2", digit)
			fmt.Println("error2", a.lhs.tp, a.lhs.value, a.rhs.tp, a.rhs.value)
			printNode(a)
			fmt.Println("error2")
		}
		return a.lhs.value
	}
	x, y := find2(a.lhs, pos), find2(a.rhs, pos)
	if x >= 0 && y >= 0 && x != y {
		fmt.Println("error3")
	}
	if x > y {
		return x
	}
	return y
}

func dfs(pos int, p *node) {
	if pos == 14 {
		ans, err := eval(p)
		fmt.Println(digit, ans)
		if !err && ans == 0 {
			fmt.Println("ans is ", digit)
		}
		return
	}
	if pos == 7 {
		fmt.Println("pos", pos)
		printNode(p)
		fmt.Println(digit)
	}
	x := find2(p, pos)
	digit[pos] = 9
	nxt := reduce(p)
	dfs(pos+1, nxt)
	if x > 0 {
		digit[pos] = x
		nxt := reduce(p)
		dfs(pos+1, nxt)
	}
	digit[pos] = 0
}

func main() {
	var arr [4]*node
	for i := 0; i < 4; i++ {
		arr[i] = &node{value: 0}
	}
	cnt := 0
	digit = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	// digit = []int{9, 9, 9, 9, 9, 9, 9, 0, 0, 0, 0, 0, 0, 0}
	for {
		var str, x, y string
		_, err := fmt.Scan(&str, &x)
		if err != nil {
			break
		}
		idx := x[0] - 'w'
		if str == "inp" {
			arr[idx] = &node{value: cnt, tp: 1, min: 1, max: 9}
			cnt++
		} else {
			fmt.Scan(&y)
			var op string
			switch str {
			case "mul":
				op = "*"
			case "mod":
				op = "%"
			case "add":
				op = "+"
			case "div":
				op = "/"
			case "eql":
				op = "=="
			}
			// fmt.Println(x, "=", x, op, y)
			num, err := strconv.Atoi(y)
			if err != nil {
				idx2 := y[0] - 'w'
				arr[idx] = &node{tp: 2, lhs: arr[idx], rhs: arr[idx2], op: op}
			} else {
				arr[idx] = &node{tp: 2, lhs: arr[idx], rhs: &node{value: num, min: num, max: num}, op: op}
			}
		}
		arr[idx] = reduce(arr[idx])
		fmt.Print(x, "=")
		printNode(arr[idx])
		fmt.Println()
	}

	// fmt.Print("y is ")
	// find(arr[3])
	// fmt.Println(found)
	// printNode(arr[3])
	// dfs(0, arr[3])
	for i := 0; i > 0; i-- {
		valid := true
		for j, k := i, 13; k >= 6; k-- {
			if j%10 == 0 {
				valid = false
				break
			}
			digit[k] = j % 10
			j /= 10
		}
		if !valid {
			continue
		}
		ret, err := eval(arr[3])
		if err == false && ret == 0 {
			fmt.Println("ans=", i)
			break
		} else {
			if i%100 == 99 {
				fmt.Println(i)
			}
		}
	}
	// printNode(arr[2])
}
