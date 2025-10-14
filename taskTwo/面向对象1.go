package main

import (
	"fmt"
	"math"
)

/*
*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
func main() {
	rectangle := Rectangle{25, 10}
	circle := Circle{10}
	fmt.Printf("矩形面积：%.2f\n", rectangle.Area())
	fmt.Printf("矩形周长：%.2f\n", rectangle.Perimeter())
	fmt.Printf("圆形面积：%.2f\n", circle.Area())
	fmt.Printf("圆形周长：%.2f\n", circle.Perimeter())
}

type Shape interface {
	Area() float64      //计算面积
	Perimeter() float64 //计算周长
}

// 矩形
type Rectangle struct {
	length, width float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// 圆形
type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	math.Pow(c.radius, 2)
	return math.Pi * math.Pow(c.radius, 2)
}
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
