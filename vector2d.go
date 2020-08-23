package main

import "math"

type Vector2d struct {
	x float64
	y float64
}

func (v Vector2d) Add(v2 Vector2d) Vector2d {
	return Vector2d{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}
}

func (v Vector2d) Sub(v2 Vector2d) Vector2d {
	return Vector2d{
		x: v.x - v2.x,
		y: v.y - v2.y,
	}
}

func (v Vector2d) Mul(v2 Vector2d) Vector2d {
	return Vector2d{
		x: v.x * v2.x,
		y: v.y * v2.y,
	}
}

func (v Vector2d) AddVal(val float64) Vector2d {
	return Vector2d{
		x: v.x + val,
		y: v.y + val,
	}
}

func (v Vector2d) MulVal(val float64) Vector2d {
	return Vector2d{
		x: v.x * val,
		y: v.y * val,
	}
}

func (v Vector2d) DivVal(val float64) Vector2d {
	return Vector2d{
		x: v.x / val,
		y: v.y / val,
	}
}

func (v Vector2d) limit(lower, upper float64) Vector2d {
	return Vector2d{
		x: math.Min(math.Max(v.x, lower), upper),
		y: math.Min(math.Max(v.y, lower), upper),
	}
}

func (v Vector2d) Dist(v2 Vector2d) float64 {
	return math.Sqrt(math.Pow(v.x-v2.x, 2) + math.Pow(v.y-v2.y, 2))
}
