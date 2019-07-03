package array_test

import "testing"

// 数组的申明
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3}

	arr1[1] = 5

	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

// 多维数组的声明
// 结果：声明多维数组时，维数必须相同，不过不同，会将少的维数补 0
func TestDimArrayInit(t *testing.T) {
	arr := [2][3]int{{1, 2}, {3, 4, 5}}
	t.Log(arr)
}

// 数组的迭代
func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 3, 4, 5}

	// 使用传统的 for 进行迭代
	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	t.Log("=======================")

	// 类似Java里面的 forEach 迭代；
	// 语法for idx<>, item<元素> := range arr {}
	for _, e := range arr {
		t.Log(e)
	}
}

// 数组的切片，分区，与 python 的方式一致
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[3:]
	t.Log(arr3_sec)
}
