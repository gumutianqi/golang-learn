package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int // 申明一个可变长度的 slice 切片

	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)

	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// 通过 make 函数申明一个 slice len 为 3，容量为 5 的切片
	s2 := make([]int, 3 /* length */, 5 /* capacity */)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 9)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

// 这个例子可以看到一个自增长的 slice, 每增加四个元素，cap 容量就 * 2
// 容量自增长跟 Java 当中 HashMap 的容量增长次相似，自增长会有性能代价
func TestSliceGrowing(t *testing.T) {
	s := []int{}

	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)
}

// 数组与切片
func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}

	// slice can only be compared to nil
	//if a == b {
	//	t.Log("...")
	//}

	// 值是可以比较的
	if a[1] == b[1] {
		t.Log("==")
	}
}
