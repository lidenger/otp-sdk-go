package otpsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// AddAccountSecret 增加账号密钥
func AddAccountSecret(account string) error {
	params := make(map[string]any)
	params["account"] = account
	// 1 启用， 2 禁用
	params["isEnable"] = 1
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return err
	}
	fullUrl, err := url.JoinPath(addr, GenAccessTokenPath)
	if err != nil {
		return err
	}
	resp, err := http.Post(fullUrl, "application/json", bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}
	result, err := readResult(resp)
	if err != nil {
		return err
	}
	fmt.Println(result)

	return nil
}

// 获取指定账号密钥

// 生成动态令牌

// 验证动态令牌
