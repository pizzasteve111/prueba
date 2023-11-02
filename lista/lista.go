package lista

type Lista[T any] interface {
	//Devuelve verdadero si la lista no tiene elementos, False en caso contrario
	EstaVacia() bool

	//Inserta elemento en la primer posicion de la lista
	InsertarPrimero(T)

	//Inserta elemento en la ultima posicion de la lista
	InsertarUltimo(T)

	//Borra elemento en la primer posicion de la lista. Si esta vacia, entra en panico con el mensaje "La lista esta vacia"
	BorrarPrimero() T

	//Devuelve primer elemento de la lista. Si esta vacia, entra en panico con el mensaje "La lista esta vacia"
	VerPrimero() T

	//Devuelve ultimo elemento de la lista. Si esta vacia, entra en panico con el mensaje "La lista esta vacia"
	VerUltimo() T

	//Devuelve el largo de la lista
	Largo() int

	//Itera sobre la lista mientras la funcion visitar devuelva true sobre el dato iterado, deja de iterar si la funcion visitar devuelve false o termina de recorrer la lista.
	Iterar(visitar func(T) bool)

	//Devuelve un iterador del tipo IteradorLista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	//Devuelve elemento actual. Si no hay elemento a ver, entra en panico con el mensaje "El iterador termino de iterar".
	VerActual() T

	//Devuelve true si hay elemento siguiente, false en caso contrario.
	HaySiguiente() bool

	//Itera el proximo elemento. Si no hay siguiente, entra en panico con el mensaje "El iterador termino de iterar".
	Siguiente()

	//Inserta un elemento entre el elemento actual y el anterior. El iterador actual pasa a estar sobre el nodo insertado.
	Insertar(T)

	//Borra y devuelve el elemento actual. Si no hay elemento a borrar, entra en panico con el mensaje "El iterador termino de iterar".
	Borrar() T
}
