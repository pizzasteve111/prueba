package usuario

import (
	"fmt"
	"os"
	"strconv"
	TDAdic "tdas/diccionario"
)

type postImp struct {
	id      int
	likes   TDAdic.DiccionarioOrdenado[string, int]
	creador Usuario
	mensaje string
}

func cmp(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func CrearPost(texto string, user Usuario, arr []Post) Post {
	return &postImp{
		id:      len(arr),
		likes:   TDAdic.CrearABB[string, int](cmp),
		creador: user,
		mensaje: texto,
	}
}

func (p *postImp) AsignarLikes(nombre_usuario string) {
	if !p.likes.Pertenece(nombre_usuario) {
		p.likes.Guardar(nombre_usuario, 0)
	}
}

func (p postImp) ImprimirPost() {
	fmt.Fprintln(os.Stdout, "Post ID "+strconv.Itoa(p.id))
	fmt.Fprintln(os.Stdout, p.creador.ObtenerNombre()+" dijo: "+p.mensaje)
	fmt.Fprintln(os.Stdout, "Likes: "+strconv.Itoa(p.likes.Cantidad()))
}

func (p postImp) MostrarLikes() {
	cant_likes := strconv.Itoa(p.likes.Cantidad())
	fmt.Fprintln(os.Stdout, "El post tiene "+cant_likes+" likes:")

	for iter := p.likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		usuario, _ := iter.VerActual()
		fmt.Fprintln(os.Stdout, usuario)
	}

}

func (p postImp) ObtenerCreador() Usuario {
	return p.creador
}
