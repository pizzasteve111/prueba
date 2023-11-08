package usuario

import (
	TDAdic "tdas/diccionario"
)

type Usuario interface {

	//Esta primitiva verifica si un usuario esta logueado o no, devolviendo true o false en respectivos casos
	EstaLogeado() bool

	//Esta primitiva loguea al usuario. En caso de que no este en el archivo de entrada, se devuelve el error pertinente
	Logearse(TDAdic.Diccionario[string, int]) error

	//Devuelve el nombre del usuario
	ObtenerNombre() string

	//Se encarga de avanzar en el feed, segun la afinidad de los usuarios que hayan publicando posts; en caso que dos posts
	//tengan sendos usuarios con misma afinidad al loggeado , se debe visualizar el post que primero se haya creado.
	//Si todo se ejecuta correctamente devuelve nil y el post, caso contrario error pertinente
	VerFeed() (Post, error)

	//El usuario crea un post que se asigna al feed del resto de los usuarios,
	//en caso de no estar logueado, se devuelve el error pertinente.
	PublicarPost(string, []Usuario, []Post) (Post, error)

	//Devuelve la posicion del usuario en el archivo de entrada.
	ObtenerPos() int

	Deslogearse() error
}
