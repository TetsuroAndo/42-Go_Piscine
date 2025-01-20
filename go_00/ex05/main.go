package main

import "github.com/01-edu/z01"

func printTwoDigits(n int) {
	// 2桁で表示する
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			printTwoDigits(i)
			z01.PrintRune(' ')
			printTwoDigits(j)
			// 最後("98 99")以外は", "を付与
			if i == 98 && j == 99 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
