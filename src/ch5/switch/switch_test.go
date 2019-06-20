package _switch

import "testing"

// switch 条件
func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Larry")
		default:
			t.Log("It is not 0-3")
		}
	}
}

// switch case ,case 表达式
func TestSwitchCaseMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0: // 偶数
			t.Log("Even")
		case i%2 == 1: // 基数
			t.Log("Old")
		default:
			t.Log("unknown")
		}
	}
}
