package main

import (
	"fmt"
	"reflect"
)

func task14(t interface{}) {
	fmt.Println(reflect.TypeOf(t))
}
