package main

import (
	"fmt"

	"github.com/mumax/3/cuda" // Importar la librería CUDA de Go
)

// Ejemplo de uso de la librería CUDA para obtener información sobre los dispositivos CUDA disponibles.
func main() {
	// Inicializar CUDA
	err := cuda.Init(0)
	if err != nil {
		fmt.Println("Error al inicializar CUDA:", err)
		return
	}
	defer cuda.Finish()

	// Obtener el número de dispositivos CUDA disponibles
	numDevices, err := cuda.DeviceCount()
	if err != nil {
		fmt.Println("Error al obtener el número de dispositivos CUDA:", err)
		return
	}

	fmt.Printf("Número de dispositivos CUDA disponibles: %d\n", numDevices)

	// Iterar sobre todos los dispositivos CUDA y mostrar información sobre cada uno
	for i := 0; i < numDevices; i++ {
		dev, err := cuda.Device(i)
		if err != nil {
			fmt.Printf("Error al obtener el dispositivo CUDA %d: %v\n", i, err)
			continue
		}

		// Imprimir información sobre el dispositivo CUDA
		fmt.Printf("Dispositivo CUDA %d:\n", i)
		fmt.Printf("  Nombre: %s\n", dev.Name())
		fmt.Printf("  Memoria global disponible: %d bytes\n", dev.TotalMemory())
		fmt.Printf("  Número máximo de hilos por bloque: %d\n", dev.MaxThreadsPerBlock())
		fmt.Printf("  Número máximo de bloques por cuadrícula: %d\n", dev.MaxGridSize(0))
	}
}
