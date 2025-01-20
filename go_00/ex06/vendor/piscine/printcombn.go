package piscine

import "ft"

var firstPrint = true

func PrintCombN(n int) {
    if n <= 0 || n >= 10 {
        return
    }
    var digits [10]rune
    firstPrint = true
    recursivePrintCombN(digits[:], 0, '0', n)
}

func recursivePrintCombN(digits []rune, index int, start rune, n int) {
    if index == n {
        printDigitsSlice(digits, n)
        return
    }
    for c := start; c <= '9'-rune(n-index)+1; c++ {
        digits[index] = c
        recursivePrintCombN(digits, index+1, c+1, n)
    }
}

func printDigitsSlice(digits []rune, n int) {
    if !firstPrint {
        ft.PrintRune(',')
        ft.PrintRune(' ')
    }
    for i := 0; i < n; i++ {
        ft.PrintRune(digits[i])
    }
    firstPrint = false
}
