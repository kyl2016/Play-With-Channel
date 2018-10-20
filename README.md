# Play-With-Golang
通过各种Demo来研究Golang

## Channel

1. make(chan int) 与 make(chan int, 1) 的区别
2. 单向 chan
3. cap与len的区别
4. 

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
