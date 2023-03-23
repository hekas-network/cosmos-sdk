package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/hekas-network/cosmos-sdk/codec"
	codectypes "github.com/hekas-network/cosmos-sdk/codec/types"
	"github.com/hekas-network/cosmos-sdk/std"
	"github.com/hekas-network/cosmos-sdk/testutil/testdata"
	sdk "github.com/hekas-network/cosmos-sdk/types"
	"github.com/hekas-network/cosmos-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
