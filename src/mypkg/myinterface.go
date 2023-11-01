package mypkg

import "fmt"

type MyInterface interface { // объявление типа интерфейса
	MethodWithoutParameters()      // тип поддерживает этот интерфейс, если содержит этот метод...
	MethodWithParameter(float64)   // ...а также этот метод (с параметром float64)...
	MethodWithReturnValue() string // ...и этот метод (с возвращаемым значением string)
}

type MyType int // объявление типа, поддерживающего myinterface

func (m MyType) MethodWithoutParameters() { // первый обязательный метод
	fmt.Println("MethodWithoutParameters called")
}
func (m MyType) MethodWithParameter(f float64) { // второй обязательный метод (с параметром float64)
	fmt.Println("MethodWithParameter called with", f)
}
func (m MyType) MethodWithReturnValue() string { // третий обязательный метод (с возвращаемым значением string)
	return "Hi from MethodWithReturnValue"
}
func (my MyType) MethodNotInInterface() { // тип может поддерживать интерфейс даже в том случае,...
	fmt.Println("MethodNotInInterface called") // ...если содержит другие методы, не входящие в этот интерфейс
}
