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
	init_digit(79997931969649)
	//
	// digit[3] - 2 == digit[4]
	// digit[5] + 6 == digit[6]
	// digit[2] - 8 == digit[7]
	// digit[0] + 2 == digit[13]
	// digit[1] - 5 == digit[12]
	// digit[9] + 3 == digit[10]
	// digit[8] - 3 == digit[11]
	// 01234567890123
	// 16931171414113
	// 79997391969649
	var w, x, y, z int
	w = read()
	x = 1
	y = w + 8
	z = w + 8

	w = read()
	z = z * 26
	y = w + 8
	z = (digit[0]+8)*26 + digit[1] + 8
	fmt.Println("z=", z)

	w = read()
	y = w + 3
	z = ((digit[0]+8)*26+digit[1]+8)*26 + digit[2] + 3
	fmt.Println("z=", z)

	w = read()
	y = w + 10
	z = (((digit[0]+8)*26+digit[1]+8)*26+digit[2]+3)*26 + digit[3] + 10
	fmt.Println("z=", z)

	w = read()
	// digit[3] - 2 == digit[4]
	x = digit[3] - 2
	z = z / 26
	x = bti(x == w)
	x = bti(x == 0)
	y = y * 0
	y = y + 25
	y = y * x
	y = y + 1
	z = z * y

	y = w + 8

	z = z + y
	z = ((digit[0]+8)*26+digit[1]+8)*26 + digit[2] + 3
	fmt.Println("z=", z)

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
	y = w + 8
	y = y * x
	z = z + y
	z = (((digit[0]+8)*26+digit[1]+8)*26+digit[2]+3)*26 + digit[5] + 8
	fmt.Println("z=", z)

	w = read()
	// digit[5] + 6 == digit[6]
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
	y = w + 8
	y = y * x
	z = z + y
	z = ((digit[0]+8)*26+digit[1]+8)*26 + digit[2] + 3
	fmt.Println("z=", z)

	w = read()
	// digit[2] + 3 - 11 == digit[7]
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
	y = w + 5
	y = y * x
	z = z + y
	z = (digit[0]+8)*26 + digit[1] + 8
	fmt.Println("z=", z)

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
	y = w + 9
	y = y * x
	z = z + y
	z = ((digit[0]+8)*26+digit[1]+8)*26 + digit[8] + 9
	fmt.Println("z=", z)

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
	y = w + 3
	y = y * x
	z = z + y
	z = (((digit[0]+8)*26+digit[1]+8)*26+digit[8]+9)*26 + digit[9] + 3
	fmt.Println("z=", z)

	w = read()
	// digit[9] + 3 == digit[10]
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
	y = w + 4
	y = y * x
	z = z + y
	z = ((digit[0]+8)*26+digit[1]+8)*26 + digit[8] + 9
	fmt.Println("z=", z)

	w = read()
	// digit[8] + 9 - 12 == digit[11]
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
	y = w + 9
	y = y * x
	z = z + y
	z = (digit[0]+8)*26 + digit[1] + 8
	fmt.Println("z=", z)

	w = read()
	// digit[1] + 8 - 13 == digit[12]
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
	y = w + 2
	y = y * x
	z = z + y
	z = digit[0] + 8
	fmt.Println("z=", z)

	w = read()
	//   digit[0] + 8 - 6 == digit[13]
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
	y = w + 7
	y = y * x
	z = z + y
	z = 0
	fmt.Println(w, x, y, z)
}
