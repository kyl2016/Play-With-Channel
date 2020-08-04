package main

/*
typedef void (*callback)(void *);

static callback _cb;
static void *_user_data;
static void register_callback(callback cb, void *user_data) {
	_cb = cb;
	_user_data = user_data;
}
static void wait_event() {
	_cb(_user_data);
}

void cb_proxy(void *v);

static void _register_callback(void *user_data) {
	register_callback(cb_proxy, user_data);
}
*/
import "C"
import (
	"fmt"
	pointer "github.com/mattn/go-pointer"
	"unsafe"
)

type Callback struct {
	Func     func(string)
	UserData string
}

func my_callback(v string) {
	fmt.Println("hello", v)
}

func main() {
	C._register_callback(pointer.Save(&Callback{
		Func:     my_callback,
		UserData: "my-callback",
	}))
	C.wait_event()
}

//export cb_proxy
func cb_proxy(v unsafe.Pointer) {
	cb := pointer.Restore(v).(*Callback)
	cb.Func(cb.UserData)
}
