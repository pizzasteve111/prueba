package funciones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAcola "tdas/cola"
	TDAdic "tdas/diccionario"
	errores "tp2/errores"
	TDAuser "tp2/usuario"
)

func LeerArchivo(ruta string) ([]TDAuser.Usuario, TDAdic.Diccionario[string, int], error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, nil, errores.ErrorLeerArchivo{}
	}
	arr := []TDAuser.Usuario{}

	lector := bufio.NewScanner(archivo)
	dicc := TDAdic.CrearHash[string, int]()
	cont := 0
	for lector.Scan() {
		linea := lector.Text()
		usuario := TDAuser.CrearUsuario(linea, cont)
		dicc.Guardar(linea, cont)
		arr = append(arr, usuario)
		cont++
	}

	fmt.Fprintln(os.Stdout, dicc.Pertenece("alan"))

	return arr, dicc, nil
}

func ComandoLogin(entrada []string, dicc_usuarios TDAdic.Diccionario[string, int], arr []TDAuser.Usuario, fila TDAcola.Cola[TDAuser.Usuario]) error {
	if len(entrada) != 2 {
		return errores.ErrorParametros{}
	}
	nombre_usuario := entrada[1]
	if !dicc_usuarios.Pertenece(nombre_usuario) {
		return errores.UsarioInexistente{}
	}
	if !fila.EstaVacia() {
		return errores.YaLogueado{}
	}
	pos_usuario := dicc_usuarios.Obtener(nombre_usuario)
	usuario := arr[pos_usuario]

	err := usuario.Logearse(dicc_usuarios)

	if err != nil {
		return err
	}

	fila.Encolar(usuario)

	return nil
}

func ComandoLogout(entrada []string, fila TDAcola.Cola[TDAuser.Usuario]) error {
	if len(entrada) != 1 {
		return errores.ErrorParametros{}
	}
	if fila.EstaVacia() {
		return errores.NoLogueado{}
	}
	usuario := fila.Desencolar()
	err := usuario.Deslogearse()

	if err != nil {
		return err
	}

	return nil
}

func ComandoPublicar(entrada []string, fila TDAcola.Cola[TDAuser.Usuario], arr_user []TDAuser.Usuario, arr_post []TDAuser.Post) (TDAuser.Post, error) {
	if len(entrada) == 1 {
		return nil, errores.ErrorParametros{}
	}
	mensaje := strings.Join(entrada[1:], " ")

	if fila.EstaVacia() {
		return nil, errores.NoLogueado{}
	}
	usuario := fila.VerPrimero()
	post, err := usuario.PublicarPost(mensaje, arr_user, arr_post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func ComandoLikear(entrada []string, fila TDAcola.Cola[TDAuser.Usuario], arr_post []TDAuser.Post) error {
	if len(entrada) != 2 {
		return errores.ErrorParametros{}
	}
	if fila.EstaVacia() {
		return errores.NoLogueado{}
	}
	usuario := fila.VerPrimero()
	id, err := strconv.Atoi(entrada[1])
	if err != nil {
		return errores.ErrorParametros{}
	}

	err2 := usuario.Likear(id, arr_post)

	if err2 != nil {
		return err2
	}
	return nil
}

func ComandoVerFeed(entrada []string, fila TDAcola.Cola[TDAuser.Usuario]) (TDAuser.Post, error) {
	if len(entrada) != 1 || entrada[0] != "ver_siguiente_feed" {
		return nil, errores.ErrorParametros{}
	}

	if fila.EstaVacia() {
		return nil, errores.NoHayMasPost{}
	}

	usuario := fila.VerPrimero()
	post, err := usuario.VerFeed()

	if err != nil {
		return nil, err
	}

	return post, nil
}

func ComandoVerLikes(entrada []string, arr_posts []TDAuser.Post) (TDAuser.Post, error) {
	if len(entrada) != 2 || entrada[0] != "mostrar_likes" {
		return nil, errores.ErrorParametros{}
	}

	id, err := strconv.Atoi(entrada[1])
	if err != nil {
		return nil, errores.ErrorParametros{}
	}

	if id < 0 || len(arr_posts) <= id {
		return nil, errores.PostInexistente{}
	}

	post := arr_posts[id]

	return post, nil

}
