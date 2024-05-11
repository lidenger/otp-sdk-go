package otpsdk

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	sign string
	addr string
	key  string
	iv   string
)

// Conf 初始化配置
// serverSign 服务标识（在otp server申请服务时填写的）
func Conf(serverSign, otpServerAddr, ServerKey, ServerIV string) {
	sign = serverSign
	addr = otpServerAddr
	key = ServerKey
	iv = ServerIV
}

type Result struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Msg  string `json:"msg"`
}

// 使用接入服务的密钥和IV生成time token
func genTimeToken() (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	now := strconv.FormatInt(time.Now().Unix(), 10)
	padData := Pad([]byte(now), aes.BlockSize)
	cipherText := make([]byte, len(padData))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(cipherText, padData)
	timeToken := hex.EncodeToString(cipherText)
	return timeToken, nil
}

// GenAccessToken 请求otp server生成 access token
// address otp server地址，例如：http://127.0.0.1:8066
// serverSign 服务标识（在otp server申请服务时填写的）
// key,iv 服务密钥和IV
func GenAccessToken() (string, error) {
	timeToken, err := genTimeToken()
	if err != nil {
		return "", err
	}
	params := make(map[string]string)
	params["serverSign"] = sign
	params["timeToken"] = timeToken
	jsonParams, err := json.Marshal(params)
	if err != nil {
		return "", err
	}
	fullUrl, err := url.JoinPath(addr, GenAccessTokenPath)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(fullUrl, "application/json", bytes.NewBuffer(jsonParams))
	if err != nil {
		return "", err
	}
	result, err := readResult(resp)
	if err != nil {
		return "", err
	}
	return result.Data, nil
}

// VerifyAccessToken 验证access token
func VerifyAccessToken(token string) error {
	fullUrl, err := url.JoinPath(addr, VerifyAccessTokenPath)
	if err != nil {
		return err
	}
	resp, err := http.Get(fullUrl + "?accessToken=" + token)
	if err != nil {
		return err
	}
	_, err = readResult(resp)
	return err
}

// VerifyAndGenAccessToken 验证token，不通过生成获取新的token
func VerifyAndGenAccessToken(token string) (string, error) {
	err := VerifyAccessToken(token)
	if err == nil {
		return token, nil
	}
	return GenAccessToken()
}
