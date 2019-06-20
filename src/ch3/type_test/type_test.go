package type_test

import "testing"

func TestImplicit(t *testing.T) {
	var a int32
	var b int64

	b = int64(a)

	t.Log(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a

	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

/**
Go 对没有赋值的变量会自动进行赋值成该变量类型的 "零值"
===================
bool : false
integer : 0
float : 0.0
string : ""
pointer, function, interface, slice, channel, map : nil
*/
func TestString(t *testing.T) {
	var s string

	t.Logf("*%s*", s)
}

/**
以下单元测试证明了：
Go 赋值是：值传递，对两个变量赋值是，是开创一个新的内存空间，然后拷贝所有元素
*/
func TestPoint2(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := a

	b[0] = 9

	t.Log(a, b)

}
