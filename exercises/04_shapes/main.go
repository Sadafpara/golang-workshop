package main

import (
	"fmt"
	"math"
)

/*
TODO: Create a shape area calculator
Define an interface called 'Shape' with a method 'Area() float64'.
Implement the 'Shape' interface for different shapes (e.g., Circle, Rectangle, Triangle).
Create a function 'CalculateTotalArea' that takes a slice of Shapes and returns the sum of their areas.
*/

type Shape interface {
	Area() float64

}

type Rectangle struct {
	width, height float64
}

type Circle struct {
    radius float64
}

type Triangle struct {
    Base, Height float64
} 


func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5 
}


func CalculateTotalArea(shapes []Shape) float64 {
    var area float64
    for _, shape := range shapes {
        area += shape.Area()
    }
    return area
}



func main() {
    shapes := []Shape{
        Circle{radius: 5},
        Rectangle{width: 10, height: 5},
        Triangle{Base: 8, Height: 4},
    }

    totalArea := CalculateTotalArea(shapes)
    fmt.Println("Total area:", totalArea)
}




