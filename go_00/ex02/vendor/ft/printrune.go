package ft

import "syscall"

func PrintRune(r rune) {
    syscall.Write(syscall.Stdout, []byte{byte(r)})
}
