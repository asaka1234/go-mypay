package go_mypay

import (
	"fmt"
	"testing"
)

func TestCryDeposit(t *testing.T) {

	//构造client
	cli := NewClient(nil, &MyPayInitParams{MERCHANT_ID, ACCESS_KEY, BACK_KEY, DEPOSIT_URL, DEPOSIT_CHECK_URL, WITHDRAW_URL, WITHDRAW_CHECK_URL, DEAL_URL})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenDepositRequestDemo() MyPayDepositReq {
	return MyPayDepositReq{
		APIUserID:     "111",
		APIAmountType: "1",
		Amount:        "200",
		LegalTenderID: "1",
		APIOrderID:    "234",
		PaymentName:   "哈哈",
		PhoneNumber:   "11111",
	}
}
