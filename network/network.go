package network

import (
	"blockchain-miner/minero"
	"encoding/json"
	"net/http"
	"sync"
)

var (
	chain     minero.CadenaBloques
	minerLock sync.Mutex
)

func HandleMineBlock(w http.ResponseWriter, r *http.Request) {
	minerLock.Lock()
	defer minerLock.Unlock()

	var blockData struct {
		Data string `json:"data"`
	}

	if err := json.NewDecoder(r.Body).Decode(&blockData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newBlock := minero.Bloque{
		Timestamp: "now",
		Data:      blockData.Data,
		PrevHash:  chain.ObtenerUltimoBloque().Hash,
	}

	minero.MinarBloque(&newBlock, 5)
	chain.AgregarBloque(newBlock)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBlock)
}

func HandleGetChain(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(chain)
}
