package main

import (
	"bufio"
	e "errors"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

//Error vs Exception

//error的优势
//简单
//考虑失败，而不是成功
//没有隐藏的控制流
//完全交给你来控制 error
//Error are values

// ErrDemo 参考：bufio.go
var (
	ErrDemo = errors.New("{包名}: invalid use of UnreadByte")
)

//创建自定义error类型
type errorMo struct {
	s string
	e error
}

func (err errorMo) Error() string {
	return string(err.s)
}

func New(text string) error {
	e2 := errorMo{s: text}
	return &e2
}

var ErrMoonus = New("EOF")
var ErrGo = e.New("EOF")

func main() {
	wrapErr()
	panicAndRecover()
	errN := New("EOF")
	//is true
	if ErrMoonus == errN {
		fmt.Println("custom ErrMoonus equal New(EOF)")
	}

	//is false
	if ErrGo == e.New("EOF") {
		fmt.Println("errors.New equal")
	}

	main2()

	err := demo()
	if err != nil {
		println(err.Error())
	}

	fmt.Println("end")
}

func wrapErr() bool {
	err := ErrDemo
	err2 := errors.Wrap(err, "warp")
	if e.Is(err2, ErrDemo) {
		return true
	}
	return false
}

func warp() error {
	//"github.com/pkg/errors"
	return errors.Wrap(errors.New("err"), "")
}

func OpaqueErr() bool {
	err := New("errorMo")
	_, ok := err.(*errorMo)
	return ok
}

func (err errorMo) Temporary() bool {
	return true
}

func (err errorMo) IsMoonusCall() bool {
	isTemporary := IsTemporary(err)
	return isTemporary
}
func panicAndRecover() {
	defer func() {
		errR := recover()
		fmt.Println(errR)
	}()

	panic("demo panic")
}

func demo() error {
	return ErrDemo
}

func appCrash() {
	//表示不可恢复的程序错误
	panic("inconceivable")
}

//Sentinel Error
//哨兵： 预定义error 参考：bufio.go

// MyError Error Types
// 参考 PathError
type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d:%s", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{"err file", "err.go", 42}
}

// Opaque errors
//不透明的错误处理-只需要返回错误而不假设其内容

type temporary interface {
	Temporary() bool
}

// IsTemporary Assert errors for behaviour, not type; 对象对外暴露行为，而不是错误类型
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

//indented flow is for errors
//无错误的正常流程代码，将会是一条直线，而不是缩进代码
func indented() {
	_, err := os.Open(os.Stdout.Name())

	//correct
	if err != nil {
		// handle err
	}

	//do stuff

	_, err = os.Open(os.Stdout.Name())
	//not good
	if err == nil {
		// do stuff
	}
	//handle err
}

func CountLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}

// 封装error
type errWriter struct {
	io.Writer
	err error
}

func (e errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}

	n := 0
	n, e.err = e.Write(buf)
	return n, nil
}

// Wrap errors
// go error 没有堆栈信息，异常从底层抛向顶层，难以排查
// you should only handle errors once.

//错误要被日志记录
//应用程序处理错误
//滞后不在报告当前错误
// 第三方库 pkg/errors
//公共组件标准库不建议依赖errors，第三方库

func WriteConfig(w io.Writer) error {
	_, err := os.Open("")
	if err != nil {
		err = errors.Wrapf(err, "open fail")
	}
	return err
}

func main2() {
	err := WriteConfig(os.Stdout)
	if err != nil {
		fmt.Printf("original error : %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace : \n %+v\n", err)
	}

}

func fileOp() error {
	_, err := os.Open("/path")
	if err != nil {
		return err
	}
	// read or write
	return nil
}
