package main

import "fmt"

var digit [15]int
var cnt int

func read() int {
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
	// init_digit(99999966993115)
	init_digit(99999969993999)
	var w, x, y, z int
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			for k := 1; k < 10; k++ {
				cnt = 0
				w, x, y, z := 0, 0, 0, 0
				digit[cnt] = i
				digit[1] = j
				digit[2] = k
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
				fmt.Println("s1", w, x, y, z)
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
				fmt.Println("s2", w, x, y, z)
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
				fmt.Println("s3", w, x, y, z)
			}
		}
	}
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
	y = y + 8
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

	y = 25*x + 1
	z = z * y
	y = (w + 2) * x
	z = z + y

	w = read()
	x = z % 26
	z = z / 26
	x = x + -6
	x = bti(x == w)
	x = bti(x == 0)

	y = 25*x + 1

	z = z * y

	y = w + 7
	y = y * x

	z = z + y
	fmt.Println(w, x, y, z)
}
