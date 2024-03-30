package minero

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
)

// Bloque representa un bloque de la cadena de bloques
type Bloque struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
	Nonce     int
}

// MinarGrupoBloques mina un grupo de bloques al mismo tiempo
func MinarGrupoBloques(c CadenaBloques, grupo []Bloque, dificultad int) {
	var wg sync.WaitGroup
	cola := make(chan Bloque, len(grupo))

	for _, bloque := range grupo {
		wg.Add(1)
		go func(b Bloque) {
			defer wg.Done()
			MinarBloque(&b, dificultad)
			cola <- b
		}(bloque)
	}

	wg.Wait()
	close(cola)

	for bloque := range cola {
		c.AgregarBloque(bloque)
	}
}

// CadenaBloques representa la cadena de bloques
type CadenaBloques struct {
	Bloques []Bloque
}

// CalcularHash calcula el hash SHA256 de un bloque
func (b *Bloque) CalcularHash() {
	registro := fmt.Sprintf("%d%s%s%d", b.Index, b.Timestamp, b.Data, b.Nonce)
	hash := sha256.New()
	hash.Write([]byte(registro))
	hashBytes := hash.Sum(nil)
	doubleHash := sha256.New()
	doubleHash.Write(hashBytes)
	b.Hash = hex.EncodeToString(doubleHash.Sum(nil))
}

// AgregarBloque agrega un nuevo bloque a la cadena de bloques
func (c *CadenaBloques) AgregarBloque(b Bloque) {
	b.CalcularHash()
	c.Bloques = append(c.Bloques, b)
}

// ObtenerUltimoBloque obtiene el último bloque de la cadena de bloques
func (c *CadenaBloques) ObtenerUltimoBloque() *Bloque {
	if len(c.Bloques) == 0 {
		return nil
	}
	return &c.Bloques[len(c.Bloques)-1]
}

// MinarBloque mina un bloque hasta encontrar un hash válido con la dificultad especificada
func MinarBloque(b *Bloque, dificultad int) {
	for {
		b.CalcularHash()
		if b.Hash[:dificultad] == "5"[:dificultad] { // Ajusta el patrón para la dificultad
			break
		}
		b.Nonce++
	}
}

// Minería representa el proceso de minería
type Minería struct {
	Éxito    int     // Contador de minerías exitosas
	Contador float64 // Contador de minerías exitosas con 10 decimales
}

// Minar realiza la minería de un bloque
func (m *Minería) Minar(b *Bloque, dificultad int) {
	MinarBloque(b, dificultad)
	m.Éxito++
}
