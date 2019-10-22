package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"unicode/utf8"
)

type Charset string

const (
	UTF8    = Charset("utf-8")
	GBK     = Charset("gbk")
	GB18030 = Charset("gb18030")
)

func main() {
	println(^uintptr(0), "\n", (^uintptr(0) >> 61), "\n", (^uintptr(0) >> 62), "\n", (^uintptr(0) >> 63), 4<<(^uintptr(0)>>63))

	println(15 >> 1)

	//fmt.Printf("%s\n\n", in)
	fmt.Printf("%x\n\n", in)
	//fmt.Printf("%s\n\n", string([]byte(in)))
	//fmt.Printf("%v\n\n", in)

	println(utf8.FullRune(in)) // 不过如果不需要考虑各种意外，且只需要区分utf8和gb18030
	println(utf8.FullRuneInString(string(in)))

	println(ConvertByte2String(in, GBK))
	println(ConvertByte2String(in, GB18030))
	println(ConvertByte2String(in, UTF8))
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GBK:
		var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

var in = []byte{
	121,
	110,
	120,
	105,
	95,
	102,
	97,
	99,
	101,
	95,
	50,
	48,
	49,
	57,
	48,
	51,
	50,
	50,
	47,
	90,
	104,
	101,
	110,
	103,
	88,
	105,
	110,
	103,
	98,
	97,
	111,
	47,
	157,
	168,
	233,
	184,
	163,
	233,
	185,
	164,
	46,
	106,
	112,
	103,
}
