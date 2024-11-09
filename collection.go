package gds

type Collection[V comparable] interface {
	First() V
	Len() int
	IsEmpty() bool
	IsNotEmpty() bool
	List() []V
}
