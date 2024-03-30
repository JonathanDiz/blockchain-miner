package main

import (
	"fmt"
)

type Temporada struct {
	Nombre string
}

type CondicionesAmbientales struct {
	Nivel string
}

type ModeloPrediccionPrecio struct{}

func (m *ModeloPrediccionPrecio) PredecirPrecioCarne(oferta, demanda, inflacion float64) float64 {
	// Inferimos el precio de la carne a partir de las ofertas y demandas de los productos y la inflación del mercado financiero en el año actual (por ejemplo, 0.05).
	precioInferido := oferta*demanda + inflacion*0.5
	return precioInferido
}

func (m *ModeloPrediccionPrecio) PredecirPrecioPollo(oferta, demanda, inflacion float64) float64 {
	// Inferimos el precio del pollo a partir de las ofertas y demandas de los productos y la inflación del mercado financiero en el año actual (por ejemplo, 0.05).
	precioInferido := oferta*demanda + inflacion*0.5
	return precioInferido
}

func (m *ModeloPrediccionPrecio) PredecirPrecioPescado(oferta, demanda, inflacion float64, temporada Temporada, condicionesAmbientales CondicionesAmbientales) float64 {
	// Inferimos el precio del pescado a partir de las ofertas y demandas de los productos y la inflación del mercado financiero en el año actual (por ejemplo, 0.05).
	precioInferido := oferta*demanda + inflacion*0.5

	if temporada.Nombre == "verano" {
		precioInferido += 0.1
	}

	if temporada.Nombre == "invierno" {
		precioInferido -= 0.1
	}

	if condicionesAmbientales.Nivel == "optimas" {
		precioInferido += 0.1
	}

	if condicionesAmbientales.Nivel == "riesgo" {
		precioInferido -= 0.1
	}

	return precioInferido
}

func main() {
	// Creamos una instancia del modelo de predicción de precio.
	modeloPrediccionPrecio := &ModeloPrediccionPrecio{}

	// Predecimos el precio de la carne.
	precioCarne := modeloPrediccionPrecio.PredecirPrecioCarne(100, 200, 0.05)
	fmt.Println("Precio de la carne:", precioCarne)

	// Predecimos el precio del pollo.
	precioPollo := modeloPrediccionPrecio.PredecirPrecioPollo(100, 200, 0.05)
	fmt.Println("Precio del pollo:", precioPollo)

	// Creamos instancias de temporada y condiciones ambientales.
	verano := Temporada{Nombre: "verano"}
	optimas := CondicionesAmbientales{Nivel: "optimas"}

	// Predecimos el precio del pescado.
	precioPescado := modeloPrediccionPrecio.PredecirPrecioPescado(100, 200, 0.05, verano, optimas)
	fmt.Println("Precio del pescado:", precioPescado)
}
