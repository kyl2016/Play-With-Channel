package atomic

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAddUint32_Negative_1(t *testing.T) {
	var i uint32 = 1
	delta := int32(-1)
	fmt.Println(delta, uint32(delta))
	r := atomic.AddUint32(&i, uint32(delta))
	if r != 0 {
		t.Error("should be 0")
	}
}

func TestAddUint32_Negative_2(t *testing.T) {
	var i uint32 = 1
	negative := -1

	// ^uint32(-N-1)
	// 我们先要把差量的绝对值减去1，然后再把得到的这个无类型的整数常量，转换为uint32类型的值，最后，在这个值之上做按位异或操作，就可以获得最终的参数值了。
	r := atomic.AddUint32(&i, ^uint32(-negative-1))
	if r != 0 {
		t.Error("should be 0")
	}
}

func TestAtomic_Load(t *testing.T) {
	var i int32 = 5
	atomic.StoreInt32(&i, 6)
	j := atomic.LoadInt32(&i)
	if j != 6 {
		t.Error("should equal")
	}

}

func TestAtomic_Value(t *testing.T) {
	type Data struct {
		Name string
		Age  uint8
	}

	var d = Data{
		Name: "Kitty",
		Age:  28,
	}

	v := atomic.Value{}
	v.Store(d)

	d.Age = 30

	d2 := v.Load()
	fmt.Println(d2.(Data))
}

func TestAtomic_CAS(t *testing.T) {
	var i int32 = 1
	swapped := atomic.CompareAndSwapInt32(&i, 1, 2)
	if !swapped {
		t.Error("should swapped")
	} else {
		fmt.Println(i)
	}
}
