package main

import (
	"Figuras/domain"
	"Figuras/service"
)

func main() {
	// Cargar figuras desde un archivo
	service.CargarFigurasDesdeArchivo("figuras.json")

	// Listar todas las figuras
	//service.ListarFiguras()

	// Obtener una figura por ID
	service.ObtenerFiguraPorID("elipse-1")
	service.ObtenerFiguraPorID("lalala")

	// Agregar nuevas figuras
	service.AgregarNuevaFigura(domain.Elipse{ID: "elipse-3", RadioA: 3.0, RadioB: 2.0})
	service.AgregarNuevaFigura(domain.Rectangulo{ID: "rectangulo-2", Base: 5.0, Altura: 4.0})
	service.AgregarNuevaFigura(domain.Triangulo{ID: "triangulo-2", Base: 6.0, Altura: 3.0})

	// Listar todas las figuras después de agregar nuevas figuras
	service.ListarFiguras()
}
