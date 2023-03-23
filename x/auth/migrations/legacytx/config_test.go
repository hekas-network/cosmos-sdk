package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/hekas-network/cosmos-sdk/codec"
	cryptoAmino "github.com/hekas-network/cosmos-sdk/crypto/codec"
	"github.com/hekas-network/cosmos-sdk/testutil/testdata"
	sdk "github.com/hekas-network/cosmos-sdk/types"
	"github.com/hekas-network/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/hekas-network/cosmos-sdk/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
