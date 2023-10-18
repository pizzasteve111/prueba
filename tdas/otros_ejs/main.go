package pruebas

import (
	"fmt"
)

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
