package rev

// 原地反转 slice 中的元素
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j+1 {
		s[i], s[j] = s[j], s[i]
	}
}
