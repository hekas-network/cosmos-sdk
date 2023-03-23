package types

import (
	"github.com/hekas-network/cosmos-sdk/codec"
	cryptocodec "github.com/hekas-network/cosmos-sdk/crypto/codec"
)

var amino = codec.NewLegacyAmino()

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
