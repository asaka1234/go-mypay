package go_mypay

import (
	"crypto/tls"
	"github.com/asaka1234/go-mypay/utils"
	"github.com/mitchellh/mapstructure"
)

// 下单
func (cli *Client) Deposit(req DepositReq) (*DepositRsp, error) {

	rawURL := cli.DepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["appId"] = cli.MerchantID //uid要参与签名
	params["apiOrderType"] = 1       //充值

	// Generate signature
	signStr, _ := utils.Sign(params, cli.AccessKey)
	params["_sign"] = signStr

	var result DepositRsp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	//fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
