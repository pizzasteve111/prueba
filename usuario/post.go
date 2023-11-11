package usuario

type Post interface {

	//Asigna un unico like por usuario al post.
	AsignarLikes(string)

	//Imprime el estado actual del post.
	ImprimirPost()

	//Devuelve al creador del Post
	ObtenerCreador() Usuario

	//Muestra los likes asignados al post.
	MostrarLikes()
}
