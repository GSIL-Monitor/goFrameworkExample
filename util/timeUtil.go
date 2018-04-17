package util

import (
	"fmt"
	"reflect"
)

func printTimeElapsed(fn interface{}) {
	var decoratedFunc, targetFunc reflect.Value

	decoratedFunc = reflect.ValueOf(decoPtr).Elem()
	targetFunc = reflect.ValueOf(fn)

	v := reflect.MakeFunc(targetFunc.Type(),
		func(in []reflect.Value) (out []reflect.Value) {
			fmt.Println("before")
			out = targetFunc.Call(in)
			fmt.Println("after")
			return
		})

	decoratedFunc.Set(v)
	return

}

func b(i int32) int32 {
	return i
}

func a() {
	printTimeElapsed(b, 1)
}
