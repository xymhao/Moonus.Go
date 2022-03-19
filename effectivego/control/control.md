Go的控制结构与C的控制结构有一定的联系，但在重要的方面又有些许不同。Go不在使用`do`和`while`做循环，只有一个更为通用的`for`;`switch`的使用更加的灵活，`if`和`switch`可以像`for`一样可以接受可选的初始化语句；
```
func main() {
   for true {
      fmt.Println("i am while")
   }

   if err := Chmod(0664); err != nil {
      log.Print(err)
   }
}

func Chmod(i int) error {
   return errors.New("1234")
}
```
### if
```
if x > 0 {
   return
}
```

在 Go 库中，你会发现当一个 if 语句不流入下一个语句，即主体以 break、 continue、 goto 或 return 结束时，不必要的 else 就被省略了。大多数情况我们都可以不使用else，过多的控制语句会加大我们的认知负担。

下例是一种常见的情况，代码必须防范一系列的错误条件。若控制流成功继续， 则说明程序已排除错误。由于出错时将以 `return` 结束， 之后的代码也就无需 `else` 了。
```
func demo2() error {
   f, err := os.Open("")
   if err != nil {
      return err
   }
   d, err := f.Stat()
   if err != nil {
      f.Close()
      return err
   }
   codeUsing2(f, d)
   return nil
}
```

### Redeclaration and reassignment 重新声明和重新分配
最后一个示例展示了短声明 := 如何使用。 调用了 os.Open 的声明为
```
f, err := os.Open(name)

d, err := f.Stat()
```
紧接着，err在第一条语句赋值后，err又被`f.Stat()`重新赋值，这种方式是合法的，`f.Start`使用的只是前面申明的`err`，他只是重新赋值了而已。

在满足下列条件时，已被声明的变量 err 可出现在:= 声明中：

- 本次声明与已声明的 err 处于同一作用域中（若 err 已在外层作用域中声明过，则此次声明会创建一个新的变量 §），
- 在初始化中与其类型相应的值才能赋予 err，且
- 在此次声明中至少另有一个变量是新声明的。

### for
Go的for循环类似于C，但不同于C。它包含`for`和`while`，不再有`do-while`。他有三种形式，但只有一种需要分号。
```
func demoFor() {
   //while
   for true {
      fmt.Println("i am while")
   }

   //regular for
   for i := 0; i < 10; i++ {

   }

   // forr
   users := make([]int, 5)
   for i, user := range users {
      fmt.Println(i, user)
   }
}
```

若你只需要该遍历中的第二个项（值），请使用**空白标识符**，即下划线来丢弃第一个值：
```
//_表示空白标识符
for _, user := range users {
   fmt.Println(user)
}
```

对于字符串，range 能够提供更多便利。它能通过解析 UTF-8， 将每个独立的 Unicode 码点分离出来。错误的编码将占用一个字节，并以符文 U+FFFD 来代替。 （名称 “符文” 和内建类型 rune 是 Go 对单个 Unicode 码点的成称谓。 详情见语言规范）。循环
```
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```
输出结果
```
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

最后，Go 没有逗号操作符，而 `++` 和 `--` 为语句而非表达式。 因此，若你想要在 `for` 中使用多个变量，应采用平行赋值的方式 （因为它会拒绝 `++` 和 `--`）
```
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

### switch
Go的swtich相较于C的更为通用，他不局限于常量、整数，case语句会自上而下的逐一进行求值匹配，若没有表达式他将匹配true，所以我们可以将`if-else-if-else`写成`switch`,这也更符合Go的风格。
```
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

`switch` 并不会自动下溯，但 `case` 可通过逗号分隔来列举相同的处理条件。
```
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```
尽管它们在 Go 中的用法和其它类 C 语言差不多，但 break 语句可以使 switch 提前终止。不仅是 switch， 有时候也必须打破层层的循环。在 Go 中，我们只需将标签放置到循环外，然后 breaking 到那里即可。下面的例子展示了二者的用法。
```
func demoBreak(src []int, size int, sizeOne, sizeTwo int, validateOnly bool) {
Loop:
   for n := 0; n < len(src); n += size {
      switch {
      case src[n] < sizeOne:
         if validateOnly {
            break
         }
         size = 1
         update(src[n])

      case src[n] < sizeTwo:
         if n+1 >= len(src) {
            break Loop
         }
         if validateOnly {
            break
         }
         size = 2
         update(src[n] + src[n+1]<<1)
      }
   }
}

func update(i int) {

}
```
`continue` 语句也接受一个可选的标签，但它只应用于循环

### Type switch
switch 也可用于判断接口变量的动态类型。如 类型选择 通过圆括号中的关键字 type 使用类型断言语法。通常我们也叫做模式匹配。
```
func typeSwitch() {
   var t interface{}
   t = functionOfSomeType()
   switch t := t.(type) {
   default:
      fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
   case bool:
      fmt.Printf("boolean %t\n", t)             // t has type bool
   case int:
      fmt.Printf("integer %d\n", t)             // t has type int
   case *bool:
      fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
   case *int:
      fmt.Printf("pointer to integer %d\n", *t) // t has type *int
   }
}

func functionOfSomeType() interface{} {
   return true
}
```
