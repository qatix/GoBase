package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io"
	"github.com/davecgh/go-spew/spew"
	"os"
	"log"
	"sync"
	"fmt"
	"strings"
	"github.com/joho/godotenv"
)

type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string //维护每个区块在链中的正确顺序
	Difficulty int
	Nonce      string
}

const difficulty = 1

var mutex = &sync.Mutex{}

var BlockChain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) +
		block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock Block

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	newBlock.Difficulty = difficulty

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		hashStr := calculateHash(newBlock)
		//spew.Dump(newBlock)
		if !isHashValid(hashStr, newBlock.Difficulty) {
			fmt.Println(hashStr,"[",i,"]do more work!")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(hashStr, " work done！")
			newBlock.Hash = hashStr
			break
		}
	}

	return newBlock, nil
}

func isHashValid(hash string, difficulty int) bool {
	prefix := strings.Repeat("0", difficulty)
	return strings.HasPrefix(hash, prefix);
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func repliaceChain(newBLocks []Block) {
	if len(newBLocks) > len(BlockChain) {
		BlockChain = newBLocks
	}
}

func run() error {
	mux := makeMuxRouter()
	httpPort := os.Getenv("PORT")
	log.Println("Listening on ", os.Getenv("PORT"))

	s := &http.Server{
		Addr:           ":" + httpPort,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(BlockChain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

type Message struct {
	BPM int
}

func handleWriteBlockchain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}

	defer r.Body.Close()

	mutex.Lock()
	newBlock, err := generateBlock(BlockChain[len(BlockChain)-1], m.BPM)
	mutex.Unlock()
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}

	if isBlockValid(newBlock, BlockChain[len(BlockChain)-1]) {
		newBlockChain := append(BlockChain, newBlock)
		repliaceChain(newBlockChain)
		spew.Dump(BlockChain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlockchain).Methods("POST")
	return muxRouter
}

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{0,t.String(),0,"",calculateHash(genesisBlock),difficulty,""}
		spew.Dump(genesisBlock)

		mutex.Lock()
		BlockChain = append(BlockChain, genesisBlock)
		mutex.Unlock()
	}()

	log.Fatal(run())

	//t := time.Now()
	//genesisBlock := Block{0, t.String(), 0, "", "", difficulty, ""}
	//spew.Dump(genesisBlock)
	//
	//BlockChain = append(BlockChain, genesisBlock)
	//
	//m := Message{BPM: 22}
	//newBlock, err := generateBlock(BlockChain[len(BlockChain)-1], m.BPM)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	spew.Dump(newBlock)
	//}

}
