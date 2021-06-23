package polymorphism

import "fmt"

// Person go 多态
type Person interface {
	Call()
}

type Student struct {
	Name string
}

type Teacher struct {
	Name string
}

func (receiver Teacher) Call() {
	fmt.Println("I am teacher:" + receiver.Name)
}

func (receiver Student) Call() {
	fmt.Println("I am student：" + receiver.Name)
}

func Say(person Person) {
	person.Call()
}

func DemoRun() {
	student := Student{Name: "Moon"}
	teacher := Teacher{Name: "Moonus"}
	Say(student)
	Say(teacher)
}
