package usuario

import (
	TDAheap "tdas/cola_prioridad"
	TDAdic "tdas/diccionario"
	err "tp2/errores"
)

type usuarioImp struct {
	nombre   string
	logueado bool
	posts    TDAheap.ColaPrioridad[Post]
	posicion int
}

func CrearUsuario(nombre_usuario string, pos int) Usuario {
	usuario := &usuarioImp{
		nombre:   nombre_usuario,
		logueado: false,
		posicion: pos,
	}
	usuario.posts = TDAheap.CrearHeap[Post](usuario.cmp)
	return usuario
}

func (u usuarioImp) cmp(post1, post2 Post) int {
	afinidad1 := calcularDistancia(u, post1.ObtenerCreador())
	afinidad2 := calcularDistancia(u, post2.ObtenerCreador())

	if afinidad1 > afinidad2 {
		return 1
	} else if afinidad1 < afinidad2 {
		return -1
	}
	return 0
}

func calcularDistancia(u1 usuarioImp, u2 Usuario) int {
	pos1 := u1.ObtenerPos()
	pos2 := u2.ObtenerPos()

	dist := pos1 - pos2
	if dist < 0 {
		dist *= -1
	}
	return dist
}

func (u usuarioImp) ObtenerPos() int {
	return u.posicion
}

func (u usuarioImp) EstaLogeado() bool {
	return u.logueado
}

func (u usuarioImp) ObtenerNombre() string {
	return u.nombre
}

func (u *usuarioImp) Logearse(dicc TDAdic.Diccionario[string, int]) error {
	if !dicc.Pertenece(u.nombre) {
		return err.UsarioInexistente{}
	}
	u.logueado = true
	return nil
}

func (u *usuarioImp) VerFeed() (Post, error) {
	if u.posts.EstaVacia() || !u.EstaLogeado() {
		return nil, err.SinPostsParaVer{}
	}

	return u.posts.Desencolar(), nil
}

func (u *usuarioImp) encolarPost(post Post) {
	u.posts.Encolar(post)
}

func (u *usuarioImp) PublicarPost(mensaje string, usuarios_logueados []Usuario, arr []Post) error {
	if !u.EstaLogeado() {
		return err.NoLogueado{}
	}
	post := CrearPost(mensaje, u, arr)
	for _, usuario := range usuarios_logueados {
		if u.nombre != usuario.ObtenerNombre() {
			usuario.encolarPost(post)
		}
	}
	arr = append(arr, post)
	return nil
}

func (u *usuarioImp) Likear(id int, arr []Post) error {
	if u.posts.EstaVacia() || !u.EstaLogeado() {
		return err.SinPostsParaVer{}
	}
	post := arr[id]
	post.AsignarLikes(u.nombre)
	return nil
}
