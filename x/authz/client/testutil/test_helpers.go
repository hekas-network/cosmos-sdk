package testutil

import (
	"github.com/hekas-network/cosmos-sdk/testutil"
	clitestutil "github.com/hekas-network/cosmos-sdk/testutil/cli"
	"github.com/hekas-network/cosmos-sdk/testutil/network"
	"github.com/hekas-network/cosmos-sdk/x/authz/client/cli"
)

func CreateGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
