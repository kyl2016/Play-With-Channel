# Play-With-Golang
通过各种Demo来研究Golang

## goroutines

to execute tasks independently, potentially in parallel

## Channel

for communication, synchronization between goroutines.

1. 长度为0和大于0的区别

buffered channel
```
ch := make(chan int, 3) // Allocates an hchan struct on the heap, returns a pointer to it.
```
- goroutine-safe
- stores up to capacity elements, and provides FIFO semantics
- sends values between goroutines
- can cause them to block, unblock

unbuffered channel
```
ch := make(chan int)
```

2. 单向 chan
3. cap与len的区别
4. 如何判断从 channel 读取数据是否超时？示例：[channel_timeout](./channel_timeout.go)
5. 如果判断 channel 是否已满，不能在向其发送数据？示例：[channel_is_empty_or_full](./channel_is_empty_or_full.go)
6. 为了避免读到的是已关闭的 channel 中的数据，建议使用：
```
if v, ok := <- ch; ok {
    // ...
}
```
## Map
1. 从map中取某项
```
if val, ok := map["name"]; ok {

}
```
## enum
```
type OperationType string

const {
    Play   OperationType = "play"
    Puase  OperationType = "pause"
    Resume OperationType = "resume"
    Stop   OperationType = "stop"
}

func IsValidOperationType(s string) (OperationType, bool) {
    switch s {
    case string(Play), string(Pause), string(Resume), string(Stop):
        return Operation(s), true
    }
      
    return Play, false
}

```

## select 
- select 就是监听IO操作，只能用于 channel 的发送和接收，如果多个 case 都满足条件，则用**伪随机（pseudo-random choice）算法**选择一个执行。示例：[channel_random_select_case](channel_random_select_case.go)
- switch 可为各种类型进行分支操作，通过i.(type)可以为接口类型进行分支判断。分支是**顺序执行**的，这和 select 不同。

## json
[convert json to struct](json_to_struct.go)

## for range
using **value semantics**
The for range is iterating over its **own copy** of the array/map/slice.
When ranging over a slice using the value semantics form, a copy of the slice header is taken.


