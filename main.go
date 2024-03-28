package main

import (
	"fmt"
	"time"

	"github.com/JonathanDiz/blockchain-miner/pkg/minero" // Importa el paquete minería utilizando la convención de importación de módulos de Go
)

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

	for i := 0; i < 100; i++ { // Corrige el límite del bucle para que solo haya 100 bloques
		datosBloque := fmt.Sprintf("Bloque %d", i+1)
		nuevoBloque := minero.Bloque{
			Index:     i + 1,
			Timestamp: time.Now().Format(time.RFC3339),
			Data:      datosBloque,
			PrevHash:  cadena.ObtenerUltimoBloque().Hash,
		}
		minero.MinarBloque(&nuevoBloque) // Llama a la función MinarBloque del paquete minero
		cadena.AgregarBloque(nuevoBloque)
	}

	// Mostrar el contador de minerías exitosas con 10 decimales
	fmt.Printf("Contador de minerías exitosas: %.10f\n", float64(minería.Éxito)/1e10)
}
