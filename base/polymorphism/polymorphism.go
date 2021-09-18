package polymorphism

import "fmt"

//多态学习

type Person interface {
	Introduction()
	GetName() string
}

type Info struct {
	Name string
	Age  int
}

type Student struct {
	Info
}

type Teacher struct {
	Info
}

func (receiver Teacher) Introduction() {
	fmt.Println("I am teacher:" + receiver.Name)

}

func (receiver Student) Introduction() {
	fmt.Println("I am student：" + receiver.Name)
}

func (info Info) GetName() string {
	return info.Name
}

// Say 多态
func Say(person Person) {
	person.Introduction()
	person.GetName()
}

func DemoRun() {
	student := Student{Info{Name: "Moon"}}
	student.Introduction()
	teacher := Teacher{Info{Name: "Moon"}}
	Say(student)
	Say(teacher)
}

type error interface {
	Error() string
}

func (receiver Teacher) Error() string {
	if receiver.Name == "xym" {
		return "xym 不能当老师"
	}

	return ""
}
