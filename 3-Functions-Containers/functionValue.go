package __Functions_Containers

import "fmt"

func FunctionValue(a, b int) (res int) {

	func(a, b *int) {
		*a = *a / *b
		*b = *a * *b
	}(&a, &b)

	res = a + b
	return

}

func FunctionValueTest(a, b int, res func(a, b int) int) {
	print(res(a, b)) // Output: 15
}

// Circle 定义结构体
type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

func sum(t ...int) (res int) {
	for _, v := range t {
		res += v
	}
	return res
}

func sumNum(t ...interface{}) (res float64) {
	for _, tmp := range t {
		switch v := tmp.(type) {
		case int, int8, int16, int32, int64:
			res += float64(v.(int))
		case uint, uint8, uint16, uint32, uint64:
			res += float64(v.(uint))
		case float32:
			res += float64(v)
		case float64:
			res += v
		case complex64:
			res += float64(real(v))
		case complex128:
			res += float64(real(v))
		default:
			// 不支持的类型可以选择返回错误或忽略
		}
	}
	return res
}
