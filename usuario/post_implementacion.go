package usuario

import (
	"fmt"
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
	fmt.Println("Post ID " + strconv.Itoa(p.id))
	fmt.Println(p.creador.ObtenerNombre() + " dijo: " + p.mensaje)
	fmt.Println("Likes: " + strconv.Itoa(p.likes.Cantidad()))
}

func (p postImp) ObtenerCreador() Usuario {
	return p.creador
}
