package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	store "go-api/abi"
	logType "go-api/model"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/gorilla/mux"
)

var pvtKey string = os.Getenv("PRIVATE_KEY")

var contract string = "0xa51d07bC2dB396Ae209D025566941BE3A6FEd54C"

var ethAddress string = os.Getenv("ETH_RINKEBY")

func main() {
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/getOwner/", getContractOwner).Methods("GET")
	router.HandleFunc("/api/getFailedTransferCredits/{address}", getFailedTransferCredits).Methods("GET")
	router.HandleFunc("/api/gameInfo/", getGameInfo).Methods("GET")

	router.HandleFunc("/api/startGame/", startGameAPI).Methods("GET")
	router.HandleFunc("/api/bid/", makeBidAPI).Methods("POST")
	router.HandleFunc("/api/restartGame/", endRoundAPI).Methods("POST")

	defer log.Fatal(http.ListenAndServe(":8090", router))
}

//  ----------------- Getter APIs ------------------- //
func getContractOwner(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status Created"
	_owner := owner()
	address := common.HexToAddress(contract)
	data := map[string]common.Address{
		"contract": address,
		"owner":    _owner,
	}
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func getFailedTransferCredits(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["address"]
	credits := failedTransferCredits(key)
	jsonResp, err := json.Marshal(credits)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func getGameInfo(w http.ResponseWriter, r *http.Request) {
	_game := gameInfo()
	jsonResp, err := json.Marshal(_game)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

//  --------- Setter APIs (Just for Demo, not need to run as Private Key is required) --------- //
func makeBidAPI(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var bid logType.Bid
	json.Unmarshal(requestBody, &bid)
	makeBid(big.NewInt(bid.BidValue.Int64()))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bid)
}

func endRoundAPI(w http.ResponseWriter, r *http.Request) {
	EndRoundAndRestart()
	data := map[string]string{
		"Message": "Game Restarted",
	}
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func startGameAPI(w http.ResponseWriter, r *http.Request) {
	StartGame()
	data := map[string]string{
		"Message": "Game Restarted",
	}
	jsonResp, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

//  ----------------- Getters ------------------- //
func failedTransferCredits(address string) (value *big.Int) {
	instance := getInstance()
	value_ := big.NewInt(0)
	address_ := common.HexToAddress(address)

	value_, err := instance.FailedTransferCredits(&bind.CallOpts{}, address_)
	if err != nil {
		log.Fatal(err)
	}
	return value_
}

func owner() common.Address {
	instance := getInstance()
	owner_, err := instance.Owner(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	return owner_
}

func gameInfo() []byte {
	instance := getInstance()
	game_, err := instance.GameInfo(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(game_)
	fmt.Println("json data", data)
	return data
}

//  ----------------- Setters ------------------- //
func makeBid(value *big.Int) {
	client, err := ethclient.Dial(ethAddress)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(pvtKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value        // in wei
	auth.GasLimit = uint64(0) // in units
	auth.GasPrice = gasPrice
	instance := getInstance()
	tx, err := instance.Bid(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func StartGame() {
	client, err := ethclient.Dial(ethAddress)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(pvtKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("from address ", fromAddress, nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice
	instance := getInstance()
	tx, err := instance.StartGame(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

func EndRoundAndRestart() {
	client, err := ethclient.Dial(ethAddress)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(pvtKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("from address ", fromAddress, nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1000000) // in units
	auth.GasPrice = gasPrice
	instance := getInstance()
	tx, err := instance.EndGame(auth)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}

// websocket RPC is required
func eventSubscribe() {
	client, err := ethclient.Dial(ethAddress)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}

func makeBidEvent(start *big.Int, end *big.Int) [][]byte {
	client, err := ethclient.Dial(ethAddress)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(contract)
	query := ethereum.FilterQuery{
		FromBlock: start,
		ToBlock:   end,
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	var vLogArray [][]byte
	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Println("Log Data:", vLog.Data)

		fmt.Printf("Log Name: Transfer\n")

		var bidEvent logType.LogBid
		err := contractAbi.UnpackIntoInterface(&bidEvent, "BidMade", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		bidEvent.Bidder = common.HexToAddress(vLog.Topics[1].Hex())
		bidEvent.Amount = big.NewInt(0)
		bidEvent.Timestamp = big.NewInt(0)

		fmt.Printf("From: %s", bidEvent.Bidder.Hex())
		fmt.Println("\nAmount: ", bidEvent.Amount)
		fmt.Println("\nTimestamp: ", bidEvent.Timestamp)
	}

	fmt.Printf("\n\n")
	return vLogArray
}

//  ----------------- Utils ------------------- //
func getInstance() *store.Store {
	client, err := ethclient.Dial(ethAddress)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress(contract)
	instance, err := store.NewStore(address, client)
	return instance
}
