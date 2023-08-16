# Go Bank

This repo is just to learn the basics of GO lang with testing and to build web api

## Create/Init Module

```shell
go mod init github.com/gunjanpatel/go-bank
go mod tidy
```

## Refer Local package (while working on dev)

Just add following line into go.mod to the practice module

```go
require github.com/gunjanpatel/go-bank v0.0.0

replace github.com/gunjanpatel/go-bank => ../go-bank
```

## Publish Package

```shell
git tag v0.1.0
git push origin v0.1.0
```

## Use published package

go get github.com/gunjanpatel/go-bank
