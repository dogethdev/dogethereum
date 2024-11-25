package father

import (
	"bytes"
	"encoding/json"
	"fmt"
	badger "github.com/dgraph-io/badger/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	DB   *badger.DB
	Rpcs []string
)

func Init() {
	log.Info("init the ethereum classic dadabase....")
	var err error
	DB, err = badger.Open(badger.DefaultOptions("./etc/"))
	if err != nil {
		log.Info(err.Error())
		os.Exit(1)
	}
	fp, err := os.Open("./ethereum_classic_rpc_list.json")
	if err != nil {
		log.Info("ethereum classic syncing error", "msg", err.Error())
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	err = decoder.Decode(&Rpcs)
	if err != nil {
		log.Info("ethereum classic syncing error", "msg", err.Error())
		os.Exit(1)
	}
	fmt.Println(Rpcs)
}

func Sync() {
	reqStruct := `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0xfff", true],"id":1}`
	number := uint64(1)
	index := 0
	for true {
		index += 1
		sendData := strings.Replace(reqStruct, "0xfff", hexutil.Uint64(number).String(), -1)
		rcpSever := Rpcs[index%len(Rpcs)] //choose a target
		rsp, err := http.Post(rcpSever, "application/json", bytes.NewBufferString(sendData))
		if err != nil {
			log.Info("ethereum classic syncing error", "msg", err.Error())
			continue
		}
		if rsp.StatusCode != http.StatusOK {
			log.Info("ethereum classic syncing error", "HTTP StatusCode", rsp.StatusCode, "rcpSever", rcpSever)
			continue
		}
		body, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			log.Info("ethereum classic syncing error", "msg", err.Error(), rsp.StatusCode, "rcpSever", rcpSever)
			continue
		}
		var result map[string]interface{}
		err = json.Unmarshal([]byte(body), &result)
		if err != nil {
			log.Info("ethereum classic syncing error", "msg", err.Error(), rsp.StatusCode, "rcpSever", rcpSever)
			continue
		}
		blockInfo, exist := result["result"].(map[string]interface{})
		if !exist {
			log.Info("ethereum classic syncing error", "msg", err.Error(), rsp.StatusCode, "rcpSever", rcpSever)
			continue
		}

		err = DB.Update(func(txn *badger.Txn) error {
			err = txn.Set([]byte(strconv.FormatUint(number, 10)), body)
			return err
		})
		if err != nil {
			log.Info("ethereum classic syncing error", "msg", err.Error(), rsp.StatusCode, "rcpSever", rcpSever)
			continue
		}

		timestamp, err := hexutil.DecodeUint64(blockInfo["timestamp"].(string))
		log.Info("ethereum classic syncing progress", "number", number, "time", common.PrettyAge(time.Unix(int64(timestamp), 0)))
		number += 1
	}

}
