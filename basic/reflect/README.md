# reflect

## Overview
The typical use is to take a value with static type interface{} and extract its dynamic type information by calling TypeOf, which returns a Type.

A call to ValueOf returns a Value representing the run-time data. Zero takes a type and returns a Value representing a zero value for that type.

Golang 的 gRPC 也是通过反射实现的。

Golang 关于类型设计的一些原则：
* 变量包括（type、value）两部分
* type 包括 static type 和 concrete type。编码时看到的 type 是 static type，runtime 系统看到的是 concrete type。
* 类型断言能否成功，取决于变量的 concrete type。

反射主要与 interface 类型相关。

interface{} 类型的变量包含了2个指针，一个指向值的类型（对应 concrete type），另一个指针指向实际的值（对应的 value）。

例如，创建类型为 *os.File 的变量，然后将其赋值给一个接口变量 r：
```
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

var r io.Reader
r = tty
```
接口变量 r 的 pair 中将记录如下信息：(tty, *os.File)，这个 pair 在接口变量的连续赋值过程中是不变的
```$xslt
var w io.Writer
w = r.(io.Writer)
```
接口变量 w 的 pair 与 r 的 pair 相同，都是（tty, *os.File），即使 w 是空接口类型，pair 也是不变的。
反射就是用来检测存储在接口变量内部（值 value；类型 concrete type）pair 的一种机制。

## reflect 的基本功能 TypeOf 和 ValueOf

- 已知原有类型【进行“强制转换”】
```$xslt
realValue := value.Interface().(已知的类型)
```
注意：
    * 如果转换的类型不完全符合，则直接 panic！ （最好用 `realValue, ok := value.Interface().(已知的类型)`）
    * 转换的时候，要区分是指针还是值
    * 反射可以将“反射类型对象”再重新转换为“接口类型对象”

- 未知原类型【遍历探测其Field】

## 使用场景
* 遍历结构体的字段、方法
* 获取结构体的tag标记的值
* Handing JSON: encoding/json等大量使用了reflect
The most common use of reflection is marshaling and unmarshaling data from a file or a network.
Whenever you specify struct tags for JSON or database mapping, you are depending on reflection.
* THe biggest hint is checking the types of the parameters.
    * If you call a function with a parameter of type interface{}, there's a good chance it's going to use reflection to examine or change the parameter's value.
* Memoization and Short Term Memory ([sample_memoization](sample/sample_memoization.go))
    * It's a lot like caching at a function level. Memoization is the process of creating a function that wraps one of these invariant functions, caching the input and output values to avoid doing unneeded work. For functions that do complex or slow things, the performance savings can be tremendous.
    * You have one service that provides a value, and another that uses that value. Because there's a network call to get the value, it takes some time. When the value doesn't change all that often, and it doesn't matter if the value is out of date by a few seconds, caching that value temporarily can give your system a significant performance gain. 

It's not a solution that you reach for all of the time, but when you have a problem that seems impossible because there's no commonality between types, or because the data is dynamic, reflection is your secret weapon.

## 不可比较 not Comparable
* Slices
* Maps
* Functions
* Structs that contains fields of type slice, map or function

## 性能
Golang 的反射很慢，和它的API设计有关。在java里面，一般这样使用：
```$xslt
Field field = classA.getField("hello");
field.get(obj1);
field.get(obj2);
```
这个取得的反射对象类型是java.lang.reflect.Field，是可以复用的。
但Golang的反射：
```$xslt
_type := reflect.Typeof(obj)
field, _ := _type.FieldByName("hello")
```
这里取出来的field对象是reflect.StructField类型，但是没办法得到对应对象上的值。如果取值：
```$xslt
value := reflect.ValueOf(obj)
fieldValue := value.FieldByName("hello")
```
这里取出来的fieldValue类型是reflect.Value，它是一个具体的值，而不是一个可复用的反射对象了，每次反射都需要malloc这个reflect.Value结构体，并且还涉及到GC。
这里取出来的fieldValue类型是reflect.Value，它是一个具体的值，而不是一个可复用的反射对象了，每次反射都需要malloc这个reflect.Value结构体，并且还涉及到GC。

慢的主要原因：
* 涉及到内存分配以及后续的GC；
* reflect实现里面有大量的枚举，也就是for循环，比如类型之类的。

## 总结
* 反射可以大大提高程序的灵活性，使得interface{}有更大的发挥余地
* 反射可以将“接口类型变量”转换为“反射类型对象”（使用TypeOf和ValueOf）
* 反射可以将”反射类型对象“转为“接口类型对象”
    * reflect.Value.Interface{}.(已知的类型)
    * 遍历reflect.Type的Field获取其Field
* 反射可以修改反射类型的对象，但是其值必须是“addressable” （pointer-interface）
* 通过反射可以“动态”调用方法
* 因为Golang本身不支持模板，因此在需要的场景可以使用反射来实现