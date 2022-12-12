package generics

type List[T any] struct {
	slice []T
}

func (l *List[T]) push(item T) {
	l.slice = append(l.slice, item)
}

func (l *List[T]) pop() T {
	item := l.slice[0]
	l.slice = l.slice[1:]
	return item
}
