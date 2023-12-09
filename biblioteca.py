from grafotp import Grafo
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
    cola = deque()
    padres = {origen : None}
    orden = {origen : 0}
    cola.append(origen)
    visitados.add(origen)
    while not len(cola) == 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                cola.append(w)        
    return padres, orden

def orden_topologico(grafo:Grafo):
    g_ent = grados_entrada(grafo)
    cola = deque()
    for v in grafo:
        if g_ent[v] == 0:
            cola.append(v)
    orden = []
    while not cola.empty():
        v = cola.popleft()
        orden.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                cola.append(w)
    return orden    

def camino_min_dfs(grafo: Grafo, origen, destino):
    visitados =set(origen)
    padres = {origen:None}
    padres = dfs_aux(grafo,origen,destino,visitados,padres)
    camino = reconstruir_camino(padres)

    return camino


def dfs_aux(grafo : Grafo, origen,destino, visitados, padres):
    if origen == destino:
        return padres
    for w in grafo.adyacentes(origen):
        if w not in visitados:
            visitados.add(w)
            dfs_aux(grafo,w,destino,visitados,padres)


def cfcs_grafo(grafo:Grafo,v):
    resultados = []
    visitados = set()
    dfc_cfcs(grafo,v,visitados,{},{},deque(),set(),resultados,[0])
    return resultados


def dfc_cfcs(grafo : Grafo,v, visitados,orden,mas_bajo,pila:deque,apilados,cfcs,contador):
    orden[v] = contador[0]
    mas_bajo[v] = orden[v]  
    contador[0]+=1

    visitados.add(v)
    apilados.add(v)
    pila.appendleft(v)
    
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfc_cfcs(grafo,w,visitados,orden,mas_bajo,pila,apilados,cfcs,contador)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v],mas_bajo[w])
            
    if orden[v]==mas_bajo[v]:
        nueva_cfc= []
        while pila:
            w = pila.popleft()            
            apilados.remove(w)
            nueva_cfc.append(w)
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
