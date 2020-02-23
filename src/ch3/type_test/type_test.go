package type_test

import "testing"

type MyInt int64

func TestTypeTry(t *testing.T) {

	var a int32 = 1
	var b int64
	//	b = a 隐示类型转换不支持
	b = int64(a)
	// 显示类型转换
	var c MyInt = 1
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T ,%T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string // 空值 不是nil
	t.Log(s)
	t.Log(len(s))
}
