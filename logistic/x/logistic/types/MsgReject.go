package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgReject{}

type MsgReject struct {
	Customer sdk.AccAddress `json:"customer" yaml:"customer"` // address of the validator operator
	OrderID  string         `json:"orderid" yaml:"orderid"`
}

func NewMsgReject(customer sdk.AccAddress, orderid string) MsgReject {
	return MsgReject{
		Customer: customer,
		OrderID:  orderid,
	}
}

const RejectConst = "Reject"

// nolint
func (msg MsgReject) Route() string { return RouterKey }
func (msg MsgReject) Type() string  { return RejectConst }
func (msg MsgReject) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Customer)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgReject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgReject) ValidateBasic() error {
	if msg.Customer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing customer address")
	}
	return nil
}
