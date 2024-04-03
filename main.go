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

func liberarDatos(datos *[]byte) {
	if datos != nil && len(*datos) <= 1024 {
		// Si los datos son menores o iguales a 1024 bytes, liberarlos
		*datos = (*datos)[:0]
		pool.Put(datos)
	}
}

func minarBloque(cadena *minero.CadenaBloques, bloque *minero.Bloque, dificultad int) {
	for {
		fmt.Printf("Minando bloque %d...\n", bloque.Index)
		minero.MinarBloque(bloque, dificultad) // Llama a la función MinarBloque del paquete minero
		cadena.AgregarBloque(*bloque)

		if verificarHash(bloque.Hash, dificultad) {
			break
		}
		bloque.Nonce++
	}
}

func verificarHash(hash string, dificultad int) bool {
	for i := 0; i < dificultad; i++ {
		if hash[i] != '0' {
			return false
		}
	}
	return true
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
		bloque := minero.Bloque{
			Index:     len(cadena.Bloques),
			Timestamp: time.Now().Format(time.RFC3339),
			Data:      fmt.Sprintf("Datos del bloque %d", len(cadena.Bloques)),
			PrevHash:  cadena.ObtenerUltimoBloque().Hash,
		}

		minarBloque(&cadena, &bloque, dificultad)

		// Muestra un mensaje indicando que se ha minado el bloque
		fmt.Printf("Bloque %d minado\n", bloque.Index)
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
