package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

type YaLogueado struct{}

func (e YaLogueado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type UsarioInexistente struct{}

func (e UsarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type NoLogueado struct{}

func (e NoLogueado) Error() string {
	return "Error: no habia usuario loggeado"
}

type SinPostsParaVer struct {
}

func (e SinPostsParaVer) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type PostInexistente struct{}

func (e PostInexistente) Error() string {
	return "Error: Post inexistente o sin likes"
}

type NoHayMasPost struct{}

func (e NoHayMasPost) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}
