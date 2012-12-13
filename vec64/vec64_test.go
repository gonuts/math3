package vec64

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		A, B, Result Vector
	}{
		{Vector{0, 0, 0}, Vector{0, 0, 0}, Vector{0, 0, 0}},
		{Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{5, 7, 9}},
	}
	for _, tt := range tests {
		if r := Add(tt.A, tt.B); r != tt.Result {
			t.Errorf("Add(%v, %v) != %v (got %v)", tt.A, tt.B, tt.Result, r)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(Vector{1, 2, 3}, Vector{4, 5, 6})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		A, B, Result Vector
	}{
		{Vector{0, 0, 0}, Vector{0, 0, 0}, Vector{0, 0, 0}},
		{Vector{1, 2, 3}, Vector{6, 5, 4}, Vector{-5, -3, -1}},
	}
	for _, tt := range tests {
		if r := Sub(tt.A, tt.B); r != tt.Result {
			t.Errorf("Sub(%v, %v) != %v (got %v)", tt.A, tt.B, tt.Result, r)
		}
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sub(Vector{1, 2, 3}, Vector{4, 5, 6})
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		A, B   Vector
		Result float64
	}{
		{Vector{0, 0, 0}, Vector{0, 0, 0}, 0},
		{Vector{1, 2, 3}, Vector{4, 5, 6}, 32},
	}
	for _, tt := range tests {
		if r := Dot(tt.A, tt.B); r != tt.Result {
			t.Errorf("Dot(%v, %v) != %v (got %v)", tt.A, tt.B, tt.Result, r)
		}
	}
}

func BenchmarkDot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dot(Vector{1, 2, 3}, Vector{4, 5, 6})
	}
}

func TestCross(t *testing.T) {
	tests := []struct {
		A, B, Result Vector
	}{
		{Vector{1, 0, 0}, Vector{0, 1, 0}, Vector{0, 0, 1}},
		{Vector{2, 0, 0}, Vector{0, 3, 0}, Vector{0, 0, 6}},
		{Vector{0, 1, 0}, Vector{1, 0, 0}, Vector{0, 0, -1}},
	}
	for _, tt := range tests {
		if r := Cross(tt.A, tt.B); r != tt.Result {
			t.Errorf("Cross(%v, %v) != %#v (got %#v)", tt.A, tt.B, tt.Result, r)
		}
	}
}

func BenchmarkCross(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cross(Vector{1, 2, 3}, Vector{4, 5, 6})
	}
}

func TestNormalize(t *testing.T) {
	tests := []struct {
		V        Vector
		Expected Vector
	}{
		{Vector{1.0, 2.0, -2.0}, Vector{1.0 / 3.0, 2.0 / 3.0, -2.0 / 3.0}},
	}
	for _, tt := range tests {
		if result := tt.V.Normalize(); result != tt.Expected {
			t.Errorf("%v.Normalize() != %v (got %v)", tt.V, tt.Expected, result)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		Vec       Vector
		Length    float64
		LengthSqr float64
	}{
		{Vector{0.0, 0.0, 0.0}, 0.0, 0.0},

		{Vector{1.0, 0.0, 0.0}, 1.0, 1.0},
		{Vector{0.0, 1.0, 0.0}, 1.0, 1.0},
		{Vector{0.0, 0.0, 1.0}, 1.0, 1.0},

		{Vector{-1.0, 0.0, 0.0}, 1.0, 1.0},
		{Vector{0.0, -1.0, 0.0}, 1.0, 1.0},
		{Vector{0.0, 0.0, -1.0}, 1.0, 1.0},

		{Vector{3.0, -4.0, 0.0}, 5.0, 25.0},
		{Vector{1.0, 2.0, -2.0}, 3.0, 9.0},
		{Vector{3.14, 20.7, 0.5}, 20.942769635365803, 438.59959999999995},
	}

	for _, tt := range tests {
		if l2 := tt.Vec.LengthSqr(); l2 != tt.LengthSqr {
			t.Errorf("%v.LengthSqr() != %v (got %#v)", tt.Vec, tt.LengthSqr, l2)
		}
		if l := tt.Vec.Length(); l != tt.Length {
			t.Errorf("%v.Length() != %v (got %#v)", tt.Vec, tt.Length, l)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		Vec      Vector
		Expected Vector
	}{
		{Vector{0, 0, 0}, Vector{0, 0, 0}},
		{Vector{1, 2, 3}, Vector{1, 2, 3}},
		{Vector{-1, -2, -3}, Vector{1, 2, 3}},
		{Vector{-1, 2, -3}, Vector{1, 2, 3}},
	}
	for _, tt := range tests {
		if result := tt.Vec.Abs(); result != tt.Expected {
			t.Errorf("%v.Abs() != %v (got %v)", tt.Vec, tt.Expected, result)
		}
	}
}

func TestNegate(t *testing.T) {
	tests := []struct {
		Vec      Vector
		Expected Vector
	}{
		{Vector{0, 0, 0}, Vector{0, 0, 0}},
		{Vector{1, 2, 3}, Vector{-1, -2, -3}},
		{Vector{-1, -2, -3}, Vector{1, 2, 3}},
		{Vector{-1, 2, -3}, Vector{1, -2, 3}},
	}
	for _, tt := range tests {
		if result := tt.Vec.Negate(); result != tt.Expected {
			t.Errorf("%v.Negate() != %v (got %v)", tt.Vec, tt.Expected, result)
		}
	}
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		Vec      Vector
		Expected bool
	}{
		{Vector{0, 0, 0}, true},
		{Vector{1, 0, 0}, false},
		{Vector{-1, 0, 0}, false},
	}
	for _, tt := range tests {
		if tt.Vec.IsZero() != tt.Expected {
			t.Errorf("%v.IsZero() != %v", tt.Vec, tt.Expected)
		}
	}
}
