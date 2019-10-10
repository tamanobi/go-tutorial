// Package checkdigit は特定の文字列
package checkdigit

import (
	"regexp"
)

// Luhn は与えられた数字列がLuhnアルゴリズムによって数字列を検証する
// https://ja.wikipedia.org/wiki/Luhn%E3%82%A2%E3%83%AB%E3%82%B4%E3%83%AA%E3%82%BA%E3%83%A0
func Luhn(str string) bool {
	if r := regexp.MustCompile(`\d+`); r.MatchString(str) == false {
		return false
	}

	l := make([]int, len(str))
	for i, char := range str {
		d := int(char - '0')
		if i%2 == 0 {
			d *= 2
		}
		l = append(l, d)
	}
	total := 0
	for _, v := range l {
		total += v
	}
	return total%10 == 0
}
