## Dogethereum
Name:&emsp;&emsp;Dogether  
Algorithm:&emsp;&emsp;Ethash  
|Reward | blockhight<2100000 | blockhight>2100000 |
|-------|--------------------|--------------------|
| miner | 40000              | 2400               |
| devs  | 4000               | 0                  |


## Build
1. install golang 1.18.10
```shell
$ git clone https://github.com/dogethdev/dogethereum.git
$ cd dogethereum
$ go mod vendor
$ go build -v ./cmd/geth
```

## Start Dogether node

```shell
$ geth --syncmode full --datadir "./nodedata" --ethash.dagdir "./nodedata" --ethstats "your_node_name:dogether@state.dogether.dog"

```
