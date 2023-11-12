package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAcola "tdas/cola"
	errores "tp2/errores"
	archivos "tp2/funciones"
	TDAuser "tp2/usuario"
)

func main() {
	parametros := os.Args[1:]

	if len(parametros) != 1 {
		err := errores.ErrorParametros{}
		err_string := err.Error()
		fmt.Fprintln(os.Stdout, err_string)
		return
	}

	fila := TDAcola.CrearColaEnlazada[TDAuser.Usuario]()
	arr_post := []TDAuser.Post{}
	ruta_archivo := parametros[0]

	arr_usuarios, dicc_usuarios, err1 := archivos.LeerArchivo(ruta_archivo)
	if err1 != nil {
		err1_string := err1.Error()
		fmt.Fprintln(os.Stdout, err1_string)
		return
	}
	entrada := bufio.NewScanner(os.Stdin)

	for entrada.Scan() {
		linea := entrada.Text()
		lista_argumentos := strings.Split(linea, " ")
		comando := lista_argumentos[0]

		switch comando {

		case "login":
			err := archivos.ComandoLogin(lista_argumentos, dicc_usuarios, arr_usuarios, fila)
			if err != nil {
				err_string := err.Error()
				fmt.Fprintln(os.Stdout, err_string)
				continue
			}
			usuario := (fila.VerPrimero()).ObtenerNombre()
			fmt.Fprintln(os.Stdout, "Hola "+usuario)

		case "logout":
			err := archivos.ComandoLogout(lista_argumentos, fila)
			if err != nil {
				err_string := err.Error()
				fmt.Fprintln(os.Stdout, err_string)
				continue
			}

			fmt.Fprintln(os.Stdout, "Adios")

		case "publicar":
			post, err := archivos.ComandoPublicar(lista_argumentos, fila, arr_usuarios, arr_post)
			if err != nil {
				err_string := err.Error()
				fmt.Fprintln(os.Stdout, err_string)
				continue
			}

			arr_post = append(arr_post, post)
			fmt.Fprintln(os.Stdout, "Post publicado")

		case "ver_siguiente_feed":
			post, err := archivos.ComandoVerFeed(lista_argumentos, fila)
			if err != nil {
				err_string := err.Error()
				fmt.Fprintln(os.Stdout, err_string)
				continue
			}
			post.ImprimirPost()

		case "likear_post":
			err := archivos.ComandoLikear(lista_argumentos, fila, arr_post)
			if err != nil {
				err_string := err.Error()
				fmt.Fprintln(os.Stdout, err_string)
				continue
			}
			fmt.Fprintln(os.Stdout, "Post likeado")
		case "mostrar_likes":
			post, err := archivos.ComandoVerLikes(lista_argumentos, arr_post)
			if err != nil {
				err_string := err.Error()
				fmt.Fprintln(os.Stdout, err_string)
				continue
			}
			err2 := post.MostrarLikes()
			if err2 != nil {
				err2_string := err2.Error()
				fmt.Fprintln(os.Stdout, err2_string)
				continue
			}

		}

	}

}
