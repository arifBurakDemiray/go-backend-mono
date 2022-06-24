package utils

type Response[T any] struct {
	status int
	data T
}