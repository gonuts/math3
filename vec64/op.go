package vec64

// Add computes the sum of two vectors.
func Add(v1, v2 Vector) Vector

func add(v1, v2 Vector) Vector {
	return Vector{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

// Sub computes the difference of two vectors.
func Sub(v1, v2 Vector) Vector

func sub(v1, v2 Vector) Vector {
	return Vector{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

// Dot computes the dot product of two vectors.
func Dot(v1, v2 Vector) float64

func dot(v1, v2 Vector) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}
