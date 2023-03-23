package nft

import (
	types "github.com/hekas-network/cosmos-sdk/codec/types"
	sdk "github.com/hekas-network/cosmos-sdk/types"
	"github.com/hekas-network/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
