package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) area() float64 {
	return r.altura * r.largura
}


type Circulo struct {
	raio float64
}

func (c Circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}
