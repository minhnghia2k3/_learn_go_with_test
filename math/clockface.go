package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

// convert point in circle
// to pixel in svg
//
//	(0,0)─────────────► X
//	   │
//	   │     {150,150} ← tâm đồng hồ ở giữa
//	   │
//	   ▼ Y
func SecondHand(tm time.Time) Point {
	p := secondHandPoint(tm)
	p = Point{p.X * 90, p.Y * 90}   // scale
	p = Point{p.X, -p.Y}            // flip
	p = Point{p.X + 150, p.Y + 150} // translate
	return p
}

func SecondInRadian(tm time.Time) float64 {
	return math.Pi / (30 / float64(tm.Second()))
}

// X = Sin = truc hoanh
// Y = Cos = truc tung
// in unit-circle
func secondHandPoint(tm time.Time) Point {
	rad := SecondInRadian(tm)
	x := math.Sin(rad)
	y := math.Cos(rad)

	return Point{x, y}
}
