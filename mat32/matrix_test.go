package mat32

import (
	"bitbucket.org/zombiezen/math3/vec32"
	"math"
	"testing"
)

func TestTranspose(t *testing.T) {
	tests := []struct {
		In  Matrix
		Out Matrix
	}{
		{Identity, Identity},
		{
			Matrix{
				{1.0, 2.0, 3.0, 4.0},
				{5.0, 6.0, 7.0, 8.0},
				{9.0, 10.0, 11.0, 12.0},
				{13.0, 14.0, 15.0, 16.0},
			},
			Matrix{
				{1.0, 5.0, 9.0, 13.0},
				{2.0, 6.0, 10.0, 14.0},
				{3.0, 7.0, 11.0, 15.0},
				{4.0, 8.0, 12.0, 16.0},
			},
		},
	}
	for _, test := range tests {
		out := test.In.Transpose()
		if out != test.Out {
			t.Errorf("%#v.Transpose() = %#v; want %#v", test.In, out, test.Out)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		A, B Matrix
		Out  Matrix
	}{
		{Identity, Identity, Identity},
		{
			Matrix{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{4, 5, 6, 1},
			},
			Matrix{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{1, 2, 3, 1},
			},
			Matrix{
				{1, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{5, 7, 9, 1},
			},
		},
		{
			Matrix{
				{1, 0, 0, 0},
				{0, 0, 1, 0},
				{0, -1, 0, 0},
				{0, 0, 0, 1},
			},
			Matrix{
				{0, 0, -1, 0},
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 1},
			},
			Matrix{
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 1},
			},
		},
	}
	for _, test := range tests {
		out := Mul(test.A, test.B)
		if !checkMatrix(out, test.Out, 0.01) {
			t.Errorf("Mul(%#v, %#v) = %#v; want %#v", test.A, test.B, out, test.Out)
		}
	}
}

func BenchmarkMul(b *testing.B) {
	m1 := Matrix{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{4.0, 5.0, 6.0, 1.0},
	}
	m2 := Matrix{
		{1.0, 0.0, 0.0, 0.0},
		{0.0, 1.0, 0.0, 0.0},
		{0.0, 0.0, 1.0, 0.0},
		{1.0, 2.0, 3.0, 1.0},
	}
	for i := 0; i < b.N; i++ {
		Mul(m1, m2)
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		In    Matrix
		Angle float32
		Axis  vec32.Vector
		Out   Matrix
	}{
		{
			Identity,
			math.Pi / 2, vec32.Vector{1, 0, 0},
			Matrix{
				{1, 0, 0, 0},
				{0, 0, 1, 0},
				{0, -1, 0, 0},
				{0, 0, 0, 1},
			},
		},
		{
			Identity,
			math.Pi / 2, vec32.Vector{0, 1, 0},
			Matrix{
				{0, 0, -1, 0},
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 1},
			},
		},
		{
			Identity,
			math.Pi / 2, vec32.Vector{0, 0, 1},
			Matrix{
				{0, 1, 0, 0},
				{-1, 0, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 1},
			},
		},
	}
	for _, test := range tests {
		out := test.In.Rotate(test.Angle, test.Axis)
		if !checkMatrix(out, test.Out, 0.01) {
			t.Errorf("%#v.Rotate(%v, %v) = %#v; want %#v", test.In, test.Angle, test.Axis, out, test.Out)
		}
	}
}

// checkMatrix returns whether m1 ~ m2, given a tolerance.
func checkMatrix(m1, m2 Matrix, tol float32) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if m2[i][j] > m1[i][j]+tol || m2[i][j] < m1[i][j]-tol {
				return false
			}
		}
	}
	return true
}
