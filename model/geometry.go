package model

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// represents a circle, including the horizontal and vertical coordinates and radius
type Circle struct {
	Point
	R float64
}