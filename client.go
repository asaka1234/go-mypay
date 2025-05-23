package go_mypay

import (
	"github.com/asaka1234/go-mypay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID int    // merchantId
	AccessKey  string // accessKey
	BackKey    string //backKey

	DepositUrl       string
	DepositCheckUrl  string
	WithdrawUrl      string
	WithdrawCheckUrl string
	DealOrderUrl     string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID int, accessKey, backKey, depositUrl, depositCheckUrl, withdrawUrl, withdrawCheckUrl, dealOrderUrl string) *Client {
	return &Client{
		MerchantID:       merchantID,
		AccessKey:        accessKey,
		BackKey:          backKey,
		DepositUrl:       depositUrl,
		DepositCheckUrl:  depositCheckUrl,
		WithdrawUrl:      withdrawUrl,
		WithdrawCheckUrl: withdrawCheckUrl,
		DealOrderUrl:     dealOrderUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}

func (cli *Client) SetMerchantInfo(merchantId int, accessKey, backKey string) {
	cli.MerchantID = merchantId
	cli.AccessKey = accessKey
	cli.BackKey = backKey
}
