package otpsdk

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type AccountSecretModel struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	SecretSeed string    `gorm:"column:secret_seed_cipher" json:"secret"` // 密钥种子密文 = AES(KEY, 密钥种子)
	Account    string    `gorm:"column:account" json:"account"`           // 账号
	DataCheck  string    `gorm:"column:data_check" json:"dataCheck"`      // 数据校验 = HMACSHA256(KEY, secret_seed_cipher + account + is_enable)
	IsEnable   uint8     `gorm:"column:is_enable" json:"isEnable"`        // 是否启用，1启用，2禁用
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

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
	fullUrl, err := url.JoinPath(addr, AddAccountSecretPath)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", fullUrl, bytes.NewBuffer(jsonParams))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	token, err := GenAccessToken()
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	_, err = readResult[string](resp)
	return err
}

// GetAccountSecret 获取指定账号密钥
func GetAccountSecret(account string) (*AccountSecretModel, error) {
	fullUrl, err := url.JoinPath(addr, GetAccountSecretPath, "/"+account)
	if err != nil {
		return nil, err
	}
	resp, err := httpGetReq(fullUrl)
	if err != nil {
		return nil, err
	}
	result, err := readResult[*AccountSecretModel](resp)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// VerifyCode 验证动态令牌
func VerifyCode(account, code string) (bool, error) {
	fullUrl := addr + VerifyCodePath + "?account=" + account + "&code=" + code
	resp, err := httpGetReq(fullUrl)
	if err != nil {
		return false, err
	}
	result, err := readResult[bool](resp)
	if err != nil {
		return false, err
	}
	return result.Data, nil
}

// GetQRCodeContent 获取密钥二维码内容
func GetQRCodeContent(account string) (string, error) {
	fullUrl, err := url.JoinPath(addr, GetAccountSecretPath, "/"+account+"/qrcode-content")
	if err != nil {
		return "", err
	}
	resp, err := httpGetReq(fullUrl)
	if err != nil {
		return "", err
	}
	result, err := readResult[string](resp)
	if err != nil {
		return "", err
	}
	return result.Data, nil
}
