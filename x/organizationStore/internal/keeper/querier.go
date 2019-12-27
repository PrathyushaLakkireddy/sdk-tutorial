package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/PrathyushaLakkireddy/sdk-tutorial/x/organizationStore/internal/types"
	)

const (
	QueryOrganizations = "orgs_list"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryOrganizations:
			return queryNames(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

func queryNames(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	//whois := keeper.GetOrganization(ctx, path[0])
	//
	//res, err := codec.MarshalJSONIndent(keeper.cdc, whois)
	//if err != nil {
	//	panic("could not marshal result to JSON")
	//}
	//
	//return res, nil

	var namesList types.QueryResOrgs

	iterator := keeper.GetOrgsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		namesList = append(namesList, string(iterator.Key()))
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
