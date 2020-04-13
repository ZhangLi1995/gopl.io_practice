package main

import "fmt"

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March",
		4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September",
		10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// append 函数
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	// 实现栈
	/*
		stack = append(stack, v)
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	*/

	// 保留顺序删除元素
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))

	// 不保留顺序删除元素
	s = []int{5, 6, 7, 8, 9}
	fmt.Println(remove2(s, 2))
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
