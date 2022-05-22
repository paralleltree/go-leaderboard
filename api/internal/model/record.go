package model

type Record[T any] struct {
	Id   string
	Item T
}
