package simulation_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"
	"github.com/hekas-network/cosmos-sdk/simapp"
	"github.com/hekas-network/cosmos-sdk/types/module"
	simtypes "github.com/hekas-network/cosmos-sdk/types/simulation"
	"github.com/hekas-network/cosmos-sdk/x/authz"
	"github.com/hekas-network/cosmos-sdk/x/authz/simulation"
)

func TestRandomizedGenState(t *testing.T) {
	app := simapp.Setup(t, false)

	s := rand.NewSource(1)
	r := rand.New(s)

	simState := module.SimulationState{
		AppParams:    make(simtypes.AppParams),
		Cdc:          app.AppCodec(),
		Rand:         r,
		NumBonded:    3,
		Accounts:     simtypes.RandomAccounts(r, 3),
		InitialStake: sdkmath.NewInt(1000),
		GenState:     make(map[string]json.RawMessage),
	}

	simulation.RandomizedGenState(&simState)
	var authzGenesis authz.GenesisState
	simState.Cdc.MustUnmarshalJSON(simState.GenState[authz.ModuleName], &authzGenesis)

	require.Len(t, authzGenesis.Authorization, len(simState.Accounts)-1)
}
