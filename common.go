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
)

func readResult(resp *http.Response) (*Result, error) {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	result := &Result{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	if result.Code != 200000 {
		return nil, errors.New(result.Msg)
	}
	return result, nil
}
