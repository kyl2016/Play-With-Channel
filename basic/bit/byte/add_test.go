package byteDemo

import (
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	b := make([]byte, 10)
	binary.BigEndian.PutUint32(b, 1)
	fmt.Println(b)

	binary.LittleEndian.PutUint32(b, 1)
	fmt.Println(b)

	binary.BigEndian.PutUint32(b, 2)
	fmt.Println(b)

	binary.LittleEndian.PutUint32(b, 2)
	fmt.Println(b)

	d := uint32(math.Pow(2, 31))
	fmt.Println(d)
	binary.BigEndian.PutUint32(b, d)
	fmt.Println(b)

	binary.LittleEndian.PutUint32(b, d)
	fmt.Println(b)
}
