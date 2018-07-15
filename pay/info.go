package pay

import (
	"cog/models"
	"github.com/ethereum/go-ethereum/common"
)

type ICogInfo interface {
	GetOwner () (models.Address)
	GetBalanceOf (addr string) (*models.BalanceInfo, error)
}

type cogInfo struct {

}

func (c cogInfo) GetOwner() (models.Address){
	owner, _ := cog.Owner(nil)
	o := models.Address{
		Address: owner.String(),
	} //aass
	return o
}

func (c cogInfo) GetBalanceOf (addr string) (*models.BalanceInfo, error){
	balance, err := cog.BalanceOf(nil, common.HexToAddress(addr))
	if err != nil{
		return nil, err
	}else{
		return &models.BalanceInfo{Addr:addr, Balance: balance}, nil
	}
}

var CogInfo ICogInfo = &cogInfo{}