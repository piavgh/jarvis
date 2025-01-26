package util

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/tranvictor/jarvis/config"
	jarvisnetworks "github.com/tranvictor/jarvis/networks"
	"github.com/tranvictor/jarvis/util/reader"
)

func ValidTxType(r *reader.EthReader, network jarvisnetworks.Network) (uint8, error) {
	if config.ForceLegacy {
		return types.LegacyTxType, nil
	}

	isDynamicFeeAvailable, err := r.CheckDynamicFeeTxAvailable()
	if err != nil {
		return 0, fmt.Errorf("couldn't check if the chain support dynamic fee: %w", err)
	}

	if !isDynamicFeeAvailable {
		return types.LegacyTxType, nil
	}

	return types.DynamicFeeTxType, nil
}
