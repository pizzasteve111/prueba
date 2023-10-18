package lista

type IteradorLista[T any] interface {
	//Devuelve elemento actual
	VerActual() T

	//Devuelve true si hay elemento siguiente, false en caso contrario.
	HaySiguiente() bool

	//Itera el proximo elemento. Si no hay siguiente, entra en panico con el mensaje "El iterador termino de iterar"
	Siguiente()

	//Inserta un elemento siguiente al elemento actual
	Insertar(T)

	//Borra y devuelve el elemento actual. Si no hay elemento a borrar, entra en panico con el mensaje "El iterador termino de iterar"
	Borrar() T
}

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

	//
	Iterar(visitar func(T) bool)

	//Devuelve un iterador del tipo IteradorLista
	Iterador() IteradorLista[T]
}
