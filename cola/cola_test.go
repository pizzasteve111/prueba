package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const mensajeError = "La cola esta vacia"

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, mensajeError, func() { cola.Desencolar() })
	require.PanicsWithValue(t, mensajeError, func() { cola.VerPrimero() })
}

func TestVolumen(t *testing.T) {
	tam := 10000
	cola := TDACola.CrearColaEnlazada[int]()

	for i := 0; i <= tam; i++ {
		cola.Encolar(i)
	}

	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 0, cola.VerPrimero())

	for i := 0; i <= tam; i++ {
		elem := cola.Desencolar()
		require.EqualValues(t, elem, i)

	}

	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestEncolarDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()

	//prueba para encolar distintos tipos de datos en una misma cola
	cola.Encolar(12)
	cola.Encolar("a")
	cola.Encolar(true)

	require.False(t, cola.EstaVacia())

	elemento1 := cola.Desencolar()
	require.IsType(t, 0, elemento1)
	require.EqualValues(t, 12, elemento1)

	elemento2 := cola.Desencolar()
	require.IsType(t, "hola", elemento2)
	require.EqualValues(t, "a", elemento2)

	elemento3 := cola.Desencolar()
	require.IsType(t, false, elemento3)
	require.EqualValues(t, true, elemento3)

	require.True(t, cola.EstaVacia())
}
