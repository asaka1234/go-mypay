package go_mypay

import (
	"crypto/tls"
	"github.com/asaka1234/go-mypay/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

func (cli *Client) Withdraw(req MyPayWithdrawReq) (*MyPayWithdrawRsp, error) {

	rawURL := cli.Params.WithdrawUrl

	// 2. Convert struct to map for signing
	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["appId"] = cli.Params.MerchantId
	params["apiOrderType"] = 2 //商户订单类型。【2：提现(兑出)】
	params["timeStamp"] = time.Now().Unix()

	// 3. Generate signature
	signStr, _ := utils.Sign(params, cli.Params.AccessKey)
	params["_sign"] = signStr

	var result MyPayWithdrawRsp

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
