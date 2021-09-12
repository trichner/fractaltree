package main

import "math"

type AABB struct {
	Origin Vector
	Size   Vector
}

func (a AABB) Union(bb AABB) AABB {
	minX := math.Min(a.Origin.X, bb.Origin.X)
	minY := math.Min(a.Origin.Y, bb.Origin.Y)
	maxX := math.Max(a.Origin.X+a.Size.X, bb.Origin.X+bb.Size.X)
	maxY := math.Max(a.Origin.Y+a.Size.Y, bb.Origin.Y+bb.Size.Y)

	origin := Vector{
		X: minX,
		Y: minY,
	}
	size := Vector{
		X: maxX - minX,
		Y: maxY - minY,
	}
	return AABB{
		Origin: origin,
		Size:   size,
	}
}

// Matrix defines A 2D matrix of the form:
// [ A  B ]
// [ C  D ]
type Matrix struct {
	A, B, C, D float64
}

func NewRotationMatrix(angle float64) Matrix {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Matrix{
		A: cos,
		B: -sin,
		C: sin,
		D: cos,
	}
}

func (m Matrix) Mult(m2 Matrix) Matrix {
	return Matrix{
		A: m.A*m2.A + m.B*m2.C,
		B: m.A*m2.B + m.B*m2.D,
		C: m.C*m2.A + m.D*m2.C,
		D: m.C*m2.B + m.D*m2.D,
	}
}

func (m Matrix) VecMult(v Vector) Vector {
	return NewVector(m.A*v.X+m.B*v.Y, m.C*v.X+m.D*v.Y)
}

type Vector struct {
	X, Y float64
}

func (v Vector) Add(b Vector) Vector {
	return NewVector(v.X+b.X, v.Y+b.Y)
}

func (v Vector) Sub(b Vector) Vector {
	return NewVector(v.X-b.X, v.Y-b.Y)
}

func (v Vector) Scale(s float64) Vector {
	return NewVector(v.X*s, v.Y*s)
}

func (v Vector) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Line struct {
	from, to Vector
}

func (l Line) Len() float64 {
	return l.to.Sub(l.from).Abs()
}

func (l Line) BoundingBox() AABB {
	maxX := math.Max(l.from.X, l.to.X)
	maxY := math.Max(l.from.Y, l.to.Y)
	minX := math.Min(l.from.X, l.to.X)
	minY := math.Min(l.from.Y, l.to.Y)
	return AABB{
		Origin: Vector{minX, minY},
		Size:   Vector{maxX - minX, maxY - minY},
	}
}

func NewVector(x, y float64) Vector {
	return Vector{X: x, Y: y}
}

func NewAngledVector(angle, length float64) Vector {
	x := math.Cos(angle) * length
	y := math.Sin(angle) * length
	return Vector{X: x, Y: y}
}
