## 引用
[Effective Go - The Go Programming Language (google.cn)](https://golang.google.cn/doc/effective_go#data)

[数据 |《高效的 Go 编程 Effective Go 2020》| Go 技术论坛 (learnku.com)](https://learnku.com/docs/effective-go/2020/data/6243)

## 使用`new`分配内存
Go提供了两种分配原语，内建函数`new`和`make`。他们所做的事情不同，所应用的类型也不同，他们可能会引起混淆，但是规则很简单。我们先来讨论一下`new`。这个是用来分配内存的内建函数，但与其他语言中的函数名不通，他不会初始化内存，只会将内存置零。

`new` 初始化对象的内存已置零，当你涉及数据结构时，对于每种类型的零值就不必进行一步初始化了。例如`bytes.Buffer`的文档中提到"零值的`Buffer`就是已准备就绪的缓冲区"。同样的，`sync.Mutex`并没有显示的构造函数`Init`方法，而是零值的`sync.Mutex`就已经是解锁状态的互斥锁。

```Go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```

```Go
// A Buffer is a variable-sized buffer of bytes with Read and Write methods.
// The zero value for Buffer is an empty buffer ready to use.
type Buffer struct {
   buf      []byte // contents are the bytes buf[off : len(buf)]
   off      int    // read at &buf[off], write at &buf[len(buf)]
   lastRead readOp // last read operation, so that Unread* can work correctly.
}
```

`SyncedBuffer` 类型的值也是在声明时就分配好内存就绪了。后续代码中， `p` 和 `v` 无需进一步处理即可正确工作。
```
func main() {
   p := new(SyncedBuffer)  // type *SyncedBuffer
   var v SyncedBuffer      // type  SyncedBuffer

   fmt.Println(p) //&{{0 0} {[] 0 0}}
   fmt.Println(v) //{{0 0} {[] 0 0}}
   fmt.Println(&v == p) //false
}
```

## 构造函数和 composite literals
有些时候零值并不那么友好，这时候我们需要初始化一个构造函数
```Go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```
上面的赋值操作显得代码过于冗长，通过a *composite literal*来简化。
```Go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```
少数情况下，若复合字面不包括任何字段，它将创建该类型的零值。表达式 `new(File)` 和 `&File{}` 是等价的。

```Go
a := [...]string{"no error", "Eio", "invalid argument"}
s := []string{"no error", "Eio", "invalid argument"}
m := map[int]string{1: "no error", 2: "Eio", 3: "invalid argument"}
```

## 使用make分配内存
使用`make(T, `*args*`)`的目的不同于`new(T)`。他只用于创建切片（slice）、map和channels。并返回类型T的一个已初始化的值。出现这种差异的原因在于，这三种本质上为引用类型，他们在使用前必须初始化。例如：切片是一个具有三种类型，包含一个指向（数组内部）数据的指针、长度以及容量， 在这三项被初始化之前，该切片为 `nil`。
```Go
type slice struct {
   array unsafe.Pointer
   len   int
   cap   int
}
```

```Go
make([]int, 10, 100)
```
会分配一个具有 100 个 int 的数组空间，接着创建一个长度为 10， 容量为 100 并指向该数组中前 10 个元素的切片结构。（生成切片时，其容量可以省略，更多信息见切片一节。） 与此相反，`new([]int)` 会返回一个指向新分配的，已置零的切片结构， 即一个指向 nil 切片值的指针。

```
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```

## Array
在详细规划内存布局时，数组是非常有用的，有时还能避免过多的内存分配， 但它们主要用作切片的构件。这是下一节的主题了，不过要先说上几句来为它做铺垫。
以下为数组在 Go 和 C 中的主要区别。在 Go 中，
- 数组是值。将一个数组赋予另一个数组会复制其所有元素。
- 特别地，若将某个数组传入某个函数，它将接收到该数组的一份**副本**而非指针。
- 数组的大小是其类型的一部分。类型 `[10]int` 和 `[20]int` 是不同的。

数组为值的属性很有用，但代价高昂；若你想要 C 那样的行为和效率，你可以传递一个指向该数组的指针。
```Go
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // Note the explicit address-of operator
```

## 切片
切片通过对数组进行封装，为数据序列提供了更通用、强大而方便的接口。 除了矩阵变换这类需要明确维度的情况外，Go 中的大部分数组编程都是通过切片来完成的。
切片保存了对底层数组的引用，若你将某个切片赋予另一个切片，它们会引用同一个数组。 若某个函数将一个切片作为参数传入，则它对该切片元素的修改对调用者而言同样可见， 这可以理解为传递了底层数组的指针。因此，Read 函数可接受一个切片实参 而非一个指针和一个计数；切片的长度决定了可读取数据的上限。以下为 os 包中 File 类型的 Read 方法签名:
```
func (f *File) Read(buf []byte) (n int, err error)
```
该方法返回读取的字节数和一个错误值（若有的话）。若要从更大的缓冲区 `b` 中读取前 32 个字节，只需对其进行**切片**即可。
```
    n, err := f.Read(buf[0:32])
```
这种切片的方法常用且高效。若不谈效率，以下片段同样能读取该缓冲区的前 32 个字节。
```go
    var n int
    var err error
    for i := 0; i < 32; i++ {
        nbytes, e := f.Read(buf[i:i+1])  // Read one byte.
        n += nbytes
        if nbytes == 0 || e != nil {
            err = e
            break
        }
    }
 ```

只要切片不超出底层数组的限制，它的长度就是可变的，只需将它赋予其自身的切片即可。 切片的容量可通过内建函数 cap 获得，它将给出该切片可取得的最大长度。 以下是将数据追加到切片的函数。若数据超出其容量，则会重新分配该切片。返回值即为所得的切片。 该函数中所使用的 len 和 cap 在应用于 nil 切片时是合法的，它会返回 0。
```Go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
```

## Two-dimensional slices 二维切片
Go 的数组和切片都是一维的。要创建等价的二维数组或切片，就必须定义一个数组的数组， 或切片的切片，就像这样：
```Go
type Transform [3][3]float64  // A 3x3 array, really an array of arrays.
type LinesOfText [][]byte     // A slice of byte slices.
```

```Go
text := LinesOfText{
   []byte("Now is the time"),
   []byte("for all good gophers"),
   []byte("to bring some fun to the party."),
}

fmt.Println(text)
/*[
   [78 111 119 32 105 115 32 116 104 101 32 116 105 109 101] 
   [102 111 114 32 97 108 108 32 103 111 111 100 32 103 111 112 104 101 114 115] 
   [116 111 32 98 114 105 110 103 32 115 111 109 101 32 102 117 110 32 116 111 32 116 104 101 32 112 97 114 116 121 46]
   ]
 */
form := Transform{
   [3]float64{1,2,3},
   [3]float64{4,5,6},
   [3]float64{7,8,9},
}
fmt.Println(form)
//[[1 2 3] [4 5 6] [7 8 9]]
```

有时必须分配一个二维数组，例如在处理像素的扫描行时，这种情况就会发生。 我们有两种方式来达到这个目的。
```Go
//为每一个切片进行独立的内存分配
func allocate() {
   // Allocate the top-level slice.
   picture := make([][]uint8, YSize) // One row per unit of y.
   fmt.Println(picture)//[[] [] [] [] [] [] [] [] [] []]

   // Loop over the rows, allocating the slice for each row.
   for i := range picture {
      picture[i] = make([]uint8, XSize)
   }

   fmt.Println(picture)
   //[[0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0]]
}

//申请一个大的切片，其他切片指向该切片
func allocate2() {
   // Allocate the top-level slice, the same as before.
   picture := make([][]uint8, YSize) // One row per unit of y.
   // Allocate one large slice to hold all the pixels.
   pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
   // Loop over the rows, slicing each row from the front of the remaining pixels slice.
   for i := range picture {
      picture[i], pixels = pixels[:XSize], pixels[XSize:]
   }
}
```

## Maps 图
Go语言中的Map其实也就是数据结构中的HashTable。
map是方便而强大的内建数据结构，它可以关联不同类型的值。其键可以是任何相等性操作符支持的类型， 如整数、浮点数、复数、字符串、指针、接口（只要其动态类型支持相等性判断）、结构以及数组。 切片不能用作映射键，因为它们的相等性还未定义。与切片一样，映射也是引用类型。 若将映射传入函数中，并更改了该映射的内容，则此修改对调用者同样可见。

map可使用一般的复合字面语法进行构建，其键 - 值对使用冒号分隔，因此可在初始化时很容易地构建它们。
```
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```
赋值于获取map中的值语法类似于数组，不同的是他的索引不一定是整数。
```
offset := timeZone["EST"]
```

若试图通过映射中不存在的键来取值，就会返回与该映射中项的类型对应的零值。 例如，若某个映射包含整数，当查找一个不存在的键时会返回零值。 集合可实现成一个值类型为 bool 的映射。将该映射中的项置为 true 可将该值放入集合中，此后通过简单的索引操作即可判断是否存在。

```
func mapDemo() {
   attended := map[string]bool{
      "Ann": true,
      "Joe": true,
   }

   if attended["Ann"] { // will be false if person is not in the map
      fmt.Println("Ann", "was at the meeting")
   }

   fmt.Println(attended)
}
```

有时你需要区分某项是不存在还是其值为零值。如对于一个值本应为零的 `"UTC"` 条目，也可能是由于不存在该项而得到零值。你可以使用多重赋值的形式来分辨这种情况。
```
var timeZone = map[string]int{
   "UTC":  0*60*60,
   "EST": -5*60*60,
   "CST": -6*60*60,
   "MST": -7*60*60,
   "PST": -8*60*60,
}

seconds, ok := timeZone["UTC"]
fmt.Println(seconds, ok)
//0 true

seconds, ok = timeZone["UFC"]
fmt.Println(seconds, ok)
//0 false
```

显然，我们可称之为 “逗号 ok” 语法。在下面的例子中，若 `tz` 存在， `seconds` 就会被赋予适当的值，且 `ok` 会被置为` true`； 若不存在，`seconds` 则会被置为零，而 ok 会被置为 false。
```
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```
若仅需判断映射中是否存在某项而不关心实际的值，可使用空白标识符`_`来代替该值的一般变量。
```
_, present := timeZone[tz]
```

要删除映射中的某项，可使用内建函数 `delete`，它以映射及要被删除的键为实参。 即便对应的键不在该映射中，此操作也是安全的。
```
delete(timeZone, "PDT")  // Now on Standard Time
```

## Printing
Go的格式化打印`printf`输出风格与C类型，但它更加的丰富。这些函数位于`fmt`的包中。如`fmt.Printf`, `fmt.Fprintf`, `fmt.Sprintf`。字符串函数(`Sprintf`)会返回一个字符串，而非填充给定的缓冲区。

```
func printing() {
   fmt.Printf("Hello %d\n", 23)
   fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
   fmt.Println("Hello", 23)
   fmt.Println(fmt.Sprint("Hello ", 23))
}
```
`fmt.Fprint` 一类的格式化打印函数可接受任何实现了 `io.Writer` 接口的对象作为第一个实参；变量 `os.Stdout` 与 `os.Stderr` 都是人们熟知的例子。

像 `%d` 这样的数值格式并不接受表示符号或大小的标记， 打印例程会根据实参的类型来决定这些属性。
```
var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
// 18446744073709551615 ffffffffffffffff; -1 -1
```
若你只想要默认的转换，如使用十进制的整数，你可以使用通用的格式 %v（对应 “值”）；其结果与 Print 和 Println 的输出完全相同。此外，这种格式还能打印任意值，甚至包括数组、结构体和映射。 以下是打印上一节中定义的时区映射的语句。
```
fmt.Printf("%v\n", timeZone)  // or just fmt.Println(timeZone)
```
输出
```
map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
```

对于映射，Printf 会自动对映射值按照键的字典顺序排序。

当然，映射中的键可能按任意顺序输出。当打印结构体时，改进的格式 %+v 会为结构体的每个字段添上字段名，而另一种格式 %#v 将完全按照 Go 的语法打印值。
```
type T struct {
    a int
    b float64
    c string
}
t := &T{ 7, -2.35, "abc\tdef" }
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)
```
输出
```
&{7 -2.35 abc   def}
&{a:7 b:-2.35 c:abc     def}
&main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}
```
（请注意其中的 & 符号）当遇到 string 或 []byte 值时， 可使用 %q 产生带引号的字符串；而格式 %#q 会尽可能使用反引号。 （%q 格式也可用于整数和符文，它会产生一个带单引号的符文常量。） 此外，%x 还可用于字符串、字节数组以及整数，并生成一个很长的十六进制字符串， 而带空格的格式（% x）还会在字节之间插入空格。

```
fmt.Printf("%qn", t)
//&{'\a' %!q(float64=-2.35) "abc\tdef"}

fmt.Printf("%#q\n", t)
//&{'\a' %!q(float64=-2.35000) `abc	def`}

fmt.Printf("%x\n", 17)
//11
```

`%T`输出一个值得类型
```
fmt.Printf("%T\n", timeZone)
//map[string]int
```

若你想控制自定义类型的默认格式，只需为该类型定义一个具有 `String() string` 签名的方法。对于我们简单的类型 `T`，可进行如下操作。

```
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)
//7/-2.35/"abc\tdef"
```
Sprintf会调用类型的String方法。所以在String类型使用Sprintf会导致无限递归
```
// 请勿通过调用 Sprintf 来构造 String 方法，因为它会无限递归你的的 String 方法
type MyString string
func (m MyString) String() string {
   return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
}
```

要解决这个问题也很简单：将该实参转换为基本的字符串类型，它没有这个方法。
```
type MyString string
func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}
```

在 初始化 一节中，我们将看到避免这种递归的另一种技术。

另一种打印技术就是将打印例程的实参直接传入另一个这样的例程。Printf 的签名为其最后的实参使用了 ...interface{} 类型，这样格式的后面就能出现任意数量，任意类型的形参了。
```
func Printf(format string, v ...interface{}) (n int, err error) {
```

还有很多关于打印知识点没有提及。详情请参阅 `godoc` 对 `fmt` 包的说明文档。

顺便一提，`...` 形参可指定具体的类型，例如从整数列表中选出最小值的函数 `min`，其形参可为 `...int` 类型。
```
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // largest int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
```

## Append
现在我们要对内建函数 `append` 的设计进行补充说明。`append` 函数的签名不同于前面我们自定义的 `Append` 函数。大致来说，它就像这样：
```
func append(slice []T, elements ...T) []T
```

其中的 *T* 为任意给定类型的占位符。实际上，你无法在 Go 中编写一个类型 `T` 由调用者决定的函数。这也就是为何 `append` 为内建函数的原因：它需要编译器的支持。C# 和 Java 支持泛型。

`append` 会在切片末尾追加元素并返回结果。我们必须返回结果， 原因与我们手写的 `Append` 一样，即底层数组可能会被改变。以下简单的例子
```
x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
```
将打印 [1 2 3 4 5 6]。因此 append 有点像 Printf 那样，可接受任意数量的实参。

但如果我们要像 Append 那样将一个切片追加到另一个切片中呢？ 很简单：在调用的地方使用 ...，就像我们在上面调用 Output 那样。以下代码片段的输出与上一个相同。
```
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)
```
