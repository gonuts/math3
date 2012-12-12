// Package vec32 operates on four-dimensional float32 vectors.
package vec32

import (
	"fmt"
	"math"
)

// Vector holds a four-dimensional vector. The default value is a zero vector.
type Vector [4]float32

// Vec3 returns a copy of the vector with the 4th component set to zero.
func (v Vector) Vec3() Vector {
	return Vector{v[0], v[1], v[2], v[3]}
}

// Normalize creates a new vector that is of unit length in the same direction as the vector.
func (v Vector) Normalize() Vector {
	vlen := v.Length()
	if vlen == 0 {
		return v
	}
	return v.Scale(1.0 / vlen)
}

// Length returns the magnitude of the vector.
func (v Vector) Length() float32 {
	return float32(math.Sqrt(float64(v.LengthSqr())))
}

// LengthSqr returns the magnitude squared of the vector.  This is cheaper to compute than Length.
func (v Vector) LengthSqr() float32 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2] + v[3]*v[3]
}

// Abs returns a new vector with all positive components.
func (v Vector) Abs() Vector {
	return Vector{
		float32(math.Abs(float64(v[0]))),
		float32(math.Abs(float64(v[1]))),
		float32(math.Abs(float64(v[2]))),
		float32(math.Abs(float64(v[3]))),
	}
}

// Negate returns a new vector in the opposite direction.
func (v Vector) Negate() Vector {
	return Vector{-v[0], -v[1], -v[2], -v[3]}
}

// Inverse returns a new vector that is the result of 1.0 / v[i] for all i.  Any zero value is left as zero.
func (v Vector) Inverse() Vector {
	var r Vector
	for axis, comp := range v {
		if comp != 0.0 {
			r[axis] = 1.0 / comp
		}
	}
	return r
}

// IsZero indicates whether the vector is the zero vector.
func (v Vector) IsZero() bool {
	return v[0] == 0 && v[1] == 0 && v[2] == 0 && v[3] == 0
}

func (v Vector) String() string {
	return fmt.Sprintf("[%.4f %.4f %.4f %.4f]", v[0], v[1], v[2], v[3])
}

func (v Vector) GoString() string {
	return fmt.Sprintf("vec32.Vector{%#v, %#v, %#v, %#v}", v[0], v[1], v[2], v[3])
}

// Add computes the sum of two vectors.
func Add(v1, v2 Vector) Vector

func add(v1, v2 Vector) Vector {
	return Vector{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

// Sum computes the sum of two or more vectors.
func Sum(v1, v2 Vector, vn ...Vector) Vector {
	result := Vector{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
	for _, u := range vn {
		result[0] += u[0]
		result[1] += u[1]
		result[2] += u[2]
		result[3] += u[3]
	}
	return result
}

// Sub computes the difference of two vectors.
func Sub(v1, v2 Vector) Vector

func sub(v1, v2 Vector) Vector {
	return Vector{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

// Scale multiplies all of a vector's components by a scalar.
func (v Vector) Scale(s float32) Vector {
	return Vector{v[0] * s, v[1] * s, v[2] * s, v[3] * s}
}

// Mul multiplies the components of two vectors together.
func Mul(v1, v2 Vector) Vector {
	return Vector{v1[0] * v2[0], v1[1] * v2[1], v1[2] * v2[2], v1[3] * v2[3]}
}

// Dot computes the dot product of two vectors.
func Dot(v1, v2 Vector) float32

func dot(v1, v2 Vector) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

// Cross computes the three-dimensional cross product of two vectors.
func Cross(v1, v2 Vector) Vector

func cross(v1, v2 Vector) Vector {
	return Vector{
		v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0],
	}
}

// CreateCS finds two normalized vectors orthogonal to the given one that can be used as a coordinate system.
//
// All calculations are done in three-dimensional space.  This is particularly
// useful for UV-mapping and the like.
func CreateCS(normal Vector) (u, v Vector) {
	if normal[0] == 0 && normal[1] == 0 {
		if normal[2] < 0 {
			u = Vector{-1.0, 0.0, 0.0}
		} else {
			u = Vector{1.0, 0.0, 0.0}
		}
		v = Vector{0.0, 1.0, 0.0}
	} else {
		d := 1.0 / float32(math.Sqrt(float64(normal[1]*normal[1]+normal[0]*normal[0])))
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
