package types

import "github.com/cosmos/cosmos-sdk/codec"

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgOrgStore{}, "organizationStore/SetOrganization", nil)
	cdc.RegisterConcrete(OrgUsers{}, "organizationStore/SetOrganizationUser", nil)
	cdc.RegisterConcrete(MsgDeleteOrganization{}, "organizationStore/DeleteOrganization", nil)
	cdc.RegisterConcrete(MsgDeleteOrgUser{}, "organizationStore/DeleteOrganizationUser", nil)
}
