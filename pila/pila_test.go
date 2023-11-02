package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

func TestVolumen(t *testing.T) {
	tam := 10000
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < tam; i++ {
		pila.Apilar(i)
	}

	require.False(t, pila.EstaVacia())
	require.EqualValues(t, tam-1, pila.VerTope())
	for i := tam - 1; i >= 0; i-- {
		elem := pila.Desapilar()
		require.EqualValues(t, elem, i)
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })

}

func TestApilarDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[any]()
	pilaEnetero := TDAPila.CrearPilaDinamica[int]()
	pilaString := TDAPila.CrearPilaDinamica[string]()

	pila.Apilar(pilaEnetero)
	pila.Apilar(8)
	pila.Apilar("hola")
	pila.Apilar(false)

	require.EqualValues(t, false, pila.Desapilar())
	require.EqualValues(t, "hola", pila.Desapilar())
	require.EqualValues(t, 8, pila.Desapilar())

	pilaEnetero.Apilar(900)
	pilaString.Apilar("b")

	require.Equal(t, 900, pilaEnetero.Desapilar())
	require.Equal(t, "b", pilaString.Desapilar())

	require.False(t, pila.EstaVacia())

	pila.Desapilar()

	require.True(t, pila.EstaVacia())
	require.True(t, pilaEnetero.EstaVacia())
	require.True(t, pilaString.EstaVacia())
}
