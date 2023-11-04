package cola_prioridad

const (
	CAPACIDAD_INI      = 7
	CRITERIO_REDUCCION = 4
	REDIMENSION        = 2
)

type heap[T any] struct {
	elementos   []T
	cantidad    int
	comparacion func(T, T) int
}

func CrearHeap[T any](f func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		elementos:   make([]T, CAPACIDAD_INI),
		cantidad:    0,
		comparacion: f,
	}
}

func swap[T any](elementos []T, i, j int) {
	elementos[j], elementos[i] = elementos[i], elementos[j]
}

func downheap[T any](elementos []T, indice_act int, comparacion func(T, T) int) {
	if indice_act >= len(elementos)-1 {
		return
	}
	pos_hijo_izq := (indice_act * 2) + 1
	pos_hijo_der := (indice_act * 2) + 2

	mas_chico := indice_act

	if pos_hijo_izq < len(elementos) && comparacion(elementos[pos_hijo_izq], elementos[mas_chico]) < 0 {
		mas_chico = pos_hijo_izq
	}
	if pos_hijo_der < len(elementos) && comparacion(elementos[mas_chico], elementos[pos_hijo_der]) > 0 {
		mas_chico = pos_hijo_der
	}

	if mas_chico != indice_act {
		swap(elementos, indice_act, mas_chico)
		downheap(elementos, mas_chico, comparacion)
	}
}

func upheap[T any](elementos []T, indice_act int, comparacion func(T, T) int) {
	if indice_act == 0 {
		return
	}
	padre := (indice_act - 1) / 2

	if padre >= 0 && comparacion(elementos[indice_act], elementos[padre]) > 0 {
		swap(elementos, indice_act, padre)
		upheap(elementos, padre, comparacion)
	}
}

func heapify[T any](arr []T, f func(T, T) int) {

	for i := (len(arr) - 1); i >= 0; i-- {
		downheap(arr, i, f)
	}
}

func CrearHeapArr[T any](arr []T, f func(T, T) int) ColaPrioridad[T] {
	heapify[T](arr, f)
	return &heap[T]{
		elementos:   arr,
		cantidad:    len(arr),
		comparacion: f,
	}
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func Heapsort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify[T](elementos, funcion_cmp)
	for i := len(elementos); i >= 0; i-- {
		swap[T](elementos, 0, i)
		downheap[T](elementos, 0, funcion_cmp)
	}
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.elementos[0]
}

func (h *heap[T]) redimensionar(capacidad_nueva int) {
	nuevo_arr := make([]T, capacidad_nueva)
	copy(h.elementos, nuevo_arr)
	h.elementos = nuevo_arr
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	borrado := h.elementos[0]
	swap[T](h.elementos, 0, (h.cantidad - 1))
	downheap[T](h.elementos, 0, h.comparacion)

	h.cantidad--

	if h.cantidad > 0 && (h.cantidad*CRITERIO_REDUCCION)/2 >= cap(h.elementos) {
		h.redimensionar(cap(h.elementos) / REDIMENSION)
	}

	return borrado
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == cap(h.elementos) {
		h.redimensionar(h.cantidad * REDIMENSION)
	}
	h.elementos[h.cantidad] = elem
	upheap[T](h.elementos, h.cantidad, h.comparacion)
	h.cantidad++
}
