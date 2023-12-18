package service

import (
	"Figuras/domain"
	"encoding/json"
	"fmt"
	"os"
)

var figuras = make(map[string]domain.Figura)

func CargarFigurasDesdeArchivo(filename string) {

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	var registros []domain.RegistroFigura
	err = json.Unmarshal(data, &registros)
	if err != nil {
		fmt.Println("Error al deserializar el archivo JSON:", err)
		return
	}

	for _, registro := range registros {
		var figura domain.Figura

		switch registro.Type {
		case "Elipse":
			figura = &domain.Elipse{ID: registro.ID, RadioA: registro.RadioA, RadioB: registro.RadioB}
		case "Rectangulo":
			figura = &domain.Rectangulo{ID: registro.ID, Base: registro.Base, Altura: registro.Altura}
		case "Triangulo":
			figura = &domain.Triangulo{ID: registro.ID, Base: registro.Base, Altura: registro.Altura}
		default:
			fmt.Println("Tipo de figura no reconocido:", registro.Type)
			continue
		}

		figuras[figura.Imprimir()] = figura
	}
}

func ListarFiguras() {
	fmt.Println("Listado de Figuras:")
	for _, figura := range figuras {
		fmt.Println(figura.Imprimir())
		fmt.Printf("Área: %f\n\n", figura.CalcularArea())
	}
}

func ObtenerFiguraPorID(findId string) {
	encontrada := false
EncontrarFigura:
	for _, figura := range figuras {
		if figura.GetId() == findId {
			fmt.Printf("Figura encontrada:\nID: %s\n", figura.GetId())
			fmt.Println(figura.Imprimir())
			fmt.Printf("Área: %f\n\n", figura.CalcularArea())
			encontrada = true
			break EncontrarFigura
		}
	}

	if !encontrada {
		fmt.Println("No se encontró ninguna figura con el ID:", findId)
	}
}

func AgregarNuevaFigura(figura domain.Figura) {
	figuras[figura.Imprimir()] = figura
	fmt.Printf("Nueva figura agregada:\n%s\nÁrea: %f\n\n", figura.Imprimir(), figura.CalcularArea())
}
