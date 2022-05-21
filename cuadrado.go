package figuras

//estructuras
type Cuadrado struct {
	Lado float64
}

//metodos para sacar el área y el perímetro

func (cua *Cuadrado) area() float64 {
	return cua.Lado * cua.Lado
}

func (cua *Cuadrado) perimetro() float64 {
	return cua.Lado * 4
}
