package types_test

import (
	"github.com/hekas-network/cosmos-sdk/simapp"
)

var (
	ecdc                  = simapp.MakeTestEncodingConfig()
	appCodec, legacyAmino = ecdc.Codec, ecdc.Amino
)
