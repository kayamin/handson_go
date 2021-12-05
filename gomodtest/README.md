# Go言語ハンズオン


## 初期化と実行

```
# パッケージ gomodtest を初期化
❯ go mod init gomodtest
go: creating new go.mod: module gomodtest

# パッケージ名と goバージョン情報のみが記載されたファイルができる
❯ cat go.mod
module gomodtest

go 1.16

# パッケージをインストール
❯ go get -u go.uber.org/zap
go: downloading go.uber.org/zap v1.17.0
go: downloading go.uber.org/atomic v1.7.0
go: downloading go.uber.org/multierr v1.6.0
go: downloading go.uber.org/atomic v1.8.0
go: downloading go.uber.org/multierr v1.7.0
go get: added go.uber.org/atomic v1.8.0
go get: added go.uber.org/multierr v1.7.0
go get: added go.uber.org/zap v1.17.0

# このパッケージで利用するパッケージが追加されていることを確認
❯ cat go.mod
module gomodtest

go 1.16

require (
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
)

# 各パッケージのハッシュ値等を管理 したファイルも生成される
❯ cat go.sum
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/pkg/errors v0.8.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.3.0/go.mod h1:M5WIy9Dh21IEIfnGCwXGc5bZfKNJtfHm1UVUgZn+9EI=
github.com/stretchr/testify v1.7.0/go.mod h1:6Fq8oRcR53rry900zMqJjRRixrwX3KX962/h/Wwjteg=
go.uber.org/atomic v1.7.0/go.mod h1:fEN4uk6kAWBTFdckzkM89CLk9XfWZrxpCo0nPH17wJc=
go.uber.org/atomic v1.8.0 h1:CUhrE4N1rqSE6FM9ecihEjRkLQu8cDfgDyoOs83mEY4=
go.uber.org/atomic v1.8.0/go.mod h1:fEN4uk6kAWBTFdckzkM89CLk9XfWZrxpCo0nPH17wJc=
go.uber.org/multierr v1.6.0/go.mod h1:cdWPpRnG4AhwMwsgIHip0KRBQjJy5kYEpYjJxpXp9iU=
go.uber.org/multierr v1.7.0 h1:zaiO/rmgFjbmCXdSYJWQcdvOCsthmdaHfr3Gm2Kx4Ec=
go.uber.org/multierr v1.7.0/go.mod h1:7EAYxJLBy9rStEaz58O2t4Uvip6FSURkq8/ppBp95ak=
go.uber.org/zap v1.17.0 h1:MTjgFu6ZLKvY6Pvaqk97GlxNBuMpV4Hy/3P6tRGlI2U=
go.uber.org/zap v1.17.0/go.mod h1:MXVU+bhUf/A7Xi2HNOnopQOrmycQ5Ih87HtOu4q5SSo=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/yaml.v2 v2.2.8/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=
gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=

# goプログラムをビルド
❯ go build main.go

# ビルドしたプログラムを実行
❯ ./main
{"level":"warn","ts":1623457341.1851578,"caller":"gomodtest/main.go:7","msg":"warning test
```

## サーバーを立ち上げて動作を確認

```
❯ go run main.go
{"level":"warn","ts":1638668916.734483,"caller":"gomodtest/main.go:16","msg":"warning test"}
[/var/folders/nm/4ktmmyq576jdk5s2z5f9fmcm0000gp/T/go-build267984337/b001/exe/main]

# 別ターミナルで
❯ curl localhost:8080/
Hello, HTTPサーバ
```
