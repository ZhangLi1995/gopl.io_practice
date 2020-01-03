package main

import (
	"fmt"
	"strconv"
)

func main() {
	//for i, r := range "Hello, 世界" {
	//	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	//}
	s := "i程序"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))

	// 整数转换为字符串
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	// 按照指定进位制格式化数字
	fmt.Println(strconv.FormatInt(int64(x), 2))

	// 将字符串转成整数
	xx, _ := strconv.Atoi("123")
	yy, _ := strconv.ParseInt("123", 10, 64)
	fmt.Println(xx, yy)

	// 常量的隐式转换类型
	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')
}
