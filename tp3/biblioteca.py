from grafotp import Grafo
from queue import Queue
import heapq
from collections import deque
import random
DAMPING_FACTOR = 0.85


def grados_entrada(grafo:Grafo):
    grados = {}
    for v in grafo.obtener_vertices():
        grados[v] = 0
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados

def grados_salida(grafo:Grafo):
    grados = {}
    for v in grafo.obtener_vertices():
        grados[v] = 0
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            grados[w] = grados[v] + 1
    return grados


def reconstruir_camino(padres, origen,destino):
    camino = []
    actual = destino
    while actual != origen:
        camino.append(actual)
        actual = padres[actual]
    camino.append(origen)
    return camino[::-1] 


def bfs(grafo: Grafo,origen):
    visitados = set()
    cola = Queue()
    padres = {origen : None}
    orden = {origen : 0}
    cola.put(origen)
    visitados.add(origen)
    while not cola.empty():
        v = cola.get()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                cola.put(w)        
    return padres, orden



def colorear_bipartito_aux( grafo : Grafo):
    colores = {}
    for v in grafo.obtener_vertices():
        if v not in colores:
            if not colorear_bipartito(grafo,v,colores):
                return False
    return True

def colorear_bipartito(grafo: Grafo, vertice_inicial,colores):
    q = Queue()
    q.put(vertice_inicial)
    colores[vertice_inicial]=0
    while not q.empty():
        v = q.get()
        for w in grafo.adyacentes(v):
            if w in colores:
                if colores[v]==colores[w]:
                    return False
            else:
                colores[w]=  1 - colores[v]
                q.put(w)
    return True

def orden_topologico(grafo:Grafo):
    g_ent = grados_entrada(grafo)
    cola = Queue()
    for v in grafo:
        if g_ent[v] == 0:
            cola.put(v)
    orden = []
    while not cola.empty():
        v = cola.get()
        orden.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                cola.put(w)
    return orden    


def dijkstra(grafo : Grafo,origen,destino):
    dist = {}
    padre= {origen:None}
    for v in grafo.vertices():
        dist[v] = float("inf")
    dist[origen]=0
    q=heapq
    q.heappush((0,origen))
    
    while not len(q)==0:
        _,v = q.heappop
        if v == destino:
            return padre, dist
        for w in grafo.adyacentes(v):
            n_dist= dist[v] + grafo.peso_arista(v,w)
            if dist[w] > n_dist:
                dist[w] = n_dist
                padre[w]= q.heappush((dist[w],w))
    return padre,dist

def camino_min_dfs(grafo: Grafo, origen, destino):
    visitados =set(origen)
    padres = {origen:None}
    padres = dfs_aux(grafo,visitados,padres,destino)
    camino = reconstruir_camino(padres)

    return camino


def dfs_aux(grafo : Grafo, v, visitados, padres,destino):
    if v == destino:
        return padres
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            dfs_aux(grafo,w,visitados,padres,destino)


def cfcs_grafo(grafo:Grafo):
    resultados = []
    visitados = set()
    for v in grafo:
        if v not in visitados:
            dfc_cfcs(grafo,v,visitados,{},{},deque(),set(),resultados,[0])

    return resultados


def dfc_cfcs(grafo : Grafo,v, visitados,orden,mas_bajo,pila:deque,apilados,cfcs,contador):
    orden[v]= mas_bajo[v] = contador[v]
    contador+=1
    visitados.add(v)
    apilados.add(v)
    
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfc_cfcs(grafo,w,visitados,orden,mas_bajo,pila,apilados,cfcs,contador)
        if w in apilados:
            mas_bajo[w] = min(mas_bajo[w],mas_bajo[v])
            
    if orden[v]==mas_bajo[v]:
        nueva_cfc= []
        while True:
            w = pila.pop()            
            apilados.remove(w)
            if w == v:
                break
        cfcs.append(nueva_cfc)



def max_freq(entradas,labels):
    frecuencias = {}
    for v in entradas:
        etiqueta = labels[v]
        frecuencias[etiqueta] = frecuencias.get(etiqueta,0) +1

    return max(frecuencias, key=frecuencias.get)

def label_propagation(grafo: Grafo):
    Label = {}
    comunidades = {}
    entradas = obtener_entaradas_vertice(grafo)
    i = 0
    
    for v in grafo.vertices:
        comunidades[i] = set()
        comunidades[i].add(v)
        Label[v] = i
        i+=1

    for _ in range(i):
        vertices = grafo.obtener_vertices()
        random.shuffle(vertices)
        for v in vertices:
            entradasV = entradas[v]
            frecuencias_max = 0
            if len(entradasV) == 0:
                frecuencias_max = Label[v]
            else:
                frecuencias_max = max_freq(entradasV, Label) 

            
            etiqueta = Label[v]
            comunidades[etiqueta].remove(v)
            
            Label[v] = frecuencias_max
    
            comunidades[frecuencias_max].add(v)

    return comunidades

def pagerank(grafo : Grafo,tol=1.0e-6):
    N = len(grafo.obtener_vertices())
    page_rank = {}
    g_sal = grados_salida(grafo)
    aux = {}
    sumas = {}

    for v in grafo.obtener_vertices():
        page_rank[v] = 1/N
        sumas[v]=0

    for v in grafo.obtener_vertices():
        pr_v = page_rank[v]
        L = g_sal[v]
        if L !=0:
            aux[v]= pr_v/L
        else:
            aux[v] = 0
     
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            sumas[w] += aux[v]
    
    for v in grafo.obtener_vertices():
        pr_v =  (1-DAMPING_FACTOR)/N + (DAMPING_FACTOR * sumas[v])
        if abs(pr_v - page_rank[v]) > tol:
            page_rank[v] = pr_v
       
    return page_rank
    
def obtener_entaradas_vertice(grafo:Grafo):
    entradas = {}
    for v in grafo.obtener_vertices():
        entradas[v] = set()
    for v in grafo.obtener_vertices():   
        for w in grafo.adyacentes(v):
            entradas[w].add(v)
            

    return entradas
