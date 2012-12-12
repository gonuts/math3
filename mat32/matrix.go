// Package mat32 provides a type for representing and manipulating a 4x4 matrix of float32.
package mat32

import (
	"bitbucket.org/zombiezen/math3/vec32"
	"fmt"
)

// Matrix holds a 4x4 matrix.  Each vector is a column of the matrix.
type Matrix [4]vec32.Vector

// Identity can be multiplied by another matrix to produce the same matrix.
var Identity = Matrix{
	{1.0, 0.0, 0.0, 0.0},
	{0.0, 1.0, 0.0, 0.0},
	{0.0, 0.0, 1.0, 0.0},
	{0.0, 0.0, 0.0, 1.0},
}

func (m Matrix) String() string {
	var result string
	for i, row := range m {
		format := "| %5.2f %5.2f %5.2f %5.2f |\n"
		switch i {
		case 0:
			format = "/ %5.2f %5.2f %5.2f %5.2f \\\n"
		case len(m) - 1:
			format = "\\ %5.2f %5.2f %5.2f %5.2f /\n"
		}
		result += fmt.Sprintf(format, row[0], row[1], row[2], row[3])
	}
	return result
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
func (m Matrix) Translate(v vec32.Vector) Matrix {
	return Mul(m, Matrix{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{v[0], v[1], v[2], 1.0},
	})
}

// Scale post-multiplies a scale and returns the result.
func (m Matrix) Scale(x, y, z float32) Matrix {
	return Mul(m, Matrix{
		{x, 0.0, 0.0, 0.0},
		{0.0, y, 0.0, 0.0},
		{0.0, 0.0, z, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	})
}

// Mul multiplies m1 by m2.
func Mul(m1, m2 Matrix) Matrix {
	var result Matrix
	m1 = m1.Transpose()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = vec32.Dot(m1[j], m2[i])
		}
	}
	return result
}

// Transform multiplies m by u.
func (m Matrix) Transform(u vec32.Vector) (v vec32.Vector) {
	for i := range v {
		for j := range u {
			v[i] += m[i][j] * u[j]
		}
	}
	return
}
