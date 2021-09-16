package domain

type Repository interface {
	Put(binary []byte) error
}