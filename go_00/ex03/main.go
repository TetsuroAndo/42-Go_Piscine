package main

import "github.com/01-edu/z01"

// IsNegative prints 'T' if n < 0, otherwise 'F'
func IsNegative(n int) {
	if n < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

func main() {
	// テスト例としていくつか呼び出してみる
	IsNegative(-1) // T
	IsNegative(0)  // F
	IsNegative(10) // F
}
