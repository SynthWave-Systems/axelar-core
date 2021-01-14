package keeper

import (
	"testing"

	tssd "github.com/axelarnetwork/tssd/pb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/axelarnetwork/axelar-core/x/tss/types"
)

func TestKeeper_StartSign_IdAlreadyInUse_ReturnError(t *testing.T) {
	s := setup(t)
	msg := types.MsgSignStart{
		Sender:    sdk.AccAddress(s.Broadcaster.LocalPrincipal),
		SigID:     "sigID",
		KeyID:     "keyID1",
		MsgToSign: []byte("message"),
	}

	_, err := s.Keeper.StartSign(s.Ctx, msg, validators)
	assert.NoError(t, err)

	msg.KeyID = "keyID2"
	msg.MsgToSign = []byte("second message")
	_, err = s.Keeper.StartSign(s.Ctx, msg, validators)
	assert.Error(t, err)
}

// Even if no session exists the keeper must not return an error, because we need to keep validators and
// non-participating nodes consistent (for non-participating nodes there should be no session)
func TestKeeper_SignMsg_NoSessionWithGivenID_Return(t *testing.T) {
	s := setup(t)

	assert.NoError(t, s.Keeper.SignMsg(s.Ctx, types.MsgSignTraffic{
		Sender:    s.Broadcaster.GetProxy(s.Ctx, s.Broadcaster.LocalPrincipal),
		SessionID: "sigID",
		Payload:   &tssd.TrafficOut{},
	}))
}