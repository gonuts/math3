package mat32

import "testing"

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
				{1.0, 0.0, 0.0, 0.0},
				{0.0, 1.0, 0.0, 0.0},
				{0.0, 0.0, 1.0, 0.0},
				{4.0, 5.0, 6.0, 1.0},
			},
			Matrix{
				{1.0, 0.0, 0.0, 0.0},
				{0.0, 1.0, 0.0, 0.0},
				{0.0, 0.0, 1.0, 0.0},
				{1.0, 2.0, 3.0, 1.0},
			},
			Matrix{
				{1.0, 0.0, 0.0, 0.0},
				{0.0, 1.0, 0.0, 0.0},
				{0.0, 0.0, 1.0, 0.0},
				{5.0, 7.0, 9.0, 1.0},
			},
		},
	}
	for _, test := range tests {
		out := Mul(test.A, test.B)
		if out != test.Out {
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
