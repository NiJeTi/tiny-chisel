package types

type Vector struct {
	X, Y float32
}

func VectorZero() Vector {
	return Vector{0, 0}
}
