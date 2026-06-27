package rtree
type Rtree struct{
	Raiz *Nodo
}
func NuevoRTree() *RTree{
	return &RTree{
		Raiz: &Nodo{
			Puntos: []Punto{}, Hijos: []*Nodo{}, EsHoja: true,
		},
	}
}
func (t *RTree ) Insertar(p Punto){
	t.Raiz.Punto = append(
		t.Raiz.Puntos,
		p,
	)
}
func (t *RTree) BuscarRango (r Rectangulo) []Punto{
	var resultado []Punto
	for _,p := range t.Raiz.Puntos{
		if r.Contiene(p){
			resultado = append(resultado,p)
		}
	}
	return resultado
}