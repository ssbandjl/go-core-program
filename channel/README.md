### Channel

Channel 主要有两种形式：

1. **有缓存 Channel（buffered channel）**，使用 `make(chan T, n)` 创建
2. **无缓存 Channel（unbuffered channel）**，使用 `make(chan T)` 创建

其中 `T` 为 Channel 传递数据的类型，`n` 为缓存的大小，这两种 Channel 的读写操作都非常简单：

```
// 创建有缓存 Channel
ch := make(chan interface{}, 10)
// 创建无缓存 Channel
ch := make(chan struct{})
// 发送
ch <- v
// 接受
v := <- ch
```



他们之间的本质区别在于其内存模型的差异，这种内存模型在 Channel 上体现为：

- 有缓存 Channel: `ch <- v` 发生在 `v <- ch` 之前
- 有缓存 Channel: `close(ch)` 发生在 `v <- ch && v == isZero(v)` 之前
- 无缓存 Channel: `v <- ch` 发生在 `ch <- v` 之前
- 无缓存 Channel: 如果 `len(ch) == C`，则从 Channel 中收到第 k 个值发生在 k+C 个值得发送完成之前

直观上我们很好理解他们之间的差异： 对于有缓存 Channel 而言，内部有一个缓冲队列，数据会优先进入缓冲队列，而后才被消费， 即向通道发送数据 `ch <- v` 发生在从通道接受数据 `v <- ch` 之前； 对于无缓存 Channel 而言，内部没有缓冲队列，即向通道发送数据 `ch <- v` 一旦出现， 通道接受数据 `v <- ch` 会立即执行， 因此从通道接受数据 `v <- ch` 发生在向通道发送数据 `ch <- v` 之前。 我们随后再根据实际实现来深入理解这一内存模型。

Go 语言还内建了 `close()` 函数来关闭一个 Channel：

```
close(ch)
```



但语言规范规定了一些要求：

- 关闭一个已关闭的 Channel 会导致 panic

- 向已经关闭的 Channel 发送数据会导致 panic

- 向已经关闭的 Channel 读取数据不会导致 panic，但读取的值为 Channel 缓存数据的零值，可以通过接受语句第二个返回值来检查 Channel 是否关闭：

  ```
  v, ok := <- ch
  if !ok {
    ... // Channel 已经关闭
  }
  ```

  

### Select

Select 语句伴随 Channel 一起出现，常见的用法是：

```
select {
case ch <- v:
	...
default:
	...
}
```



或者：

```go
select {
case v := <- ch:
	...
default:
	...
}
```