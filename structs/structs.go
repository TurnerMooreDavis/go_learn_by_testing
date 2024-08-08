package structs

import (
	"math"
	"reflect"
)

type Rectangle struct {
	Width  float32
	Height float32
}

type Shape interface {
	// total area of the shape
	Area() float32
	// get the type of the struct
	getStruct() reflect.Type
}

func (r Rectangle) Perimeter() (ret float32) {
	ret = 2 * (r.Width + r.Height)
	return
}
func (r Rectangle) Area() (ret float32) {
	ret = r.Width * r.Height
	return
}
func (r Rectangle) getStruct() (ret reflect.Type) {
	ret = reflect.TypeOf(r)
	return
}

type Circle struct {
	Radius float32
}

func (r Circle) Area() (ret float32) {
	ret = 0
	ret = r.Radius * r.Radius * math.Pi
	return
}

func (r Circle) getStruct() (ret reflect.Type) {
	ret = reflect.TypeOf(r)
	return
}

func RectArea(r Rectangle) (ret float32) {
	ret = r.Width * r.Height
	return
}
