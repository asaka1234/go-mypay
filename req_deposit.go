package go_mypay

import (
	"crypto/tls"
	"github.com/asaka1234/go-mypay/utils"
	"github.com/mitchellh/mapstructure"
)

// 下单
func (cli *Client) Deposit(req MyPayDepositReq) (*MyPayDepositRsp, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["appId"] = cli.Params.MerchantId
	params["apiOrderType"] = 1 //充值

	// Generate signature
	signStr, _ := utils.Sign(params, cli.Params.AccessKey)
	params["_sign"] = signStr

	var result MyPayDepositRsp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	//fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
