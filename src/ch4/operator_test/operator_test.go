package operator_test

import "testing"

//常量 连续常量
const (
	Readable   = 1 << iota //1 << 0 which is 0000 0001
	Writable               // 1 << 1 which is 0000 0010
	Executable             // 1 << 2 which is 0000 0100
)

func TestBitClear(t *testing.T) {
	a := 7 //0000 0111
	//0111 & 0001 = 0001
	//go语言中按位取反写法是^, 所以 a&^b 其实是 a&(^b) 利用运算符优先级省略掉括号的写法而已
	//其它语言中关键字不同写法可能是 a&(~b), 但算法是一样的
	// ^0001 = 1110  0111 & 1110 = 0110
	// 0000 0111  &^ 0000 0001 = 0000 0110
	// 0000 0111  &^ 0000 0010 = 0000 0101
	// 0000 0111  &^ 0000 0100 = 0000 0011

	// 1&^0 1
	// 1&^1 0
	// 0&^1 0
	// 0&^0 0
	a = a &^ Readable //去掉可读权限
	//a := 1
	t.Log(a&Readable == Readable)
	t.Log(a&Writable == Writable)
	t.Log(a&Executable == Executable)
}

func TestCompareArray(t *testing.T) {

	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//c := [...]int{1, 2, 3, 4, 5, 6}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//invalid operation: a == c (mismatched types [4]int and [6]int)
	//t.Log(a == c)
	t.Log(a == d)

}
