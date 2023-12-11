import sys 
from grafotp import Grafo
from collections import deque
import biblioteca
import heapq

MAX_NAVEGACION = 20

def conectividad(grafo : Grafo, pagina):
    sys.setrecursionlimit(50000)
    componentes = biblioteca.cfcs_grafo(grafo,pagina)
    
    for componente in componentes:
        if pagina in componente:
            return componente
    return []        

def camino(grafo: Grafo,origen,destino):
    camino = biblioteca.camino_min(grafo,origen,destino)
    if not camino:
        return None
    return camino

def diametro(grafo:Grafo):
    maximo = -1
    inicio = None
    destino = None
    padres = None

    for v in grafo.vertices:
        padre,orden = biblioteca.bfs(grafo,v)
        for w, dist in orden.items():
            if maximo<dist:
                inicio = v
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
    set_links=set()
    for link in links:
        set_links.add(link)

    cola = deque()
    g_sal = biblioteca.grados_salida(grafo,set_links)
    orden = []
    for v in links:
        if g_sal[v] == 0:
            cola.append(v)

    entradas = biblioteca.obtener_entradas_vertice_lect(grafo,set_links)

    while len(cola) > 0:
        v = cola.popleft()
        orden.append(v)
        for w in entradas[v]:
            if w in set_links:
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
            return format(0,'.3f')
        
        aristas_ady = 0
    
        for w in adyacentes:
            for j in adyacentes:
                if j != w  and w != pagina and grafo.estan_unidos(w,j):
                    aristas_ady +=1
                
        coef = (aristas_ady)/(len(adyacentes)*(len(adyacentes)-1))
        return (format(coef,'.3f'))

    else:
        coef_total = 0
        for v in grafo.obtener_vertices():  
            adyacentes = grafo.adyacentes(v)
            if len(adyacentes) < 2:
                continue
        
            aristas_ady = 0
    
            for w in adyacentes:
                for j in adyacentes:
                    if j != w  and w != v and grafo.estan_unidos(w,j):
                        aristas_ady +=1
                
            coef = (aristas_ady)/(len(adyacentes)*(len(adyacentes)-1))
            coef_total += coef
        return (format((coef_total/len(grafo.obtener_vertices())),'.3f'))


def ciclos(grafo : Grafo, n , pagina):
    sys.setrecursionlimit(75000)
    visitados = set()
    res = []
    return biblioteca.dfs_ciclos(grafo,visitados,pagina,pagina,res,n)
    

def comunidades(grafo:Grafo,pagina):
    comunidades = biblioteca.label_propagation(grafo)
    for set_comunidad in comunidades.values():
        if pagina in set_comunidad:
            return set_comunidad
        

def mas_importantes(grafo: Grafo,n):
    N = len(grafo.obtener_vertices())
    page_rank = {}

    for v in grafo.obtener_vertices():
        page_rank[v]=1/N

    g_sal = biblioteca.grados_salida(grafo)
    corte = True
    while corte:
        corte = biblioteca.pagerank(grafo,N,page_rank,g_sal)

    heap_min = []
    mas_imp = []
    i = 0
    for v in grafo.obtener_vertices():

        if i < n:
            heapq.heappush(heap_min,(page_rank[v],v))
            i+=1
        else:
            if page_rank[v]> heap_min[0][0]:
                heapq.heappop(heap_min)
                heapq.heappush(heap_min,(page_rank[v],v))
    
    for i in range(n):
        mas_imp.append(heapq.heappop(heap_min)[1])

    return mas_imp[::-1]


