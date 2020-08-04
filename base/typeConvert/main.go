package main

import "fmt"

func main() {
	var src = int16(-255)
	dst := int8(src)
	fmt.Println(dst)

	// 字符'�'的 Unicode 代码点是 U+FFFD，是 Unicode 标准中定义的 Replacement Character，
	// 专用于替换哪些未知的、不被认可以及无法展示的字符
	fmt.Println(string(-1), string(48), string(49))
}

// 整数是以补码的形式存储的。
// 补码是原码各位求反再加 1。
// int16 类型的值-255的补码是1111111100000001，转换为 int8，会将高 8 位直接截掉，从而得到00000001。
// 由于其最左边一位是 0，表示它是个正整数，（正整数的补码等于其原码），所以 dst 值就是 1。
