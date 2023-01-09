// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/utils"
	multisigexported "github.com/axelarnetwork/axelar-core/x/multisig/exported"
	snapshotexported "github.com/axelarnetwork/axelar-core/x/snapshot/exported"
	"github.com/btcsuite/btcd/btcec"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
	"time"
)

// Ensure, that SigHandlerMock does implement multisigexported.SigHandler.
// If this is not the case, regenerate this file with moq.
var _ multisigexported.SigHandler = &SigHandlerMock{}

// SigHandlerMock is a mock implementation of multisigexported.SigHandler.
//
//	func TestSomethingThatUsesSigHandler(t *testing.T) {
//
//		// make and configure a mocked multisigexported.SigHandler
//		mockedSigHandler := &SigHandlerMock{
//			HandleCompletedFunc: func(ctx sdk.Context, sig utils.ValidatedProtoMarshaler, moduleMetadata codec.ProtoMarshaler) error {
//				panic("mock out the HandleCompleted method")
//			},
//			HandleFailedFunc: func(ctx sdk.Context, moduleMetadata codec.ProtoMarshaler) error {
//				panic("mock out the HandleFailed method")
//			},
//		}
//
//		// use mockedSigHandler in code that requires multisigexported.SigHandler
//		// and then make assertions.
//
//	}
type SigHandlerMock struct {
	// HandleCompletedFunc mocks the HandleCompleted method.
	HandleCompletedFunc func(ctx sdk.Context, sig utils.ValidatedProtoMarshaler, moduleMetadata codec.ProtoMarshaler) error

	// HandleFailedFunc mocks the HandleFailed method.
	HandleFailedFunc func(ctx sdk.Context, moduleMetadata codec.ProtoMarshaler) error

	// calls tracks calls to the methods.
	calls struct {
		// HandleCompleted holds details about calls to the HandleCompleted method.
		HandleCompleted []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// Sig is the sig argument value.
			Sig utils.ValidatedProtoMarshaler
			// ModuleMetadata is the moduleMetadata argument value.
			ModuleMetadata codec.ProtoMarshaler
		}
		// HandleFailed holds details about calls to the HandleFailed method.
		HandleFailed []struct {
			// Ctx is the ctx argument value.
			Ctx sdk.Context
			// ModuleMetadata is the moduleMetadata argument value.
			ModuleMetadata codec.ProtoMarshaler
		}
	}
	lockHandleCompleted sync.RWMutex
	lockHandleFailed    sync.RWMutex
}

// HandleCompleted calls HandleCompletedFunc.
func (mock *SigHandlerMock) HandleCompleted(ctx sdk.Context, sig utils.ValidatedProtoMarshaler, moduleMetadata codec.ProtoMarshaler) error {
	if mock.HandleCompletedFunc == nil {
		panic("SigHandlerMock.HandleCompletedFunc: method is nil but SigHandler.HandleCompleted was just called")
	}
	callInfo := struct {
		Ctx            sdk.Context
		Sig            utils.ValidatedProtoMarshaler
		ModuleMetadata codec.ProtoMarshaler
	}{
		Ctx:            ctx,
		Sig:            sig,
		ModuleMetadata: moduleMetadata,
	}
	mock.lockHandleCompleted.Lock()
	mock.calls.HandleCompleted = append(mock.calls.HandleCompleted, callInfo)
	mock.lockHandleCompleted.Unlock()
	return mock.HandleCompletedFunc(ctx, sig, moduleMetadata)
}

// HandleCompletedCalls gets all the calls that were made to HandleCompleted.
// Check the length with:
//
//	len(mockedSigHandler.HandleCompletedCalls())
func (mock *SigHandlerMock) HandleCompletedCalls() []struct {
	Ctx            sdk.Context
	Sig            utils.ValidatedProtoMarshaler
	ModuleMetadata codec.ProtoMarshaler
} {
	var calls []struct {
		Ctx            sdk.Context
		Sig            utils.ValidatedProtoMarshaler
		ModuleMetadata codec.ProtoMarshaler
	}
	mock.lockHandleCompleted.RLock()
	calls = mock.calls.HandleCompleted
	mock.lockHandleCompleted.RUnlock()
	return calls
}

// HandleFailed calls HandleFailedFunc.
func (mock *SigHandlerMock) HandleFailed(ctx sdk.Context, moduleMetadata codec.ProtoMarshaler) error {
	if mock.HandleFailedFunc == nil {
		panic("SigHandlerMock.HandleFailedFunc: method is nil but SigHandler.HandleFailed was just called")
	}
	callInfo := struct {
		Ctx            sdk.Context
		ModuleMetadata codec.ProtoMarshaler
	}{
		Ctx:            ctx,
		ModuleMetadata: moduleMetadata,
	}
	mock.lockHandleFailed.Lock()
	mock.calls.HandleFailed = append(mock.calls.HandleFailed, callInfo)
	mock.lockHandleFailed.Unlock()
	return mock.HandleFailedFunc(ctx, moduleMetadata)
}

// HandleFailedCalls gets all the calls that were made to HandleFailed.
// Check the length with:
//
//	len(mockedSigHandler.HandleFailedCalls())
func (mock *SigHandlerMock) HandleFailedCalls() []struct {
	Ctx            sdk.Context
	ModuleMetadata codec.ProtoMarshaler
} {
	var calls []struct {
		Ctx            sdk.Context
		ModuleMetadata codec.ProtoMarshaler
	}
	mock.lockHandleFailed.RLock()
	calls = mock.calls.HandleFailed
	mock.lockHandleFailed.RUnlock()
	return calls
}

// Ensure, that KeyMock does implement multisigexported.Key.
// If this is not the case, regenerate this file with moq.
var _ multisigexported.Key = &KeyMock{}

// KeyMock is a mock implementation of multisigexported.Key.
//
//	func TestSomethingThatUsesKey(t *testing.T) {
//
//		// make and configure a mocked multisigexported.Key
//		mockedKey := &KeyMock{
//			GetBondedWeightFunc: func() sdk.Uint {
//				panic("mock out the GetBondedWeight method")
//			},
//			GetHeightFunc: func() int64 {
//				panic("mock out the GetHeight method")
//			},
//			GetMinPassingWeightFunc: func() sdk.Uint {
//				panic("mock out the GetMinPassingWeight method")
//			},
//			GetParticipantsFunc: func() []sdk.ValAddress {
//				panic("mock out the GetParticipants method")
//			},
//			GetPubKeyFunc: func(valAddress sdk.ValAddress) (multisigexported.PublicKey, bool) {
//				panic("mock out the GetPubKey method")
//			},
//			GetSnapshotFunc: func() snapshotexported.Snapshot {
//				panic("mock out the GetSnapshot method")
//			},
//			GetStateFunc: func() multisigexported.KeyState {
//				panic("mock out the GetState method")
//			},
//			GetTimestampFunc: func() time.Time {
//				panic("mock out the GetTimestamp method")
//			},
//			GetWeightFunc: func(valAddress sdk.ValAddress) sdk.Uint {
//				panic("mock out the GetWeight method")
//			},
//		}
//
//		// use mockedKey in code that requires multisigexported.Key
//		// and then make assertions.
//
//	}
type KeyMock struct {
	// GetBondedWeightFunc mocks the GetBondedWeight method.
	GetBondedWeightFunc func() sdk.Uint

	// GetHeightFunc mocks the GetHeight method.
	GetHeightFunc func() int64

	// GetMinPassingWeightFunc mocks the GetMinPassingWeight method.
	GetMinPassingWeightFunc func() sdk.Uint

	// GetParticipantsFunc mocks the GetParticipants method.
	GetParticipantsFunc func() []sdk.ValAddress

	// GetPubKeyFunc mocks the GetPubKey method.
	GetPubKeyFunc func(valAddress sdk.ValAddress) (multisigexported.PublicKey, bool)

	// GetSnapshotFunc mocks the GetSnapshot method.
	GetSnapshotFunc func() snapshotexported.Snapshot

	// GetStateFunc mocks the GetState method.
	GetStateFunc func() multisigexported.KeyState

	// GetTimestampFunc mocks the GetTimestamp method.
	GetTimestampFunc func() time.Time

	// GetWeightFunc mocks the GetWeight method.
	GetWeightFunc func(valAddress sdk.ValAddress) sdk.Uint

	// calls tracks calls to the methods.
	calls struct {
		// GetBondedWeight holds details about calls to the GetBondedWeight method.
		GetBondedWeight []struct {
		}
		// GetHeight holds details about calls to the GetHeight method.
		GetHeight []struct {
		}
		// GetMinPassingWeight holds details about calls to the GetMinPassingWeight method.
		GetMinPassingWeight []struct {
		}
		// GetParticipants holds details about calls to the GetParticipants method.
		GetParticipants []struct {
		}
		// GetPubKey holds details about calls to the GetPubKey method.
		GetPubKey []struct {
			// ValAddress is the valAddress argument value.
			ValAddress sdk.ValAddress
		}
		// GetSnapshot holds details about calls to the GetSnapshot method.
		GetSnapshot []struct {
		}
		// GetState holds details about calls to the GetState method.
		GetState []struct {
		}
		// GetTimestamp holds details about calls to the GetTimestamp method.
		GetTimestamp []struct {
		}
		// GetWeight holds details about calls to the GetWeight method.
		GetWeight []struct {
			// ValAddress is the valAddress argument value.
			ValAddress sdk.ValAddress
		}
	}
	lockGetBondedWeight     sync.RWMutex
	lockGetHeight           sync.RWMutex
	lockGetMinPassingWeight sync.RWMutex
	lockGetParticipants     sync.RWMutex
	lockGetPubKey           sync.RWMutex
	lockGetSnapshot         sync.RWMutex
	lockGetState            sync.RWMutex
	lockGetTimestamp        sync.RWMutex
	lockGetWeight           sync.RWMutex
}

// GetBondedWeight calls GetBondedWeightFunc.
func (mock *KeyMock) GetBondedWeight() sdk.Uint {
	if mock.GetBondedWeightFunc == nil {
		panic("KeyMock.GetBondedWeightFunc: method is nil but Key.GetBondedWeight was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetBondedWeight.Lock()
	mock.calls.GetBondedWeight = append(mock.calls.GetBondedWeight, callInfo)
	mock.lockGetBondedWeight.Unlock()
	return mock.GetBondedWeightFunc()
}

// GetBondedWeightCalls gets all the calls that were made to GetBondedWeight.
// Check the length with:
//
//	len(mockedKey.GetBondedWeightCalls())
func (mock *KeyMock) GetBondedWeightCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetBondedWeight.RLock()
	calls = mock.calls.GetBondedWeight
	mock.lockGetBondedWeight.RUnlock()
	return calls
}

// GetHeight calls GetHeightFunc.
func (mock *KeyMock) GetHeight() int64 {
	if mock.GetHeightFunc == nil {
		panic("KeyMock.GetHeightFunc: method is nil but Key.GetHeight was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetHeight.Lock()
	mock.calls.GetHeight = append(mock.calls.GetHeight, callInfo)
	mock.lockGetHeight.Unlock()
	return mock.GetHeightFunc()
}

// GetHeightCalls gets all the calls that were made to GetHeight.
// Check the length with:
//
//	len(mockedKey.GetHeightCalls())
func (mock *KeyMock) GetHeightCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetHeight.RLock()
	calls = mock.calls.GetHeight
	mock.lockGetHeight.RUnlock()
	return calls
}

// GetMinPassingWeight calls GetMinPassingWeightFunc.
func (mock *KeyMock) GetMinPassingWeight() sdk.Uint {
	if mock.GetMinPassingWeightFunc == nil {
		panic("KeyMock.GetMinPassingWeightFunc: method is nil but Key.GetMinPassingWeight was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetMinPassingWeight.Lock()
	mock.calls.GetMinPassingWeight = append(mock.calls.GetMinPassingWeight, callInfo)
	mock.lockGetMinPassingWeight.Unlock()
	return mock.GetMinPassingWeightFunc()
}

// GetMinPassingWeightCalls gets all the calls that were made to GetMinPassingWeight.
// Check the length with:
//
//	len(mockedKey.GetMinPassingWeightCalls())
func (mock *KeyMock) GetMinPassingWeightCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetMinPassingWeight.RLock()
	calls = mock.calls.GetMinPassingWeight
	mock.lockGetMinPassingWeight.RUnlock()
	return calls
}

// GetParticipants calls GetParticipantsFunc.
func (mock *KeyMock) GetParticipants() []sdk.ValAddress {
	if mock.GetParticipantsFunc == nil {
		panic("KeyMock.GetParticipantsFunc: method is nil but Key.GetParticipants was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetParticipants.Lock()
	mock.calls.GetParticipants = append(mock.calls.GetParticipants, callInfo)
	mock.lockGetParticipants.Unlock()
	return mock.GetParticipantsFunc()
}

// GetParticipantsCalls gets all the calls that were made to GetParticipants.
// Check the length with:
//
//	len(mockedKey.GetParticipantsCalls())
func (mock *KeyMock) GetParticipantsCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetParticipants.RLock()
	calls = mock.calls.GetParticipants
	mock.lockGetParticipants.RUnlock()
	return calls
}

// GetPubKey calls GetPubKeyFunc.
func (mock *KeyMock) GetPubKey(valAddress sdk.ValAddress) (multisigexported.PublicKey, bool) {
	if mock.GetPubKeyFunc == nil {
		panic("KeyMock.GetPubKeyFunc: method is nil but Key.GetPubKey was just called")
	}
	callInfo := struct {
		ValAddress sdk.ValAddress
	}{
		ValAddress: valAddress,
	}
	mock.lockGetPubKey.Lock()
	mock.calls.GetPubKey = append(mock.calls.GetPubKey, callInfo)
	mock.lockGetPubKey.Unlock()
	return mock.GetPubKeyFunc(valAddress)
}

// GetPubKeyCalls gets all the calls that were made to GetPubKey.
// Check the length with:
//
//	len(mockedKey.GetPubKeyCalls())
func (mock *KeyMock) GetPubKeyCalls() []struct {
	ValAddress sdk.ValAddress
} {
	var calls []struct {
		ValAddress sdk.ValAddress
	}
	mock.lockGetPubKey.RLock()
	calls = mock.calls.GetPubKey
	mock.lockGetPubKey.RUnlock()
	return calls
}

// GetSnapshot calls GetSnapshotFunc.
func (mock *KeyMock) GetSnapshot() snapshotexported.Snapshot {
	if mock.GetSnapshotFunc == nil {
		panic("KeyMock.GetSnapshotFunc: method is nil but Key.GetSnapshot was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetSnapshot.Lock()
	mock.calls.GetSnapshot = append(mock.calls.GetSnapshot, callInfo)
	mock.lockGetSnapshot.Unlock()
	return mock.GetSnapshotFunc()
}

// GetSnapshotCalls gets all the calls that were made to GetSnapshot.
// Check the length with:
//
//	len(mockedKey.GetSnapshotCalls())
func (mock *KeyMock) GetSnapshotCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetSnapshot.RLock()
	calls = mock.calls.GetSnapshot
	mock.lockGetSnapshot.RUnlock()
	return calls
}

// GetState calls GetStateFunc.
func (mock *KeyMock) GetState() multisigexported.KeyState {
	if mock.GetStateFunc == nil {
		panic("KeyMock.GetStateFunc: method is nil but Key.GetState was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetState.Lock()
	mock.calls.GetState = append(mock.calls.GetState, callInfo)
	mock.lockGetState.Unlock()
	return mock.GetStateFunc()
}

// GetStateCalls gets all the calls that were made to GetState.
// Check the length with:
//
//	len(mockedKey.GetStateCalls())
func (mock *KeyMock) GetStateCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetState.RLock()
	calls = mock.calls.GetState
	mock.lockGetState.RUnlock()
	return calls
}

// GetTimestamp calls GetTimestampFunc.
func (mock *KeyMock) GetTimestamp() time.Time {
	if mock.GetTimestampFunc == nil {
		panic("KeyMock.GetTimestampFunc: method is nil but Key.GetTimestamp was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetTimestamp.Lock()
	mock.calls.GetTimestamp = append(mock.calls.GetTimestamp, callInfo)
	mock.lockGetTimestamp.Unlock()
	return mock.GetTimestampFunc()
}

// GetTimestampCalls gets all the calls that were made to GetTimestamp.
// Check the length with:
//
//	len(mockedKey.GetTimestampCalls())
func (mock *KeyMock) GetTimestampCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetTimestamp.RLock()
	calls = mock.calls.GetTimestamp
	mock.lockGetTimestamp.RUnlock()
	return calls
}

// GetWeight calls GetWeightFunc.
func (mock *KeyMock) GetWeight(valAddress sdk.ValAddress) sdk.Uint {
	if mock.GetWeightFunc == nil {
		panic("KeyMock.GetWeightFunc: method is nil but Key.GetWeight was just called")
	}
	callInfo := struct {
		ValAddress sdk.ValAddress
	}{
		ValAddress: valAddress,
	}
	mock.lockGetWeight.Lock()
	mock.calls.GetWeight = append(mock.calls.GetWeight, callInfo)
	mock.lockGetWeight.Unlock()
	return mock.GetWeightFunc(valAddress)
}

// GetWeightCalls gets all the calls that were made to GetWeight.
// Check the length with:
//
//	len(mockedKey.GetWeightCalls())
func (mock *KeyMock) GetWeightCalls() []struct {
	ValAddress sdk.ValAddress
} {
	var calls []struct {
		ValAddress sdk.ValAddress
	}
	mock.lockGetWeight.RLock()
	calls = mock.calls.GetWeight
	mock.lockGetWeight.RUnlock()
	return calls
}

// Ensure, that MultiSigMock does implement multisigexported.MultiSig.
// If this is not the case, regenerate this file with moq.
var _ multisigexported.MultiSig = &MultiSigMock{}

// MultiSigMock is a mock implementation of multisigexported.MultiSig.
//
//	func TestSomethingThatUsesMultiSig(t *testing.T) {
//
//		// make and configure a mocked multisigexported.MultiSig
//		mockedMultiSig := &MultiSigMock{
//			GetKeyIDFunc: func() multisigexported.KeyID {
//				panic("mock out the GetKeyID method")
//			},
//			GetPayloadHashFunc: func() multisigexported.Hash {
//				panic("mock out the GetPayloadHash method")
//			},
//			GetSignatureFunc: func(p sdk.ValAddress) (btcec.Signature, bool) {
//				panic("mock out the GetSignature method")
//			},
//			ValidateBasicFunc: func() error {
//				panic("mock out the ValidateBasic method")
//			},
//		}
//
//		// use mockedMultiSig in code that requires multisigexported.MultiSig
//		// and then make assertions.
//
//	}
type MultiSigMock struct {
	// GetKeyIDFunc mocks the GetKeyID method.
	GetKeyIDFunc func() multisigexported.KeyID

	// GetPayloadHashFunc mocks the GetPayloadHash method.
	GetPayloadHashFunc func() multisigexported.Hash

	// GetSignatureFunc mocks the GetSignature method.
	GetSignatureFunc func(p sdk.ValAddress) (btcec.Signature, bool)

	// ValidateBasicFunc mocks the ValidateBasic method.
	ValidateBasicFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// GetKeyID holds details about calls to the GetKeyID method.
		GetKeyID []struct {
		}
		// GetPayloadHash holds details about calls to the GetPayloadHash method.
		GetPayloadHash []struct {
		}
		// GetSignature holds details about calls to the GetSignature method.
		GetSignature []struct {
			// P is the p argument value.
			P sdk.ValAddress
		}
		// ValidateBasic holds details about calls to the ValidateBasic method.
		ValidateBasic []struct {
		}
	}
	lockGetKeyID       sync.RWMutex
	lockGetPayloadHash sync.RWMutex
	lockGetSignature   sync.RWMutex
	lockValidateBasic  sync.RWMutex
}

// GetKeyID calls GetKeyIDFunc.
func (mock *MultiSigMock) GetKeyID() multisigexported.KeyID {
	if mock.GetKeyIDFunc == nil {
		panic("MultiSigMock.GetKeyIDFunc: method is nil but MultiSig.GetKeyID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetKeyID.Lock()
	mock.calls.GetKeyID = append(mock.calls.GetKeyID, callInfo)
	mock.lockGetKeyID.Unlock()
	return mock.GetKeyIDFunc()
}

// GetKeyIDCalls gets all the calls that were made to GetKeyID.
// Check the length with:
//
//	len(mockedMultiSig.GetKeyIDCalls())
func (mock *MultiSigMock) GetKeyIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetKeyID.RLock()
	calls = mock.calls.GetKeyID
	mock.lockGetKeyID.RUnlock()
	return calls
}

// GetPayloadHash calls GetPayloadHashFunc.
func (mock *MultiSigMock) GetPayloadHash() multisigexported.Hash {
	if mock.GetPayloadHashFunc == nil {
		panic("MultiSigMock.GetPayloadHashFunc: method is nil but MultiSig.GetPayloadHash was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPayloadHash.Lock()
	mock.calls.GetPayloadHash = append(mock.calls.GetPayloadHash, callInfo)
	mock.lockGetPayloadHash.Unlock()
	return mock.GetPayloadHashFunc()
}

// GetPayloadHashCalls gets all the calls that were made to GetPayloadHash.
// Check the length with:
//
//	len(mockedMultiSig.GetPayloadHashCalls())
func (mock *MultiSigMock) GetPayloadHashCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPayloadHash.RLock()
	calls = mock.calls.GetPayloadHash
	mock.lockGetPayloadHash.RUnlock()
	return calls
}

// GetSignature calls GetSignatureFunc.
func (mock *MultiSigMock) GetSignature(p sdk.ValAddress) (btcec.Signature, bool) {
	if mock.GetSignatureFunc == nil {
		panic("MultiSigMock.GetSignatureFunc: method is nil but MultiSig.GetSignature was just called")
	}
	callInfo := struct {
		P sdk.ValAddress
	}{
		P: p,
	}
	mock.lockGetSignature.Lock()
	mock.calls.GetSignature = append(mock.calls.GetSignature, callInfo)
	mock.lockGetSignature.Unlock()
	return mock.GetSignatureFunc(p)
}

// GetSignatureCalls gets all the calls that were made to GetSignature.
// Check the length with:
//
//	len(mockedMultiSig.GetSignatureCalls())
func (mock *MultiSigMock) GetSignatureCalls() []struct {
	P sdk.ValAddress
} {
	var calls []struct {
		P sdk.ValAddress
	}
	mock.lockGetSignature.RLock()
	calls = mock.calls.GetSignature
	mock.lockGetSignature.RUnlock()
	return calls
}

// ValidateBasic calls ValidateBasicFunc.
func (mock *MultiSigMock) ValidateBasic() error {
	if mock.ValidateBasicFunc == nil {
		panic("MultiSigMock.ValidateBasicFunc: method is nil but MultiSig.ValidateBasic was just called")
	}
	callInfo := struct {
	}{}
	mock.lockValidateBasic.Lock()
	mock.calls.ValidateBasic = append(mock.calls.ValidateBasic, callInfo)
	mock.lockValidateBasic.Unlock()
	return mock.ValidateBasicFunc()
}

// ValidateBasicCalls gets all the calls that were made to ValidateBasic.
// Check the length with:
//
//	len(mockedMultiSig.ValidateBasicCalls())
func (mock *MultiSigMock) ValidateBasicCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockValidateBasic.RLock()
	calls = mock.calls.ValidateBasic
	mock.lockValidateBasic.RUnlock()
	return calls
}
