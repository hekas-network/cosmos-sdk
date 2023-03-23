package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/hekas-network/cosmos-sdk/simapp"
	sdk "github.com/hekas-network/cosmos-sdk/types"
	"github.com/hekas-network/cosmos-sdk/types/kv"
	"github.com/hekas-network/cosmos-sdk/x/authz"
	"github.com/hekas-network/cosmos-sdk/x/authz/keeper"
	"github.com/hekas-network/cosmos-sdk/x/authz/simulation"
	banktypes "github.com/hekas-network/cosmos-sdk/x/bank/types"
)

func TestDecodeStore(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Codec
	dec := simulation.NewDecodeStore(cdc)

	now := time.Now().UTC()
	e := now.Add(1)
	grant, _ := authz.NewGrant(now, banktypes.NewSendAuthorization(sdk.NewCoins(sdk.NewInt64Coin("foo", 123))), &e)
	grantBz, err := cdc.Marshal(&grant)
	require.NoError(t, err)
	kvPairs := kv.Pairs{
		Pairs: []kv.Pair{
			{Key: []byte(keeper.GrantKey), Value: grantBz},
			{Key: []byte{0x99}, Value: []byte{0x99}},
		},
	}

	tests := []struct {
		name        string
		expectErr   bool
		expectedLog string
	}{
		{"Grant", false, fmt.Sprintf("%v\n%v", grant, grant)},
		{"other", true, ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectErr {
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			} else {
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
