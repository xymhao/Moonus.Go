package Mock

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	// ...
	println(f.Bar(99))
	println(f.Bar(98))
}
