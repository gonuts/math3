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

func TestNormalize(t *testing.T) {
	comps := []float64{1.0, 2.0, -2.0}
	length := 3.0
	v := Vector{comps[0], comps[1], comps[2]}
	vn := v.Normalize()
	if vn.Length() != 1.0 {
		t.Error("Length is not 1")
	}
	for axis := 0; axis <= 2; axis++ {
		if vn[axis] != comps[axis]/length {
			t.Error("0 component is incorrect")
		}
	}
}

func TestLength(t *testing.T) {
	type lengthTest struct {
		Vec       Vector
		Length    float64
		LengthSqr float64
	}

	tests := []lengthTest{
		lengthTest{Vector{0.0, 0.0, 0.0}, 0.0, 0.0},

		lengthTest{Vector{1.0, 0.0, 0.0}, 1.0, 1.0},
		lengthTest{Vector{0.0, 1.0, 0.0}, 1.0, 1.0},
		lengthTest{Vector{0.0, 0.0, 1.0}, 1.0, 1.0},

		lengthTest{Vector{-1.0, 0.0, 0.0}, 1.0, 1.0},
		lengthTest{Vector{0.0, -1.0, 0.0}, 1.0, 1.0},
		lengthTest{Vector{0.0, 0.0, -1.0}, 1.0, 1.0},

		lengthTest{Vector{3.0, -4.0, 0.0}, 5.0, 25.0},
		lengthTest{Vector{1.0, 2.0, -2.0}, 3.0, 9.0},
		lengthTest{Vector{3.14, 20.7, 0.5}, 20.942769635365803, 438.59959999999995},
	}

	for _, ltest := range tests {
		if ltest.Vec.LengthSqr() != ltest.LengthSqr {
			t.Errorf("LengthSqr failed for %v (wanted %.2f, got %.2f)", ltest.Vec, ltest.LengthSqr, ltest.Vec.LengthSqr())
		}
		if ltest.Vec.Length() != ltest.Length {
			t.Errorf("Length failed for %v (wanted %.2f, got %.2f)", ltest.Vec, ltest.Length, ltest.Vec.Length())
		}
	}
}

func TestAbs(t *testing.T) {
	var v Vector

	v = Vector{0, 0, 0}.Abs()
	if v[0] != 0 || v[1] != 0 || v[2] != 0 {
		t.Error("Zero vector incorrect")
	}

	v = Vector{1, 2, 3}.Abs()
	if v[0] != 1 || v[1] != 2 || v[2] != 3 {
		t.Error("All positive vector incorrect")
	}

	v = Vector{-1, -2, -3}.Abs()
	if v[0] != 1 || v[1] != 2 || v[2] != 3 {
		t.Error("All negative vector incorrect")
	}

	v = Vector{-1, 2, -3}.Abs()
	if v[0] != 1 || v[1] != 2 || v[2] != 3 {
		t.Error("Mixed vector incorrect")
	}
}

func TestNegate(t *testing.T) {
	var v Vector

	v = Vector{0, 0, 0}.Negate()
	if v[0] != 0 || v[1] != 0 || v[2] != 0 {
		t.Error("Zero vector incorrect")
	}

	v = Vector{1, 2, 3}.Negate()
	if v[0] != -1 || v[1] != -2 || v[2] != -3 {
		t.Error("All positive vector incorrect")
	}

	v = Vector{-1, -2, -3}.Negate()
	if v[0] != 1 || v[1] != 2 || v[2] != 3 {
		t.Error("All negative vector incorrect")
	}

	v = Vector{-1, 2, -3}.Negate()
	if v[0] != 1 || v[1] != -2 || v[2] != 3 {
		t.Error("Mixed vector incorrect")
	}
}

func TestIsZero(t *testing.T) {
	var v Vector

	v = Vector{0, 0, 0}
	if !v.IsZero() {
		t.Error("Zero vector is not zero")
	}

	v = Vector{1, 0, 0}
	if v.IsZero() {
		t.Error("Positive 0 vector is zero")
	}

	v = Vector{-1, 0, 0}
	if v.IsZero() {
		t.Error("Negative 0 vector is zero")
	}
}
