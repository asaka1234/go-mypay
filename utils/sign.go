package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"sort"
	"strings"
)

func Sign(params map[string]interface{}, key string) (string, error) {
	// 1. Validate key
	if key == "" {
		return "", errors.New("APP_KEY 参数为空，请填写")
	}

	// 2. Get and sort keys
	keys := lo.Keys(params)
	sort.Strings(keys) // ASCII ascending order

	// 3. Build sign string
	var sb strings.Builder
	for _, k := range keys {
		value := cast.ToString(params[k])
		if k != "_sign" && value != "" {
			//只有非空才可以参与签名
			sb.WriteString(fmt.Sprintf("%s=%s&", k, value))
		}
	}
	sb.WriteString(fmt.Sprintf("key=%s", key))
	signStr := sb.String()

	// 4. Generate MD5
	hash := md5.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	// Debug print (optional)
	fmt.Printf("验签str: %s\n结果: %s\n", signStr, signResult)

	return signResult, nil
}

func Verify(params map[string]interface{}, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["_sign"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "_sign")

	// Generate current signature
	currentSignature, err := Sign(params, signKey)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature.(string) == currentSignature, nil
}
