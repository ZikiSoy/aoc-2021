package main

import "fmt"
import "strconv"

var ans int

func parsePacket(bin string, tot int) (int, []int) {
	ret := []int{}
	var i, j int
	for i, j = 0, 0; i+7 < len(bin) && j < tot; j++ {
		fmt.Println(bin[i:])
		ver, _ := strconv.ParseInt(bin[i:i+3], 2, 64)
		ans += int(ver)
		i += 3
		typeId, _ := strconv.ParseInt(bin[i:i+3], 2, 64)
		fmt.Println("ver=", ver, " type=", typeId)
		i += 3
		if typeId == 4 {
			iter := 0
			for {
				val, _ := strconv.ParseInt(bin[i+1:i+5], 2, 64)
				fmt.Println(iter, val)
				i += 5
				iter = (iter << 4) + int(val)
				if bin[i-5] == '0' {
					ret = append(ret, iter)
					break
				}
			}
		} else {
			var arr []int
			var delta int
			if bin[i] == '0' {
				l, _ := strconv.ParseInt(bin[i+1:i+16], 2, 64)
				fmt.Println("t0 len=", l)
				delta, arr = parsePacket(bin[i+16:i+16+int(l)], 100000)
				i += 16 + int(l)
			} else {
				sz, _ := strconv.ParseInt(bin[i+1:i+12], 2, 64)
				fmt.Println("t1 sz=", int(sz))
				delta, arr = parsePacket(bin[i+12:], int(sz))
				i += 12 + delta
			}
			fmt.Println("arr = ", arr)
			var sum int
			switch typeId {
			case 0:
				sum = 0
				for _, v := range arr {
					sum += v
				}
			case 1:
				sum = 1
				for _, v := range arr {
					sum *= v
				}
			case 2:
				sum = arr[0]
				for _, v := range arr {
					if sum > v {
						sum = v
					}
				}
			case 3:
				sum = arr[0]
				for _, v := range arr {
					if sum < v {
						sum = v
					}
				}
			case 5:
				if arr[0] > arr[1] {
					sum = 1
				} else {
					sum = 0
				}
			case 6:
				if arr[0] < arr[1] {
					sum = 1
				} else {
					sum = 0
				}
			case 7:
				if arr[0] == arr[1] {
					sum = 1
				} else {
					sum = 0
				}
			}
			ret = append(ret, sum)
		}
	}
	fmt.Println("i, ret = ", i, ret)
	return i, ret
}

func main() {
	for {
		bin := ""
		var (
			a string
		)
		_, err := fmt.Scanf("%s", &a)
		if err != nil {
			fmt.Println(err)
			break
		}
		for i := 0; i < len(a); i++ {
			val, _ := strconv.ParseInt(a[i:i+1], 16, 64)
			bi := strconv.FormatInt(val, 2)
			bin += fmt.Sprintf("%04s", bi)
		}
		ans = 0
		fmt.Println(bin)
		_, arr := parsePacket(bin, 100000)

		fmt.Println("ans=", arr)
	}
}
