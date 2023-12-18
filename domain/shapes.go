package domain

import (
	"fmt"
	"math"
)

type Figura interface {
	CalcularArea() float64
	Imprimir() string
	GetId() string
}

type RegistroFigura struct {
	Type   string  `json:"type"`
	ID     string  `json:"ID"`
	RadioA float64 `json:"RadioA,omitempty"`
	RadioB float64 `json:"RadioB,omitempty"`
	Base   float64 `json:"Base,omitempty"`
	Altura float64 `json:"Altura,omitempty"`
}

// Elipse representa una figura elíptica
type Elipse struct {
	ID     string
	RadioA float64
	RadioB float64
}

// CalcularArea calcula el área de la elipse
func (e Elipse) CalcularArea() float64 {
	return math.Pi * e.RadioA * e.RadioB
}

// Imprimir imprime los detalles de la elipse
func (e Elipse) Imprimir() string {
	return fmt.Sprintf("Elipse - ID: %s, RadioA: %f, RadioB: %f", e.ID, e.RadioA, e.RadioB)
}
func (e Elipse) GetId() string {
	return e.ID
}

// Rectangulo representa una figura rectangular
type Rectangulo struct {
	ID     string
	Base   float64
	Altura float64
}

// CalcularArea calcula el área del rectángulo
func (r Rectangulo) CalcularArea() float64 {
	return r.Base * r.Altura
}

// Imprimir imprime los detalles del rectángulo
func (r Rectangulo) Imprimir() string {
	return fmt.Sprintf("Rectángulo - ID: %s, Base: %f, Altura: %f", r.ID, r.Base, r.Altura)
}

func (r Rectangulo) GetId() string {
	return r.ID
}

// Triangulo representa una figura triangular
type Triangulo struct {
	ID     string
	Base   float64
	Altura float64
}

// CalcularArea calcula el área del triángulo
func (t Triangulo) CalcularArea() float64 {
	return 0.5 * t.Base * t.Altura
}

// Imprimir imprime los detalles del triángulo
func (t Triangulo) Imprimir() string {
	return fmt.Sprintf("Triángulo - ID: %s, Base: %f, Altura: %f", t.ID, t.Base, t.Altura)
}

func (t Triangulo) GetId() string {
	return t.ID
}
