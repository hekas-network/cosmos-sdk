package codec

import (
	"github.com/hekas-network/cosmos-sdk/codec"
	cryptocodec "github.com/hekas-network/cosmos-sdk/crypto/codec"
	sdk "github.com/hekas-network/cosmos-sdk/types"
)

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	cryptocodec.RegisterCrypto(Amino)
	codec.RegisterEvidences(Amino)
	sdk.RegisterLegacyAminoCodec(Amino)
}
