package word2

import "unicode"

// IsPalindrome 判断一个字符串是否是回文串，忽略大小写以及非字母字符
func IsPalindrome(s string) bool {
	//var letters []rune
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
