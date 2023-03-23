package types

import (
	sdkerrors "github.com/hekas-network/cosmos-sdk/types/errors"
)

const StoreCodespace = "store"

var ErrInvalidProof = sdkerrors.Register(StoreCodespace, 2, "invalid proof")
