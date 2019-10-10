package calc

import "fmt"

func Plus(a, b int) int {
	return a + b
}

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	} else {
		return 1
	}
}

func Div(a, b int) (float64, error) {
	if b == 0 {
		// 0で割り算しようとしているときはエラーにする
		// エラーメッセージは小文字で始める(see https://github.com/golang/go/wiki/CodeReviewComments#error-strings)
		return 0, fmt.Errorf("zero division")
	}

	return float64(a) / float64(b), nil
}
