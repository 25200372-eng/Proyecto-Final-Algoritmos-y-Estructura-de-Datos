package main
import ("fmt" "RTreeProyecto/rtree")
func main(){
	arbol := rtree.NuevoRTree()
	arbol.Insertar(
		rtree.Punto{
			Nombre: "Lima",
			Latitud: -12.0464,
			longitud: -77.0428,
		},
	)
	arbol.Insertar(
		rtree.Punto{
			Nombre: "Cusco",
			Latitud: -13.5319,
			longitud: -71.9675,
		},
	)
	zona := rtree.Rectangulo{
		MinimoX: -15,
		MaximoX: -10,
		MinimoY: -80,
		MaximoY: -72
	}
	resultado := arbol.BuscarRango(zona)
	fmt.Println("Ciudades encontradas:")
	for _, Ciudades := range resultado{
		fmt.Println(Ciudades.Nombre)
	}
}