package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"encoding/json"
	"io"
	"github.com/davecgh/go-spew/spew"
	"os"
	"log"
	"sync"
	"strings"
	"github.com/joho/godotenv"
	"net"
	"math/rand"
	"bufio"
	"strconv"
	"fmt"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string //维护每个区块在链中的正确顺序
	Validator string
}

var mutex = &sync.Mutex{}

var BlockChain []Block
var tempBlocks []Block

var candidateBlocks = make(chan Block)
var announcements = make(chan string)
var validators = make(map[string]int)

func calculateHash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func calculateBlockHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) +
		block.PrevHash
	return calculateHash(record)
}

func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {
	var newBlock Block

	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateBlockHash(newBlock)
	newBlock.Validator = address

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

	if calculateBlockHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func repliaceChain(newBLocks []Block) {
	if len(newBLocks) > len(BlockChain) {
		BlockChain = newBLocks
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()
	genesisBlock := Block{}
	genesisBlock = Block{0, t.String(), 0, "", calculateBlockHash(genesisBlock), ""}
	spew.Dump(genesisBlock)

	BlockChain = append(BlockChain, genesisBlock)

	httpPort := os.Getenv("PORT")
	server, err := net.Listen("tcp", ":"+httpPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Http Server Listening on ", os.Getenv("PORT"))
	defer server.Close()

	go func() {
		for candidate := range candidateBlocks {
			mutex.Lock()
			tempBlocks = append(tempBlocks, candidate)
			mutex.Unlock()
		}
	}()

	go func() {
		for {
			pickWinner()
		}
	}()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func pickWinner() {
	fmt.Println("pickWinner func")
	time.Sleep(30 * time.Second)
	mutex.Lock()
	tempBlocksCopy := tempBlocks
	mutex.Unlock()
	fmt.Println("pickWinner start")

	//奖池
	lotteryPool := []string{}

	if len(tempBlocksCopy) > 0 {
	OUTER:
	//遍历tempblocks中的全部block
		for _, block := range (tempBlocksCopy) {
			for _, node := range lotteryPool {
				if block.Validator == node {
					continue OUTER
				}
			}

			// lock list of validators to prevent data race
			mutex.Lock()
			setValidators := validators
			mutex.Unlock()

			k, ok := setValidators[block.Validator]
			if ok {
				for i := 0; i < k; i++ {
					lotteryPool = append(lotteryPool, block.Validator)
				}
			}
		}

		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		//从奖池中随机选出一个优胜者，什么意思，PoS？
		lotterWinner := lotteryPool[r.Intn(len(lotteryPool))]

		//将优胜区块加入区块链，并且让所有其它节点知道
		for _, block := range tempBlocksCopy {
			if block.Validator == lotterWinner {
				mutex.Lock()
				BlockChain = append(BlockChain, block)
				mutex.Unlock()

				//告知其它节点
				for _ = range validators {
					announcements <- "\nwinning validator: " + lotterWinner + "\n"
				}
				break
			}
		}
	}

	mutex.Lock()
	tempBlocks = []Block{}
	mutex.Unlock()
}

func handleConn(conn net.Conn) {
	fmt.Println("handleConn func")
	defer conn.Close()

	go func() {
		for {
			msg := <-announcements
			io.WriteString(conn, msg)
		}
	}()

	var address string
	io.WriteString(conn, "Enter token balance:")
	scanBalance := bufio.NewScanner(conn)
	for scanBalance.Scan() {
		balance, err := strconv.Atoi(scanBalance.Text())
		if err != nil {
			log.Printf("%v not a number:%v", scanBalance.Text(), err)
			return
		}else{
			log.Println("get token balance:%d",balance)
		}
		t := time.Now()
		address = calculateHash(t.String())
		validators[address] = balance
		fmt.Println(validators)
		break
	}

	io.WriteString(conn, "\nEnter a new BPM:")
	scanBPM := bufio.NewScanner(conn)
	go func() {
		for {
			for scanBPM.Scan() {
				bpm, err := strconv.Atoi(scanBPM.Text())
				if err != nil {
					log.Printf("%v not a number:%v", scanBPM.Text(), err)
					delete(validators, address)
					conn.Close()
				}else{
					log.Printf("get bpm:%d\n",bpm)
				}

				mutex.Lock()
				oldLastIndex := BlockChain[len(BlockChain)-1]
				mutex.Unlock()

				newBlock, err := generateBlock(oldLastIndex, bpm, address)
				if err != nil {
					log.Println(err)
					continue
				}else{
					fmt.Println("get a block")
					spew.Dump(newBlock)
				}

				if isBlockValid(newBlock, oldLastIndex) {
					candidateBlocks <- newBlock
				}
				io.WriteString(conn, "\nEnter a new BPM:")
			}
		}
	}()

	for {
		time.Sleep(time.Minute)
		mutex.Lock()
		output, err := json.Marshal(BlockChain)
		mutex.Unlock()

		if err != nil {
			log.Fatal(err)
		}

		io.WriteString(conn, string(output)+"\n")
	}
}
