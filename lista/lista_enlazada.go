package lista

// defino mi struct Nodo
type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

type iteradorLista[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
	lista    *lista_enlazada[T]
}

type lista_enlazada[T any] struct {
	inicio  *nodo[T]
	fin     *nodo[T]
	tamanio int
}

func crear_nodo[T any](elemento T) *nodo[T] {
	return &nodo[T]{
		dato: elemento,
		prox: nil,
	}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &lista_enlazada[T]{
		inicio:  nil,
		fin:     nil,
		tamanio: 0,
	}
}
func (l *lista_enlazada[T]) EstaVacia() bool {
	if l.inicio == nil && l.fin == nil && l.tamanio == 0 {
		return true
	}
	return false
}

func (l *lista_enlazada[T]) InsertarPrimero(elemento T) {
	nodo := crear_nodo[T](elemento)
	if l.EstaVacia() {
		l.fin = nodo
	} else {
		nodo.prox = l.inicio
	}

	l.inicio = nodo
	l.tamanio++

}

func (l *lista_enlazada[T]) InsertarUltimo(elemento T) {
	nodo := crear_nodo[T](elemento)
	if l.EstaVacia() {
		l.inicio = nodo
	} else {
		l.fin.prox = nodo
	}

	l.fin = nodo
	l.tamanio++
}

func (l *lista_enlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	res := l.inicio.dato
	l.inicio = l.inicio.prox
	l.tamanio--

	if l.inicio == nil {
		l.fin = nil
	}

	return res
}

func (l *lista_enlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.inicio.dato
}

func (l *lista_enlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.fin.dato
}

func (l *lista_enlazada[T]) Largo() int {
	return l.tamanio
}

func (l *lista_enlazada[T]) Iterar(visitar func(T) bool) {
	act := l.inicio
	for act != nil {
		if !visitar(act.dato) {
			break
		}
		act = act.prox
	}
}

func (l *lista_enlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{
		actual:   l.inicio,
		anterior: nil,
		lista:    l,
	}
}

func (iter *iteradorLista[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorLista[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.prox
}

func (iter *iteradorLista[T]) Insertar(item T) {
	nodo := crear_nodo[T](item)

	if iter.lista.EstaVacia() {
		iter.lista.inicio = nodo
		iter.lista.fin = nodo
		iter.actual = iter.lista.inicio
	} else if iter.anterior != nil {
		iter.anterior.prox = nodo
		nodo.prox = iter.actual
		iter.actual = nodo

		if iter.actual.prox == nil {
			iter.lista.fin = nodo
		}

	} else {
		nodo.prox = iter.lista.inicio
		//iter.lista.inicio = nodo
		iter.actual = nodo
		iter.lista.inicio = iter.actual
	}
	iter.actual = nodo
	iter.lista.tamanio++

}

func (iter *iteradorLista[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	res := iter.actual.dato
	iter.lista.tamanio--

	if iter.actual == iter.lista.inicio {
		iter.lista.inicio = iter.lista.inicio.prox
		iter.actual = iter.lista.inicio
		iter.anterior = iter.lista.inicio
		if iter.lista.inicio == nil {
			iter.lista.fin = nil
		}

	} else if iter.actual == iter.lista.fin {
		iter.anterior.prox = nil
		iter.lista.fin = iter.anterior
		iter.actual = nil
	} else {
		iter.actual = iter.actual.prox
		iter.anterior.prox = iter.actual
	}

	return res
}
