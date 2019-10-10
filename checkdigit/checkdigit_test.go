package checkdigit

import "testing"

func TestHello(t *testing.T) {
	got := func() int { return 1 }()
	expected := 1
	if got != expected {
		t.Errorf(`Hello() = %d`, expected)
	}
}

func TestLuhn(t *testing.T) {
	got1 := Luhn("4242424242424242")
	if got1 != true {
		t.Errorf(`Luhn(%s) = %t`, "4242424242424242", got1)
	}
	got2 := Luhn("4242422242424242")
	if got2 != false {
		t.Errorf(`Luhn(%s) = %t; want false`, "4242422242424242", got2)
	}
	got3 := Luhn("abc")
	if got3 != false {
		t.Errorf(`Luhn(79%s) = %t; want false`, "abc", got3)
	}
	got4 := Luhn("")
	if got3 != false {
		t.Errorf(`Luhn(%s) = %t; want false`, "", got4)
	}
}
