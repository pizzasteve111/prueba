package funciones

import (
	"bufio"
	"os"
	TDAdic "tdas/diccionario"
	errores "tp2/errores"
)

func LeerArchivo(ruta string) (TDAdic.Diccionario[string, int], error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, errores.ErrorLeerArchivo{}
	}

	lector := bufio.NewScanner(archivo)
	dicc := TDAdic.CrearHash[string, int]()
	cont := 1

	for lector.Scan() {
		linea := lector.Text()
		dicc.Guardar(linea, cont)
		cont++
	}
	return dicc, nil
}
