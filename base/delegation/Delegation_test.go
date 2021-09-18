package delegation

import (
	"fmt"
	"testing"
)

func TestLabel(t *testing.T) {
	label := Label{Widget{10, 10}, "State", 100}

	// X=100, Y=10, Text=State, Widget.X=10
	fmt.Printf("X=%d, Y=%d, Text=%s Widget.X=%d\n",
		label.X, label.Y, label.Text,
		label.Widget.X)
	fmt.Println()
	// {Widget:{X:10 Y:10} Text:State X:100}
	// {{10 10} State 100}
	fmt.Printf("%+v\n%v\n", label, label)

	label.Paint()
}

func TestNewButton(t *testing.T) {

	label := Label{Widget{10, 10}, "State", 100}

	button1 := Button{Label{Widget{10, 70}, "OK", 10}}
	button2 := NewButton(50, 70, "Cancel")
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	fmt.Println()
	//[0xc4200142d0] - Label.Paint("State")
	//[0xc420014300] - ListBox.Paint(["AL" "AK" "AZ" "AR"])
	//[0xc420014330] - Button.Paint("OK")
	//[0xc420014360] - Button.Paint("Cancel")
	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	fmt.Println()
	//[0xc420014450] - ListBox.Click()
	//[0xc420014480] - Button.Click()
	//[0xc4200144b0] - Button.Click()
	for _, widget := range []interface{}{label, listBox, button1, button2} {
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
	}
}
func TestName(t *testing.T) {
	user := map[int]string{1: "xym", 2: "sq"}
	for i, name := range user {
		println(i)
		println(name)

	}
}
