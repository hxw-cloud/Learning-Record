package __Basic_Knowledge

import (
	"fmt"
)

// Judge is a function to check if a number is positive, negative or zero.
func judge(num int) bool {

	if num < 0 {
		// 数字小于0
		fmt.Println("Negative number is not allowed")
		return false
	} else if num == 0 {
		// 数字等于0
		fmt.Println("Zero is not allowed")
		return true
	} else {
		// 数字大于0
		fmt.Println("Positive number is allowed")
		return true
	}
}

// Decide is a function to check the type of a variable.
func decide(num interface{}) {
	switch num.(type) {
	case int:
		fmt.Println("num is int", num.(int))
	case string:
		fmt.Println("num is string", num.(string))
	default:
		fmt.Println("num is other type")
	}
}
