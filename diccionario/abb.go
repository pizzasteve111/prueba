package diccionario

import TDAPila "tdas/pila"

type nodo[K comparable, V any] struct {
	hijo_izq *nodo[K, V]
	hijo_der *nodo[K, V]
	valor    V
	clave    K
}

type abb[K comparable, V any] struct {
	raiz        *nodo[K, V]
	comparacion func(K, K) int
	cantidad    int
}
type iteradorABB[K comparable, V any] struct {
	abb   *abb[K, V]
	desde *K
	hasta *K
	pila  TDAPila.Pila[*nodo[K, V]]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		raiz:        nil,
		comparacion: funcion_cmp,
		cantidad:    0,
	}
}
func crearNodo[K comparable, V any](clave K, valor V) *nodo[K, V] {
	return &nodo[K, V]{
		hijo_izq: nil,
		hijo_der: nil,
		valor:    valor,
		clave:    clave,
	}
}

func (a *abb[K, V]) buscar(clave K) (*nodo[K, V], *nodo[K, V]) {
	return a.wrapper_buscar(clave, a.raiz, nil)
}

func (a *abb[K, V]) wrapper_buscar(clave K, nodo *nodo[K, V], padre *nodo[K, V]) (*nodo[K, V], *nodo[K, V]) {
	if nodo == nil {
		return nil, padre
	}
	comparado := a.comparacion(clave, nodo.clave)

	if comparado > 0 {
		return a.wrapper_buscar(clave, nodo.hijo_der, nodo)
	} else if comparado < 0 {
		return a.wrapper_buscar(clave, nodo.hijo_izq, nodo)
	} else {
		return nodo, padre
	}
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := a.buscar(clave)
	return nodo != nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodo, _ := a.buscar(clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.valor
}

func (a *abb[K, V]) Guardar(clave K, valor V) {

	nodo_actual, padre := a.buscar(clave)

	if nodo_actual != nil {
		nodo_actual.valor = valor
	} else {
		nodo := crearNodo[K, V](clave, valor)
		if a.raiz == nil {
			a.raiz = nodo
		} else if a.comparacion(clave, padre.clave) < 0 {
			padre.hijo_izq = nodo
		} else {
			padre.hijo_der = nodo
		}
		a.cantidad++
	}
}

func contarHijos[K comparable, V any](nodo *nodo[K, V]) int {
	cont := 0
	if nodo.hijo_der != nil {
		cont++
	}
	if nodo.hijo_izq != nil {
		cont++
	}

	return cont
}
func (n *nodo[K, V]) reemplazante() *nodo[K, V] {

	return n.wrapper_reemplazante()
}

func (n *nodo[K, V]) wrapper_reemplazante() *nodo[K, V] {
	if n.hijo_der == nil {
		return n
	}
	return n.hijo_der.wrapper_reemplazante()
}

func (a *abb[K, V]) Borrar(clave K) V {
	nodo, padre := a.buscar(clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	cant_hijos := contarHijos[K, V](nodo)
	dato := nodo.valor

	if cant_hijos == 0 {
		if padre == nil {
			a.raiz = nil
		} else if a.comparacion(padre.clave, nodo.clave) < 0 {
			padre.hijo_der = nil
		} else {
			padre.hijo_izq = nil
		}
	} else if cant_hijos == 1 {
		if nodo.hijo_izq == nil {
			if padre == nil {
				a.raiz = nodo.hijo_der
			} else if a.comparacion(padre.clave, nodo.clave) < 0 {
				padre.hijo_der = nodo.hijo_der
			} else {
				padre.hijo_izq = nodo.hijo_der
			}
		} else {
			if padre == nil {
				a.raiz = nodo.hijo_izq
			} else if a.comparacion(padre.clave, nodo.clave) < 0 {
				padre.hijo_der = nodo.hijo_izq
			} else {
				padre.hijo_izq = nodo.hijo_izq
			}
		}
	} else {
		reemplazo := nodo.reemplazante()
		clave_reemp := reemplazo.clave
		valor_reemp := a.Borrar(clave_reemp)
		nodo.clave, nodo.valor = clave_reemp, valor_reemp
		a.cantidad++
	}

	a.cantidad--
	return dato
}

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	a.iterarWraper(desde, hasta, visitar, a.raiz)

}

func (a *abb[K, V]) iterarWraper(desde *K, hasta *K, visitar func(clave K, dato V) bool, nodo *nodo[K, V]) bool {
	if nodo == nil {
		return true
	}

	if desde != nil && a.comparacion(nodo.clave, *desde) < 0 {
		return a.iterarWraper(desde, hasta, visitar, nodo.hijo_der)

	}
	if hasta != nil && a.comparacion(nodo.clave, *hasta) > 0 {
		return a.iterarWraper(desde, hasta, visitar, nodo.hijo_izq)

	}

	variable_bool := a.iterarWraper(desde, hasta, visitar, nodo.hijo_izq)
	if !variable_bool {
		return false
	}

	if !visitar(nodo.clave, nodo.valor) {
		return false
	}
	return a.iterarWraper(desde, hasta, visitar, nodo.hijo_der)
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iteradorABB[K, V]{
		abb:   a,
		pila:  TDAPila.CrearPilaDinamica[*nodo[K, V]](),
		desde: desde,
		hasta: hasta,
	}
	iter.iterApilar(iter.abb.raiz)
	return iter
}

func (iter *iteradorABB[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() || iter.pila.EstaVacia() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().valor
}

func (iter *iteradorABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.Desapilar()
	if nodo.hijo_der != nil {
		iter.iterApilar(nodo.hijo_der)
	}
}

func (iter *iteradorABB[K, V]) iterApilar(nodo *nodo[K, V]) {
	if nodo == nil {
		return
	}

	if iter.desde != nil && iter.abb.comparacion(nodo.clave, *iter.desde) < 0 {
		iter.iterApilar(nodo.hijo_izq)
		return

	} else if iter.hasta != nil && iter.abb.comparacion(nodo.clave, *iter.hasta) > 0 {
		iter.iterApilar(nodo.hijo_der)
		return
	}

	iter.pila.Apilar(nodo)
	iter.iterApilar(nodo.hijo_izq)

}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

func (a *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	a.IterarRango(nil, nil, f)
}
