package store

// Store ...
type Store interface {
	user() UserRepository
}