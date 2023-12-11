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

def grados_salida(grafo, lecturas = None): 
    grados = {}
    if lecturas:
        for v in lecturas:
            grados[v] = 0
            for w in grafo.adyacentes(v):
                if w in lecturas:
                    grados[v]+=1
    else:
        for v in grafo.obtener_vertices():
            grados[v] = 0
            for w in grafo.adyacentes(v):
                grados[v]+=1
    return grados


def reconstruir_camino(padres, origen,destino):
    camino = []
    actual = destino
    while actual != origen:
        camino.append(actual)
        actual = padres[actual]
    camino.append(origen)
    return camino[::-1] 


def bfs(grafo: Grafo,origen,destino = None):
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
                if destino and w == destino:
                    return padres,orden      
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


def camino_min(grafo:Grafo,origen,destino):
    padres,orden= bfs(grafo,origen,destino)
    if not destino in orden:
        return None
    
    return reconstruir_camino(padres, origen, destino)


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
    entradas = obtener_entradas_vertice(grafo)
    i = 0
    
    for v in grafo.vertices:
        comunidades[i] = set()
        comunidades[i].add(v)
        Label[v] = i
        i+=1

    for _ in range(1000):
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

def pagerank(grafo : Grafo,N,page_rank,g_sal,tol=1.0e-6):
    
    aux = {}
    sumas = {}
    for v in grafo.obtener_vertices():
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
    corte = False
    for v in grafo.obtener_vertices():

        pr_v =  (1-DAMPING_FACTOR)/N + (DAMPING_FACTOR * sumas[v])
        if abs(pr_v - page_rank[v]) > tol:
            page_rank[v] = pr_v
            corte = True
        
    return corte
   
    
def obtener_entradas_vertice(grafo:Grafo):
    entradas = {}
    for v in grafo.obtener_vertices():
        entradas[v] = set()
    for v in grafo.obtener_vertices():   
        for w in grafo.adyacentes(v):
            entradas[w].add(v)
            
    return entradas

def obtener_entradas_vertice_lect(grafo:Grafo,lecturas):
    entradas = {}
    for v in lecturas:
        entradas[v] = set()
    for v in lecturas:   
        for w in grafo.adyacentes(v):
            if w in lecturas:
                entradas[w].add(v)
            
    return entradas

def dfs_ciclos(grafo: Grafo, visitados,v,pagina,res,n):

    if len(res)==n:
        if v ==pagina:
            res.append(v)
            return res
        return None
    
    if v == pagina and len(res) != 0:
        return None
    
    visitados.add(v)
    res.append(v)
    
    for w in grafo.adyacentes(v):
        if w not in visitados or w == pagina:
            resp = dfs_ciclos(grafo,visitados,w,pagina,res,n)
            if resp:
                return resp
            
    res.pop()
    visitados.remove(v)
    return None
