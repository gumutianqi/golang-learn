package operator_test

import "testing"

/**
 * 数组的比较，必须保证两个数组的维数相同才能比较
 */
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	//c := [...]int{1, 2, 3}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

// 测试对变量的按位置零操作符 &^
func TestBitClear(t *testing.T) {

	// 结果是 1 2 4，貌似刚好对应上 Linux 上的[rwx] read=1, write=2, execution=4
	t.Log(Readable, Writable, Executable)

	a := 7            // 0111	 7 包含读写执行全部权限
	a = a &^ Readable // 清楚 a 的读权限

	t.Log("a=", a)

	t.Log("a&Readable=", a&Readable, "a&Writable=", a&Writable, "a&Executable=", a&Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

}
