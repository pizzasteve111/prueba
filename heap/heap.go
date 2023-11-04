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

func downheap[T any](elementos []T, indice_act, cantidad int, comparacion func(T, T) int) {
	if indice_act >= cantidad {
		return
	}
	pos_hijo_izq := (indice_act * 2) + 1
	pos_hijo_der := (indice_act * 2) + 2

	mas_grande := indice_act

	if pos_hijo_izq < cantidad && comparacion(elementos[pos_hijo_izq], elementos[mas_grande]) > 0 {
		mas_grande = pos_hijo_izq
	}
	if pos_hijo_der < cantidad && comparacion(elementos[pos_hijo_der], elementos[mas_grande]) > 0 {
		mas_grande = pos_hijo_der
	}

	if mas_grande != indice_act {
		swap(elementos, indice_act, mas_grande)
		downheap(elementos, mas_grande, cantidad, comparacion)
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

func heapify[T any](arr []T, cantidad int, f func(T, T) int) {
	for i := cantidad - 1; i >= 0; i-- {
		downheap(arr, i, cantidad, f)
	}
}

func CrearHeapArr[T any](arr []T, f func(T, T) int) ColaPrioridad[T] {
	heapify(arr, len(arr), f)
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

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify[T](elementos, len(elementos), funcion_cmp)
	for i := len(elementos) - 1; i >= 0; i-- {
		swap[T](elementos, 0, i)
		downheap[T](elementos, 0, i, funcion_cmp)
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
	copy(nuevo_arr, h.elementos)
	h.elementos = nuevo_arr
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	borrado := h.elementos[0]
	swap(h.elementos, 0, (h.cantidad - 1))
	h.cantidad--
	downheap(h.elementos, 0, h.cantidad, h.comparacion)

	if h.cantidad > 0 && (h.cantidad*CRITERIO_REDUCCION) <= cap(h.elementos) {
		h.redimensionar(cap(h.elementos) / REDIMENSION)
	}

	return borrado
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == cap(h.elementos) {
		h.redimensionar(h.cantidad * REDIMENSION)
	}
	h.elementos[h.cantidad] = elem
	upheap(h.elementos, h.cantidad, h.comparacion)
	h.cantidad++
}
