package cola_prioridad

type heap[T any] struct {
	elementos   []T
	cantidad    int
	comparacion func(T, T) int
}

func CrearHeap[T any](f func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		elementos:   make([]T, 0),
		cantidad:    0,
		comparacion: f,
	}
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}
