# [Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines)

Informally, a pipeline is a series of stages connected by *channels*, where each stage is a group of goroutines running the same function. In each stage, the goroutines
* receive values from *upstream* via *inbound* channels
* perform some function on that data, usually producing new values
* send values *downstream* via *outbound* channels

The first stage is sometimes called the *source* or *producer*, the last stage, the *sink* or *consumer*.

## Fan-out, fan-in
Multiple functions can read from the same channels until that channel is closed; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that's closed when all the inputs are closed. This is called *fan-in*.

[fanOut](fanOut_test.go)

use merge1, output:
```$xslt
=== RUN   TestFanOut
9

or

=== RUN   TestFanOut
4
 
```
closure use the reference of c, so the two goroutines read from the same channel.

There is a pattern to our pipeline functions:
* stages close their outbound channels when all the send operations are done.
* stages keep receiving values from inbound channels until those channels are closed.

In our example pipeline, if a stage fails to consume all inbound values, the goroutines attempting to send those values will block indefinitely:
```$xslt
// Consume the first value from the output.
out := merge(c1, c2)
fmt.Println(<-out) // 4 or 9
return
// Since we didn't receive the second value from out, one of the output goroutines is hung attempting to send it.
```
**Goroutines are not garbage collected; they must exit on their own**.

