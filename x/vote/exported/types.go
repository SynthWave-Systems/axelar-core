package exported

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
)

// VotingData is needed so that the amino codec can (un)marshal the voting data correctly
type VotingData interface {
	codec.ProtoMarshaler
}

// NewPollMeta constructor for PollMeta without nonce
func NewPollMeta(module string, id string) PollMeta {
	return PollMeta{
		Module: module,
		ID:     id,
	}
}

func (m PollMeta) String() string {
	return fmt.Sprintf("%s_%s", m.Module, m.ID)
}

// Validate performs a stateless validity check to ensure PollMeta has been properly initialized
func (m PollMeta) Validate() error {
	if m.Module == "" {
		return fmt.Errorf("missing module")
	}

	if m.ID == "" {
		return fmt.Errorf("missing poll ID")
	}

	return nil
}
