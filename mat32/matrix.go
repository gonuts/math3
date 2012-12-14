// Package mat32 provides a type for representing and manipulating a 4x4 matrix of float32.
package mat32

import (
	"bitbucket.org/zombiezen/math3/vec32"
	"fmt"
	"math"
)

// Matrix holds a 4x4 matrix.  Each vector is a column of the matrix.
type Matrix [4]vec32.Vector

// Identity can be multiplied by another matrix to produce the same matrix.
var Identity = Matrix{
	{1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, 0, 1},
}

func (m Matrix) String() string {
	var result string
	for i, row := range m.Transpose() {
		format := "| %5.2f %5.2f %5.2f %5.2f |\n"
		switch i {
		case 0:
			format = "/ %5.2f %5.2f %5.2f %5.2f \\\n"
		case len(m) - 1:
			format = "\\ %5.2f %5.2f %5.2f %5.2f /"
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
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{v[0], v[1], v[2], 1},
	})
}

// Rotate post-multiplies a rotation around an axis. The angle is in radians and
// the axis will be normalized.
func (m Matrix) Rotate(angle float32, axis vec32.Vector) Matrix {
	axis = axis.Normalize()
	x, y, z := axis[0], axis[1], axis[2]
	sin, cos := float32(math.Sin(float64(angle))), float32(math.Cos(float64(angle)))
	return Mul(m, Matrix{
		{cos + x*x*(1-cos), y*x*(1-cos) + z*sin, z*x*(1-cos) - y*sin, 0},
		{x*y*(1-cos) - z*sin, cos + y*y*(1-cos), z*y*(1-cos) + x*sin, 0},
		{x*z*(1-cos) + y*sin, y*z*(1-cos) - x*sin, cos + z*z*(1-cos), 0},
		{0, 0, 0, 1},
	})
}

// Scale post-multiplies a scale and returns the result.
func (m Matrix) Scale(scale vec32.Vector) Matrix {
	return Mul(m, Matrix{
		{scale[0], 0, 0, 0},
		{0, scale[1], 0, 0},
		{0, 0, scale[2], 0},
		{0, 0, 0, 1},
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
