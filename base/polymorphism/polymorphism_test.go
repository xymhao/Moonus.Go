package polymorphism_test

import (
	. "Moonus.Go/base/polymorphism"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestReturnTeacherName(t *testing.T) {
	p := Student{Name: "xym"}
	name := p.GetName()
	assert.Equal(t, name, "xym")
	Say(p)

	var person Person
	person = Teacher{Name: "xym"}
	person.Introduction()
}
