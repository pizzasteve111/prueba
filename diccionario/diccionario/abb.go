package diccionario

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
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccinarioOrdenado[K, V] {
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

	if comparado == 0 {
		return nodo, padre
	} else if comparado < 0 {
		return a.wrapper_buscar(clave, nodo.hijo_izq, nodo)
	} else {
		return a.wrapper_buscar(clave, nodo.hijo_der, nodo)
	}
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := a.buscar(clave)
	return nodo == nil
}

func (a *abb[K, V]) Obtener(clave K) V {
	nodo, _ := a.buscar(clave)
	return nodo.valor
}

func (a *abb[K, V]) Guardar(clave K, valor V) {
	nodo := crearNodo[K, V](clave, valor)
	if a.raiz == nil {
		a.raiz = nodo
	}
	nodo_actual, padre := a.buscar(clave)

	if nodo_actual != nil {
		nodo_actual.valor = valor
	} else {
		if a.comparacion(clave, padre.clave) < 0 {
			padre.hijo_izq = nodo
		} else {
			padre.hijo_der = nodo
		}
	}
	a.cantidad++
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
	for n.hijo_izq != nil {
		n = n.hijo_der
	}
	return n
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
		}
		if a.comparacion(padre.clave, nodo.clave) < 0 {
			padre.hijo_der = nil
		} else {
			padre.hijo_izq = nil
		}
	} else if cant_hijos == 1 {
		if nodo.hijo_izq == nil {
			if padre == nil {
				a.raiz = nodo.hijo_der
			}
			if a.comparacion(padre.clave, nodo.clave) < 0 {
				padre.hijo_der = nodo.hijo_der
			} else {
				padre.hijo_izq = nodo.hijo_der
			}
		} else {
			if padre == nil {
				a.raiz = nodo.hijo_izq
			}
			if a.comparacion(padre.clave, nodo.clave) < 0 {
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

func (a *abb[K, V]) CrearIterador(desde *K, hasta *K) *IteradorRango[K, V] {
	return &iteradorABB[K, V]{
		abb:   a,
		desde: desde,
		hasta: hasta,
	}
}
