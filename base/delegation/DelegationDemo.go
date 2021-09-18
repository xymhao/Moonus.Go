package delegation

import (
	"fmt"
)

//结构体嵌入
//我们把 Widget嵌入到了 Label 中

type Widget struct {
	X, Y int
}

type Label struct {
	Widget //  (delegation)
	Text   string
	X      int
}

func (label Label) Paint() {
	fmt.Printf("[%p] - Label.Paint(%q)\n", &label, label.Text)
}

type Button struct {
	Label // Embedding (delegation)
}

func NewButton(x, y int, text string) Button {
	return Button{Label{Widget{x, y}, text, x}}
}

// Paint Button.Paint() 接口可以通过 Label 的嵌入带到新的结构体，如果 Button.Paint() 不实现的话，会调用 Label.Paint() ，所以，在 Button 中声明 Paint() 方法，相当于 Override。
func (button Button) Paint() { // Override
	fmt.Printf("[%p] - Button.Paint(%q)\n",
		&button, button.Text)
}
func (button Button) Click() {
	fmt.Printf("[%p] - Button.Click()\n", &button)
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

func (listBox ListBox) Paint() {
	fmt.Printf("[%p] - ListBox.Paint(%q)\n",
		&listBox, listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("[%p] - ListBox.Click()\n", &listBox)
}

type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}
