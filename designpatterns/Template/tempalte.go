package Template

import "fmt"

type Cooker interface {
	fire()
	cook()
	outFire()
}

type CookTemplate struct {
}

type CookTemplate2 struct {
}

func (CookTemplate2) fire() {
	fmt.Println("fire2")
}

func (CookTemplate) fire() {
	fmt.Println("fire")
}

func (CookTemplate) cook() {
	fmt.Println("cook")

}

func (CookTemplate) outFire() {
	fmt.Println("outFire")
}

type ChineseCook struct {
	CookTemplate
	CookTemplate2
}

func (ChineseCook) fire() {
	fmt.Println("chinese fire")
}

func (ChineseCook) cook() {
	fmt.Println("chinese cook")

}

func (ChineseCook) outFire() {
	fmt.Println("chinese outFire")
}

type ForeignCook struct {
	CookTemplate
}

func (c ForeignCook) cook() {
	println("ForeignCook cook")
}

//func (c ForeignCook) outFire() {
//	println("ForeignCook outFire")
//
//}
//
//func (ForeignCook) fire()  {
//	fmt.Println("ForeignCook fire")
//}

func DoCook(cook Cooker) {
	cook.fire()
	cook.cook()
	cook.outFire()
}
