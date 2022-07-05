package testutils

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/math"

	"github.com/axelarnetwork/axelar-core/testutils/rand"
	"github.com/axelarnetwork/axelar-core/utils"
	"github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	"github.com/axelarnetwork/utils/slices"
)

// Snapshot returns random snapshot based on the given parameters
func Snapshot(participantCount uint64, threshold utils.Threshold) (snapshot exported.Snapshot) {
	for {
		participantsWeight := sdk.ZeroUint()
		participants := slices.Expand(func(int) exported.Participant {
			weight := sdk.NewUint(uint64(rand.I64Between(1, 100)))
			participantsWeight = participantsWeight.Add(weight)

			return exported.NewParticipant(rand.ValAddr(), weight)
		},
			int(participantCount),
		)

		bondedWeight := sdk.NewUint(uint64(rand.I64Between(
			participantsWeight.BigInt().Int64(),
			math.MaxInt64(
				participantsWeight.BigInt().Int64()+1,
				participantsWeight.MulUint64(uint64(threshold.Denominator)).QuoUint64(uint64(threshold.Numerator)).BigInt().Int64()),
		),
		))

		snapshot = exported.NewSnapshot(time.Now(), rand.I64Between(1, 1000), participants, bondedWeight)
		if err := snapshot.ValidateBasic(); err == nil {
			return snapshot
		}
	}
}
