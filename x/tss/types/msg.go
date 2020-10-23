package types

import (
	"fmt"

	broadcast "github.com/axelarnetwork/axelar-core/x/broadcast/exported"
	tssd "github.com/axelarnetwork/tssd/pb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// golang stupidity: ensure interface compliance at compile time
var (
	_ sdk.Msg                = &MsgKeygenStart{}
	_ broadcast.ValidatorMsg = &MsgTSS{}
)

// MsgKeygenStart indicate the start of keygen
type MsgKeygenStart struct {
	Sender    sdk.AccAddress
	NewKeyID  string
	Threshold int
}

// MsgTSS protocol message for either keygen or sign
type MsgTSS struct {
	Sender    sdk.AccAddress
	SessionID string
	Payload   *tssd.MessageOut // TODO probably should not be a pointer; it's serialized by cosmos
}

// NewMsgKeygenStart TODO unnecessary method; delete it?
func NewMsgKeygenStart(newKeyID string, threshold int) MsgKeygenStart {
	return MsgKeygenStart{
		NewKeyID:  newKeyID,
		Threshold: threshold,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgKeygenStart) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
// naming convention follows x/staking/types/msg.go
func (msg MsgKeygenStart) Type() string { return "keygen_start" }

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgKeygenStart) ValidateBasic() error {
	if msg.Sender == nil {
		return sdkerrors.Wrap(ErrTss, "sender must be set")
	}
	if msg.NewKeyID == "" {
		return sdkerrors.Wrap(ErrTss, "new key id must be set")
	}
	if msg.Threshold < 1 {
		return sdkerrors.Wrap(ErrTss, fmt.Sprintf("invalid threshold [%d]", msg.Threshold))
	}
	// TODO enforce a maximum length for msg.SessionID?
	return nil
}

// GetSignBytes implements the sdk.Msg interface.
func (msg MsgKeygenStart) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements the sdk.Msg interface.
func (msg MsgKeygenStart) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// NewMsgTSS TODO unnecessary method; delete it?
func NewMsgTSS(sessionID string, payload *tssd.MessageOut) *MsgTSS {
	return &MsgTSS{
		SessionID: sessionID,
		Payload:   payload,
	}
}

// Route implements the sdk.Msg interface.
func (msg MsgTSS) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
// naming convention follows x/staking/types/msg.go
func (msg MsgTSS) Type() string { return "in" }

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgTSS) ValidateBasic() error {
	if msg.Sender == nil {
		return sdkerrors.Wrap(ErrTss, "sender must be set")
	}
	if msg.SessionID == "" {
		return sdkerrors.Wrap(ErrTss, "session id must be set")
	}
	if !msg.Payload.IsBroadcast && len(msg.Payload.ToPartyUid) == 0 {
		return sdkerrors.Wrap(ErrTss, "non-broadcast message must specify recipient")
	}
	if msg.Payload.IsBroadcast && len(msg.Payload.ToPartyUid) != 0 {
		return sdkerrors.Wrap(ErrTss, "broadcast message must not specify recipient")
	}
	// TODO enforce a maximum length for msg.SessionID?
	return nil
}

// GetSignBytes implements the sdk.Msg interface
func (msg MsgTSS) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements the sdk.Msg interface
func (msg MsgTSS) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// SetSender implements the broadcast.ValidatorMsg interface
func (msg *MsgTSS) SetSender(sender sdk.AccAddress) {
	msg.Sender = sender
}
