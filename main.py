#!/usr/bin/python3

import sys
from netstats import *

def main(grafo):
    
    for linea in sys.stdin:
        comando_parametros = linea.rstrip()
        comando = comando_parametros.split(" ")[0]
        parametros = " ".join(comando_parametros.split(" ")[1:]).split(",")

        if comando == "camino":
            origen,destino = parametros[0],parametros[1]
            road,distancia = camino(grafo,origen,destino)
            if distancia == 0:
                print("No se encontro recorrido")
            else:
                print(" -> ".join(road))
                print("Costo: ",distancia)
        
        elif comando == "mas_importantes":
            n = int(parametros[0])
            lista_nimp = mas_importantes(grafo,n)

            print(", ".join(lista_nimp))

        elif comando == "conectados":
            pagina = parametros[0]
            set_conect = conectividad(grafo,pagina)
            print(", ".join(set_conect))

        elif comando == "ciclo":
            pagina, n = parametros[0], parametros[1]
            ciclo = ciclos(grafo,n,pagina)
            if ciclo :
                print(" -> ".join(ciclo))  
            else:
                print("no se encontro ningun ciclo")

        elif comando == "lectura":
            lista_lec = lectura(grafo,parametros)

            if not lista_lec:
                print("No existe forma de leer las paginas en orden")
            else:
                print(", ".join(lista_lec))

        elif comando == "diametro":
            diametr = diametro(grafo)
            print(" -> ".join(diametr))
            print("Costo: ",len(diametr)-1)

        elif comando == "rango":
            pagina, n = parametros[0],parametros[1]
            print(rango(grafo,pagina,int(n)))
        
        elif comando == "comunidad":
            pagina = parametros[0]
            set_com = comunidades(grafo,pagina)
            print(", ".join(set_com))

        elif comando == "navegacion":
            origen = parametros[0]
            paginas = navegacion(grafo,origen)
            print(" -> ".join(paginas))

        elif comando == "clustering":
            dato = None
            if len(parametros)>0:
                pagina = parametros[0]
                dato = clustering(grafo,pagina)
            else:
                dato = clustering()
            print(dato)

       


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Cantidad de agrumentos inválida")
    
    else:
        nuevo = Grafo(drigido=True)
        ruta = sys.argv[1]

        with open(ruta) as arcihvo:
            for linea in arcihvo:
                linea = linea.rstrip().split("\t")
                pagina = linea[0]
                links = linea[1:]

                nuevo.agregar_vertice(pagina)
                for link in links:
                    nuevo.agregar_arista(pagina,link)
    main(nuevo)           
