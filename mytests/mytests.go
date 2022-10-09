package main

import (
	"fmt"
	"strings"
)

type Shape interface {
	Area() float64
	Test() string
}

type Skin interface {
	Color() float64
}

type Cube struct {
	side   float64
	radius float64
}

func (c Cube) Area() float64 {
	fmt.Printf("memory address of c is %p\n", &c)
	return c.side * c.side * c.side * c.radius
}

func (c *Cube) Test() string {
	fmt.Printf("memory address of *c is %p\n", c)

	return "test"
}

func main() {
	rr := Cube{side: 2, radius: 3}
	var r Shape = &rr
	fmt.Printf("memory address of r is %p\n", &r)
	r.Area()
	r.Test()

	// var s Shape = &r
	// s.Area()
	// s.Test()
	// explain(s, "s")

	// v, ok := s.(Cube)
	// fmt.Println("v", "ok:", ok)
	// explain(v, "v")

}

func explain(v interface{}, key string) {
	fmt.Printf("%v value is %v\n", key, v)
	fmt.Printf("%v type is %T\n", key, v)
	switch v.(type) {
	case string:
		fmt.Printf("%v in uppercase if string", strings.ToUpper(v.(string)))
	case Cube:
		fmt.Printf("%v is a cube", v.(Cube))
	}

}
