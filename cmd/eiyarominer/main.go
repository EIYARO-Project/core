package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"core/api"
	"core/consensus"
	"core/consensus/difficulty"
	"core/protocol/bc"
	"core/protocol/bc/types"
	"core/util"
)

const (
	maxNonce = ^uint64(0) // 2^64 - 1
	isCrazy  = true
	esHR     = 1 //estimated Hashrate
)

var (
	lastNonce  = ^uint64(0)
	lastHeight = uint64(0)
	wg         sync.WaitGroup
	id         uint64
)

// do proof of work
func doWork(id uint64, bh *types.BlockHeader, seed *bc.Hash, startNonce uint64) bool {
	log.Printf("(%5d)Start from nonce: %d", id, startNonce)
	nonce := startNonce
	for i := startNonce; i <= uint64(startNonce+consensus.TargetSecondsPerBlock*esHR) && i <= startNonce+9; i++ {
		nonce = i
		//log.Printf("Current nonce = %v\n", i)
		headerHash := bh.Hash()
		if difficulty.CheckProofOfWork(&headerHash, seed, bh.Bits) {
			log.Printf("(%5d)Mining succeed! Proof hash: %v\n", id, headerHash.String())
			return true
		}
	}
	log.Printf("(%5d)Stop at nonce: %d", id, nonce)
	//lastNonce = bh.Nonce
	return false
}

func getBlockHeaderByHeight(height uint64) {
	type Req struct {
		BlockHeight uint64 `json:"block_height"`
	}

	type Resp struct {
		BlockHeader *types.BlockHeader `json:"block_header"`
		Reward      uint64             `json:"reward"`
	}

	data, _ := util.ClientCall("/get-block-header", Req{BlockHeight: height})
	rawData, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	resp := &Resp{}
	if err = json.Unmarshal(rawData, resp); err != nil {
		log.Fatalln(err)
	}
	log.Println("Reward:", resp.Reward)
}

func job(id uint64, resp *api.GetWorkResp, nonce uint64) {
	defer wg.Done()
	if doWork(id, resp.BlockHeader, resp.Seed, nonce) {
		util.ClientCall("/submit-work", &api.SubmitWorkReq{BlockHeader: resp.BlockHeader})
		getBlockHeaderByHeight(resp.BlockHeader.Height)
	}
}

func main() {
	for {
		data, _ := util.ClientCall("/get-work", nil)
		if data == nil {
			os.Exit(1)
		}
		rawData, err := json.Marshal(data)
		if err != nil {
			log.Fatalln(err)
		}
		resp := &api.GetWorkResp{}
		if err = json.Unmarshal(rawData, resp); err != nil {
			log.Fatalln(err)
		}

		log.Println("Mining at height:", resp.BlockHeader.Height)
		if lastHeight != resp.BlockHeader.Height {
			//lastNonce = ^uint64(0)
			lastNonce = uint64(0)
			id = 1
		}
		for t := 0; t < 30; t++ {
			wg.Add(1)
			go job(id, resp, lastNonce)
			lastNonce += 10
			id++
		}
		wg.Wait()

		lastHeight = resp.BlockHeader.Height
		if !isCrazy {
			return
		}
	}
}
