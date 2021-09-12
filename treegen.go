//go:build js && wasm
// +build js,wasm

package main

import (
	"math"
	"math/rand"
)

type TreeGenerator struct {
	angle, scale, angleVariation, scaleVariation float64
	seed                                         int64
}

func (t *TreeGenerator) SetAngle(angle float64) *TreeGenerator {
	t.angle = angle
	return t
}

func (t *TreeGenerator) Angle() float64 {
	return t.angle
}

func (t *TreeGenerator) SetScale(scale float64) *TreeGenerator {
	t.scale = scale
	return t
}

func (t *TreeGenerator) Scale() float64 {
	return t.scale
}

func (t *TreeGenerator) SetSeed(seed int64) *TreeGenerator {
	t.seed = seed
	return t
}

func (t *TreeGenerator) Seed() int64 {
	return t.seed
}

func (t *TreeGenerator) AngleVariation() float64 {
	return t.angleVariation
}

func (t *TreeGenerator) ScaleVariation() float64 {
	return t.scaleVariation
}

func NewDefaultTreeGenerator() *TreeGenerator {

	//empirically determined 'good' values
	return &TreeGenerator{
		angle:          1.0 / 8 * math.Pi,
		scale:          0.8,
		angleVariation: 1.0 / 30,
		scaleVariation: 1.0 / 20,
	}
}

func (t *TreeGenerator) Generate() []Line {
	rng := rand.New(rand.NewSource(t.seed))
	return t.generateTreeRec(Vector{}, Vector{
		X: 0,
		Y: -2,
	}, rng)
}

func (t *TreeGenerator) generateTreeRec(origin Vector, stem Vector, rng *rand.Rand) []Line {

	if stem.Abs() < 0.5 {
		return []Line{}
	}

	var lines []Line

	newOrigin := origin.Add(stem)
	lines = append(lines, Line{from: origin, to: newOrigin})

	origin = newOrigin

	left := t.branch(t.angle, t.scale, stem, rng)
	right := t.branch(-t.angle, t.scale, stem, rng)

	lines = append(lines, Line{
		from: origin,
		to:   origin.Add(left),
	})
	lines = append(lines, Line{
		from: origin,
		to:   origin.Add(right),
	})

	lines = append(lines, t.generateTreeRec(origin.Add(left), left, rng)...)
	lines = append(lines, t.generateTreeRec(origin.Add(right), right, rng)...)

	return lines
}

func (t *TreeGenerator) branch(angle, scale float64, stem Vector, rng *rand.Rand) Vector {

	angleRandom := rng.NormFloat64() * t.angleVariation * math.Pi
	scaleRandom := 1 + rng.NormFloat64()*t.scaleVariation

	return NewRotationMatrix(angle).Mult(NewRotationMatrix(angleRandom)).VecMult(stem).Scale(scale * scaleRandom)
}
