package gocvKit

import "image"

type (
	Point struct {
		image.Point
	}
)

func NewPoint(x, y int) *Point {
	return &Point{
		Point: image.Point{X: x, Y: y},
	}
}
