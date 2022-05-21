package figuras

import "fmt"

//interface
type Geometria interface {
	area() float64
	perimetro() float64
}

//funcion para mostrar mensaje
func Medidas(g Geometria) {
	fmt.Println("Medidas:", g)
	fmt.Println("El área es:", g.area())
	fmt.Println("El perímetro es:", g.perimetro())
}
