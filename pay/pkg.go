package pay

import (
	"cog/contract"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	addrContract = "0xFdC1FeF0976268eaFb72112E321fCb50d9e70f89"
	addrOwner = "0x3A120D99ac41D6e0630382752F9714D8772c5432"
	// import account to unlock in geth, file keyjson in keystore
	key = `{"address":"3a120d99ac41d6e0630382752f9714d8772c5432","crypto":{"cipher":"aes-128-ctr","ciphertext":"a844aa4d9aea200f8b914b8b55f7301a789dd306cefbe9770b114aa2949bfa9d","cipherparams":{"iv":"f488e7f5c5de84f18f9978280e428aa1"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ad69abe52dfa549b5d6c444cef3da08a8167e2f639c4e1d0d68a591039d205ed"},"mac":"690b5c57a539a5074abae6f87dd8d1094d6f76c52be971286e0698550d24d77d"},"id":"be3303af-83c3-4555-a10b-1bcf4920b659","version":3}`
	//pathIPC = "/home/vanthanhbk/.ethereum/rinkeby/geth.ipc"
	pathIPC = "/home/phanvanthanh_mrt/.ethereum/rinkeby/geth.ipc"
)

var cog *contract.CogNetwork

func InitializeContract() error {
	client, err := ethclient.Dial(pathIPC)
	_cog, err := contract.NewCogNetwork(common.HexToAddress(addrContract), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	cog = _cog
	return nil
}
