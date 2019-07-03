package map_test

import "testing"

// Map 的初始化
func TestMapInit(t *testing.T) {
	m1 := map[string]int{"one": 1, "tow": 2, "three": 3}
	t.Log(m1["one"])
	t.Logf("len m1=%d", len(m1))

	//
	m2 := map[int]int{}
	m2[4] = 16 // 赋值
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10 /* initial capacity */)
	t.Logf("len m3=%d", len(m3))
}

// 如何判断 map 中的 key 是 0 值还是不存在
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // 访问一个不存在的key，输出 0

	m1[2] = 0
	t.Log(m1[2]) // 访问一个存在但是值为 0 的 key，输出也是 0

	// 如何区分不存在和 0 值呢？
	m1[3] = 0

	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d.", v)
	} else {
		t.Logf("key 3 is not existing")
	}
}

// Map 迭代直接使用： for k,v := range <map>
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
