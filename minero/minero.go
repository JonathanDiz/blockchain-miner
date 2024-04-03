package minero

import (
	"blockchain-miner/sha256"
	"encoding/hex"
	"fmt"
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

// MinarBloque mina un bloque hasta encontrar un hash válido con la dificultad especificada
func MinarBloque(b *Bloque, dificultad int) {
	for {
		b.CalcularHash()
		if verificarHash(b.Hash, dificultad) {
			break
		}
		b.Nonce++
	}
}

// verificarHash verifica si el hash cumple con la dificultad especificada
func verificarHash(hash string, dificultad int) bool {
	for i := 0; i < dificultad; i++ {
		if hash[i] != '0' {
			return false
		}
	}
	return true
}

// CalcularHash calcula el hash SHA256 de un bloque utilizando el paquete sha256
func (b *Bloque) CalcularHash() {
	registro := fmt.Sprintf("%d%s%s%d", b.Index, b.Timestamp, b.Data, b.Nonce)
	hash := sha256.Hash([]byte(registro)) // Usa la función Hash del paquete sha256
	b.Hash = hex.EncodeToString(hash[:])
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

// CadenaBloques representa la cadena de bloques
type CadenaBloques struct {
	Bloques []Bloque
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
