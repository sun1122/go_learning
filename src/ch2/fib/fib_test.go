package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	//var temp int
	//fmt.Print(a)
	t.Log(" ", a)
	// := 是声明并赋值，并且系统自动推断类型，不需要var关键字
	for i := 0; i < 5; i++ {
		t.Log(" ", b)

		temp := a
		a = b
		b = temp + a
		// a, temp = temp, temp+a
		// t.Log(" ", temp)

	}
	fmt.Println()
}

func TestExchange(z *testing.T) {
	a := 1
	b := 2

	// temp := a
	// a = b
	// b = temp + a

	a, b = b, a
	z.Log(a, b)

}
