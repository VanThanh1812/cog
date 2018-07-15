package pay

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"log"
	"regexp"
	"cog/models"
	"math/big"
	"fmt"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"cog/contract"
)

type ICogWallet interface {
	Transfer(proofkey string, rq models.TransferRequest)(*types.Transaction, error)
	TransferFrom(proofkey string, rq models.TransferFromRequest)(*types.Transaction, error)
	Earn(proofkey string, eq models.EarnRequest)(*types.Transaction, error)
}

type cogWallet struct {

}


var (
	auth *bind.TransactOpts
	err error
	proofKey = "24e890c939b2f19273a2843f0ba1836c"
	pass = "thanhpv1234"
	session contract.CogNetworkSession
)

func init(){
	auth, err = bind.NewTransactor(strings.NewReader(key), pass)
	session = contract.CogNetworkSession{
		Contract: cog,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From: auth.From,
			Signer: auth.Signer,
			GasLimit: 2000000,
			GasPrice: big.NewInt(21000),
		},
	}

	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
}

func (c cogWallet) Transfer(proofKey string, rq models.TransferRequest)(*types.Transaction, error) {
	if !checkProofKey(proofKey) {
		return nil, errors.New("Proof key error.")
	}

	if !checkAddress(rq.AddrTo){
		return nil, errors.New("Address To wrong.")
	}

	if rq.Amount <= 0 {
		return nil, errors.New("Amount require more than 0")
	}

	tx, err := session.Transfer(common.HexToAddress(rq.AddrTo), big.NewInt(rq.Amount), rq.Content)

	if err != nil {
		log.Fatalf("Failed to request token transfer: %v", err)
		return nil, err
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	return tx, nil
}

func (c cogWallet) TransferFrom(proofkey string, rq models.TransferFromRequest)(*types.Transaction, error){
	return nil, nil
}

func (c cogWallet) Earn(proofKey string, eq models.EarnRequest)(*types.Transaction, error){
	return nil, nil
}

func checkAddress(addr string) bool {
	match, _ := regexp.MatchString("^0x[a-fA-F0-9]{40}$", addr)
	return match
}

func checkProofKey(key string) bool {
	if key == proofKey {
		return true
	}else{
		return false
	}
}

var CogWallet ICogWallet = &cogWallet{}