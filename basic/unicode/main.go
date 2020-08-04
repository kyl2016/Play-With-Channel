package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var s string = "Go爱好者杨"
	fmt.Printf("The string: %q\n", s)
	fmt.Printf("	=> runes(char): %q\n", []rune(s))
	fmt.Printf("	=> runes(hex): %x\n", []rune(s))
	fmt.Printf("	=> runes(dec): %d\n", []rune(s))
	fmt.Printf("	=> bytes(hex): [% x]\n", []byte(s))
	fmt.Printf("	=> bytes(hex): [% d]\n", []byte(s))

	fmt.Println(utf8.RuneCount([]byte(s)))

	for i, c := range s {
		fmt.Printf("%d %q [% x]\n", i, c, []byte(string(c)))

		if len([]byte(string(c))) > 1 {
			fmt.Println("中文")
		}
	}
	s = ""
	require := false
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		i += size
		// Officially, ZIP uses CP-437, but many readers use the system's
		// local character encoding. Most encoding are compatible with a large
		// subset of CP-437, which itself is ASCII-like.
		//
		// Forbid 0x7e and 0x5c since EUC-KR and Shift-JIS replace those
		// characters with localized currency and overline characters.
		if r < 0x20 || r > 0x7d || r == 0x5c {
			if !utf8.ValidRune(r) || (r == utf8.RuneError && size == 1) {
				println("false", "false")
			}
			require = true
		}
	}
	println("true", require)

	//for i := 'a'; i <= 'z'; i++ {
	//	fmt.Printf("%d %s\n", i, string(i))
	//}
}
