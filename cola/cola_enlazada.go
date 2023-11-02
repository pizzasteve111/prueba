package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{nil, nil}
}

func crearNodoCola[T any](elem T) *nodoCola[T] {
	return &nodoCola[T]{elem, nil}
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.ultimo == nil && cola.primero == nil
}

func (cola colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elem T) {
	nuevo_nodo := crearNodoCola(elem)

	if cola.ultimo == nil {
		cola.primero = nuevo_nodo
	} else {
		cola.ultimo.prox = nuevo_nodo
	}

	cola.ultimo = nuevo_nodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	elem := cola.primero.dato
	cola.primero = cola.primero.prox

	if cola.primero == nil {
		cola.ultimo = nil
	}

	return elem
}
