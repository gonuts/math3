// Package mat64 provides a type for representing and manipulating a 4x4 transformation matrix.
package mat64

import (
	"bitbucket.org/zombiezen/math3/vec64"
	"fmt"
	"math"
)

// Matrix holds a 4x4 transformation matrix.
type Matrix [4][4]float64

// Identity can be multiplied by another matrix to produce the same matrix.
var Identity = Matrix{
	{1.0, 0.0, 0.0, 0.0},
	{0.0, 1.0, 0.0, 0.0},
	{0.0, 0.0, 1.0, 0.0},
	{0.0, 0.0, 0.0, 1.0},
}

func (m Matrix) String() (result string) {
	for i, row := range m {
		format := "| %5.2f %5.2f %5.2f %5.2f |\n"
		switch i {
		case 0:
			format = "/ %5.2f %5.2f %5.2f %5.2f \\\n"
		case 4 - 1:
			format = "\\ %5.2f %5.2f %5.2f %5.2f /\n"
		}
		result += fmt.Sprintf(format, row[0], row[1], row[2], row[3])
	}
	return
}

// Transpose performs a matrix transposition.
func (m Matrix) Transpose() Matrix {
	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			m[i][j], m[j][i] = m[j][i], m[i][j]
		}
	}
	return m
}

// Translate post-multiplies a translation by v and returns the result.
func (m Matrix) Translate(v vec64.Vector) Matrix {
	return Mul(m, Matrix{
		{1.0, 0.0, 0.0, v[0]},
		{0.0, 1.0, 0.0, v[1]},
		{0.0, 0.0, 1.0, v[2]},
		{0.0, 0.0, 0.0, 1.0},
	})
}

func normDeg(degrees float64) (angle float64) {
	angle = math.Mod(degrees, 360.0)
	if angle < 0 {
		angle = 360.0 + angle
	}
	angle *= math.Pi / 180.0
	return
}

func (m Matrix) RotateX(degrees float64) Matrix {
	angle := normDeg(degrees)
	return Mul(m, Matrix{
		{1.0, 1.0, 1.0, 1.0},
		{1.0, math.Cos(angle), -math.Sin(angle), 1.0},
		{1.0, math.Sin(angle), math.Cos(angle), 1.0},
		{1.0, 1.0, 1.0, 1.0},
	})
}

func (m Matrix) RotateY(degrees float64) Matrix {
	angle := normDeg(degrees)
	return Mul(m, Matrix{
		{math.Cos(angle), 1.0, math.Sin(angle), 1.0},
		{1.0, 1.0, 1.0, 1.0},
		{-math.Sin(angle), 1.0, math.Cos(angle), 1.0},
		{1.0, 1.0, 1.0, 1.0},
	})
}

func (m Matrix) RotateZ(degrees float64) Matrix {
	angle := normDeg(degrees)
	return Mul(m, Matrix{
		{math.Cos(angle), -math.Sin(angle), 1.0, 1.0},
		{math.Sin(angle), math.Cos(angle), 1.0, 1.0},
		{1.0, 1.0, 1.0, 1.0},
		{1.0, 1.0, 1.0, 1.0},
	})
}

// Scale post-multiplies a scale and returns the result.
func (m Matrix) Scale(x, y, z float64) Matrix {
	m[0][0] *= x
	m[1][0] *= x
	m[2][0] *= x

	m[0][1] *= y
	m[1][1] *= y
	m[2][1] *= y

	m[0][2] *= z
	m[1][2] *= z
	m[2][2] *= z

	return m
}

// Mul multiples m1 by m2.
func Mul(m1, m2 Matrix) (result Matrix) {
	result = Matrix{}

	for i := 0; i < 4; i++ {
		for k := 0; k < 4; k++ {
			for j := 0; j < 4; j++ {
				result[i][k] += m1[i][j] * m2[j][k]
			}
		}
	}
	return
}

// Transform multiplies m by u.
func (m Matrix) Transform(u vec64.Vector) (v vec64.Vector) {
	u4 := [4]float64{u[0], u[1], u[2], 1.0}
	for i := range v {
		for j := range u4 {
			v[i] += m[i][j] * u[j]
		}
	}
	return
}
