package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 函数向表示十进制数的非负整数的字符串中插入逗号
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

const (
	test1 = "12345"
	test2 = "-12345.12345"
	test3 = "asdfghjkl123"
	test4 = "asdflkjgh321"
	test5 = "asdfghjkl;'"
)

func main() {
	fmt.Println(comma(test1))
	fmt.Println(comma1(test1))
	fmt.Println(comma2(test2))
	fmt.Println(judge(test3, test4))
	fmt.Println(judge(test3, test5))
}

// 练习 3.10：编写一个非递归的 comma 函数，运用 bytes.Buffer，而不是简单的字符串拼接
func comma1(s string) string {
	var buf bytes.Buffer
	l := len(s)
	count := (len(s) - 1) / 3
	remain := l - count*3
	for i, j := 0, 3-remain; i < l; {
		if j%3 == 0 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
		i++
		j++
	}

	for i := len(s) - 1; i >= 0; i-- {

	}
	return buf.String()
}

// 练习 3.11：增强 comma 函数的功能，让其能够正确处理浮点数，以及带有可选正负号的数字
func comma2(s string) string {
	var buf bytes.Buffer
	var flag string
	if s[0] == '+' || s[0] == '-' {
		flag = string(s[0])
	}
	buf.WriteString(flag)
	stringArrays := strings.Split(s[1:], ".")
	intPart := comma1(stringArrays[0])
	buf.WriteString(intPart + ".")
	floatPart := comma1(stringArrays[1])
	buf.WriteString(floatPart)
	return buf.String()
}

// 练习 3.12：编写一个函数判断两个字符串是否同文异构，也就是，它们都含有相同的字符但是排列顺序不同
const ASCII = 256

func judge(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	bucket1 := make([]byte, ASCII)
	bucket2 := make([]byte, ASCII)
	l := len(s1)
	for i := 0; i < l; i++ {
		bucket1[s1[i]]++
		bucket2[s2[i]]++
	}
	for i := 0; i < ASCII; i++ {
		if bucket1[i] != bucket2[i] {
			return false
		}
	}
	return true
}
