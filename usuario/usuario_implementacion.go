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
}

func CrearUsuario(nombre_usuario string) Usuario {
	return &usuarioImp{
		nombre:   nombre_usuario,
		logueado: false,
		posts:    TDAheap.CrearHeap[Post](),
	}
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

func (u *usuarioImp) PublicarPost(mensaje string, usuarios_logueados []usuarioImp, arr []Post) error {
	if !u.EstaLogeado() {
		return err.NoLogueado{}
	}
	post := CrearPost(mensaje, u.nombre, arr)
	for _, usuario := range usuarios_logueados {
		if u.nombre != usuario.ObtenerNombre() {
			usuario.encolarPost(post)
		}
	}
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

//holis
