package condition

import "testing"

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log(a)
	}

	// v，err 是函数的返回值
	//if v, err = someFunc(); err == nil {
	//	t.Log("...")
	//}
}

// 测试 if 表达式的写法
func TestMyCondition(t *testing.T) {
	a := 1
	b := 2

	if a > 0 && b > 0 {
		t.Log("true")
	}

	if c := a > 0 && b > 0; c {
		t.Log(c)
	}
}
