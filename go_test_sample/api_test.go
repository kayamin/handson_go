// httptest を用いてAPI サーバーのテストを行う

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer(t *testing.T) {
	// httptest を用いてサーバーを一時的に立ち上げる
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close() // テスト関数から return する際に実行するように設定する

	// 立ち上げたサーバーに対してリクエストをする
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	got, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	// API からのレスポンスを用いてなんならの処理をする

	if res.StatusCode != 200 {
		t.Errorf("GET %s: expected status code = %d; got %d", ts.URL, 200, res.StatusCode)
	}
	if string(got) != "Hello, client\n" {
		t.Errorf("expected body %v; got %v", "Hello, client", string(got))
	}
}

// リクエストのハンドラーの実装をテストしたい場合
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}

func TestHelloHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/dummy", nil)
	w := httptest.NewRecorder()

	helloHandler(w, r)

	resp := w.Result()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("cannot read test response: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("got = %d, want = 200", resp.StatusCode)
	}

	if string(body) != "hello world!" {
		t.Errorf("got = %s, want = hello world!", body)
	}
}
