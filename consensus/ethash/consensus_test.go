package ethash

import (
	"fmt"
	"math/big"
	"testing"
)

func TestSuitableForDual(t *testing.T) {
	blockNum := big.NewInt(123456000)
	for {
		ret := isSuitableForDual(blockNum)

		fmt.Printf("blockNum:%s ,%d\n", blockNum, ret)
		blockNum.Add(blockNum, big.NewInt(1))
		if blockNum.Cmp(big.NewInt(123456000+1000)) == 1 {
			break
		}
	}

}
