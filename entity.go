package go_mypay

type MyPayInitParams struct {
	MerchantId int    `json:"merchantId" mapstructure:"merchantId" config:"merchantId" yaml:"merchantId"` // merchantId
	AccessKey  string `json:"accessKey" mapstructure:"accessKey" config:"accessKey" yaml:"accessKey"`     // accessKey
	BackKey    string `json:"backKey" mapstructure:"backKey" config:"backKey" yaml:"backKey"`             //backKey

	DepositUrl       string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl" yaml:"depositUrl"`
	DepositCheckUrl  string `json:"depositCheckUrl" mapstructure:"depositCheckUrl" config:"depositCheckUrl" yaml:"depositCheckUrl"`
	WithdrawUrl      string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl" yaml:"withdrawUrl"`
	WithdrawCheckUrl string `json:"withdrawCheckUrl" mapstructure:"withdrawCheckUrl" config:"withdrawCheckUrl" yaml:"withdrawCheckUrl"`
	DealOrderUrl     string `json:"dealOrderUrl" mapstructure:"dealOrderUrl" config:"dealOrderUrl" yaml:"dealOrderUrl"`
}

//----------------------------------

type MyPayDepositReq struct {
	APIUserID     string `json:"apiUserId" mapstructure:"apiUserId"`         //商户下不同用户唯一识别 Id
	APIAmountType string `json:"apiAmountType" mapstructure:"apiAmountType"` //币种, 【1：RMB(默认)；2：USDT】
	Amount        string `json:"amount" mapstructure:"amount"`               //充值金额，为 RMB 时必须为整数
	LegalTenderID string `json:"legaltenderid" mapstructure:"legaltenderid"` //法币类型, 1-人民币；2-美元（美元商户必传）3-越南盾； 4-印尼盾； 5-泰铢， 6-印度卢比 8-韩元
	APIOrderID    string `json:"apiOrderId" mapstructure:"apiOrderId"`       //商户订单唯一识别 Id
	PaymentName   string `json:"paymentName" mapstructure:"paymentName"`     //付款人名称 中文姓名不超过 4 个汉字
	PhoneNumber   string `json:"phoneNumber" mapstructure:"phoneNumber"`     //用户手机号码 在收银台需验证后四位(legaltenderid=1 时需要传入验证)
	//sdk来提供的
	//AppID string `json:"appId" mapstructure:"appId"` //商户的appId
	//sdk计算
	//Signature  string `json:"_sign" mapstructure:"_sign"` //参数里
	//值是写死的1
	//APIOrderType  string `json:"apiOrderType" mapstructure:"apiOrderType"`   //商户订单类型。【1：充值】
}

type MyPayDepositRsp struct {
	Code   int                 `json:"code"`
	Msg    string              `json:"msg"`
	Exist  int                 `json:"exist"`
	Result *MyPayDepositResult `json:"result"`
}

type MyPayDepositResult struct {
	TradeID       int    `json:"tradeId"`
	URL           string `json:"url"`
	SwiftCode     string `json:"swiftcode"`
	APIAmountType string `json:"apiAmountType"`
}

//------------------------------------

type MyPayDepositBackReq struct {
	AppID       string `json:"appId" mapstructure:"appId"`             //商户的 appId
	APIUserID   string `json:"apiUserId" mapstructure:"apiUserId"`     //商户用户 Id
	TradeID     string `json:"tradeId" mapstructure:"tradeId"`         //平台订单唯一识别号 id
	APIOrderID  string `json:"apiOrderId" mapstructure:"apiOrderId"`   //商户订单唯一识别 Id
	TradeStatus string `json:"tradeStatus" mapstructure:"tradeStatus"` //订单状态:0：订单创建中；1 待支付；2 已支付(待放行) ；3 成功(流程完成)；4失败
	Amount      string `json:"amount" mapstructure:"amount"`           //支付金额（四舍五入取整）
	Price       string `json:"price" mapstructure:"price"`             //USDT 价格（1USDT=?RMB）
	AmountUSDT  string `json:"amountUSDT" mapstructure:"amountUSDT"`   //平台商户账号获取 USDT 的数量
	Fee         string `json:"fee" mapstructure:"fee"`                 //手续费（USDT）
	CoinCode    string `json:"coinCode" mapstructure:"coinCode"`       //支付金额的币种类型（RMB）
	PayType     string `json:"payType" mapstructure:"payType"`         //支付方式 1：银行卡；2：支付宝；3：微信
	TimeStamp   string `json:"timeStamp" mapstructure:"timeStamp"`     //Notice: 只有deposit的回调会有该字段
	Signature   string `json:"_sign" mapstructure:"_sign"`
}

// 商户返回 success 说明已收到通知, 否则将再重复 4 次请求

//================================================

type MyPayWithdrawReq struct {
	APIUserID     string `json:"apiUserId" mapstructure:"apiUserId"`         //商户下不同用户唯一识别 Id
	APIAmountType string `json:"apiAmountType" mapstructure:"apiAmountType"` //币种, 【1：RMB(默认)；2：USDT】
	Amount        string `json:"amount" mapstructure:"amount"`               //充值金额，为 RMB 时必须为整数
	LegalTenderID string `json:"legaltenderid" mapstructure:"legaltenderid"` //法币类型, 1-人民币；2-美元（美元商户必传）3-越南盾； 4-印尼盾； 5-泰铢， 6-印度卢比 8-韩元
	TradeType     string `json:"tradeType" mapstructure:"tradeType"`         //交易类型【1：匹配下单；2 抢单】
	APIOrderID    string `json:"apiOrderId" mapstructure:"apiOrderId"`       //商户订单唯一识别 Id
	PayType       string `json:"payType" mapstructure:"payType"`             //支付方式, 1：银行卡；6：支付宝扫码；7：微信扫码；8-momo 电子钱包 9-Zalo 电子钱包；10-TrueMoney 电子钱包
	// NotifyURL     string `json:"notifyUrl" mapstructure:"notifyUrl"` //回调地址(为空，则以商户后台设置回调)
	BankUserName string `json:"bankUserName" mapstructure:"bankUserName"` //收款人名称
	BankCardID   string `json:"bankCardId" mapstructure:"bankCardId"`     //收款人银行卡号/支付宝账户
	BankName     string `json:"bankName" mapstructure:"bankName"`         //收款银行名称
	BankBranch   string `json:"bankBranch" mapstructure:"bankBranch"`     //收款银行支行名称
	//sdk提供
	//AppID string `json:"appId" mapstructure:"appId"` //商户的 appId
	//Signature  string `json:"_sign" mapstructure:"_sign"` //参数里
	////值是写死的1
	//APIOrderType  string `json:"apiOrderType" mapstructure:"apiOrderType"`   //商户订单类型。【2：提现(兑出)】
	//直接sdk设置为当前
	//TimeStamp int64 `json:"timeStamp" mapstructure:"timeStamp"` //时间戳，精确到秒
}

type MyPayWithdrawRsp struct {
	Code   int                  `json:"code"`
	Msg    string               `json:"msg"`
	Result *MyPayWithdrawResult `json:"result"`
}

type MyPayWithdrawResult struct {
	TradeID     string `json:"tradeId"`
	Status      int    `json:"status"`
	TradeStatus int    `json:"tradeStatus"`
}

// --------------------------------------------------

type MyPayWithdrawBackReq struct {
	AppID       string `json:"appId" mapstructure:"appId"`             //商户的 appId
	APIUserID   string `json:"apiUserId" mapstructure:"apiUserId"`     //商户用户 Id
	TradeID     string `json:"tradeId" mapstructure:"tradeId"`         //平台订单唯一识别号 id
	APIOrderID  string `json:"apiOrderId" mapstructure:"apiOrderId"`   //商户订单唯一识别 Id
	TradeStatus int    `json:"tradeStatus" mapstructure:"tradeStatus"` //订单状态:0：订单创建中；1 待支付；2 已支付(待放行) ；3 成功(流程完成)；4失败
	Amount      string `json:"amount" mapstructure:"amount"`           //支付金额（四舍五入取整）
	Price       string `json:"price" mapstructure:"price"`             //USDT 价格（1USDT=?RMB）
	AmountUSDT  int    `json:"amountUSDT" mapstructure:"amountUSDT"`   //平台商户账号获取 USDT 的数量
	Fee         string `json:"fee" mapstructure:"fee"`                 //手续费（USDT）
	CoinCode    string `json:"coinCode" mapstructure:"coinCode"`       //支付金额的币种类型（RMB）
	PayType     string `json:"payType" mapstructure:"payType"`         //支付方式 1：银行卡；2：支付宝；3：微信
	Signature   string `json:"_sign" mapstructure:"_sign"`
}

// 回调通知，商户返回 success 说明已收到通知, 否则将再重复 4 次请求

//============================

// 验证接口(回调成功后必须调用此接口进行验证)->deposit/withdraw的验证都是同一个req和resp
type MyPayCommonCheckReq struct {
	OutTradeNo string `json:"out_trade_no" mapstructure:"out_trade_no"` //商户订单唯一识别 Id
	Amount     string `json:"amount" mapstructure:"amount"`             //交易金额
	//sdk自己设置和计算
	//AppID     string `json:"appId" mapstructure:"appId"` //商户的 appId
	//Signature string `json:"_sign" mapstructure:"_sign"`
}

type MyPayCommonCheckRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//============================

// 验证接口(回调成功后必须调用此接口进行验证)
type MyPayDealOrderReq struct {
	TradeID   string `json:"tradeId" mapstructure:"tradeId"`     //支付平台订单唯一识别号 id
	DealType  string `json:"dealType" mapstructure:"dealType"`   //处理类型: 3：放行（兑出）；4：强制放行（兑出）
	PayerName string `json:"payerName" mapstructure:"payerName"` //付款人名称
	//sdk自己设置和计算
	//AppID     string `json:"appId" mapstructure:"appId"` //商户的 appId
	//Signature string `json:"_sign" mapstructure:"_sign"`
}

type MyPayDealOrderRsp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
