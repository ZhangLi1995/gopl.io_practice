package main

import "fmt"

// 函数向表示十进制数的非负整数的字符串中插入逗号
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	fmt.Println(comma("12345"))
}
