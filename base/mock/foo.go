package mock

type foo interface {
	Bar(x int) int
}

func sut(f foo) {
	// ...
	println(f.Bar(99))
	println(f.Bar(98))
}
