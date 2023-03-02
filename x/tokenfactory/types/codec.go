package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "cosmwasm/tokenfactory/create-denom", nil)
	cdc.RegisterConcrete(&MsgMint{}, "cosmwasm/tokenfactory/mint", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "cosmwasm/tokenfactory/burn", nil)
	// cdc.RegisterConcrete(&MsgForceTransfer{}, "cosmwasm/tokenfactory/force-transfer", nil)
	cdc.RegisterConcrete(&MsgChangeAdmin{}, "cosmwasm/tokenfactory/change-admin", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgMint{},
		&MsgBurn{},
		// &MsgForceTransfer{},
		&MsgChangeAdmin{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	// Note: these 3 are inlines from authz/codec in 0.46 so we can be compatible with 0.45
	sdk.RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	codec.RegisterEvidences(amino)

	amino.Seal()
}
