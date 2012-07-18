// +build !amd64

package vec64

// Add computes the sum of two vectors.
func Add(v1, v2 Vector) Vector {
	return Vector{v1[X] + v2[X], v1[Y] + v2[Y], v1[Z] + v2[Z]}
}

// Sub computes the difference of two vectors.
func Sub(v1, v2 Vector) Vector {
	return Vector{v1[X] - v2[X], v1[Y] - v2[Y], v1[Z] - v2[Z]}
}

// Dot computes the dot product of two vectors.
func Dot(v1, v2 Vector) float64 {
	return v1[X]*v2[X] + v1[Y]*v2[Y] + v1[Z]*v2[Z]
}
