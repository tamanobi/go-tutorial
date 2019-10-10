package calc

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlus(t *testing.T) {
	got := Plus(1, 2)
	want := 1 + 2
	if got != want {
		t.Errorf(`got = %d, want = %d`, got, want)
	}
}

func TestAbsPlusOne(t *testing.T) {
	got := Abs(1)
	want := 1
	if got != want {
		t.Errorf(`got = %d, want = %d`, got, want)
	}
}

func TestAbsMinusOne(t *testing.T) {
	t.Skip("未実装")
	// 実装してみよう
}

func TestDivZeroDivision(t *testing.T) {
	// テーブルドリブンテストで入出力を表のように管理できる
	// https://github.com/golang/go/wiki/CodeReviewComments#useful-test-failures
	var tests = []struct {
		in1 int
		in2 int
		out float64
		err error
	}{
		{2, 2, 1.0, nil},
		{2, 0, 0.0, fmt.Errorf("zero division")},
	}

	for _, tt := range tests {
		got, err := Div(tt.in1, tt.in2)
		var testErr error
		if err != nil {
			testErr = tt.err
		}
		// errorは内容を比較するため、reflectパッケージを使っている
		if got != tt.out || !reflect.DeepEqual(testErr, err) {
			t.Errorf("Div(%d, %d) = %v, %v want %v, %v",
				tt.in1, tt.in2, got, err, tt.out, testErr)
		}
	}
}
