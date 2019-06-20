package loop

import "testing"

// 学习 Go 里面的条件与循环

// 简单的 for 循环
func TestWhileLoop(t *testing.T) {
	for i := 0; i < 3; i++ {
		t.Log(i)
	}
}

// while 循环
func TestWhile(t *testing.T) {
	n := 0
	for n < 3 {
		n++
		t.Log(n)
	}
}

// while(true) 无限循环
func TestWhileTrue(t *testing.T) {
	n := 0
	for {
		n++
		t.Log(n)
		if n == 3 {
			break
		}
	}
}
