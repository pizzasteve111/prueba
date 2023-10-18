package cola

// defino mi struct Nodo
type nodo[T any] struct {
	dato T
	prox *nodo[T]
}

//defino struct cola

type cola_Enlazada[T any] struct {
	inicio *nodo[T]
	fin    *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return new(cola_Enlazada[T])
}
func (c *cola_Enlazada[T]) EstaVacia() bool {
	return c.inicio == nil
}
func (c *cola_Enlazada[T]) VerPrimero() T {

	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.inicio.dato
}
func (c *cola_Enlazada[T]) Encolar(item T) {
	nuevo := &nodo[T]{dato: item, prox: nil}
	//la cola esta vacia, el primer y ultimo nodo es el "mismo" por lo que comparten el mismo nodo
	if c.fin == nil {
		c.inicio = nuevo

	} else {
		//ahora fin va a apuntar al nuevo nodo con el dato que quiero encolar
		c.fin.prox = nuevo
		//el ultimo nodo de la lista se va a convertir en este mismo nodo, asi la proxima vez que encole se hace desde el nuevo nodo, tomandolo como el ultimo

	}
	c.fin = nuevo

}

func (c *cola_Enlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	//res es el dato que devuelvo
	res := c.inicio.dato
	//ahora el primer nodo de la cola va a ser el que le sigue al nodo que desencolo
	c.inicio = c.inicio.prox
	if c.inicio == nil {
		c.fin = nil
	}
	//devuelvo el dato
	return res

}
