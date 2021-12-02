package types

import (
	"fmt"

	"github.com/axelarnetwork/axelar-core/x/vote/exported"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewGenesisState is the constructor for GenesisState
func NewGenesisState(params Params, pollMetadatas []exported.PollMetadata) *GenesisState {
	return &GenesisState{
		Params:        params,
		PollMetadatas: pollMetadatas,
	}
}

// DefaultGenesisState represents the default genesis state
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultParams(), []exported.PollMetadata{})
}

// Validate validates the genesis state
func (m GenesisState) Validate() error {
	if err := m.Params.Validate(); err != nil {
		return getValidateError(err)
	}

	for _, pollMetadata := range m.PollMetadatas {
		if pollMetadata.Is(exported.Pending) {
			return getValidateError(fmt.Errorf("state of poll metadata %s is pending", pollMetadata.Key.String()))
		}

		if err := pollMetadata.Validate(); err != nil {
			return getValidateError(err)
		}
	}

	return nil
}

func getValidateError(err error) error {
	return sdkerrors.Wrapf(err, "genesis state for module %s is invalid", ModuleName)
}
