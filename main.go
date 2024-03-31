package main

import (
	"blockchain-miner/minero"
	"fmt"
	"sync"
	"time"
)

var pool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024) // Ejemplo de inicialización de un recurso
	},
}

func obtenerDatos() []byte {
	datos := pool.Get()
	if datos == nil {
		// Si no hay datos disponibles en el pool, crear nuevos datos
		return make([]byte, 1024)
	}
	return datos.([]byte)
}

func liberarDatos(datos *[]byte) {
	if datos != nil && len(*datos) <= 1024 {
		// Si los datos son menores o iguales a 1024 bytes, liberarlos
		*datos = (*datos)[:0]
		pool.Put(datos)
	}
}

func minarGrupo(cadena *minero.CadenaBloques, grupoBloques []minero.Bloque, wg *sync.WaitGroup, dificultad int) {
	defer wg.Done()
	for i, bloque := range grupoBloques {
		fmt.Printf("Minando bloque %d del grupo...\n", i)
		minero.MinarBloque(&bloque, dificultad) // Llama a la función MinarBloque del paquete minero
		cadena.AgregarBloque(bloque)

		datos := obtenerDatos()
		defer liberarDatos(&datos)
	}
}

func main() {
	cadena := minero.CadenaBloques{}
	minería := minero.Minería{}

	bloqueGenesis := minero.Bloque{
		Index:     0,
		Timestamp: time.Now().Format(time.RFC3339),
		Data:      "Genesis Block",
		PrevHash:  "",
	}

	bloqueGenesis.CalcularHash()
	cadena.AgregarBloque(bloqueGenesis)

	target := 0.0000001000 // Define el objetivo de la minería por sesión de minado

	dificultad := 5 // Dificultad inicial

	for minería.Contador < target {
		grupoBloques := make([]minero.Bloque, 0)
		for i := 0; i < 20; i++ {
			bloque := minero.Bloque{
				Index:     i,
				Timestamp: time.Now().Format(time.RFC3339),
				Data:      fmt.Sprintf("Datos del bloque %d", i),
				PrevHash:  cadena.ObtenerUltimoBloque().Hash,
			}
			grupoBloques = append(grupoBloques, bloque)
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go minarGrupo(&cadena, grupoBloques, &wg, dificultad)

		// Espera a que se complete el minado del grupo
		wg.Wait()

		// Muestra un mensaje indicando que se ha minado el grupo
		fmt.Println("Grupo de bloques minado")
		for _, bloque := range grupoBloques {
			fmt.Printf("Grupo %d\n", bloque.Index)
		}
		fmt.Println("--------------------")

		// Aumenta el contador de éxito
		minería.Éxito++
		fmt.Printf("Éxito: %d\n", minería.Éxito)
		fmt.Println("--------------------")

		// Ajusta la dificultad dinámicamente si es necesario
		// Aquí puedes implementar tu lógica para ajustar la dificultad
		// basándote en el tiempo de minado promedio, la tasa de éxito, etc.

		// Verifica si se ha alcanzado el objetivo de minería por sesión
		minería.Contador = float64(minería.Éxito) / 1e10
		if minería.Contador >= target {
			fmt.Println("Se alcanzó el objetivo por sesión de minado. Deteniendo la minería.")
			break
		}
	}
}
