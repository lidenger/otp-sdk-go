package otpsdk

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	GenAccessTokenPath    = "/v1/access-token"
	VerifyAccessTokenPath = "/v1/access-token/verify"

	AddAccountSecretPath = "/v1/secret"
	GetAccountSecretPath = "/v1/secret"
	VerifyCodePath       = "/v1/secret/valid"
)

func readResult[T any](resp *http.Response) (*Result[T], error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	result := &Result[T]{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if result.Code != 200000 {
		return nil, errors.New(result.Msg)
	}
	return result, nil
}

func httpGetReq(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	token, err := GenAccessToken()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}
