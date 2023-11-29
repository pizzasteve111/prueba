class Grafo:
    def __init__(self, drigido = False, vertices = []):
        self.dirigido = drigido
        self.vertices = {}
        if len(vertices) != 0:
            for vertice in vertices:
                self.vertices[vertice]= {}

    def agregar_vertice(self,v):
        if v in self.vertices:
            raise ValueError(f"ya existe este vertice en el grafo.")
        
        self.vertices[v] = {}

    def eliminar_vertice(self,v):
        if v not in self.vertices:
            raise ValueError(f"No esta el vertice en el grafo.")
        
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
        if v not in self.vertices or w not in self.vertices:
            raise ValueError(f"vertice/s no esta en el grafo")

        if w in self.vertices[v]:
            raise ValueError(f"ya esta la arista en el grafo")
        
        self.vertices[v][w] = p

        if not self.dirigido:
            self.vertices[w][v] = p
    
    def eliminar_arista(self,v,w,p = 1):
        if v not in self.vertices or w not in self.vertices:
            raise ValueError(f"vertice/s no esta en el grafo")
        
        if w not in self.obtener_vertices[v]:
            raise ValueError(f"Arista inexistente en el grafo")
        
        #revisar esta parte:
        self.vertices[v].pop(w)

        if not self.dirigido:
            self.vertices[w].pop(v)

    def estan_unidos(self,v,w):
        if v not in self.vertices or w not in self.vertices:
            return False
        
        if w not in self.obtener_vertices[v]:
            return False
        
        return True

    def peso_arista(self,v,w):
        if v not in self.vertices or w not in self.vertices:
            raise ValueError(f"vertice/s no esta en el grafo")
        
        if w not in self.obtener_vertices[v]:
            raise ValueError(f"Arista inexistente en el grafo")
        
        return self.vertices[v][w]
    
    def vertice_aleatorio(self):
        return self.vertices[0]
    
    def adyacentes(self,v):
        res = []

        for w in self.vertices[v].keys():
            res.append(w)
        return res