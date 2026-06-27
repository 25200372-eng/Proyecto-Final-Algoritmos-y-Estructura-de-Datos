Package rtree
type Rectangulo struct{
	MinimoX float64
	MinimoY float64

	MaximoX float64
	MaximoY float64
}
func (r Rectangulo) Contiene(p Punto) bool{
	return p.Latitud >= r.MinimoX &&
			p.Latitud <= r.MaximoX &&
			p.longitud >= r.MinimoX &&
			p.longitud <= r.MaximoY
}