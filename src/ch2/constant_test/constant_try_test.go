package constant_test

import "testing"

//常量 连续常量
const (
	Monday = iota + 1
	Tuesday
	Wendesday
)

const (
	Readable   = 1 << iota //1 << 0 which is 0000 0001
	Writable               // 1 << 1 which is 0000 0010
	Executable             // 1 << 2 which is 0000 0100
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry2(t *testing.T) {
	//a := 7 //0000 0111
	a := 1
	t.Log(a&Readable == Readable)
	t.Log(a&Writable == Writable)
	t.Log(a&Executable == Executable)
}
