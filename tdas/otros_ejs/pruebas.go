package main

import (
	"fmt"
)

/*//TDAPilita "tdas/ejs" // Asegúrate de importar el paquete correcto aquí

type Pilita[T any] interface {

	// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
	EstaVacia() bool

	// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
	// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
	VerTope() T

	// Apilar agrega un nuevo elemento a la pila.
	Apilar(T)

	// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
	// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
	Desapilar() T

	Miltitop(int) []any

	Aplicar(f func(T) T) Pilita[T]
}
type pilita[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilita[T any]() Pilita[T] {
	//crea una pila que tiene su arreglo(datos) y la cantidad inicial de elementos que es 0
	pila := &pilita[T]{
		datos:    make([]T, 0),
		cantidad: 0,
	}

	return pila

}
func (p *pilita[T]) Aplicar(f func(T) T) Pilita[T] {
	copia := CrearPilita[T]()
	res := CrearPilita[T]()
	for p.EstaVacia() != true {
		dato := p.Desapilar()
		copia.Apilar(f(dato))
	}
	for copia.EstaVacia() != true {
		res.Apilar(copia.Desapilar())
	}
	return res
}
func (p *pilita[T]) EstaVacia() bool {
	//devuelve booleano indicando si la pila esta vacia. True para vacio, False en caso contrario
	return 0 == p.cantidad

}
func (p *pilita[T]) VerTope() T {
	if p.EstaVacia() != false {
		panic("La pila esta vacia")
	}
	//devuelve el ultimo elemento almacenado en la pila, es -1 porque se cuenta desde el 0
	return p.datos[p.cantidad-1]

}

func (p *pilita[T]) Desapilar() T {
	orden := "desapilar"
	if p.EstaVacia() != false {
		panic("La pila esta vacia")
	}

	res := p.datos[p.cantidad-1]
	//res va a ser la pila sin el ultimo elemento
	p.cantidad--
	//se actualiza la cantidad de elementos
	p.redimensionar(orden)

	return res
}

// defino funcion maximo para usar en Apilar
func maximo(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func (p *pilita[T]) redimensionar(orden string) {
	if orden == "desapilar" {
		if p.cantidad <= len(p.datos)/4 && len(p.datos) > 4 {
			//ajusto la capacidad del arreglo segun esta indicado en la pagina
			nueva_capacidad := len(p.datos) / 2
			if nueva_capacidad < 4 {
				nueva_capacidad = 4
			}
			//creo arreglo vacio para datos T con la capacidad actualizada
			nuevo := make([]T, nueva_capacidad)
			//a ese nuevo le meto los datos previos
			copy(nuevo, p.datos)

			p.datos = nuevo
		}

	} else if orden == "apilar" {
		capacidad := len(p.datos)
		//ajusto la capacidad segun como esta indicado en la pagina
		if p.cantidad == capacidad {
			nueva_capacidad := maximo(len(p.datos)*2, 1)
			nuevo := make([]T, nueva_capacidad)
			copy(nuevo, p.datos)
			//ahora la pila conserva los mismos elementos,pero el doble de capacidad
			p.datos = nuevo

		}

	}
}

func (p *pilita[T]) Apilar(item T) {
	orden := "apilar"
	p.redimensionar(orden)
	//se agrega elemento en el espacio libre
	p.datos[p.cantidad] = item
	p.cantidad++
}
func (p *pilita[T]) Miltitop(n int) []any {
	cant := min(n, p.cantidad)
	res := make([]any, cant)
	for i := 0; i < cant; i++ {
		dato := p.Desapilar()
		res[i] = dato
	}
	return res
}

type Cola[T any] interface {

	// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
	EstaVacia() bool

	// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
	// "La cola esta vacia".
	VerPrimero() T

	// Encolar agrega un nuevo elemento a la cola, al final de la misma.
	Encolar(T)

	// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
	Desencolar() T

	Multiprimeros(k int) []T
}

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
	cola := &cola_Enlazada[T]{
		inicio: nil,
		fin:    nil,
	}
	return cola
}
func (c *cola_Enlazada[T]) EstaVacia() bool {
	return c.inicio == nil
}
func (c *cola_Enlazada[T]) VerPrimero() T {

	if c.EstaVacia() == true {
		panic("La cola esta vacia")
	}
	return c.inicio.dato
}
func (c *cola_Enlazada[T]) Encolar(item T) {
	nuevo := &nodo[T]{dato: item, prox: nil}
	//la cola esta vacia, el primer y ultimo nodo es el "mismo" por lo que comparten el mismo nodo
	if c.fin == nil {
		c.inicio = nuevo
		c.fin = nuevo
	} else {
		//ahora fin va a apuntar al nuevo nodo con el dato que quiero encolar
		c.fin.prox = nuevo
		//el ultimo nodo de la lista se va a convertir en este mismo nodo, asi la proxima vez que encole se hace desde el nuevo nodo, tomandolo como el ultimo
		c.fin = nuevo
	}

}

func (c *cola_Enlazada[T]) Desencolar() T {
	if c.EstaVacia() == true {
		panic("La cola esta vacia")
	}
	//res es el dato que devuelvo
	res := c.inicio.dato
	//ahora el primer nodo de la cola va a ser el que le sigue al nodo que desencolo
	c.inicio = c.inicio.prox
	//devuelvo el dato
	return res

}
func (c *cola_Enlazada[T]) Multiprimeros(k int) []T {
	res := []T{}
	if c.EstaVacia() == true {
		return res
	}
	cont := 0
	for c.EstaVacia() != true {
		if cont <= k {
			res = append(res, c.Desencolar())
		}
		cont++

	}
	return res
	//es de orden O(n)
}
//DIV Y CONQUISTA

//EJ 2
/* el algoritmo planteado no es de division y conquista pues no funciona de manera recursiva. Es decir que se divide el arreglo en partes mas sencillas
y se le aplica la misma funcion para que a esa porcion la vuelva a sub dividir, asi hasta llegar al caso base */

func min_arreglo(arreglo []int) int {
	fmt.Println("hola")
	return wrapper_min(arreglo, 0, len(arreglo)-1)
}
func wrapper_min(arreglo []int, inicio int, fin int) int {
	if inicio == fin {
		return arreglo[fin]
	}
	medio := (inicio + fin) / 2
	min_izq := wrapper_min(arreglo, inicio, medio)
	min_der := wrapper_min(arreglo, medio+1, fin)
	if min_der < min_izq {
		return min_der
	}
	return min_izq

}

func ordenado(arreglo []int, inicio int, fin int) bool {
	if len(arreglo) <= 1 {
		return true
	}
	medio := (inicio + fin) / 2

	izquierdaOrdenado := ordenado(arreglo, inicio, medio)
	derechaOrdenado := ordenado(arreglo, medio+1, fin)

	return izquierdaOrdenado == true && derechaOrdenado == true && (arreglo[medio] <= arreglo[medio+1])

}
func pruebas() {
	arreglo1 := []int{1, 2, 3, 4, 5}
	arreglo2 := []int{5, 4, 3, 2, 1}
	arreglo3 := []int{1, 3, 2, 4, 5}

	fmt.Println(ordenado(arreglo1, 0, len(arreglo1)-1)) // Debe imprimir true
	fmt.Println(ordenado(arreglo2, 0, len(arreglo2)-1)) // Debe imprimir false
	fmt.Println(ordenado(arreglo3, 0, len(arreglo3)-1)) // Debe imprimir false
}
