package args

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	option := BooleanOption{}
	fmt.Println(option)
	valueOf := reflect.ValueOf(&option)
	valueOfOpt := valueOf.Elem()
	field := valueOfOpt.FieldByName("Logging")
	field.SetBool(true)
	fmt.Println(option)
}

//-l -p 8080 -d /usr/logs
//single option
//todo  -Bool -l
func TestBool_return_true_when_exist_l(t *testing.T) {
	option := BooleanOption{}
	result := parse(reflect.TypeOf(option), "-l").(*BooleanOption)
	assert.Equal(t, true, result.Logging)
}

//todo  -int -p 8080
//todo  -string -d /user/logs
//todo  -multi :-l -p 8080 -d /usr/logs

func TestParseArgs_(t *testing.T) {
	//arguments := parse("-l", "-p", "8080", "-d", "/usr/log")
	//assert.Equal(t, true, arguments.logging)
}
