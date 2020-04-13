package main

import (
	"unicode"
	"unicode/utf8"
)

// WARN: 这里之所有不是传递的指针，是因为仅仅反转，不会改变切片的长度或者容量，因此变动是可以传回给原变量的

// 原地反转 slice 中的元素
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// slice 左移 n 个元素
func leftMove(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

// 练习4.3：重写函数 reverse，使用数组指针作为参数而不是 slice
func reverse2(s *[10]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 练习4.4：编写一个函数 rotate，实现一次遍历就可以完成元素旋转
func rotate(s []int, n int) []int {
	if n <= 0 || n >= len(s) {
		return s
	}
	r := make([]int, len(s))
	for i, v := range s {
		r[(len(s)-n+i)%len(s)] = v
	}
	return r
}

// 练习4.5：编写一个就地处理函数，用于去除 []string slice 中相邻的重复字符串元素
func dupSlice(s []string) []string {
	if len(s) <= 1 {
		return s
	}
	index := 0
	for i := 1; i < len(s); i++ {
		if s[index] != s[i] {
			index++
			s[index] = s[i]
		}
	}
	return s[:index+1]
}

// 练习4.6：编写一个就地处理函数，用于将一个 UTF-8 编码的字节 slice 中所有相邻的 Unicode 空白
// 字符（查看 unicode.IsSpace）缩减为一个ASCII 空白字符
func dupSlice2(bytes []byte) []byte {
	var i int
	for l := 0; l < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		l += size
		if unicode.IsSpace(r) {
			// byte(32) 表示 ASCII 码的空格
			if i > 0 && bytes[i-1] == byte(32) {
				copy(bytes[i:], bytes[i+size:])
			} else {
				bytes[i] = byte(32)
				copy(bytes[i+1:], bytes[i+size:])
				i++
			}
		} else {
			i += size
		}
	}
	return bytes[0:i]
}

// 练习4.7：修改函数 reverse，来翻转一个 UTF-8 编码的字符串中的字符元素，传入参数是该字符串对应
// 的字节 slice 类型（[]byte）。你可以做到不需要重新分配内存就实现该功能吗？
func reverse3(bytes []byte) []byte {
	var ret []byte
	for i := len(bytes); i > 0; {
		r, size := utf8.DecodeLastRune(bytes[:i])
		ret = append(ret, []byte(string(r))...)
		i -= size
	}
	return ret
}
