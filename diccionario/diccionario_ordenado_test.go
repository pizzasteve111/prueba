package diccionario_test

import (
	"fmt"
	"math/rand"
	TDADiccionario "tdas/diccionario"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDiccVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestDiccClaveDefault(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("") })

	abbNum := TDADiccionario.CrearABB[int, string](compararNumeros)
	require.False(t, abbNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](compararCadenas)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("B"))
	require.EqualValues(t, 10, abb.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("B") })
}

func TestDiccGuardar(t *testing.T) {
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	abb := TDADiccionario.CrearABB[string, string](compararCadenas)
	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))

	require.False(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[1], valores[1])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))

	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], abb.Obtener(claves[2]))
}

func TestReemplazarDato(t *testing.T) {
	clave := "Perro"
	clave2 := "Gato"
	abb := TDADiccionario.CrearABB[string, string](compararCadenas)
	abb.Guardar(clave, "guau")
	abb.Guardar(clave2, "miau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, "guau", abb.Obtener(clave))
	require.EqualValues(t, "miau", abb.Obtener(clave2))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(clave, "miu")
	abb.Guardar(clave2, "baubau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener(clave))
	require.EqualValues(t, "baubau", abb.Obtener(clave2))
}

func TestReemplazarDatoHopscotch(t *testing.T) {

	abb := TDADiccionario.CrearABB[int, int](compararNumeros)
	for i := 0; i < 500; i++ {
		abb.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		abb.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = abb.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccBorrar(t *testing.T) {
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](compararCadenas)

	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[0]) })
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[0]) })

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[1]) })
}

func TestReutilizacionDeBorrados(t *testing.T) {

	abb := TDADiccionario.CrearABB[string, string](compararCadenas)
	clave := "hola"
	abb.Guardar(clave, "mundo!")
	abb.Borrar(clave)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(clave))
	abb.Guardar(clave, "mundooo!")
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, "mundooo!", abb.Obtener(clave))
}

func TestClavesNumericas(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](compararNumeros)
	clave := 10
	valor := "Gatito"

	abb.Guardar(clave, valor)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, valor, abb.Obtener(clave))
	require.EqualValues(t, valor, abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestClaveVacia1(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, string](compararCadenas)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}

func TestValNulo(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, *int](compararCadenas)
	clave := "Pez"
	abb.Guardar(clave, nil)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, (*int)(nil), abb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func mezclar(arr1 []string, arr2 []int) {
	//mezcla aleatoriamente dos arreglos
	for i := len(arr1) - 1; i > 0; i-- {
		j := generarNumeroAleatorio(len(arr1) - 1)
		arr1[i], arr1[j] = arr1[j], arr1[i]
		arr2[i], arr2[j] = arr2[j], arr2[i]
	}
}

func ejecutarPruebaVolumen1(b *testing.B, n int) {
	dicc := TDADiccionario.CrearHash[string, int]()

	claves := make([]string, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
	}

	mezclar(claves, valores)

	for i := range claves {
		dicc.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dicc.Cantidad(), "La cantidad de elementos es incorrecta")

	ok := true
	for i := 0; i < n; i++ {
		ok = dicc.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dicc.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dicc.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dicc.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dicc.Cantidad())
}

func BenchmarkDicc(b *testing.B) {
	for _, n := range VOLUMENES {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen1(b, n)
			}
		})
	}
}

func TestIterarDiccVacio(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](compararCadenas)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccIterar(t *testing.T) {
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	valor1 := "guau"
	valor2 := "miau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](compararCadenas)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	iter := abb.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, Buscar(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, Buscar(segundo, claves))
	require.EqualValues(t, valores[Buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, Buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIterNoLlegaAlFinal(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, string](compararCadenas)
	claves := []string{"B", "A", "C"}
	abb.Guardar(claves[0], "")
	abb.Guardar(claves[1], "")
	abb.Guardar(claves[2], "")

	abb.Iterador()
	iter2 := abb.Iterador()
	iter2.Siguiente()
	iter3 := abb.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, Buscar(primero, claves))
	require.NotEqualValues(t, -1, Buscar(segundo, claves))
	require.NotEqualValues(t, -1, Buscar(tercero, claves))
}

func TestIterInternoClaves(t *testing.T) {
	clave1 := "Perro"
	clave2 := "Gato"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	abb := TDADiccionario.CrearABB[string, *int](compararCadenas)
	abb.Guardar(claves[0], nil)
	abb.Guardar(claves[1], nil)
	abb.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	abb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, Buscar(cs[0], claves))
	require.NotEqualValues(t, -1, Buscar(cs[1], claves))
	require.NotEqualValues(t, -1, Buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIterInternoValores(t *testing.T) {
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](compararCadenas)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func generarNumeroAleatorio(N int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(N + 1)
}

func TestVolumenIterCorte(t *testing.T) {

	abb := TDADiccionario.CrearABB[int, int](compararNumeros)

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < 10000; i++ {
		numero := generarNumeroAleatorio(10000)
		abb.Guardar(numero, numero)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	abb.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIteradorRangoDefinido(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](compararCadenas)
	abb.Guardar("D", 15)
	abb.Guardar("B", 4)
	abb.Guardar("S", 90)
	abb.Guardar("A", 4)
	abb.Guardar("C", 5000)
	desde := "A"
	hasta := "D"
	iter := abb.IteradorRango(&desde, &hasta)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, "A", primero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	segundo, _ := iter.VerActual()
	require.EqualValues(t, "B", segundo)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, "C", tercero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, "D", cuarto)

}

func TestIteradorRangoSinDesde(t *testing.T) {
	abb := TDADiccionario.CrearABB[string, int](compararCadenas)
	abb.Guardar("D", 15)
	abb.Guardar("B", 4)
	abb.Guardar("S", 90)
	abb.Guardar("A", 4)
	abb.Guardar("C", 5000)
	hasta := "D"
	iter := abb.IteradorRango(nil, &hasta)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, "A", primero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	segundo, _ := iter.VerActual()
	require.EqualValues(t, "B", segundo)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, "C", tercero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, "D", cuarto)

	iter.Siguiente()

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

}

func TestIteradorRangoSinHasta(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararNumeros)
	abb.Guardar(7, 7)
	abb.Guardar(3, 3)
	abb.Guardar(8, 8)
	abb.Guardar(2, 2)
	abb.Guardar(5, 5)
	desde := 3
	iter := abb.IteradorRango(&desde, nil)

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.EqualValues(t, 3, primero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	segundo, _ := iter.VerActual()
	require.EqualValues(t, 5, segundo)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.EqualValues(t, 7, tercero)

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	cuarto, _ := iter.VerActual()
	require.EqualValues(t, 8, cuarto)

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

}

func TestIteradorInternoRangos(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararNumeros)
	abb.Guardar(7, 7)
	abb.Guardar(3, 3)
	abb.Guardar(8, 8)
	abb.Guardar(2, 2)
	abb.Guardar(5, 5)
	desde := 3
	hasta := 7
	factorial := 1
	ptrFactorial := &factorial
	abb.IterarRango(&desde, &hasta, func(_ int, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 105, factorial)
}

func TestIteradorInternoRangosSinHasta(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararNumeros)
	abb.Guardar(7, 7)
	abb.Guardar(3, 3)
	abb.Guardar(8, 8)
	abb.Guardar(2, 2)
	abb.Guardar(5, 5)
	desde := 5
	factorial := 1
	ptrFactorial := &factorial
	abb.IterarRango(&desde, nil, func(_ int, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 280, factorial)
}

func TestPruebaABBborrarConDosHijos(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](compararNumeros)
	abb.Guardar(7, 7)
	abb.Guardar(3, 3)
	abb.Guardar(8, 8)

	elem := abb.Borrar(7)
	require.Equal(t, 7, elem)

	abb.Guardar(elem, elem)
	iter := abb.Iterador()
	iter.Siguiente()
	guardado1, guardado2 := iter.VerActual()
	require.Equal(t, guardado1, elem)
	require.Equal(t, guardado2, elem)

}

var VOLUMENES = []int{12500, 25000, 50000, 100000, 200000, 400000}

func compararCadenas(cadena1, cadena2 string) int {
	if cadena1 > cadena2 {
		return 1
	} else if cadena1 < cadena2 {
		return -1
	} else {
		return 0
	}
}

func compararNumeros(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func Buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}
