package Strategy

type IStrategy interface {
	do(int, int) int
}

type add struct {
}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct {
}

func (*reduce) do(a, b int) int {
	return a - b
}

type Operator struct {
	Strategy IStrategy
}

func (operator Operator) Calculator(a, b int) int {
	return operator.Strategy.do(a, b)
}
