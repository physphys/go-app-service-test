## 概要

DDDの勉強用

## 図の作成

パッケージ依存関係図の作成

```shell
go install github.com/kisielk/godepgraph@latest
godepgraph -s go-app-service-test/application/usecase | dot -Tpng -o package_dependency.png 
```

パッケージ構成図の作成

```shell
go get github.com/jfeliu007/goplantuml/parser
go install github.com/jfeliu007/goplantuml/cmd/goplantuml@latest
goplantuml -recursive application domain > packages.puml
```

## 静的解析

LCOMの計測

```shell
go install github.com/yahoojapan/lcom4go/cmd/lcom4@latest
go vet -vettool=`which lcom4` ./...
```

golangcilint

```shell
golangci-lint run ./...
```


