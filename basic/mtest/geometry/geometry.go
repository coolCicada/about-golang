// geometry.go
package main

import (
	"fmt"
	"log"
	"mtest/geometry/rectangle"
)

var rectLen, rectWidth float64 = -6, -7

func init() {
	println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {

	fmt.Println("Geometrical shape properties")
	fmt.Println(rectangle.Area(rectLen, rectWidth))
	fmt.Println(rectangle.Diagonal(rectLen, rectWidth))
}
