package pila_test

import (
	"fmt"
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	// mas pruebas para este caso...
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})

}
func Test_OrdenPila(t *testing.T) {
	fmt.Println("PRIMER TEST")
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("primero")
	pila.Apilar("segundo")
	pila.Apilar("tercero")
	pila.Apilar("ultimo")
	//el tope de la pila debe ser "ultimo"
	require.Equal(t, "ultimo", pila.VerTope())
	//la pila NO esta vacia
	require.Equal(t, false, pila.EstaVacia())
	var verificacion []string
	//comparo que el orden de la pila desapilada sea el inverso de la misma
	for pila.EstaVacia() != true {

		verificacion = append(verificacion, pila.Desapilar())
	}
	require.Equal(t, []string{"ultimo", "tercero", "segundo", "primero"}, verificacion)
	//luego del ciclo, la pila deberia estar vacia
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})

	require.Equal(t, true, pila.EstaVacia())
}
func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	elementos := 10000
	for i := 0; i < elementos; i++ {
		dato := i
		pila.Apilar(dato)

	}
	require.Equal(t, false, pila.EstaVacia())

	for f := elementos - 1; f >= 0; f-- {
		//verifico que al desapilar los elementos inversos se correspondan
		if pila.EstaVacia() != true {
			dato2 := f
			desapilado := pila.Desapilar()
			require.Equal(t, dato2, desapilado)
		}

	}
	require.Equal(t, true, pila.EstaVacia())

}
func Test_Apilar_Desapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	elementos := 100
	for i := 0; i < elementos; i++ {
		dato := i
		pila.Apilar(dato)
		require.Equal(t, false, pila.EstaVacia())
		pila.Desapilar()
		require.Equal(t, true, pila.EstaVacia())

	}

}
func Test_PilaVaciada(t *testing.T) {
	//creo una pila y le agrego elementos
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	//ahora la vacio y verifico que cumpla las condiciones de una pila vacia.
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})
}
func Test_PilaNueva(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	//tiene que cumplir las mismas condiciones que una pila vacia
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})
}

func Test_PilaVariada(t *testing.T) {
	//creo una pila a la que le puedo apilar cualquier tipo de dato, desde ints,string,vectores o mismas pilas
	pila := TDAPila.CrearPilaDinamica[any]()
	chiquita := TDAPila.CrearPilaDinamica[any]()
	vector := []int{1, 2, 3, 4}
	pila.Apilar(1)
	pila.Apilar("segundo")
	pila.Apilar(vector)
	pila.Apilar(chiquita)
	//comp 1
	require.Equal(t, false, pila.EstaVacia())
	dato_pil := pila.Desapilar()
	//comp 2
	require.Equal(t, dato_pil, chiquita)
	dato_vtr := pila.Desapilar()
	//comp3
	require.Equal(t, vector, dato_vtr)
	dato_str := pila.Desapilar()
	//comp 4
	require.Equal(t, dato_str, "segundo")
	//comp 5
	dato_int := pila.Desapilar()
	require.Equal(t, dato_int, 1)
	//comp 6
	require.Equal(t, true, pila.EstaVacia())
	//comp 7
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.VerTope()
	})
	//comp 8
	require.PanicsWithValue(t, "La pila esta vacia", func() {
		pila.Desapilar()
	})

}
