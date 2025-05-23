package go_mypay

import (
	"crypto/tls"
	"github.com/asaka1234/go-mypay/utils"
	"github.com/mitchellh/mapstructure"
)

// outType = 1, deposit ,    =2, withdraw
func (cli *Client) Check(req CommonCheckReq, outType int) (*CommonCheckRsp, error) {

	rawURL := ""
	if outType == 1 {
		//deposit
		rawURL = cli.DepositCheckUrl
	} else if outType == 2 {
		//withdraw
		rawURL = cli.WithdrawCheckUrl
	}

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["appId"] = cli.MerchantID

	// Generate signature
	signStr, _ := utils.Sign(params, cli.AccessKey)
	params["_sign"] = signStr

	var result CommonCheckRsp

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
