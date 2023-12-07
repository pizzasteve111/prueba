from grafotp import Grafo
from collections import deque
import biblioteca
import heapq

MAX_NAVEGACION = 20

def conectividad(grafo : Grafo, pagina):
    componentes = biblioteca.cfcs_grafo(grafo)
    
    for componente in componentes:
        if pagina in componente:
            return componente
            

def camino(grafo: Grafo,origen,destino):
    camino = biblioteca.camino_min_dfs(grafo,origen,destino)
    return camino, len(camino)-1

def diametro(grafo:Grafo):
    maximo = -1
    inicio = None
    destino = None
    padres = None

    for v in grafo.vertices:
        padre,orden = biblioteca.bfs(grafo,v)
        for w, dist in orden.items():
            if maximo<dist:
                maximo = dist
                destino = w
                padres = padre
     
    return biblioteca.reconstruir_camino(padres,inicio,destino)

def rango(grafo: Grafo,v,n):
    res = set()
    _,ordenes = biblioteca.bfs(grafo,v)
    for w,orden in ordenes.items():
        if orden == n:
            res.add(w)

    return len(res)

def navegacion(grafo : Grafo, origen):
    paginas = []
    paginas.append(origen)
    
    while len(grafo.adyacentes(origen))>0:
        origen = grafo.adyacentes(origen)[0]
        paginas.append(origen)
        if len(paginas) == MAX_NAVEGACION:
            break
        
    return paginas

def lectura(grafo:Grafo,links):
    cola = deque()
    g_sal = biblioteca.grados_salida(grafo)
    orden = []
    for v in links:
        if g_sal[v] == 0:
            cola.append(v)

    entradas = biblioteca.obtener_entaradas_vertice(grafo)

    while len(cola)>0:
        v = cola.popleft()
        orden.append(v)
        for w in entradas[v]:
            g_sal[w]-=1
            if g_sal[w]==0:
                cola.append(w)
    
    if len(orden) != len(links):
        return  None

    return orden

def clustering(grafo: Grafo, pagina = None):
    if pagina:
        adyacentes = grafo.adyacentes(pagina)
        if len(adyacentes) < 2:
            return 0
        
        aristas_ady = 0
    
        for w in adyacentes:
            for j in adyacentes:
                if grafo.estan_unidos(w,j):
                    aristas_ady +=1
                
        coef = float((aristas_ady)/(len(adyacentes)*(len(adyacentes)-1)))
        return round(coef,3)

    else:
        coef_total = 0
        for v in grafo.obtener_vertices():  
            adyacentes = grafo.adyacentes(v)
            if len(adyacentes) < 2:
                return 0
        
            aristas_ady = 0
    
            for w in adyacentes:
                for j in adyacentes:
                    if grafo.estan_unidos(w,j):
                        aristas_ady +=1
                
            coef = float((aristas_ady)/(len(adyacentes)*(len(adyacentes)-1)))
            coef_total += round(coef,3)
        return float(round((coef_total/len(grafo.obtener_vertices)),3))


def ciclos(grafo : Grafo, n , pagina):
    visitados = set()
    res = []
    return dfs_ciclos(grafo,visitados,pagina,pagina,res,n)
    

def dfs_ciclos(grafo: Grafo, visitados,v,pagina,res,n):
    visitados.add(v)
    res.append(v)

    if len(res)==n and res[-1]==pagina:
        return res
    
    for w in grafo.adyacentes(v):
        if w not in visitados:
            res = dfs_ciclos(grafo,visitados,w,pagina,res,n)
            if res:
                return res
            
    res.pop()
    visitados.remove(v)
    return None

def comunidades(grafo:Grafo,pagina):
    comunidades = biblioteca.label_propagation(grafo)
    for set_comunidad in comunidades.values():
        if pagina in set_comunidad:
            return set_comunidad
        

def mas_importantes(grafo: Grafo,n):
    page_ranks = biblioteca.pagerank(grafo)
    heap_min = []
    mas_imp = []
    i = 0

    for v in grafo.obtener_vertices():

        if i < n:
            heapq.heappush(heap_min,(page_ranks[v],v))
            i+=1
        else:
            if page_ranks[v]> heap_min[0][0]:
                heapq.heappop(heap_min)
                heapq.heappush(heap_min,(page_ranks[v],v))
    
    for i in range(n):
        mas_imp.append(heapq.heappop(heap_min))

    return mas_imp.reverse()


def mostrar_comandos():
    return 
    