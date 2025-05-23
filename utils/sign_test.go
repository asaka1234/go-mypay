package utils

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	
	accessKey := "6308afb129ea00301bd7c79621d07591"

	params := map[string]interface{}{
		"uid":        1254879,
		"amount":     300,
		"coinName":   "USDT",
		"orderId":    "12345678910",
		"protocol":   "ERC20",
		"uniqueCode": 11256,
	}

	signStr, _ := Sign(params, accessKey)

	fmt.Println(signStr)

}
