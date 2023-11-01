package diccionario

import (
	"fmt"
	TDALISTA "tdas/lista"
)

const (
	capacidad_inicial      = 7
	criterio_redimensionar = 3
	criterio_reducir       = 2
	redimension            = 2
)

type claveValor[K comparable, V any] struct {
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct {
	tamanio        int
	arreglo_listas []TDALISTA.Lista[*claveValor[K, V]]
	capacidad      int
}

type iteradorDiccionario[K comparable, V any] struct {
	hash   *hashAbierto[K, V]
	indice int
	lista  TDALISTA.IteradorLista[*claveValor[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &hashAbierto[K, V]{
		tamanio:        0,
		capacidad:      capacidad_inicial,
		arreglo_listas: make([]TDALISTA.Lista[*claveValor[K, V]], capacidad_inicial),
	}

	crear_tabla(&hash.arreglo_listas)

	return hash
}
func crear_tabla[K comparable, V any](arreglo *[]TDALISTA.Lista[*claveValor[K, V]]) {
	for i := 0; i < len(*arreglo); i++ {
		(*arreglo)[i] = TDALISTA.CrearListaEnlazada[*claveValor[K, V]]()
	}
}

func (h *hashAbierto[K, V]) buscar(clave K) (TDALISTA.IteradorLista[*claveValor[K, V]], bool) {
	pertenece := false

	indice := int(funcionDeHashing(convertirABytes[K](clave))) % h.capacidad
	lista := h.arreglo_listas[indice]
	iter_lista := lista.Iterador()

	for iter_lista.HaySiguiente() {
		actual := iter_lista.VerActual().clave
		if actual == clave {
			pertenece = true
			break
		}
		iter_lista.Siguiente()
	}
	return iter_lista, pertenece
}

func (h *hashAbierto[K, V]) redimensionar(nueva_cap int) {
	h.capacidad = nueva_cap
	nueva_tabla := make([]TDALISTA.Lista[*claveValor[K, V]], h.capacidad)

	crear_tabla(&nueva_tabla)

	for _, lista := range h.arreglo_listas {
		iterador := lista.Iterador()
		for iterador.HaySiguiente() {
			actual := iterador.VerActual()
			indice := int(funcionDeHashing(convertirABytes(actual.clave))) % h.capacidad
			nueva_tabla[indice].InsertarUltimo(actual)
			iterador.Siguiente()
		}
	}
	h.arreglo_listas = nueva_tabla
}

func (iter *iteradorDiccionario[K, V]) indicePosicionNoVacia(pos int) {
	var indice int
	encontrado := false
	for i := pos; i < iter.hash.capacidad; i++ {
		if !iter.hash.arreglo_listas[i].EstaVacia() {
			indice = i
			encontrado = true
			break
		}
	}

	if !encontrado {
		iter.lista = nil
	} else {

		iter.indice = indice
		iter.lista = iter.hash.arreglo_listas[indice].Iterador()
	}

}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	//LOGICA REDIMENSION
	factor_carga := h.tamanio / h.capacidad
	if factor_carga >= criterio_redimensionar {
		h.redimensionar(h.capacidad * redimension)
	}

	iter, pertenece := h.buscar(clave)

	if pertenece {
		actual := iter.VerActual()
		actual.valor = dato
	} else {
		iter.Insertar(&claveValor[K, V]{clave, dato})
		h.tamanio++
	}
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	_, pertenece := h.buscar(clave)
	return pertenece
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	iter, pertenece := h.buscar(clave)
	if !pertenece {
		panic("La clave no pertenece al diccionario")
	}

	return iter.VerActual().valor
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {

	iter, pertenece := h.buscar(clave)

	if !pertenece {
		panic("La clave no pertenece al diccionario")
	}

	borrado := iter.Borrar().valor

	h.tamanio--

	factor_carga := h.tamanio / h.capacidad
	if factor_carga <= criterio_reducir && h.capacidad > capacidad_inicial {
		h.redimensionar(h.capacidad / redimension)
	}

	return borrado
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.tamanio
}

func (h *hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, lista := range h.arreglo_listas {

		iter := lista.Iterador()
		for iter.HaySiguiente() {
			actual := iter.VerActual()
			if !f(actual.clave, actual.valor) {
				return
			}
			iter.Siguiente()

		}
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iteradorDiccionario[K, V]{
		hash:   h,
		indice: 0,
		lista:  nil,
	}
	iter.indicePosicionNoVacia(iter.indice)

	return iter
}

func (iter *iteradorDiccionario[K, V]) HaySiguiente() bool {
	return (iter.lista != nil && iter.lista.HaySiguiente()) || (iter.indice+1 < iter.hash.capacidad && !iter.hash.arreglo_listas[iter.indice+1].EstaVacia())
}

func (iter *iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.lista.VerActual()
	return actual.clave, actual.valor
}

func (iter *iteradorDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iter.lista.Siguiente()

	if !iter.lista.HaySiguiente() {
		iter.indicePosicionNoVacia(iter.indice + 1)

	}

}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// funcion de hashing de jenkins que obtuvimos de internet
func funcionDeHashing(data []byte) uint32 {
	var hash uint32
	for _, c := range data {
		hash += uint32(c)
		hash += (hash << 10)
		hash ^= (hash >> 6)
	}
	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)
	return hash

}
