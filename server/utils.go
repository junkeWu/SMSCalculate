package server

import (
	"math"
	"regexp"
	"unicode"
)

func CalculateContent(content string) int {
	if content == "" {
		return 0
	}
	// 传入字符串根据纯英文和非纯英文计算（汉字、字母、数字、标点符号（不区分全角/半角）以及空格等，都按1个字计算）
	char := isChineseChar(content)
	if char {
		return int(math.Ceil(float64(len(content)) / 70))
	}
	return int(math.Ceil(float64(len(content)) / 160))
}

// isChineseChar 是否包含中文或者中文字符
func isChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[。；，：“”（）、？《》]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
