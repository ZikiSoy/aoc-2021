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
}

func reduce(a *node) {
	if a.tp == 0 {
		return
	}
	if a.tp == 2 {
		if a.lhs.tp == 0 && a.rhs.tp == 0 {
			a.tp = 0
		}
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
			a.value = a.lhs.value == a.rhs.value
		}
	}
	return
}
func printNode(a *node) {
	if a.tp == 0 {
		fmt.Print(a.value)
	}
	if a.tp == 1 {
		fmt.Printf("read(%d)", a.value)
	}
	fmt.Print("(")
	printNode(a.lhs)
	fmt.Print(op)
	printNode(a.rhs)
	fmt.Print(")")
}

func main() {
	arr[4] * node
	for i := 0; i < 4; i++ {
		arr[i] = &node{value: 0}
	}
	cnt := 0
	for {
		var str, x, y string
		_, err := fmt.Scan(&str)
		if err != nil {
			break
		}
		if str == "inp" {
			fmt.Scan(&x)
			idx := x[0] - 'w'
			arr[idx] = &node{value: cnt, tp: 1}
			cnt++
		} else {
			fmt.Scan(&x, &y)
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
			idx := x[0] - 'w'
			arr[idx] = &node{value: cnt, tp: 1}
			// fmt.Println(x, "=", x, op, y)
			num, err = strconv.Atoi(y)
			if err != nil {
				idx2 = y[0] - 'w'
				arr[idx] = &node{tp: 2, lhs: arr[idx], rhs: arr[idx2], op: op}
			} else {
				arr[idx] = &node{tp: 2, lhs: arr[idx], rhs: &node{value: num}, op: op}
			}
			reduce(arr)
		}
	}
	printNode(arr[2])
	ans := 0
	fmt.Println("ans=", ans)
}
