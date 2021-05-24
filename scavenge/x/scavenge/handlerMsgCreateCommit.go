package scavenge

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/taan02991/scavenge/x/scavenge/keeper"
	"github.com/taan02991/scavenge/x/scavenge/types"
)

func handleMsgCreateCommit(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateCommit) (*sdk.Result, error) {
	var commit = types.Commit{
		Scavenger:             msg.Creator,
		SolutionHash:          msg.SolutionHash,
		SolutionScavengerHash: msg.SolutionScavengerHash,
	}
	k.CreateCommit(ctx, commit)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
