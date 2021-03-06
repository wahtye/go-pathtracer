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
	return NewVector(
		rand.Float64()*2.-1.,
		rand.Float64()*2.-1.,
		rand.Float64(),
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

func (vec *Vector) Multiply(vec2 *Vector) *Vector {
	return NewVector(
		vec.X*vec2.X,
		vec.Y*vec2.Y,
		vec.Z*vec2.Z,
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

func (vec *Vector) Sq() *Vector {
	return vec.Multiply(vec)
}

func (vec *Vector) Sqrt() *Vector {
	return NewVector(
		math.Sqrt(vec.X),
		math.Sqrt(vec.Y),
		math.Sqrt(vec.Z),
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

func (vec *Vector) Invert() *Vector {
	return NewVector(
		1./vec.X,
		1./vec.Y,
		1./vec.Z,
	)
}

func (vec *Vector) GetAxis(axis Axis) float64 {
	switch axis {
	case AxisX:
		return vec.X
	case AxisY:
		return vec.Y
	case AxisZ:
		return vec.Z
	}

	return 0
}

func (vec *Vector) RotateTowards(normal *Vector) *Vector {
	upVector := NewVector(0, 0, 1)

	if normal.Z > .999999 {
		return NewVector(vec.X, vec.Y, math.Abs(vec.Z))
	} else if normal.Z < -.999999 {
		return NewVector(vec.X, vec.Y, -math.Abs(vec.Z))
	}

	a1 := upVector.Cross(normal).Normalize()
	a2 := a1.Cross(normal).Normalize()

	p1 := a1.MultiplyScalar(vec.X)
	p2 := a2.MultiplyScalar(vec.Y)
	p3 := normal.MultiplyScalar(vec.Z)

	rotatedVector := p1.Add(p2).Add(p3).Normalize()
	return rotatedVector
}

func (vec *Vector) Min(vec2 *Vector) *Vector {
	return NewVector(
		math.Min(vec.X, vec2.X),
		math.Min(vec.Y, vec2.Y),
		math.Min(vec.Z, vec2.Z),
	)
}

func (vec *Vector) Max(vec2 *Vector) *Vector {
	return NewVector(
		math.Max(vec.X, vec2.X),
		math.Max(vec.Y, vec2.Y),
		math.Max(vec.Z, vec2.Z),
	)
}

func (vec *Vector) Average(vec2 *Vector) *Vector {
	return NewVector(
		(vec.X+vec2.X)/2.,
		(vec.Y+vec2.Y)/2.,
		(vec.Z+vec2.Z)/2.,
	)
}
