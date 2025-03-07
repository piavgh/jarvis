package config

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/tranvictor/jarvis/accounts/types"
	"github.com/tranvictor/jarvis/networks"
)

var (
	Debug     bool = true
	DegenMode bool
)

var (
	cachedNetwork networks.Network
	mu            sync.Mutex
)

func Network() networks.Network {
	if cachedNetwork != nil {
		return cachedNetwork
	}

	err := SetNetwork(NetworkString)
	if err != nil {
		panic(err)
	}

	return cachedNetwork
}

func SetNetwork(networkStr string) error {
	mu.Lock()
	defer mu.Unlock()

	var err error
	var inited bool

	if cachedNetwork != nil {
		inited = true
	}

	cachedNetwork, err = networks.GetNetwork(networkStr)
	if err != nil {
		return err
	}

	if inited {
		fmt.Printf("Switched to network: %s\n", cachedNetwork.GetName())
	} else {
		fmt.Printf("Network: %s\n", cachedNetwork.GetName())
	}
	return nil
}

var NetworkString string

var (
	GasPrice      float64
	ExtraGasPrice float64
	TipGas        float64
	ExtraTipGas   float64
	GasLimit      uint64
	ExtraGasLimit uint64
	Nonce         uint64
	From          string
	FromAcc       types.AccDesc
	To            string
	Value         *big.Int
	RawValue      string
	MethodIndex   uint64
	PrefillMode   bool
	PrefillStr    string
	PrefillParams []string
	NoFuncCall    bool
	Tx            string
	TxType        uint8

	AllZeroParamsMethods bool
	AtBlock              int64

	MsigValue float64
	MsigTo    string

	DontBroadcast     bool
	DontWaitToBeMined bool
	ForceERC20ABI     bool
	RetryBroadcast    bool
	YesToAllPrompt    bool
	ForceLegacy       bool

	CustomABI      string
	JSONOutputFile string

	Simulate bool
)
