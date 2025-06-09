package go_mypay

import (
	"fmt"
	"testing"
)

func TestCryWithdraw(t *testing.T) {

	//构造client
	cli := NewClient(nil, &MyPayInitParams{MERCHANT_ID, ACCESS_KEY, BACK_KEY, DEPOSIT_URL, DEPOSIT_CHECK_URL, WITHDRAW_URL, WITHDRAW_CHECK_URL, DEAL_URL})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() MyPayWithdrawReq {
	return MyPayWithdrawReq{
		APIUserID:     "111",
		APIAmountType: "1",
		Amount:        "5000",
		LegalTenderID: "1",
		TradeType:     "1",
		APIOrderID:    "23",
		PayType:       "1",
		BankUserName:  "哈哈",
		BankCardID:    "1791791",
		BankName:      "呵呵呵",
		BankBranch:    "支行",
	}
}
