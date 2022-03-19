## 引用
[Effective Go - The Go Programming Language (google.cn)](https://golang.google.cn/doc/effective_go#defer)

[函数 |《高效的 Go 编程 Effective Go 2020》| Go 技术论坛 (learnku.com)](https://learnku.com/docs/effective-go/2020/function/6242)

### 多返回值
Go不同寻常的特性它支持多返回值，在其他语言中，如果存在多返回值，我们可能需要封装一个对象，或者是通过地址传参修改实参。

在C中，写入操作发生的错误会用一个负数标记，而错误码会隐藏在某个不确定的位置(可能是一个全局变量)。而在 Go 中，`Write` 会返回写入的字节数以及一个错误： “是的，您写入了一些字节，但并未全部写入，因为设备已满”。 在 `os` 包中，`File.Write` ：
```
// Write writes len(b) bytes to the File.
// It returns the number of bytes written and an error, if any.
// Write returns a non-nil error when n != len(b).
func (f *File) Write(b []byte) (n int, err error) {
   if err := f.checkValid("write"); err != nil {
      return 0, err
   }
   n, e := f.write(b)
   if n < 0 {
      n = 0
   }
   if n != len(b) {
      err = io.ErrShortWrite
   }

   epipecheck(f, e)

   if e != nil {
      err = f.wrapErr("write", e)
   }

   return n, err
}
```
如文档所述，他会返回写入的字节数，并且会在`n != len(b)`时返回一个非nil的`error`错误值。这是Go常见的异常处理方法。

### 返回值参数命名

Go函数的返回参数可以被命名，并可以作为常规变量使用，就像传入的形参一样。初始化为其类型的默认值；
```
func nextInt(b []byte, i int) (int, int)

func nextInt(b []byte, pos int) (value, nextPos int) 
```
若该函数执行了一条不带实参的 `return` 语句，则结果形参的当前值将被返回。
```
func err() (a, b int) {
   return

   return 0, 0
}
```

由于被命名的结果已经初始化，且已经关联至无参数的返回，它们就能让代码简单而清晰。 下面的 `io.ReadAtLeast` 就是个很好的例子：
```
// ReadAtLeast reads from r into buf until it has read at least min bytes.
// It returns the number of bytes copied and an error if fewer bytes were read.
// The error is EOF only if no bytes were read.
// If an EOF happens after reading fewer than min bytes,
// ReadAtLeast returns ErrUnexpectedEOF.
// If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer.
// On return, n >= min if and only if err == nil.
// If r returns an error having read at least min bytes, the error is dropped.
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) {
   if len(buf) < min {
      return 0, ErrShortBuffer
   }
   for n < min && err == nil {
      var nn int
      nn, err = r.Read(buf[n:])
      n += nn
   }
   if n >= min {
      err = nil
   } else if n > 0 && err == EOF {
      err = ErrUnexpectedEOF
   }
   return
}
```
### defer 延迟函数
Go 语言的`defer`语句用于预设一个延迟执行的函数，这个函数会在执行`defer`的函数返回之前立即执行。类似于Java、C#中的`try-finally`。例如无论以何种路径返回，都必须释放资源的函数。 典型的例子就是解锁互斥和关闭文件。

关闭文件
```
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}

```
推迟诸如 Close 之类的函数调用有两点优势：第一，它能确保你不会忘记关闭文件。如果你以后又为该函数添加了新的返回路径时， 这种情况往往就会发生。第二，它意味着 “关闭” 离 “打开” 很近， 这总比将它放在函数结尾处要清晰明了。

被推迟函数的实参（如果该函数为方法则还包括接收者）在推迟执行时就会求值， 而不是在调用执行时才求值。这样不仅无需担心变量值在函数执行时被改变， 同时还意味着单个已推迟的调用可推迟多个函数的执行。下面是个简单的例子。

```
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```
被推迟的函数按照后进先出（LIFO）的顺序执行，因此以上代码在函数返回时会打印 `4 3 2 1 0`。一个更具实际意义的例子是通过一种简单的方法， 用程序来跟踪函数的执行。

解锁
```
type Foo struct {
   mu    sync.Mutex
   count int
}

func (f *Foo) Bar2() {
   f.mu.Lock()
   //返回时 释放锁
   defer f.mu.Unlock()

   if f.count < 1000 {
      f.count += 3
      return
   }
   f.count++
   return
}
```
We can do better by exploiting the fact that arguments to deferred functions are evaluated when the `defer` executes. The tracing routine can set up the argument to the untracing routine. This example:
```
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```
输出
```
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```
对于习惯其它语言中块级资源管理的程序员，defer 似乎有点怪异， 但它最有趣而强大的应用恰恰来自于其基于函数而非块的特点。在 `panic` 和 `recover` 这两节中，我们将看到关于它可能性的其它例子。
