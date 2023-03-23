package client

import (
	govclient "github.com/hekas-network/cosmos-sdk/x/gov/client"
	"github.com/hekas-network/cosmos-sdk/x/upgrade/client/cli"
)

var (
	LegacyProposalHandler       = govclient.NewProposalHandler(cli.NewCmdSubmitLegacyUpgradeProposal)
	LegacyCancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitLegacyCancelUpgradeProposal)
)
