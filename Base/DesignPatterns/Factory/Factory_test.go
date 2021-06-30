package Factory

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPerson(t *testing.T) {
	newPerson := NewPerson("张三", 4)
	person2 := NewPerson2("李四", 5)
	fmt.Println(newPerson, person2)
	fmt.Println(*newPerson, person2)

	newPerson.Greet()
	person2.Greet2()

	println(person2.name)
}

func TestUpdate(t *testing.T) {
	person := NewPerson("xym", 99)
	fmt.Println(person)
	updateAge(person)
	fmt.Println(person.Age)
	assert.Equal(t, 0, person.Age)

	updateName(*person)
	fmt.Println(person.Name)
}

func updateAge(person *Person) {
	person.Age = 0
}

func updateName(person Person) {
	person.Name = "update"
}

func TestUpdate2(t *testing.T) {
	person2 := NewPerson2("xym", 12)
	updateAge2(&person2)
	assert.Equal(t, 0, person2.age)

	updateName2(person2)

	assert.Equal(t, "xym", person2.name)
}

func updateAge2(person *Person2) {
	person.age = 0
}

func updateName2(person Person2) {
	person.name = "update"
}

func TestNewHTTPClient(t *testing.T) {
	//client := NewHTTPClient()
	//httpClient := mockHTTPClient{}
}

func TestQueryUser(t *testing.T) {
	doer := NewMockHTTPClient()
	if err := QueryUser(doer); err != nil {
		t.Errorf("QueryUser failed, err: %v", err)
	}
}

func TestNewPersonFactory(t *testing.T) {
	newBaby := NewPersonFactory(1)
	baby := newBaby("john")
	println(baby.name, baby.age)

	newTeenager := NewPersonFactory(16)
	teen := newTeenager("jill")
	println(teen.name, teen.age)

}
