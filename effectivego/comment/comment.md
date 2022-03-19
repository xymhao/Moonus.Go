# 代码注释

[Effective Go - The Go Programming Language (google.cn)](https://golang.google.cn/doc/effective_go#commentary)
[代码注释 |《高效的 Go 编程 Effective Go 2020》| Go 技术论坛 (learnku.com)](https://learnku.com/docs/effective-go/2020/code-annotation/6238)

Go提供了C语言风格的代码块注释方式`/**/`和C++风格的行注释`//`。行注释是我们经常使用的。块注释通常用于组件的注释。通常也用于注释一大段代码。

`godoc`是一个程序，也是一个Web应用服务。它对Go的源文件进行处理，导出包中的文档。顶层的注释申明，作为该条目的说明文档，这些注释的类型和风格决定了 `godoc` 生成的文档质量

每个包都应该有包注释，对于包含多个文件，包注释只需要保存在任意一个文件中即可。包注释应在整体上对该包进行介绍，并提供包的相关信息。它将出现在 `godoc` 页面中的最上面，并为紧随其后的内容建立详细的文档。

![image.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/ec34e89c17704bf99bbc566e28f34546~tplv-k3u1fbpfcp-watermark.image?)

Go's declaration syntax allows grouping of declarations. A single doc comment can introduce a group of related constants or variables. Since the whole declaration is presented, such a comment can often be perfunctory.

Go 的声明语法允许成组声明。单个文档注释应介绍一组相关的常量或变量。 由于是整体声明，这种注释往往较为笼统。
```
// Error codes returned by failures to parse an expression.
var (
    ErrInternal      = errors.New("regexp: internal error")
    ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
    ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
    ...
)
```

Grouping can also indicate relationships between items, such as the fact that a set of variables is protected by a mutex.
即便是对于私有名称，也可通过成组声明来表明各项间的关系，例如某一组由互斥体保护的变量。
```
var (
    countLock   sync.Mutex
    inputCount  uint32
    outputCount uint32
    errorCount  uint32
)
```