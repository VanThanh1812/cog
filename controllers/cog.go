package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"cog/models"
	"cog/pay"
)

//  CogController operations for Cog
type CogController struct {
	beego.Controller
}

// URLMapping ...
func (c *CogController) URLMapping() {
	c.Mapping("GetOwner", c.GetOwner)
	c.Mapping("GetBalanceOf", c.GetBalanceOf)
	c.Mapping("Transfer", c.Transfer)
}


// GetOwner ...
// @Title GetOwner
// @Description get Owner Contract
// @Success 200 {object} models.Address
// @Failure 403  is empty
// @router /owner [get]
func (c *CogController) GetOwner(){
	res := models.Response{
		Data: pay.CogInfo.GetOwner(),
		Error: models.ErrorResponse{
			Message:"Success",
			Code:200,
		},
	}
	c.Data["json"] = res
	c.ServeJSON()
}

// GetBalanceOf by address ...
// @Title GetBalanceof
// @Description get Balance of Address
// @Param	address		query 	string	true		"the address you want to get"
// @Success 200 {object} models.Response
// @Failure 403  is error
// @router /balanceof [get]
func (c *CogController) GetBalanceOf(){
	addr := c.Ctx.Input.Query("address")

	if addr == "" {
		return
	}

	balance, err := pay.CogInfo.GetBalanceOf(addr)

	var res models.Response

	if err == nil{
		res = models.Response{
			Data: balance,
			Error: models.ErrorResponse{
				Message:"Success",
				Code:200,
			},
		}
	}else{
		res = models.Response{
			Data: 0,
			Error: models.ErrorResponse{
				Message: "Success",
				Code:403,
			},
		}
	}

	c.Data["json"] = res
	c.ServeJSON()
}

// Post ...
// @Title Send amount CNC
// @Description create request tranfer
// @Param	key		query 	string	true		"The key for req"
// @Param	body		body 	models.TransferRequest	true		"body for Payment content"
// @Success 201 {object} models.Response
// @Failure 403 body is empty
// @router /transfer [post]
func (c *CogController) Transfer() {
	var rq models.TransferRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &rq)

	var proofkey string
	proofkey = c.Ctx.Input.Query("key")

	var res models.Response
	tx, err := pay.CogWallet.Transfer(proofkey, rq)

	if err == nil{
		res = models.Response{
			Data: tx,
			Error: models.ErrorResponse{
				Message:"Success",
				Code:200,
			},
		}
	}else{
		res = models.Response{
			Data: nil,
			Error: models.ErrorResponse{
				Message: err.Error(),
				Code: 403,
			},
		}
	}

	c.Data["json"] = res
	c.ServeJSON()
}
