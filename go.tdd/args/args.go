package args

import (
	"fmt"
	"reflect"
)

type Options struct {
	logging  bool   `option:"-l"`
	port     int    `option:"p"`
	director string `option:"d"`
}

type BooleanOption struct {
	Logging bool `option:"-l"`
}

func parse(typeOpt reflect.Type, args ...string) interface{} {
	valueOpt := reflect.New(typeOpt)
	num := typeOpt.NumField()
	for i := 0; i < num; i++ {
		field := typeOpt.Field(i)
		fieldVal := valueOpt.Elem().Field(i)
		fmt.Println(field.Name, field.Tag.Get("option"))
		if args[i] == field.Tag.Get("option") {
			if field.Type.Kind() == reflect.Bool {
				fieldVal.SetBool(true)
			}
		}

	}
	return valueOpt.Interface()
}
