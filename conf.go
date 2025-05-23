package go_mypay

const (
	//参数
	MERCHANT_ID = 111
	ACCESS_KEY  = "111"
	BACK_KEY    = "1111"

	//法币
	DEPOSIT_URL        = "https://api.mypayment.pro/wallet/chat/generateOrders"
	DEPOSIT_CHECK_URL  = "https://api.mypayment.pro/wallet/chat/check"
	WITHDRAW_URL       = "https://api.mypayment.pro/wallet/api/createOrder"
	WITHDRAW_CHECK_URL = "https://api.mypayment.pro/wallet/api/check"
	DEAL_URL           = "https://api.mypayment.pro/wallet/api/dealOrder"
)
