package cola_prioridad_test

import (
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func comparacionNum(a, b int) int {
	return a - b
}
func compararStr(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap(comparacionNum)
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
}

func TestHeapUnElemento(t *testing.T) {
	heap := TDAHeap.CrearHeap(comparacionNum)
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })

	heap.Encolar(1)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 1, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())
}

func TestHeapVariosElem(t *testing.T) {
	heap := TDAHeap.CrearHeap(comparacionNum)

	heap.Encolar(1)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 1, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())

	heap.Encolar(12)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 12, heap.VerMax())
	require.EqualValues(t, 2, heap.Cantidad())

	heap.Encolar(13)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 13, heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())

	heap.Encolar(5)
	require.EqualValues(t, 13, heap.VerMax())
	require.EqualValues(t, 4, heap.Cantidad())
}

func TestDesencolar(t *testing.T) {
	heap := TDAHeap.CrearHeap[string](compararStr)

	heap.Encolar("gato")
	heap.Encolar("vaca")
	heap.Encolar("perro")

	require.Equal(t, "vaca", heap.VerMax())
	require.Equal(t, 3, heap.Cantidad())

	require.Equal(t, "vaca", heap.Desencolar())
	require.Equal(t, 2, heap.Cantidad())

	require.Equal(t, "perro", heap.Desencolar())
	require.Equal(t, 1, heap.Cantidad())

	require.Equal(t, "gato", heap.Desencolar())
	require.Equal(t, 0, heap.Cantidad())

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })

}
func TestVolumen(t *testing.T) {
	tam := 10000
	heap := TDAHeap.CrearHeap[int](comparacionNum)

	for i := 0; i < tam; i++ {
		heap.Encolar(i)
		require.EqualValues(t, i, heap.VerMax())
		require.EqualValues(t, i+1, heap.Cantidad())
	}

	require.False(t, heap.EstaVacia())
	require.EqualValues(t, tam-1, heap.VerMax())

	for i := tam - 1; i >= 0; i-- {
		elem := heap.Desencolar()
		require.EqualValues(t, elem, i)
	}

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestHeapsort(t *testing.T) {
	arr := []string{"perro", "gato", "vaca", "guau", "miau", "moo", "arbol"}
	ordenado := []string{"arbol", "gato", "guau", "miau", "moo", "perro", "vaca"}
	TDAHeap.HeapSort(arr, compararStr)
	require.Equal(t, ordenado, arr)
}

func TestHeapDesdeArreglo(t *testing.T) {
	arr := []int{12, 4, 7, 1, 3, 2, 90}
	arr_esperado := []int{90, 12, 7, 4, 3, 2, 1}
	heap := TDAHeap.CrearHeapArr(arr, comparacionNum)

	require.Equal(t, 90, heap.VerMax())
	require.False(t, heap.EstaVacia())
	require.Equal(t, 7, heap.Cantidad())

	for _, esperado := range arr_esperado {
		elem := heap.Desencolar()
		require.Equal(t, esperado, elem)
	}

	require.Equal(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
}
