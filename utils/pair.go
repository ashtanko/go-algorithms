package utils

type Pair[L interface{}, R interface{}] struct {
	Left  L
	Right R
}
