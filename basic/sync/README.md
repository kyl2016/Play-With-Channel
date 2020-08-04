# sync

## Initialization

If a package p imports package q, the completion of q's init functions happens before the start of any of p's.

The start of the function main.main happens after all init functions have finished.

## Goroutine creation

## 同步

### channel communication

Channel communication is the main method of synchronization between goroutines. Each send on a particular channel is matched to a corresponding receive from that channel, usually in a different goroutine.

A send on a channel happens before the corresponding receive from that channel completes.

[sample](chan/main.go)

### Locks

For any sync.Mutex or sync.RWMutex variable l and n < m, call n of l.Unlock() happens before call m of l.Lock() returns.

[sample](lock/main.go)

### Once
Multiple threads can execute once.Do(f) for a particular f, but only one will run f(), and the other calls block until f() has returned.

[sample](once/once.go)

Double-checked locking is an attempt to avoid the overhead of synchronization.
[double_check](once_double_check/once_double_check.go)

### Incorret syncronization

As before, there is no guarantee that, in main, observing the write to done implies observing the write to a, so this program could print an empty string too. Worse, there is no guarantee that the write to done will ever be observed by main, since there are no synchronization events between the two threads. The loop in main is not guaranteed to finish.

[incorrect](incorrect/main.go)

## refers

[The Go Memory Model](https://golang.org/ref/mem)
[进程内协同：同步、互斥与通讯](https://time.geekbang.org/column/article/96994)
