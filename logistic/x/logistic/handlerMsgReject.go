package logistic

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/earth2378/logistic/x/logistic/keeper"
	"github.com/earth2378/logistic/x/logistic/types"
)

func handlerMsgReject(ctx sdk.Context, k keeper.Keeper, msg types.MsgReject) (*sdk.Result, error) {
	currentDeal, err := k.GetDeal(ctx, msg.OrderID)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal does not exists")
	}

	if msg.Customer.String() != currentDeal.Customer.String() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Fail receive, invalid Customer")
	}

	if currentDeal.State != types.InTransit {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid state")
	}

	if !currentDeal.Cancelable {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Deal can't be rejected")
	}

	currentDeal.State = types.Cancelled
	k.SetDeal(ctx, currentDeal)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeReceive),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Customer.String()),
			sdk.NewAttribute(types.AttributeOrderID, msg.OrderID),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
