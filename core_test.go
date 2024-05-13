package otpsdk

import "testing"

func TestAddAccountSecret(t *testing.T) {
	err := AddAccountSecret("liweiyi2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAccountSecret(t *testing.T) {
	secret, err := GetAccountSecret("liweiyi")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", secret)
}

func TestVerifyCode(t *testing.T) {
	success, err := VerifyCode("liweiyi2", "431775")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("验证结果: %+v", success)
}

func TestGetQRCodeContent(t *testing.T) {
	content, err := GetQRCodeContent("liweiyi2")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("QR content: %+v", content)
}
