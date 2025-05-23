package go_mypay

import (
	"encoding/json"
	"errors"
	"github.com/asaka1234/go-mypay/utils"
	"github.com/mitchellh/mapstructure"
	"log"
)

func (cli *Client) WithdrawCallback(req WithdrawBackReq, processor func(WithdrawBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	// Verify signature
	flag, err := utils.Verify(params, cli.BackKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		return err
	}
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		log.Printf("MyPay back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
