package staking_test

import (
	"math/big"
	"testing"

	"cosmossdk.io/math"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/hekas-network/cosmos-sdk/codec"
	"github.com/hekas-network/cosmos-sdk/crypto/keys/ed25519"
	"github.com/hekas-network/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/hekas-network/cosmos-sdk/simapp"
	sdk "github.com/hekas-network/cosmos-sdk/types"
	"github.com/hekas-network/cosmos-sdk/x/staking/keeper"
	"github.com/hekas-network/cosmos-sdk/x/staking/types"
)

func init() {
	sdk.DefaultPowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

// nolint:deadcode,unused,varcheck
var (
	priv1 = secp256k1.GenPrivKey()
	addr1 = sdk.AccAddress(priv1.PubKey().Address())
	priv2 = secp256k1.GenPrivKey()
	addr2 = sdk.AccAddress(priv2.PubKey().Address())

	valKey  = ed25519.GenPrivKey()
	valAddr = sdk.AccAddress(valKey.PubKey().Address())

	commissionRates = types.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())

	PKs = simapp.CreateTestPubKeys(500)
)

// getBaseSimappWithCustomKeeper Returns a simapp with custom StakingKeeper
// to avoid messing with the hooks.
func getBaseSimappWithCustomKeeper(t *testing.T) (*codec.LegacyAmino, *simapp.SimApp, sdk.Context) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	appCodec := app.AppCodec()

	app.StakingKeeper = keeper.NewKeeper(
		appCodec,
		app.GetKey(types.StoreKey),
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(types.ModuleName),
	)
	app.StakingKeeper.SetParams(ctx, types.DefaultParams())

	return codec.NewLegacyAmino(), app, ctx
}

// generateAddresses generates numAddrs of normal AccAddrs and ValAddrs
func generateAddresses(app *simapp.SimApp, ctx sdk.Context, numAddrs int, accAmount math.Int) ([]sdk.AccAddress, []sdk.ValAddress) {
	addrDels := simapp.AddTestAddrsIncremental(app, ctx, numAddrs, accAmount)
	addrVals := simapp.ConvertAddrsToValAddrs(addrDels)

	return addrDels, addrVals
}
