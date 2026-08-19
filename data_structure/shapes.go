package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float64 {
	return math.Pi * math.Pow(r.Radius, 2)
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * 2
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
