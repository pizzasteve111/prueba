package diccionario

import (
	"fmt"
	TDALISTA "tdas/lista"
)

const (
	capacidad_inicial      = 7
	criterio_redimensionar = 0.7
	criterio_reducir       = 0.2
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
	actual *claveValor[K, V]
	indice int
	lista  TDALISTA.IteradorLista[*claveValor[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &hashAbierto[K, V]{
		tamanio:        0,
		capacidad:      capacidad_inicial,
		arreglo_listas: make([]TDALISTA.Lista[*claveValor[K, V]], capacidad_inicial),
	}

	for i := 0; i < capacidad_inicial; i++ {
		hash.arreglo_listas[i] = TDALISTA.CrearListaEnlazada[*claveValor[K, V]]()
	}

	return hash
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

	for i := 0; i < h.capacidad; i++ {
		nueva_tabla[i] = TDALISTA.CrearListaEnlazada[*claveValor[K, V]]()
	}
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

func (h hashAbierto[K, V]) indicePosicionNoVacia(pos int) int {
	if pos > h.capacidad {
		return pos
	}
	indice := pos
	for i := range h.arreglo_listas {
		if i <= pos {
			continue
		}
		if !h.arreglo_listas[i].EstaVacia() {
			indice = i
		}
	}
	return indice
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	//LOGICA REDIMENSION
	factor_carga := float32(h.tamanio / h.capacidad)
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
	if !h.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	iter, _ := h.buscar(clave)

	return iter.VerActual().valor
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	if !h.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	iter, _ := h.buscar(clave)

	borrado := iter.Borrar().valor

	h.tamanio--

	factor_carga := float32(h.tamanio / h.capacidad)
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
		if lista.EstaVacia() {
			continue
		}
		iter := lista.Iterador()
		for iter.HaySiguiente() {
			actual := iter.VerActual()
			if !f(actual.clave, actual.valor) {
				return
			} else {
				iter.Siguiente()
			}
		}
	}
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iteradorDiccionario[K, V]{
		hash:   h,
		actual: nil,
		indice: 0,
		lista:  nil,
	}
	/*
		indice := h.indicePosicionNoVacia(-1)
		if indice == -1 {
			return iter
		}
		iter.actual = h.arreglo_listas[indice].VerPrimero()
		iter.lista = h.arreglo_listas


	*/
	return iter
}

func (iter *iteradorDiccionario[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.clave, iter.actual.valor
}

func (iter *iteradorDiccionario[K, V]) Siguiente() {

	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	if iter.lista != nil {
		iter.lista.Siguiente()
		if iter.lista.HaySiguiente() && iter.indice < iter.hash.tamanio {
			iter.indice++
			iter.lista = iter.hash.arreglo_listas[iter.indice].Iterador()
		}

	} else if iter.indice < iter.hash.tamanio {
		iter.indice++
		iter.lista = iter.hash.arreglo_listas[iter.indice].Iterador()
		if !iter.lista.HaySiguiente() {
			iter.Siguiente()
		}
	} else {
		iter.lista = nil
	}

}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func funcionDeHashing(data []byte) uint32 {
	/*var hash uint32
	for _, b := range data {
		hash = uint32(b) + ((hash << 5) - hash)
	}
	return hash
	*/

	var hash uint32 = 5381

	for _, char := range data {
		hash = (hash << 5) + hash + uint32(char)
	}

	return hash
}
