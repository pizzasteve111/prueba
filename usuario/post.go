package usuario

type Post interface {

	//Asigna un unico like por usuario al post.
	AsignarLikes(string)

	//Imprime el estado actual del post.
	ImprimirPost()

	//Devuelve al creador del Post
	ObtenerCreador() Usuario

	//Muestra los likes asignados al post, devuelve error si el post no existe o no tiene likes.
	MostrarLikes() error
}
