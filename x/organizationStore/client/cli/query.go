package cli

import (
	"github.com/spf13/cobra"
	"github.com/PrathyushaLakkireddy/sdk-tutorial/x/organizationStore/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"fmt"
	)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the organization module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	nameserviceQueryCmd.AddCommand(client.GetCommands(
		GetCmdOrganizations(storeKey, cdc),
		GetCmdOrganizationUsers(storeKey,cdc),
		GetCmdOrganizationDetails(storeKey,cdc),
	)...)
	return nameserviceQueryCmd
}


func GetCmdOrganizationDetails(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "org_details [orgName]",
		Short: "Get details of an organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			orgName := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/org_details/%s", queryRoute,orgName),nil)
			if err != nil {
				fmt.Printf("could not query organizations\n", err)
				return nil
			}

			fmt.Printf(string(res))

			//var out types.OrgUsers
			//cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(nil)
		},
	}
}

func GetCmdOrganizationUsers(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "org_users_list [orgName]",
		Short: "Userslist of an organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			orgName := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/org_users_list/%s", queryRoute,orgName),nil)
			if err != nil {
				fmt.Printf("could not query organizations\n", err)
				return nil
			}

			fmt.Printf(string(res))

			//var out types.OrgUsers
			//cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(nil)
		},
	}
}

// GetCmdOrgs queries a list of all orgs
func GetCmdOrganizations(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "orgs_list",
		Short: "Organization List",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/orgs_list", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query names\n")
				return nil
			}

			var out types.QueryResOrgs
			cdc.MustUnmarshalJSON(res, &out)

			return cliCtx.PrintOutput(out)
		},
	}
}