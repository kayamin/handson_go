package main

import (
	"testing"
	"time"
)

func add(a, b int) int {
	time.Sleep(time.Duration(a+b) * time.Second)
	return a + b
}

// シンプルなテスト
func TestAddSimple(t *testing.T) {
	if add(1, 2) != 3 {
		t.Errorf("add() = %v, want %v", add(1, 2), 3)
	}
}

// Table Driven Test 形式 (入力と期待する出力を配列の形で与えでまとめてテストする）
func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal_1", args: args{a: 1, b: 2}, want: 3},
		{name: "normal_2", args: args{a: 0, b: 2}, want: 2},
	}
	for _, tt := range tests {
		tt := tt // Goではループで用いられる変数は同じアドレスを使うため，並列実行する場合は tt をループ内のローカル変数として定義したほうが良い

		t.Run(tt.name, func(t *testing.T) { // 与えている name はテストケース（サブテストと呼ぶ）の名称であり，-run オプションで指定することで実行するサブテストを絞り込める
			// 例： go test -v -run Add/mal_1
			t.Parallel() // テストの並列実行を設定するメソッド
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
