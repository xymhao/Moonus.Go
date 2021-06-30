package polymorphism

import "fmt"

//多态学习

type Person interface {
	Call()
	MyName() string
}

type error interface {
	Error() string
}

type Student struct {
	Name string
}

type Teacher struct {
	Name string
}

func (receiver Teacher) Error() string {
	if receiver.Name == "xym" {
		return "xym 不能当老师"
	}

	return ""
}

func (receiver Teacher) Call() {
	s := receiver.Error()
	fmt.Println("I am teacher:"+receiver.Name, s)

}

func (receiver Student) Call() {
	fmt.Println("I am student：" + receiver.Name)
}

func (receiver Student) MyName() string {
	return receiver.Name
}

func (receiver Teacher) MyName() string {
	return receiver.Name
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
