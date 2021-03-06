# Goでのテストサンプル

[Goのテストに入門してみよう！](https://future-architect.github.io/articles/20200601/#API%E3%82%B5%E3%83%BC%E3%83%90%E3%81%AB%E3%82%A2%E3%82%AF%E3%82%BB%E3%82%B9%E3%81%99%E3%82%8B%E3%83%86%E3%82%B9%E3%83%88%E3%82%92%E3%81%97%E3%81%9F%E3%81%84)

```
# 必要なパッケージを取得
❯ go mod tidy

# 全てのテストを実行
❯ go test -v
=== RUN   TestApiServer
--- PASS: TestApiServer (0.00s)
=== RUN   TestHelloHandler
--- PASS: TestHelloHandler (0.00s)
=== RUN   TestAddSimple
--- PASS: TestAddSimple (3.00s)
=== RUN   TestAdd
=== RUN   TestAdd/normal_1
=== PAUSE TestAdd/normal_1
=== RUN   TestAdd/normal_2
=== PAUSE TestAdd/normal_2
=== CONT  TestAdd/normal_1
=== CONT  TestAdd/normal_2
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/normal_2 (2.00s)
    --- PASS: TestAdd/normal_1 (3.00s)
=== RUN   TestMakeGatewayInfoDeepEqual
    struct_test.go:88: MakeGatewayInfo() got = {CoffeeShopWiFi 192.168.0.1 ffff0000 [{ristretto 192.168.0.116 0001-01-01 00:00:00 +0000 UTC} {aribica 192.168.0.104 2009-11-10 23:06:32 +0000 UTC} {macchiato 192.168.0.153 2009-11-10 23:39:43 +0000 UTC} {espresso 192.168.0.121 0001-01-01 00:00:00 +0000 UTC} {latte 192.168.0.219 2009-11-10 23:00:23 +0000 UTC} {americano 192.168.0.188 2009-11-10 23:03:05 +0000 UTC}]}, want {CoffeeShopWiFi 192.168.0.2 ffff0000 [{ristretto 192.168.0.116 0001-01-01 00:00:00 +0000 UTC} {aribica 192.168.0.104 2009-11-10 23:06:32 +0000 UTC} {macchiato 192.168.0.153 2009-11-10 23:39:43 +0000 UTC} {espresso 192.168.0.121 0001-01-01 00:00:00 +0000 UTC} {latte 192.168.0.221 2009-11-10 23:00:23 +0000 UTC}]}
--- FAIL: TestMakeGatewayInfoDeepEqual (0.00s)
=== RUN   TestMakeGatewayInfoGoCmp
    struct_test.go:97: MakeGatewayInfo() mismatch (-want +got):
          main.Gateway{
                SSID:      "CoffeeShopWiFi",
        -       IPAddress: s"192.168.0.2",
        +       IPAddress: s"192.168.0.1",
                NetMask:   {0xff, 0xff, 0x00, 0x00},
                Clients: []main.Client{
                        ... // 2 identical elements
                        {Hostname: "macchiato", IPAddress: s"192.168.0.153", LastSeen: s"2009-11-10 23:39:43 +0000 UTC"},
                        {Hostname: "espresso", IPAddress: s"192.168.0.121"},
                        {
                                Hostname:  "latte",
        -                       IPAddress: s"192.168.0.221",
        +                       IPAddress: s"192.168.0.219",
                                LastSeen:  s"2009-11-10 23:00:23 +0000 UTC",
                        },
        +               {
        +                       Hostname:  "americano",
        +                       IPAddress: s"192.168.0.188",
        +                       LastSeen:  s"2009-11-10 23:03:05 +0000 UTC",
        +               },
                },
          }
--- FAIL: TestMakeGatewayInfoGoCmp (0.00s)
FAIL
exit status 1
FAIL    go_test_sample  6.021s

# 特定のテスト，サブテストのみ実行する
❯ go test -v -run Add/mal_1
=== RUN   TestAddSimple
--- PASS: TestAddSimple (3.00s)
=== RUN   TestAdd
=== RUN   TestAdd/normal_1
=== PAUSE TestAdd/normal_1
=== CONT  TestAdd/normal_1
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/normal_1 (3.00s)
PASS
ok      go_test_sample  6.019s
```

## モックを利用したテスト

[gomockでGoのインターフェースのmockを作成してテストを実行する](https://www.asobou.co.jp/blog/web/gomock)

```
go get github.com/golang/mock/mockgen
cd mock_test
mockgen -source=sample.go -destination mock/mock_sample.go

# 生成したファイルを利用するようにテストコードを修正

# テストを実行
go test -v -run Mock/
=== RUN   TestWithSelfCreatedMock
--- PASS: TestWithSelfCreatedMock (0.00s)
=== RUN   TestWithGeneratedMock
--- PASS: TestWithGeneratedMock (0.00s)
PASS
ok      go_test_sample/mock_test        0.006s

```