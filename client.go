package go_mypay

import (
	"github.com/asaka1234/go-mypay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params MyPayInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params MyPayInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}

func (cli *Client) SetMerchantInfo(merchantId int, accessKey, backKey string) {
	cli.Params.MerchantId = merchantId
	cli.Params.AccessKey = accessKey
	cli.Params.BackKey = backKey
}
