package main

import (
	"fmt"
)

// ej composicion de funciones
/*
type ComposicionFunciones struct {
	funciones Pilita[func(float64) float64]
}

func CrearComposicion() ComposicionFunciones {
	return ComposicionFunciones{
		funciones: CrearPilita[func(float64) float64](),
	}
}
func (c *ComposicionFunciones) AgregarFuncion(f func(float64) float64) {
	c.funciones.Apilar(f)
}
func (c *ComposicionFunciones) Aplicar(x float64) float64 {
	dato := x
	for c.funciones.EstaVacia() != true {
		f := c.funciones.Desapilar()
		dato = f(dato)
	}
	return dato
}
func ascendencia(p Pilita[int]) bool {
	if p.EstaVacia() == true {
		return true
	}

	copia := CrearPilita[int]()
	tope := p.Desapilar()
	copia.Apilar(tope)
	for p.EstaVacia() != true {
		comp := p.Desapilar()
		if tope > comp {
			return false
		}
		copia.Apilar(comp)
		tope = comp
	}
	for copia.EstaVacia() != true {
		dato := copia.Desapilar()
		p.Apilar(dato)
	}
	return true
	//el rendimiento es O(n)
}
*/

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

	Invertir()
}

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

// DIV Y CONQUISTA
func MergePilas(p1 Pilita[int], p2 Pilita[int]) []int {
	resultado := []int{}
	copia1 := CrearPilita[int]()
	copia2 := CrearPilita[int]()

	for p1.EstaVacia() != true {
		copia1.Apilar(p1.Desapilar())
	}
	for p2.EstaVacia() != true {
		copia2.Apilar(p2.Desapilar())
	}
	for copia1.EstaVacia() != true && copia2.EstaVacia() != true {
		dato1 := copia1.VerTope()
		dato2 := copia2.VerTope()
		if dato1 <= dato2 {
			resultado = append(resultado, dato1)

			copia1.Desapilar()
		} else if dato2 <= dato1 {
			resultado = append(resultado, dato2)

			copia2.Desapilar()
		}
	}
	for copia1.EstaVacia() != true {
		resultado = append(resultado, copia1.Desapilar())
	}
	for copia2.EstaVacia() != true {
		resultado = append(resultado, copia2.Desapilar())
	}
	return resultado
}

func min_arreglo(arreglo []int) int {
	fmt.Println("hola")
	if len(arreglo) == 0 {
		// Si el arreglo está vacío, no hay mínimo que retornar
		panic("El arreglo está vacío")
	}
	return wrapper_min(arreglo, 0, len(arreglo)-1)
}
func wrapper_min(arreglo []int, inicio int, fin int) int {
	if inicio >= fin {
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
	if fin-inicio == 0 {
		// Si el rango tiene 0 o 1 elemento, está ordenado por definición
		return true
	}
	medio := (fin + inicio) / 2

	izquierdaOrdenado := ordenado(arreglo, inicio, medio)
	derechaOrdenado := ordenado(arreglo, medio+1, fin)

	return izquierdaOrdenado == true && derechaOrdenado == true && (arreglo[medio] <= arreglo[medio+1])
}
func encuentraElementoFueraDeLugar(arr []int) int {
	return buscaFueraDeLugar(arr, 0, len(arr)-1)
}

func buscaFueraDeLugar(arr []int, izquierda, derecha int) int {
	if izquierda > derecha {
		// No debería ocurrir, pero por precaución.
		return -1
	}

	if izquierda == derecha {
		return arr[izquierda]
	}

	medio := (izquierda + derecha) / 2

	if arr[medio] == medio {
		// El elemento en medio está en su posición correcta.
		// El elemento fuera de lugar está en la mitad derecha.
		return buscaFueraDeLugar(arr, medio+1, derecha)
	}

	// El elemento fuera de lugar está en la mitad izquierda.
	return buscaFueraDeLugar(arr, izquierda, medio)
}

func BuscarCero(arreglo []int) int {
	return wrapper_cero(arreglo, 0, len(arreglo)-1)
}

func wrapper_cero(arreglo []int, inicio int, fin int) int {
	if inicio == fin {
		if arreglo[inicio] == 0 {
			return inicio
		}
		return -1
	}
	medio := (inicio + fin) / 2
	if medio-1 >= 0 && arreglo[medio-1] == 1 && arreglo[medio] == 0 {
		return medio
	} else if arreglo[medio] == 1 {
		return wrapper_cero(arreglo, medio+1, fin)
	} else {
		return wrapper_cero(arreglo, inicio, medio)
	}

}
func PosicionPico(v []int, ini, fin int) int {
	if ini == fin {
		return v[ini]
	}
	medio := (ini + fin) / 2
	if v[medio] > v[medio-1] && v[medio] > v[medio+1] {
		return medio
	}
	if v[medio] > v[medio+1] {
		return PosicionPico(v, ini, medio)
	} else {
		return PosicionPico(v, medio+1, fin)
	}
}
func buscar_raiz(f func(int) int, a int, b int) int {
	if a == b {
		return f(a)
	}
	medio := (a + b) / 2
	if f(medio) == 0 {
		return medio
	} else if (f(medio) * f(a)) < 0 {
		return buscar_raiz(f, a, medio)
	} else {
		return buscar_raiz(f, medio+1, b)
	}

}
func saliente(arr []int) int {
	return wrapper_saliente(arr, 0, len(arr)-1)
}
func wrapper_saliente(arr []int, ini int, fin int) int {
	if ini == fin {
		return arr[ini]
	}
	medio := (fin + ini) / 2
	if arr[medio-1] > arr[medio] || arr[medio+1] < arr[medio] {
		return arr[medio]
	}
	if arr[medio-1] < arr[medio] {
		return wrapper_saliente(arr, medio+1, fin)
	} else {
		return wrapper_saliente(arr, ini, medio)
	}
}
func ta_ordenado(arr []int) bool {
	return wrapper_orden(arr, 0, len(arr)-1)
}
func wrapper_orden(arr []int, ini int, fin int) bool {
	if ini == fin {
		return true
	}
	medio := (ini + fin) / 2
	if arr[medio-1] > arr[medio] || arr[medio] > arr[medio+1] {
		return false
	}
	if arr[medio] > arr[medio-1] {
		return wrapper_orden(arr, medio+1, fin)
	} else {
		return wrapper_orden(arr, ini, medio)
	}

}

/*
	func leiva_joyas(arr[]int)int {
		return wrapper_joyas(arr,0,len(arr)-1)
	}

	func wrapper_joyas(arr []int,ini int, fin int) int{
		if ini==fin{
			return arr[ini]
		}
		medio:=(ini+fin)/2
		grupo1:=arr[ini:medio]
		grupo2:=arr[medio+1:fin]
		if balanza(grupo1,grupo2)>0{
			return wrapper_joyas(arr,ini,medio)
		} else if balanza(grupo1,grupo2)==0{
			return 0
		}else {
			return wrapper_joyas(arr,medio+1,fin)
		}
	}
*/

func indicePrimerCero(arr []int) int {
	return wrapper_indice(arr, 0, len(arr)-1)
}
func wrapper_indice(arr []int, ini int, fin int) int {
	if ini == fin {
		if arr[ini] != 0 {
			return -1
		} else {
			return ini
		}
	}
	medio := (ini + fin) / 2
	if medio-1 >= 0 && arr[medio] == 0 && arr[medio-1] == 1 {
		return medio
	}
	if arr[medio] == 1 {
		return wrapper_indice(arr, medio+1, fin)
	} else {
		return wrapper_indice(arr, ini, medio)
	}
}

func multiplicar(arr []float64) float64 {
	return wrapper_mult(arr, 0, len(arr)-1)
}
func wrapper_mult(arr []float64, ini int, fin int) float64 {
	if ini == fin {
		return arr[ini]
	}
	medio := (ini + fin) / 2
	return wrapper_mult(arr, ini, medio) * wrapper_mult(arr, medio+1, fin)
}

type Nacionalidad int
type Persona struct {
	Edad         int
	Nombre       int
	Nacionalidad Nacionalidad
}

/*
func ordenarclientes(arr []Persona) []Persona {
	//creo un arreglo de 32 que son listas de personas que comparten nacionalidad
	menores := []Persona{}
	adultos := make([][]Persona, 32)
	for _, persona := range arr {
		if persona.Edad <= 12 {
			menores = append(menores, persona)

		} else {
			adultos[persona.Nacionalidad] = append(adultos[persona.Nacionalidad], persona)
		}
	}
	//ordeno a los menores de menor a mayor
	menores = Ordenar(menores)

	var adultosordenados []Persona

	for _, grupos := range adultos {
		//ordeno por edad al grupo de personas con misma nacionalidad
		grupo = Ordenar(grupos)

		adultosordenados = append(adultosordenados, grupo...)
	}
	res := append(menores, adultosordenados...)
	return res

}
*/
/*
func buscarDiaFalla(dia int) int{
	return wrapper_dia(0,dia)
}
func wrapper_dia(ini int, fin int) int{
	if ini==fin{
		return ini
	}
	medio:=(ini+fin)/2
	if !todoOkElDia(medio) && todoOkElDia(medio-1)==true{
		return medio
	}
	if todoOkElDia(medio)==true{
		return wrapper_dia(medio+1,fin)
	}else{
		return wrapper_dia(ini,medio)
	}
}*/
/*
func prueba_volumen(n int) int{
	return wrapper_vol(0,n)
}
func wrapper_vol(ini int, fin int) int{
	if ini==fin{
		return ini
	}
	medio:=(ini+fin)/2
	if !testVoluemn(medio) && testVolumen(medio-1){
		return medio
	}
	if testVoluemn(medio) {
		return wrapper_vol(medio +1,fin)

	}else{
		return wrapper_vol(ini,medio)
	}
}
*/
func parto(pila Pilita[int], mitad int) Pilita[int] {
	auxiliar := CrearPilita[int]()
	res := CrearPilita[int]()
	for i := 0; i <= mitad; i++ {
		auxiliar.Apilar(pila.Desapilar())
	}
	for !auxiliar.EstaVacia() {
		res.Apilar(auxiliar.Desapilar())
	}
	return res
}

/*
func (p *pila_enlazada[T])CantidadApariciones(elem T) int{
	if p.tope==nil{
		return 0
	}
	act:=p.tope
	cont:=0
	for p.tope!=nil{
		if p.tope.dato==elem{
			cont++
		}
		act=act.anterios

	}
	return cont
}
*/

type Alumno struct {
	nombre string
	nota1  int
	nota2  int
	nota3  int
}

/*
func ordenPromedio(arreglo []Alumno) []Alumno{
	res:=[]Alumno{}
	//promedios contiene 10 arreglos para cada promedio en especifico
	promedios:=make([][]Alumno,10)
	for _,alumno := range arreglo{

		nota_total:=alumno.nota1+alumno.nota2+alumno.nota3
		promedio:=nota_total/3
		promedios[promedio]=append(promedios[promedio],alumno )
	}
	//
	for _,grupos:=range promedios{
		grupo=Ordenar_Raditz(grupos)
		res=append(res, grupo...)
	}
	return res

}

func ordenAlumno(arreglo []Alumno) []Alumno{
	res:=[]Alumno{}
	//llamo a la funcion de ordenar para que vaya ordenando a los alumnos segun la nota del parcialito
	arreglo=Ordenar_counting(arreglo)
	arreglo=Ordenar_counting(arreglo)
	arreglo=Ordenar_counting(arreglo)
	for _,alumno := range(arreglo){
		promedio:=(alumno.nota1+alumno.nota2+alumno.nota3)/3

	}

}*/

func bool_ordenado(arr []int) bool {
	return wrapper_bool(arr, 0, len(arr)-1)
}

func wrapper_bool(arr []int, ini int, fin int) bool {
	if ini == fin {
		return true
	}
	medio := (ini + fin) / 2
	if medio-1 <= 0 && arr[medio] > arr[medio+1] {
		return false
	}
	if arr[medio] < arr[medio-1] || arr[medio] > arr[medio+1] {
		return false
	}
	if arr[medio] > arr[medio-1] {
		return wrapper_bool(arr, medio+1, fin)
	} else {
		return wrapper_bool(arr, ini, medio)
	}
}

func encontrar_raiz(n int) int {
	return wrapper_raiz(0, n)
}
func wrapper_raiz(ini int, fin int) int {
	if ini == fin {
		return ini
	}

	medio := (ini + fin) / 2
	if medio*medio == fin {
		return medio
	}
	if medio*medio < fin {
		return wrapper_raiz(medio+1, fin)
	} else {
		return wrapper_raiz(ini, medio)
	}
}

type Elemento[T any] struct {
	Dato      T
	Siguiente *Elemento[T]
}

// Cola representa una cola.
type Cola[T any] struct {
	Inicio *Elemento[T]
	Final  *Elemento[T]
	Tamaño int
}

// CrearCola crea una nueva cola vacía.
func CrearCola[T any]() Cola[T] {
	return Cola[T]{}
}

// Encolar agrega un elemento al final de la cola.
func (c *Cola[T]) Encolar(dato T) {
	nuevoElemento := &Elemento[T]{Dato: dato}

	if c.Inicio == nil {
		c.Inicio = nuevoElemento
		c.Final = nuevoElemento
	} else {
		c.Final.Siguiente = nuevoElemento
		c.Final = nuevoElemento
	}

	c.Tamaño++
}

// Desencolar elimina y devuelve el elemento al frente de la cola.
func (c *Cola[T]) Desencolar() T {
	if c.Inicio == nil {
		panic("la cola está vacía")
	}

	dato := c.Inicio.Dato
	c.Inicio = c.Inicio.Siguiente
	c.Tamaño--

	if c.Inicio == nil {
		c.Final = nil
	}

	return dato
}

// EstaVacia verifica si la cola está vacía.
func (c *Cola[T]) EstaVacia() bool {
	return c.Tamaño == 0
}
func partiras[T any](org Cola[T], k int) []Cola[T] {
	res := []Cola[T]{}
	for i := 0; i < k; i++ {
		cola := CrearCola[T]()
		res = append(res, cola)
	}
	cont := 0
	posicion := 0
	for org.EstaVacia() != true {
		dato := org.Desencolar()
		if cont <= k {
			res[posicion].Encolar(dato)
		}
		if cont == k {
			cont = 0
			posicion++
		}
		cont++
	}
	return res
}
func multiplico(arr []float64) float64 {
	return wrapper_multiplico(arr, 0, len(arr)-1)
}
func wrapper_multiplico(arr []float64, ini int, fin int) float64 {
	if ini == fin {
		return arr[ini]
	}
	medio := (ini + fin) / 2
	izq := wrapper_multiplico(arr, ini, medio)

	der := wrapper_multiplico(arr, medio+1, fin)
	return izq * der
}
func orden(pila Pilita[int]) bool {
	aux := CrearPilita[int]()

	for pila.EstaVacia() != true {
		dato := pila.Desapilar()
		if pila.EstaVacia() != true && dato > pila.VerTope() {
			return false
		}
		aux.Apilar(dato)
	}
	for aux.EstaVacia() != true {
		pila.Apilar(aux.Desapilar())
	}
	return true
}
func multiprimeros(cola Cola[int], k int) []int {
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, cola.Desencolar())
	}
	return res

}
func mergear(p1 Pilita[int],p2 Pilita[int]) []int{
	res:=[]int{}
	aux1:=CrearPilita[int]()
	aux2:=CrearPilita[int]()
	tam_min:=0
	for !p1.EstaVacia(){
		aux1.Apilar(p1.Desapilar())

	}
	for !p2.EstaVacia(){
		aux2.Apilar(p2.Desapilar())
	}
	for !aux1.EstaVacia() && !aux2.EstaVacia(){
		dato1:=aux1.VerTope()
		dato2:=aux2.VerTope()
		if dato1<dato2{
			if tam_min
		}
	}
}

func main() {
	cola := CrearCola[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	fmt.Println(multiprimeros(cola, 2))
}
