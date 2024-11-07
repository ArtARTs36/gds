package gds

type Collection[V comparable] interface {
	IsEmpty() bool
	IsNotEmpty() bool
	List() []V
}
