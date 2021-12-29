package main

import "fmt"

var digit [15]int
var cnt int

var w, x, y, z int

func read() int {
	fmt.Println(w, x, y, z)
	ret := digit[cnt]
	cnt++
	return ret
}

func init_digit(num int) {
	for s := 13; s >= 0; s-- {
		digit[s] = num % 10
		num /= 10
	}
}
func bti(x bool) int {
	if x {
		return 1
	}
	return 0
}

func main() {
	// init_digit(79997391969649)
	init_digit(16931171414113)
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 14
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 8
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 13
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 8
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 13
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 3
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 12
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 10
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	fmt.Println("x=", x)
	z = z / 26
	x = x + -12
	fmt.Println("x=", x)
	fmt.Println("w=", w)
	x = bti(x == w)
	x = bti(x == 0)
	fmt.Println(x, z)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 8
	y = y * x
	z = z + y
	fmt.Println("z2", x, z)
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 12
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 8
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -2
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 8
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -11
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 5
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 13
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 9
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 1
	x = x + 14
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 3
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + 0
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 4
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -12
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 9
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -13
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 2
	y = y * x
	z = z + y
	w = read()
	x = x * 0
	x = x + z
	x = x % 26
	z = z / 26
	x = x + -6
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y
	y = y * 0
	y = y + w
	y = y + 7
	y = y * x
	z = z + y
	fmt.Println(w, x, y, z)
}
