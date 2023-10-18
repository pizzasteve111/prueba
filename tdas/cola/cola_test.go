package cola_test

import (
	"fmt"
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ColaVacia(t *testing.T) {
	fmt.Println("Test 1")
	cola := TDACola.CrearColaEnlazada[any]()
	require.Equal(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})
}

func Test_Volumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()
	elementos := 1000
	for i := 0; i < elementos; i++ {
		dato := i
		cola.Encolar(dato)
	}
	require.Equal(t, false, cola.EstaVacia())
	for f := 0; f < elementos; f++ {

		dato2 := f
		dato_comp := cola.Desencolar()
		require.Equal(t, dato2, dato_comp)
	}
	require.Equal(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})

}
func Test_Cola_Variada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()
	copia := TDACola.CrearColaEnlazada[any]()
	vector := []int{1, 2, 3, 4}
	cola.Encolar(copia)
	cola.Encolar("segundo")
	cola.Encolar(vector)
	cola.Encolar(4)
	require.Equal(t, false, cola.EstaVacia())
	dato_col := cola.Desencolar()
	require.Equal(t, dato_col, copia)

	dato_str := cola.Desencolar()
	require.Equal(t, dato_str, "segundo")
	dato_vec := cola.Desencolar()
	require.Equal(t, dato_vec, vector)
	dato_int := cola.Desencolar()
	require.Equal(t, dato_int, 4)
	//verifico que la cola queda correctamente vacia
	require.Equal(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})

}
func Test_Encolar_Desencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()
	cola.Encolar("prueba")
	//verifico que la cola no esta vacia
	require.False(t, cola.EstaVacia())
	require.Equal(t, "prueba", cola.VerPrimero())
	//ahora la vacio, se debe comportar como una cola recien creada
	cola.Desencolar()
	require.Equal(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})

}
func Test_ColaNueva(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()
	require.Equal(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})
}

func Test_ColaVaciada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()
	cola.Encolar(1)
	cola.Desencolar()

	require.Equal(t, true, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.Desencolar()
	})
	require.PanicsWithValue(t, "La cola esta vacia", func() {
		cola.VerPrimero()
	})
}
