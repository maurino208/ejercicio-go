package main

import (
	"Figuras/domain"
	"Figuras/service"
)

func main() {

	service.CargarFigurasDesdeArchivo("figuras.json")

	service.ListarFiguras()

	service.ObtenerFiguraPorID("elipse-1")
	service.ObtenerFiguraPorID("lalala")

	service.AgregarNuevaFigura(domain.Elipse{ID: "elipse-3", RadioA: 3.0, RadioB: 2.0})
	service.AgregarNuevaFigura(domain.Rectangulo{ID: "rectangulo-2", Base: 5.0, Altura: 4.0})
	service.AgregarNuevaFigura(domain.Triangulo{ID: "triangulo-2", Base: 6.0, Altura: 3.0})

	service.ListarFiguras()
}
