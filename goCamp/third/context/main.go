package main

import (
	"context"
	"fmt"
)

func main() {
	background := context.Background()
	ctxA := context.WithValue(background, "a", "a")
	ctxB := context.WithValue(ctxA, "b", "b")

	fmt.Println("ctx background:", background.Value("a"), background.Value("b")) //ctx background: <nil> <nil>
	fmt.Println("ctx a:", ctxA.Value("a"), ctxA.Value("b"))                      //ctx a: a <nil>
	fmt.Println("ctx b:", ctxB.Value("a"), ctxB.Value("b"))                      //ctx b: a b

}
