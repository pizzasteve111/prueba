package pila

const criterioRedimension = 2
const criterioReduccion = 4
const capacidadInicial = 10

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := &pilaDinamica[T]{
		datos:    make([]T, capacidadInicial),
		cantidad: 0,
	}

	return pila
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.cantidad == 0 {
		panic("La pila esta vacia")
	}

	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) redimensionar(capacidad_nueva int) {
	nuevo := make([]T, capacidad_nueva)
	copy(nuevo, pila.datos)
	pila.datos = nuevo
}

func (pila *pilaDinamica[T]) Apilar(elem T) {

	if pila.cantidad == cap(pila.datos) {
		nueva_capacidad := criterioRedimension * cap(pila.datos)

		pila.redimensionar(nueva_capacidad)
	}

	pila.datos[pila.cantidad] = elem
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	valor := pila.datos[pila.cantidad-1]
	pila.cantidad--

	if pila.cantidad > 0 && (pila.cantidad*criterioReduccion) <= cap(pila.datos) {
		nueva_capacidad := cap(pila.datos) / criterioRedimension

		pila.redimensionar(nueva_capacidad)
	}

	return valor
}
