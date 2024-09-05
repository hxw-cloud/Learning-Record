package __Basic_Knowledge

import "fmt"

func Cycle() {
L1:
	fmt.Println("hello world")
	for i := 0; i < 15; i++ {
		if i == 5 {
			goto L1
		}
		fmt.Println(i) // Output: 0 1 2 3 4 5
		if i == 13 {
			break
		} else if i == 6 {
			continue
		} else if i%9 == 0 {
			break
		}
	}
}

func RangeCycle() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for k, v := range nums {
		fmt.Printf("key: %v , value: %v \n", k, v) // Output: 0 1 2 3 4 5 6 7 8 9
		if v%9 == 0 {
			break
		} else if v%3 == 0 {
			continue
		}
	}
	for _, v := range nums {
		fmt.Printf("value: %v \n", v)
	}
}
