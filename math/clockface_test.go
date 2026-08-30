package clockface

import (
	"math"
	"testing"
	"time"
)

// Cong thuc:
// 360° = 2π 	= 60s
//
//	π/30	= 1s
//
// => độ = giây * π/30
// SecondsInRadian take time, and convert the second to radian
func TestSecondInRadian(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, c := range cases {

		t.Run(testName(c.time), func(t *testing.T) {
			got := SecondInRadian(c.time)

			if got != c.angle {
				t.Errorf("got %v, want %v", got, c.angle)
			}
		})

	}
}

func TestSecondHand(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			got := secondHandPoint(c.time)
			if !roughlyEqualPoint(got, c.point) {
				t.Errorf("got %v, want %v", got, c.point)
			}

		})
	}

}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(tm time.Time) string {
	return tm.Format("15:04:05")
}

func roughlyEqualFloat64(a, b float64) bool {
	equalThreshold := 1e-7 // 0,0000001 (hay 1 × 10⁻⁷)
	return math.Abs(a-b) < equalThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
