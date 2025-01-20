package main

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}
	comb := make([]rune, n) // n桁分の組み合わせを保持するスライス

	var backtrack func(start int, depth int)
	backtrack = func(start int, depth int) {
		if depth == n {
			// comb に格納した n 桁を表示
			for i := 0; i < n; i++ {
				z01.PrintRune(comb[i])
			}
			// 最後の組み合わせ(例:n=3 なら "789")かどうか判定
			if string(comb) == lastComb(n) {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			return
		}

		// start から 9 まで走査し、comb[depth] に代入して再帰
		for digit := start; digit <= 9; digit++ {
			comb[depth] = rune(digit + '0')
			backtrack(digit+1, depth+1)
		}
	}

	backtrack(0, 0)
}

// 最後の組み合わせを文字列として返すためのヘルパー。
// 例: n=3 → "789", n=1 → "9", n=2 → "89" etc.
func lastComb(n int) string {
	r := make([]rune, n)
	current := 9
	for i := n - 1; i >= 0; i-- {
		r[i] = rune(current + '0')
		current--
	}
	return string(r)
}

func main() {
	// テスト例として n=3 を呼び出す
	PrintCombN(3)
}
