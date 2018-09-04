package geometry

import (
	"math"
	"math/rand"
)

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func NewZeroVector() *Vector {
	return NewVector(0, 0, 0)
}

func NewHemisphereVector() *Vector {
	phi := rand.Float64() * 2. * math.Pi
	rd := rand.Float64()
	r := math.Sqrt(rd)

	return NewVector(
		math.Cos(phi)*r,
		math.Sin(phi)*r,
		math.Sqrt(1.-rd),
	).Normalize()
}

func (vec *Vector) Add(vec2 *Vector) *Vector {
	return NewVector(
		vec.X+vec2.X,
		vec.Y+vec2.Y,
		vec.Z+vec2.Z,
	)
}

func (vec *Vector) Subtract(vec2 *Vector) *Vector {
	return NewVector(
		vec.X-vec2.X,
		vec.Y-vec2.Y,
		vec.Z-vec2.Z,
	)
}

func (vec *Vector) MultiplyScalar(multiplier float64) *Vector {
	return NewVector(
		vec.X*multiplier,
		vec.Y*multiplier,
		vec.Z*multiplier,
	)
}

func (vec *Vector) DivideScalar(divider float64) *Vector {
	return NewVector(
		vec.X/divider,
		vec.Y/divider,
		vec.Z/divider,
	)
}

func (vec *Vector) Cross(vec2 *Vector) *Vector {
	return NewVector(
		vec.Y*vec2.Z-vec.Z*vec2.Y,
		vec.Z*vec2.X-vec.X*vec2.Z,
		vec.X*vec2.Y-vec.Y*vec2.X,
	)
}

func (vec *Vector) Dot(vec2 *Vector) float64 {
	return vec.X*vec2.X + vec.Y*vec2.Y + vec.Z*vec2.Z

}

func (vec *Vector) Normalize() *Vector {
	magnitude := vec.Magnitude()
	return vec.DivideScalar(magnitude)
}

func (vec *Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.X, 2.) + math.Pow(vec.Y, 2.) + math.Pow(vec.Z, 2.))
}

func (vec *Vector) RotateTowards(normal *Vector) *Vector {
	upVector := NewVector(0, 0, 1)

	if normal.Z > .999999 {
		return NewVector(vec.X, vec.Y, math.Abs(vec.Z))
	} else if normal.Z < -.999999 {
		return NewVector(vec.X, vec.Y, -math.Abs(vec.Z))
	}

	a1 := upVector.Cross(normal)
	a2 := a1.Cross(normal)

	p1 := a1.MultiplyScalar(vec.X)
	p2 := a2.MultiplyScalar(vec.Y)
	p3 := normal.MultiplyScalar(vec.Z)

	return p1.Add(p2).Add(p3).Normalize()
}
