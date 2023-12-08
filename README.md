## Dogethereum
Official Golang implementation of the Ethereum protocol.

## Build
1. install golang 1.18.10
```shell
$ cd dogethereum
$ go mod vendor
$ go build -v ./cmd/geth
```

## Start Dogether node
### Linux
```shell
$ geth --http --http.addr 0.0.0.0 --datadir "./nodedata" --ethash.dagdir "./nodedata"
```
Save above code to run.sh
#### run
```shell
$ ./run.sh
```

### Windwos
```shell
geth --http --http.addr 0.0.0.0 --datadir "./nodedata" --ethash.dagdir "./nodedata"
pause
```
Save above code to run.bat and click it
