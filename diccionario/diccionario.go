package diccionario

type Diccionario[K comparable, V any] interface {
	//Guarda un valor asociado a una clave en el diccionario
	Guardar(clave K, dato V)

	//verifica si la clave pasada por parametro esta en ek diccionario, devolviendo true o false.
	Pertenece(clave K) bool

	//Esta funcion obtiene el/los valores asociados a la cave recibida por parametro. En caso de no existir dicha clave,
	//entra en pánico con el mensaje 'La clave no pertenece al diccionario'.
	Obtener(clave K) V

	//Esta funcion elimina el/los valores asociados a la clave recibida por parametro.
	//En caso de no existir la clave, entra en pánico con el mensaje 'La clave no pertenece al diccionario'.
	Borrar(clave K) V

	//Devuelve la cantidad de claves asociadas al diccionario
	Cantidad() int

	//Itera sobre el diccionario mientras la funcion pasada por parametro devuelva true sobre el dato iterado, deja de iterar si la funcion devuelve false o termina de recorrer el diccionario.
	Iterar(func(clave K, dato V) bool)

	//Devuelve un iterador del tipo IteradorDiccionario
	Iterador() IterDiccionario[K, V]
}

type IterDiccionario[K comparable, V any] interface {
	//Devuelve True en el caso que haya un elemento siguiente en el diccionario, False en caso contrario
	HaySiguiente() bool

	//devuelve la clave y valor sobre la que estamos iterando. En caso de que no se este iterando entra en panico con el mensaje 'El iterador termino de iterar'
	VerActual() (K, V)

	//itera el siguiente elemento, si no hay siguiente entra en panico con el mensaje 'El iterador termino de iterar'.
	Siguiente()
}
