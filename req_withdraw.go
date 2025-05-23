package go_mypay

import (
	"crypto/tls"
	"github.com/asaka1234/go-mypay/utils"
	"github.com/mitchellh/mapstructure"
	"time"
)

func (cli *Client) Withdraw(req WithdrawReq) (*WithdrawRsp, error) {

	rawURL := cli.WithdrawUrl

	// 2. Convert struct to map for signing
	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["appId"] = cli.MerchantID //uid要参与签名
	params["apiOrderType"] = 2       //商户订单类型。【2：提现(兑出)】
	params["timeStamp"] = time.Now().Unix()

	// 3. Generate signature
	signStr, _ := utils.Sign(params, cli.AccessKey)
	params["_sign"] = signStr

	var result WithdrawRsp

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
