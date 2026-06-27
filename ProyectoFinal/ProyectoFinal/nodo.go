package rtree 
type Nodo struct{
	Puntos []Punto
	Hijos []*Nodo
	EsHoja bool
}