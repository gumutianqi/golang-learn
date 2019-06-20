package contant_test

import (
	"testing"
)

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

const (
	Readable = 1 << iota
	Writable
	Executable
)

/*
 * 扩展阅读：https://studygolang.com/articles/7412
 */
func TestConstantTry1(t *testing.T) {
	t.Log(Readable, Writable, Executable)

	a := 7 // 0111
	t.Log("a&Readable=", a&Readable, "a&Writable=", a&Writable, "a&Executable=", a&Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

	a = 1 // 0001
	t.Log("a&Readable=", a&Readable, "a&Writable=", a&Writable, "a&Executable=", a&Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)

}
