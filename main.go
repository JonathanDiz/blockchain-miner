package main

import (
	"blockchain-miner/minero"  // Importa el paquete minero
	"blockchain-miner/network" // Importa el paquete network
	"fmt"
	"net/http"
	"sync"
	"time"
)

func minarGrupo(cadena *minero.CadenaBloques, grupoBloques []minero.Bloque, wg *sync.WaitGroup) {
	defer wg.Done()
	for i, bloque := range grupoBloques {
		fmt.Printf("Minando bloque %d del grupo...\n", i)
		minero.MinarBloque(&bloque, 0) // Llama a la función MinarBloque del paquete minero
		cadena.AgregarBloque(bloque)
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

	// Configurar rutas HTTP
	http.HandleFunc("/mine", network.HandleMineBlock)
	http.HandleFunc("/chain", network.HandleGetChain)

	// Iniciar servidor HTTP
	go func() {
		fmt.Println("Servidor escuchando en el puerto 8080...")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Error en el servidor:", err)
		}
	}()

	target := 0.0000001000 // Define el objetivo de la minería por sesión de minado

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
		go minarGrupo(&cadena, grupoBloques, &wg)

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

		// Verifica si se ha alcanzado el objetivo de minería por sesión
		minería.Contador = float64(minería.Éxito) / 1e10
		if minería.Contador >= target {
			fmt.Println("Se alcanzó el objetivo por sesión de minado. Deteniendo la minería.")
			break
		}
	}
}
