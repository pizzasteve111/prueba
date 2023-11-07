package usuario

type Post interface {

	//Asigna un unico like por usuario al post.
	AsignarLikes(string)

	//Imprime el estado actual del post.
	ImprimirPost()

	ObtenerCreador() Usuario
}
