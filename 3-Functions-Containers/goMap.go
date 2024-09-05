package __Functions_Containers

import (
	"fmt"
)

func FuncMap() {

	var m = make(map[int]int)

	m[10] = 12
	m[11] = 13
	delete(m, 1)

	for k := range m {
		fmt.Printf("key is %d and value is %d \n", k, m[k])
	}
	for k, v := range m {
		fmt.Printf("key is %d and value is %d \n", k, v)
	}
	// 如果键不存在，ok 的值为 false，value 的值为该类型的零值
	if value, ok := m[1]; ok {
		fmt.Println(1, "存在，值为：", value)
	} else {
		fmt.Println(1, " 不存在")
	}

}
