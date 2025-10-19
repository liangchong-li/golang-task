package main

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
// 实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。

import (
	"fmt"
	"math"
)

type Shape interface {
	// 求面积
	Area() float64
	// 求周长
	Perimeter() float64
}

// 矩形
type Rectangle struct {
	length float64
	width float64
}
func (rectangle *Rectangle) Area() float64 {
	return rectangle.length * rectangle.width
}
func (rectangle *Rectangle) Perimeter() float64 {
	return 2 * (rectangle.length + rectangle.width)
}

// 圆
type Circle struct {
	radius float64
}
func (circle *Circle) Area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func (circle *Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.radius
}

func main() {
	rectangle := Rectangle {
		length: 10,
		width: 4,
	}
	areaOfRectangle := rectangle.Area()
	fmt.Println("areaOfRectangle:", areaOfRectangle)
	perimeterOfRectangle := rectangle.Perimeter()
	fmt.Println("perimeterOfRectangle:", perimeterOfRectangle)

	circle := Circle {
		radius: 5,
	}
	areaOfCircle := circle.Area()
	fmt.Println("areaOfCircle:", areaOfCircle)
	perimeterOfCircle := circle.Perimeter()
	fmt.Println("perimeterOfCircle:", perimeterOfCircle)
}
