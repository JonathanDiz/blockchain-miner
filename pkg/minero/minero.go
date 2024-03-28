package minero

import (
	"crypto/sha256"
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

// CadenaBloques representa la cadena de bloques
type CadenaBloques struct {
	Bloques []Bloque
}

// CalcularHash calcula el hash SHA256 de un bloque
func (b *Bloque) CalcularHash() {
	registro := fmt.Sprintf("%d%s%s%d", b.Index, b.Timestamp, b.Data, b.Nonce)
	hash := sha256.New()
	hash.Write([]byte(registro))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
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

// MinarBloque mina un bloque hasta encontrar un hash válido
func MinarBloque(b *Bloque) {
	for {
		b.CalcularHash()
		if b.Hash[0:4] == "0000" {
			break
		}
		b.Nonce++
	}
}

// Minería representa el proceso de minería
type Minería struct {
	Éxito int // Contador de minerías exitosas
}

// Minar realiza la minería de un bloque
func (m *Minería) Minar(b *Bloque) {
	MinarBloque(b)
	m.Éxito++
}
