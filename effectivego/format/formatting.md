
[Effective Go - The Go Programming Language (google.cn)](https://golang.google.cn/doc/effective_go#formatting)

[格式化 |《高效的 Go 编程 Effective Go 2020》| Go 技术论坛 (learnku.com)](https://learnku.com/docs/effective-go/2020/format/6237)

>格式化问题是备受争议最多的一个话题，每个人可以适应不通的编码风格，若所有人都遵循相同的编码风格，在这类问题上浪费的时间会减少。

Go语言提供了`go fmt`，让机器来处理风格缩进、对齐、保留注释并在需要时重新格式化。

![image.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/f908ffc6bd174a6c93e51b04efd1e51a~tplv-k3u1fbpfcp-watermark.image?)

借助IDE（GoLand），通过快捷键也能快速及时帮我们格式化代码。
![image.png](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/80bdc6250bf345fa8d239dc8c2b0c6a4~tplv-k3u1fbpfcp-watermark.image?)

## 缩进

我们使用制表符（tab）缩进，gofmt 默认也使用它。在你认为确实有必要时再使用空格。

## 行的长度

![image.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/07f5c72118af4aac959f8dc0ad32cc79~tplv-k3u1fbpfcp-watermark.image?)
Go 对行的长度没有限制，别担心打孔纸不够长。如果一行实在太长，也可进行折行并插入适当的 tab 缩进。


## 括号

比起 C 和 Java，Go 所需的括号更少：控制结构（if、for 和 switch）在语法上并不需要圆括号。
如下实例，Go的语法结构比Java、C、C#在控制语句中省略了括号。
```
func control(x int) {
   if x == 1 {
      fmt.Println("1")
   }

   switch x {
   case 1:
      fmt.Println("switch 1")
   case 2:
      fmt.Println("switch 2")
   }

   for x < 10 {
      x++
   }

   for i := 0; i < 10; i++ {
      
   }
}
```
此外，操作符优先级处理变得更加简洁。（个人不确定这个优化处理是否妥当。）
```
func main() {
   x := 1
   var result = x<<8 + 1
   fmt.Println(result)
}
```
符号优先级中，`<<`的优先级是低于`+`的。在其他语言中，该结果会是`512`，而Go的简化，结果为`256`.