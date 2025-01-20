// printrune.go
package ft

import "syscall"

// PrintRune は 1文字 r を標準出力へ書き込みます。
func PrintRune(r rune) {
    // 'syscall.Stdout' は通常 1 (Unixのfile descriptor 1)
    // ASCII文字であれば1バイトでOK
    syscall.Write(syscall.Stdout, []byte{byte(r)})
}
