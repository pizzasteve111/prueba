package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PrimerElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)

	require.Equal(t, lista.VerPrimero(), 1)

	iter := lista.Iterador()

	iter.Insertar(3)

	require.Equal(t, lista.VerPrimero(), 3)
	require.Equal(t, iter.VerActual(), 3)
}

func Test_UltimoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

	iter := lista.Iterador()
	iter.Insertar(1)
	iter.Insertar(2)
	require.False(t, lista.EstaVacia())
	iter.Siguiente()
	//[2,1]
	require.Equal(t, 1, iter.VerActual())

	require.Equal(t, lista.VerUltimo(), 1)
}

func Test_MedioElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })

	iter := lista.Iterador()
	iter.Insertar(1)
	iter.Insertar(2)
	iter.Insertar(3)
	iter.Insertar(4)

	require.Equal(t, 4, lista.Largo())
	//[4,3,2,1]
	medio := 4 / 2
	cont := 0

	for i := 0; i <= medio && iter.HaySiguiente(); i++ {

		iter.Siguiente()

		cont++
	}

	iter.Insertar(5)
	require.Equal(t, 5, lista.Largo())
	//[4,3,5,2,1]

	require.Equal(t, 5, iter.VerActual())
	require.Equal(t, lista.Largo()/2, lista.Largo()-cont)
	require.Equal(t, 1, lista.VerUltimo())

}

func Test_EliminarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 3, lista.Largo())

	iterador := lista.Iterador()
	//[1,2,3]
	dato := iterador.Borrar()

	require.Equal(t, 2, lista.Largo())
	require.Equal(t, 1, dato)
	require.Equal(t, 2, lista.VerPrimero())
	require.Equal(t, 3, lista.VerUltimo())

}

func Test_EliminarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	lista.InsertarPrimero("a")
	lista.InsertarUltimo("b")
	lista.InsertarUltimo("c")

	iterador := lista.Iterador()
	for i := 1; i < 3; i++ {
		iterador.Siguiente()
	}

	require.Equal(t, "c", iterador.VerActual())
	dato := iterador.Borrar()
	require.Equal(t, "c", dato)
	require.Equal(t, "b", lista.VerUltimo())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })

}

func Test_BorrarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	lista.InsertarPrimero("a")
	lista.InsertarUltimo("b")
	lista.InsertarUltimo("c")
	lista.InsertarUltimo("d")
	lista.InsertarUltimo("e")

	iterador := lista.Iterador()

	medio := 5 / 2

	for i := 0; i < medio; i++ {
		iterador.Siguiente()
	}
	require.Equal(t, "c", iterador.VerActual())
	dato := iterador.Borrar()
	require.Equal(t, "c", dato)
	require.Equal(t, "d", iterador.VerActual())

}

func Test_IteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)

	cont_par := 0

	lista.Iterar(func(elemento int) bool {
		if elemento%2 == 0 {
			cont_par++
		}
		return true
	})
	require.Equal(t, 4, cont_par)

}

func Test_IteradorInternoEncontrarElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(4)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)

	elem_buscado := 5
	pos := 0
	encontrado := false
	lista.Iterar(func(n int) bool {
		if n == elem_buscado {
			encontrado = true
			return false
		}
		pos++
		return true
	})
	require.Equal(t, 4, pos)
	require.Equal(t, true, encontrado)
}
