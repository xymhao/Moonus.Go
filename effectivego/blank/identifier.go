package main

import (
	"encoding/json"
	"fmt"
	"io"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

type pd struct {
	name string
	age  int
}

func (receiver pd) MarshalJSON() ([]byte, error) {
	return []byte("demo"), nil
}
func main() {
	//fd, err := os.Open("test.go")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// TODO: use fd.
	//_ = fd

	var p = pd{name: "123", age: 2}
	funcName(&p)
	var person = []pd{{name: "xym", age: 22}}
	personJson, err := json.Marshal(person)
	fmt.Println(personJson, err)
	structJson()

}

func funcName(p interface{}) {
	m, ok := p.(json.Marshaler)
	fmt.Println(m, ok)
}

type person struct {
	Name string `json:"name"` // 字段解释，可指json 字符串的名字
	Age  int    `json: age`
}

func structJson() {
	jsonStr, err := json.Marshal(person{"xym", 25})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(err, jsonStr)

	var obj person
	err = json.Unmarshal(jsonStr, &obj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(obj)
}
