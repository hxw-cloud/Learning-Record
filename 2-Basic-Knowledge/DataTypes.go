package __Basic_Knowledge

import (
	"fmt"
	"strconv"
)

// 变量声明
func Variable(a, b int) {

	fmt.Println(a)
	fmt.Println(b)
	var (
		name = "hxw"
		age  = 13
		sex  = "man"
	)

	fmt.Println(name + strconv.Itoa(age) + sex)

}

func Constant() {

	const name = "myName"
	const age int = 32
	const a1, a2, a3 = 1, 2, "43"
	fmt.Println(name)
	//const 出现时被重置为0，每出现一次自动加1
	const (
		F = iota
		G = iota
		H = iota
	)

	const (
		I = iota
		J
		K
	)

}

func TypeConversion() {
	var aInt int = 32
	// 一般用这种方式强制转
	fmt.Printf("转float64 %f  \n", float64(aInt))
	fmt.Printf("转string %v  \n", strconv.Itoa(aInt))
	fmt.Printf("转float32 %f  \n", float32(aInt))
}
