package vec64

import (
	"fmt"
	"math"
)

// Vector holds a three-dimensional vector. The default value is a zero vector.
type Vector [3]float64

// Normalize creates a new vector that is of unit length in the same direction as the vector.
func (v Vector) Normalize() Vector {
	vlen := v.Length()
	if vlen == 0 {
		return v
	}
	return v.Scale(1.0 / vlen)
}

// Length returns the magnitude of the vector.
func (v Vector) Length() float64 {
	return math.Sqrt(v.LengthSqr())
}

// LengthSqr returns the magnitude squared of the vector.  This is cheaper to compute than Length.
func (v Vector) LengthSqr() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}

// Abs returns a new vector with all positive components.
func (v Vector) Abs() Vector {
	return Vector{math.Abs(v[0]), math.Abs(v[1]), math.Abs(v[2])}
}

// Negate returns a new vector in the opposite direction.
func (v Vector) Negate() Vector {
	return Vector{-v[0], -v[1], -v[2]}
}

// Inverse returns a new vector that is the result of 1.0 / v[i] for all i.  Any zero value is left as zero.
func (v Vector) Inverse() (r Vector) {
	for axis, comp := range v {
		if comp != 0.0 {
			r[axis] = 1.0 / comp
		}
	}
	return
}

// IsZero indicates whether the vector is the zero vector.
func (v Vector) IsZero() bool {
	return v[0] == 0 && v[1] == 0 && v[2] == 0
}

func (v Vector) String() string {
	return fmt.Sprintf("<%.4f, %.4f, %.4f>", v[0], v[1], v[2])
}

func (v Vector) GoString() string {
	return fmt.Sprintf("vector.Vector{%#v, %#v, %#v}", v[0], v[1], v[2])
}

// Sum computes the sum of two or more vectors.
func Sum(v1, v2 Vector, vn ...Vector) Vector {
	result := Vector{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
	for _, u := range vn {
		result[0] += u[0]
		result[1] += u[1]
		result[2] += u[2]
	}
	return result
}

// AddScalar adds a scalar to all of a vector's components.
func (v Vector) AddScalar(s float64) Vector {
	return Vector{v[0] + s, v[1] + s, v[2] + s}
}

// Scale multiplies all of a vector's components by a scalar.
func (v Vector) Scale(s float64) Vector {
	return Vector{v[0] * s, v[1] * s, v[2] * s}
}

// Mul multiplies the components of two vectors together.
func Mul(v1, v2 Vector) Vector {
	return Vector{v1[0] * v2[0], v1[1] * v2[1], v1[2] * v2[2]}
}

// Cross computes the cross product of two vectors.
func Cross(v1, v2 Vector) Vector {
	return Vector{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0],
	}
}

// CreateCS finds two normalized vectors orthogonal to the given one that can be used as a coordinate system.
//
// This is particularly useful for UV-mapping and the like.
func CreateCS(normal Vector) (u, v Vector) {
	if normal[0] == 0 && normal[1] == 0 {
		if normal[2] < 0 {
			u = Vector{-1.0, 0.0, 0.0}
		} else {
			u = Vector{1.0, 0.0, 0.0}
		}
		v = Vector{0.0, 1.0, 0.0}
	} else {
		d := 1.0 / math.Sqrt(normal[1]*normal[1]+normal[0]*normal[0])
		u = Vector{normal[1] * d, -normal[0] * d, 0.0}
		v = Cross(normal, u)
	}
	return
}

// Reflect calculates a reflection of a vector based on a normal.
func Reflect(v, n Vector) Vector {
	vn := Dot(v, n)
	if vn < 0 {
		return v.Negate()
	}
	return Sub(n.Scale(2*vn), v)
}
