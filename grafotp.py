class Grafo:
    def __init__(self, drigido = False, vertices = []):
        self.dirigido = drigido
        self.vertices = {}
        if len(vertices) != 0:
            for vertice in vertices:
                self.vertices[vertice]= {}

    def agregar_vertice(self,v):
        if v not in self.vertices:
            self.vertices[v] = {}

    def eliminar_vertice(self,v):        
        self.vertices.pop(v)
        
        for adyacentes in self.vertices.values():
            for vertice in adyacentes.keys():
                if vertice==v:
                    adyacentes.pop(v) 
    
    def obtener_vertices(self):
        res = []
        for v in self.vertices.keys():
            res.append(v)
        return res
    
    def agregar_arista(self,v,w,p = 1):
      
        self.vertices[v][w] = p

        if not self.dirigido:
            self.vertices[w][v] = p
    
    def eliminar_arista(self,v,w,p = 1):
        if self.estan_unidos(v,w):
            self.vertices[v].pop(w)
            if not self.dirigido:
                self.vertices[w].pop(v)

    def estan_unidos(self,v,w):
        if v not in self.vertices or w not in self.vertices:
            return False
        
        if w not in self.vertices[v]:
            return False
        
        return True

    def peso_arista(self,v,w):
      
        return self.vertices[v][w]
    
    def vertice_aleatorio(self):
        return self.vertices[0]
    
    def adyacentes(self,v):
        res = []
        if v in self.vertices:
            for w in self.vertices[v].keys():
                res.append(w)
            return res
    
    def aristas(self):
        visitados = set()
        aristas = []
        for v in self.vertices.keys():
            if v not in visitados:
                visitados.add(v)
            for w in  self.vertices.values():
                if w not in visitados:
                    aristas.append((v,w,self.peso_arista(v,w)))
                visitados.add(w)
        return aristas