package main

import (
	"testing"

	mock_main "go_test_sample/mock_test/mock"

	"github.com/golang/mock/gomock"
)

/////////////////////////
// mockを自作する場合
/////////////////////////
type ApiClientMock struct{}

func (a *ApiClientMock) Request(data string) (string, error) {
	return data, nil
}

func TestWithSelfCreatedMock(t *testing.T) {
	d := &DataRegister{}
	d.client = &ApiClientMock{} // mockを登録し利用する
	expected := "bar"
	res, err := d.Register(expected)
	if err != nil {
		t.Fatal("Register error!", err)
	}
	if res != expected {
		t.Fatal("Value does not match.")
	}
}

/////////////////////////////////////////////////
// mockgen によって生成された mock のコードを用いる場合
/////////////////////////////////////////////////
func TestWithGeneratedMock(t *testing.T) {
	// mockのコントローラを作成します
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// ApiClientインターフェイスのmockを作成する
	mockApiClient := mock_main.NewMockApiClient(ctrl)

	// 作成したmockに対して期待する呼び出しと返り値を定義する
	mockApiClient.EXPECT(). // EXPECT()では呼び出されたかどうか
				Request("bar").    // Request()ではそのメソッド名が指定した引数で呼び出されたかどうか
				Return("bar", nil) // Return()では返り値を指定します

	d := &DataRegister{}
	d.client = mockApiClient // mockを登録
	expected := "bar"

	res, err := d.Register(expected)
	if err != nil {
		t.Fatal("Register error!", err)
	}
	if res != expected {
		t.Fatal("Value does not match.")
	}
}
