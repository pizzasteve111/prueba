package diccionario

import (
	"fmt"
	TDALISTA "tdas/lista"
)

const capacidad_inicial = 7

type claveValor[K comparable, V any] struct {
	clave K
	valor V
}

type hashAbierto[K comparable, V any] struct {
	tamanio        int
	arreglo_listas []TDALISTA.Lista[claveValor[K, V]]
	capacidad      int
}

type iteradorDiccionario[K comparable, V any] struct {
	hash   *hashAbierto[K, V]
	actual *claveValor[K, V]
	indice int
	lista  TDALISTA.IteradorLista[claveValor[K, V]]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &hashAbierto[K, V]{
		tamanio:        0,
		capacidad:      capacidad_inicial,
		arreglo_listas: make([]TDALISTA.Lista[claveValor[K, V]], capacidad_inicial),
	}

	for i := 0; i < capacidad_inicial; i++ {
		hash.arreglo_listas[i] = TDALISTA.CrearListaEnlazada[claveValor[K, V]]()
	}

	return hash
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	h.tamanio++

}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {

	return false
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	var valor V
	return valor
}

func (h *hashAbierto[K, V]) Borrar(clave K) V

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.tamanio
}

func (h *hashAbierto[K, V]) Iterar(func(clave K, dato V) bool)

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	//return &iteradorDiccionario{
	//	hash : h,
	//	actual :nil,
	//}
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
	var hash uint32
	for _, b := range data {
		hash = uint32(b) + ((hash << 5) - hash)
	}
	return hash
}
