package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	cap_inicial := 5
	//crea una pila que tiene su arreglo(datos) y la cantidad inicial de elementos que es 0
	pila := &pilaDinamica[T]{
		datos:    make([]T, cap_inicial),
		cantidad: 0,
	}

	return pila

}

func (p *pilaDinamica[T]) EstaVacia() bool {
	//devuelve booleano indicando si la pila esta vacia. True para vacio, False en caso contrario
	return 0 == p.cantidad

}
func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	//devuelve el ultimo elemento almacenado en la pila, es -1 porque se cuenta desde el 0
	return p.datos[p.cantidad-1]

}

func (p *pilaDinamica[T]) Desapilar() T {

	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	p.cantidad--
	//se actualiza la cantidad de elementos
	p.redimensionar()

	return p.datos[p.cantidad]
}

// defino funcion maximo para usar en Apilar
func maximo(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func (p *pilaDinamica[T]) redimensionar() {
	num_reducir := 4
	num_ampliar := 2
	nueva_capacidad := len(p.datos)

	if p.cantidad*num_reducir <= len(p.datos)/num_reducir && len(p.datos) > num_reducir {
		//ajusto la capacidad del arreglo segun esta indicado en la pagina
		nueva_capacidad /= num_ampliar

	}

	capacidad := len(p.datos)
	//ajusto la capacidad segun como esta indicado en la pagina
	if p.cantidad == capacidad {
		nueva_capacidad *= num_ampliar

	}
	if nueva_capacidad != len(p.datos) {
		nuevo := make([]T, nueva_capacidad)
		//a ese nuevo le meto los datos previos
		copy(nuevo, p.datos)

		p.datos = nuevo
	}

}

func (p *pilaDinamica[T]) Apilar(item T) {

	p.redimensionar()
	//se agrega elemento en el espacio libre
	p.datos[p.cantidad] = item
	p.cantidad++
}

func AgregarAlFondo(p pila[T], elem T){
	if p.EstaVacia(){
		p.Apilar(elem)
	}
	dato:=p.Desapilar()
	if !p.EstaVacia(){
		AgregarAlFondo(p, elem)
	}
	p.Apilar(dato)
}