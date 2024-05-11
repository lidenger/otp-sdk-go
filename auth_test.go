package otpsdk

import "testing"

func conf() {
	address := "http://127.0.0.1:8066"
	serverSign := "server1"
	serverKey := "0c8441ba0ec011efbb1e2cf05daf3fe5"
	serverIV := "0c8441ba0ec011ef"
	Conf(serverSign, address, serverKey, serverIV)
}

func TestGenAccessToken(t *testing.T) {
	conf()
	token, err := GenAccessToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestVerifyAccessToken(t *testing.T) {
	conf()
	token := "2edefa52af1e848c56a2749d25c653ac8ca23818cf11a5ff9e4a9cc088b0b5ff8a413078138d519818be5fdf60275c5ffa92fa70cbfd3d1ce30bb61f7c496c4fafb0dc72acfa4b3031d4780275802544"
	err := VerifyAccessToken(token)
	if err != nil {
		t.Log("verify token fail")
	} else {
		t.Log("verify token success")
	}

	token2 := "xxx"
	err = VerifyAccessToken(token2)
	if err == nil {
		t.Fatal("和预期不一致，应该失败")
	} else {
		t.Log("符合预期")
	}
}
