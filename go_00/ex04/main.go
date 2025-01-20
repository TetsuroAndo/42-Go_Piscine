package main

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				// 最後("789")以外は", "を付与
				if i == '7' && j == '8' && k == '9' {
					z01.PrintRune('\n')
				} else {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
