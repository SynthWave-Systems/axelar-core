// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	context "context"
	tofnd "github.com/axelarnetwork/axelar-core/x/tss/tofnd"
	tsstypes "github.com/axelarnetwork/axelar-core/x/tss/types"
	exported1 "github.com/axelarnetwork/axelar-core/x/vote/exported"
	votetypes "github.com/axelarnetwork/axelar-core/x/vote/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"sync"
)

// Ensure, that TofndClientMock does implement tsstypes.TofndClient.
// If this is not the case, regenerate this file with moq.
var _ tsstypes.TofndClient = &TofndClientMock{}

// TofndClientMock is a mock implementation of tsstypes.TofndClient.
//
// 	func TestSomethingThatUsesTofndClient(t *testing.T) {
//
// 		// make and configure a mocked tsstypes.TofndClient
// 		mockedTofndClient := &TofndClientMock{
// 			KeygenFunc: func(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_KeygenClient, error) {
// 				panic("mock out the Keygen method")
// 			},
// 			SignFunc: func(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_SignClient, error) {
// 				panic("mock out the Sign method")
// 			},
// 		}
//
// 		// use mockedTofndClient in code that requires tsstypes.TofndClient
// 		// and then make assertions.
//
// 	}
type TofndClientMock struct {
	// KeygenFunc mocks the Keygen method.
	KeygenFunc func(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_KeygenClient, error)

	// SignFunc mocks the Sign method.
	SignFunc func(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_SignClient, error)

	// calls tracks calls to the methods.
	calls struct {
		// Keygen holds details about calls to the Keygen method.
		Keygen []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// Sign holds details about calls to the Sign method.
		Sign []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
	}
	lockKeygen sync.RWMutex
	lockSign   sync.RWMutex
}

// Keygen calls KeygenFunc.
func (mock *TofndClientMock) Keygen(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_KeygenClient, error) {
	if mock.KeygenFunc == nil {
		panic("TofndClientMock.KeygenFunc: method is nil but TofndClient.Keygen was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockKeygen.Lock()
	mock.calls.Keygen = append(mock.calls.Keygen, callInfo)
	mock.lockKeygen.Unlock()
	return mock.KeygenFunc(ctx, opts...)
}

// KeygenCalls gets all the calls that were made to Keygen.
// Check the length with:
//     len(mockedTofndClient.KeygenCalls())
func (mock *TofndClientMock) KeygenCalls() []struct {
	Ctx  context.Context
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}
	mock.lockKeygen.RLock()
	calls = mock.calls.Keygen
	mock.lockKeygen.RUnlock()
	return calls
}

// Sign calls SignFunc.
func (mock *TofndClientMock) Sign(ctx context.Context, opts ...grpc.CallOption) (tofnd.GG20_SignClient, error) {
	if mock.SignFunc == nil {
		panic("TofndClientMock.SignFunc: method is nil but TofndClient.Sign was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockSign.Lock()
	mock.calls.Sign = append(mock.calls.Sign, callInfo)
	mock.lockSign.Unlock()
	return mock.SignFunc(ctx, opts...)
}

// SignCalls gets all the calls that were made to Sign.
// Check the length with:
//     len(mockedTofndClient.SignCalls())
func (mock *TofndClientMock) SignCalls() []struct {
	Ctx  context.Context
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}
	mock.lockSign.RLock()
	calls = mock.calls.Sign
	mock.lockSign.RUnlock()
	return calls
}

// Ensure, that TofndKeyGenClientMock does implement tsstypes.TofndKeyGenClient.
// If this is not the case, regenerate this file with moq.
var _ tsstypes.TofndKeyGenClient = &TofndKeyGenClientMock{}

// TofndKeyGenClientMock is a mock implementation of tsstypes.TofndKeyGenClient.
//
// 	func TestSomethingThatUsesTofndKeyGenClient(t *testing.T) {
//
// 		// make and configure a mocked tsstypes.TofndKeyGenClient
// 		mockedTofndKeyGenClient := &TofndKeyGenClientMock{
// 			CloseSendFunc: func() error {
// 				panic("mock out the CloseSend method")
// 			},
// 			ContextFunc: func() context.Context {
// 				panic("mock out the Context method")
// 			},
// 			HeaderFunc: func() (metadata.MD, error) {
// 				panic("mock out the Header method")
// 			},
// 			RecvFunc: func() (*tofnd.MessageOut, error) {
// 				panic("mock out the Recv method")
// 			},
// 			RecvMsgFunc: func(m interface{}) error {
// 				panic("mock out the RecvMsg method")
// 			},
// 			SendFunc: func(messageIn *tofnd.MessageIn) error {
// 				panic("mock out the Send method")
// 			},
// 			SendMsgFunc: func(m interface{}) error {
// 				panic("mock out the SendMsg method")
// 			},
// 			TrailerFunc: func() metadata.MD {
// 				panic("mock out the Trailer method")
// 			},
// 		}
//
// 		// use mockedTofndKeyGenClient in code that requires tsstypes.TofndKeyGenClient
// 		// and then make assertions.
//
// 	}
type TofndKeyGenClientMock struct {
	// CloseSendFunc mocks the CloseSend method.
	CloseSendFunc func() error

	// ContextFunc mocks the Context method.
	ContextFunc func() context.Context

	// HeaderFunc mocks the Header method.
	HeaderFunc func() (metadata.MD, error)

	// RecvFunc mocks the Recv method.
	RecvFunc func() (*tofnd.MessageOut, error)

	// RecvMsgFunc mocks the RecvMsg method.
	RecvMsgFunc func(m interface{}) error

	// SendFunc mocks the Send method.
	SendFunc func(messageIn *tofnd.MessageIn) error

	// SendMsgFunc mocks the SendMsg method.
	SendMsgFunc func(m interface{}) error

	// TrailerFunc mocks the Trailer method.
	TrailerFunc func() metadata.MD

	// calls tracks calls to the methods.
	calls struct {
		// CloseSend holds details about calls to the CloseSend method.
		CloseSend []struct {
		}
		// Context holds details about calls to the Context method.
		Context []struct {
		}
		// Header holds details about calls to the Header method.
		Header []struct {
		}
		// Recv holds details about calls to the Recv method.
		Recv []struct {
		}
		// RecvMsg holds details about calls to the RecvMsg method.
		RecvMsg []struct {
			// M is the m argument value.
			M interface{}
		}
		// Send holds details about calls to the Send method.
		Send []struct {
			// MessageIn is the messageIn argument value.
			MessageIn *tofnd.MessageIn
		}
		// SendMsg holds details about calls to the SendMsg method.
		SendMsg []struct {
			// M is the m argument value.
			M interface{}
		}
		// Trailer holds details about calls to the Trailer method.
		Trailer []struct {
		}
	}
	lockCloseSend sync.RWMutex
	lockContext   sync.RWMutex
	lockHeader    sync.RWMutex
	lockRecv      sync.RWMutex
	lockRecvMsg   sync.RWMutex
	lockSend      sync.RWMutex
	lockSendMsg   sync.RWMutex
	lockTrailer   sync.RWMutex
}

// CloseSend calls CloseSendFunc.
func (mock *TofndKeyGenClientMock) CloseSend() error {
	if mock.CloseSendFunc == nil {
		panic("TofndKeyGenClientMock.CloseSendFunc: method is nil but TofndKeyGenClient.CloseSend was just called")
	}
	callInfo := struct {
	}{}
	mock.lockCloseSend.Lock()
	mock.calls.CloseSend = append(mock.calls.CloseSend, callInfo)
	mock.lockCloseSend.Unlock()
	return mock.CloseSendFunc()
}

// CloseSendCalls gets all the calls that were made to CloseSend.
// Check the length with:
//     len(mockedTofndKeyGenClient.CloseSendCalls())
func (mock *TofndKeyGenClientMock) CloseSendCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCloseSend.RLock()
	calls = mock.calls.CloseSend
	mock.lockCloseSend.RUnlock()
	return calls
}

// Context calls ContextFunc.
func (mock *TofndKeyGenClientMock) Context() context.Context {
	if mock.ContextFunc == nil {
		panic("TofndKeyGenClientMock.ContextFunc: method is nil but TofndKeyGenClient.Context was just called")
	}
	callInfo := struct {
	}{}
	mock.lockContext.Lock()
	mock.calls.Context = append(mock.calls.Context, callInfo)
	mock.lockContext.Unlock()
	return mock.ContextFunc()
}

// ContextCalls gets all the calls that were made to Context.
// Check the length with:
//     len(mockedTofndKeyGenClient.ContextCalls())
func (mock *TofndKeyGenClientMock) ContextCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockContext.RLock()
	calls = mock.calls.Context
	mock.lockContext.RUnlock()
	return calls
}

// Header calls HeaderFunc.
func (mock *TofndKeyGenClientMock) Header() (metadata.MD, error) {
	if mock.HeaderFunc == nil {
		panic("TofndKeyGenClientMock.HeaderFunc: method is nil but TofndKeyGenClient.Header was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHeader.Lock()
	mock.calls.Header = append(mock.calls.Header, callInfo)
	mock.lockHeader.Unlock()
	return mock.HeaderFunc()
}

// HeaderCalls gets all the calls that were made to Header.
// Check the length with:
//     len(mockedTofndKeyGenClient.HeaderCalls())
func (mock *TofndKeyGenClientMock) HeaderCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHeader.RLock()
	calls = mock.calls.Header
	mock.lockHeader.RUnlock()
	return calls
}

// Recv calls RecvFunc.
func (mock *TofndKeyGenClientMock) Recv() (*tofnd.MessageOut, error) {
	if mock.RecvFunc == nil {
		panic("TofndKeyGenClientMock.RecvFunc: method is nil but TofndKeyGenClient.Recv was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRecv.Lock()
	mock.calls.Recv = append(mock.calls.Recv, callInfo)
	mock.lockRecv.Unlock()
	return mock.RecvFunc()
}

// RecvCalls gets all the calls that were made to Recv.
// Check the length with:
//     len(mockedTofndKeyGenClient.RecvCalls())
func (mock *TofndKeyGenClientMock) RecvCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRecv.RLock()
	calls = mock.calls.Recv
	mock.lockRecv.RUnlock()
	return calls
}

// RecvMsg calls RecvMsgFunc.
func (mock *TofndKeyGenClientMock) RecvMsg(m interface{}) error {
	if mock.RecvMsgFunc == nil {
		panic("TofndKeyGenClientMock.RecvMsgFunc: method is nil but TofndKeyGenClient.RecvMsg was just called")
	}
	callInfo := struct {
		M interface{}
	}{
		M: m,
	}
	mock.lockRecvMsg.Lock()
	mock.calls.RecvMsg = append(mock.calls.RecvMsg, callInfo)
	mock.lockRecvMsg.Unlock()
	return mock.RecvMsgFunc(m)
}

// RecvMsgCalls gets all the calls that were made to RecvMsg.
// Check the length with:
//     len(mockedTofndKeyGenClient.RecvMsgCalls())
func (mock *TofndKeyGenClientMock) RecvMsgCalls() []struct {
	M interface{}
} {
	var calls []struct {
		M interface{}
	}
	mock.lockRecvMsg.RLock()
	calls = mock.calls.RecvMsg
	mock.lockRecvMsg.RUnlock()
	return calls
}

// Send calls SendFunc.
func (mock *TofndKeyGenClientMock) Send(messageIn *tofnd.MessageIn) error {
	if mock.SendFunc == nil {
		panic("TofndKeyGenClientMock.SendFunc: method is nil but TofndKeyGenClient.Send was just called")
	}
	callInfo := struct {
		MessageIn *tofnd.MessageIn
	}{
		MessageIn: messageIn,
	}
	mock.lockSend.Lock()
	mock.calls.Send = append(mock.calls.Send, callInfo)
	mock.lockSend.Unlock()
	return mock.SendFunc(messageIn)
}

// SendCalls gets all the calls that were made to Send.
// Check the length with:
//     len(mockedTofndKeyGenClient.SendCalls())
func (mock *TofndKeyGenClientMock) SendCalls() []struct {
	MessageIn *tofnd.MessageIn
} {
	var calls []struct {
		MessageIn *tofnd.MessageIn
	}
	mock.lockSend.RLock()
	calls = mock.calls.Send
	mock.lockSend.RUnlock()
	return calls
}

// SendMsg calls SendMsgFunc.
func (mock *TofndKeyGenClientMock) SendMsg(m interface{}) error {
	if mock.SendMsgFunc == nil {
		panic("TofndKeyGenClientMock.SendMsgFunc: method is nil but TofndKeyGenClient.SendMsg was just called")
	}
	callInfo := struct {
		M interface{}
	}{
		M: m,
	}
	mock.lockSendMsg.Lock()
	mock.calls.SendMsg = append(mock.calls.SendMsg, callInfo)
	mock.lockSendMsg.Unlock()
	return mock.SendMsgFunc(m)
}

// SendMsgCalls gets all the calls that were made to SendMsg.
// Check the length with:
//     len(mockedTofndKeyGenClient.SendMsgCalls())
func (mock *TofndKeyGenClientMock) SendMsgCalls() []struct {
	M interface{}
} {
	var calls []struct {
		M interface{}
	}
	mock.lockSendMsg.RLock()
	calls = mock.calls.SendMsg
	mock.lockSendMsg.RUnlock()
	return calls
}

// Trailer calls TrailerFunc.
func (mock *TofndKeyGenClientMock) Trailer() metadata.MD {
	if mock.TrailerFunc == nil {
		panic("TofndKeyGenClientMock.TrailerFunc: method is nil but TofndKeyGenClient.Trailer was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTrailer.Lock()
	mock.calls.Trailer = append(mock.calls.Trailer, callInfo)
	mock.lockTrailer.Unlock()
	return mock.TrailerFunc()
}

// TrailerCalls gets all the calls that were made to Trailer.
// Check the length with:
//     len(mockedTofndKeyGenClient.TrailerCalls())
func (mock *TofndKeyGenClientMock) TrailerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTrailer.RLock()
	calls = mock.calls.Trailer
	mock.lockTrailer.RUnlock()
	return calls
}

// Ensure, that TofndSignClientMock does implement tsstypes.TofndSignClient.
// If this is not the case, regenerate this file with moq.
var _ tsstypes.TofndSignClient = &TofndSignClientMock{}

// TofndSignClientMock is a mock implementation of tsstypes.TofndSignClient.
//
// 	func TestSomethingThatUsesTofndSignClient(t *testing.T) {
//
// 		// make and configure a mocked tsstypes.TofndSignClient
// 		mockedTofndSignClient := &TofndSignClientMock{
// 			CloseSendFunc: func() error {
// 				panic("mock out the CloseSend method")
// 			},
// 			ContextFunc: func() context.Context {
// 				panic("mock out the Context method")
// 			},
// 			HeaderFunc: func() (metadata.MD, error) {
// 				panic("mock out the Header method")
// 			},
// 			RecvFunc: func() (*tofnd.MessageOut, error) {
// 				panic("mock out the Recv method")
// 			},
// 			RecvMsgFunc: func(m interface{}) error {
// 				panic("mock out the RecvMsg method")
// 			},
// 			SendFunc: func(messageIn *tofnd.MessageIn) error {
// 				panic("mock out the Send method")
// 			},
// 			SendMsgFunc: func(m interface{}) error {
// 				panic("mock out the SendMsg method")
// 			},
// 			TrailerFunc: func() metadata.MD {
// 				panic("mock out the Trailer method")
// 			},
// 		}
//
// 		// use mockedTofndSignClient in code that requires tsstypes.TofndSignClient
// 		// and then make assertions.
//
// 	}
type TofndSignClientMock struct {
	// CloseSendFunc mocks the CloseSend method.
	CloseSendFunc func() error

	// ContextFunc mocks the Context method.
	ContextFunc func() context.Context

	// HeaderFunc mocks the Header method.
	HeaderFunc func() (metadata.MD, error)

	// RecvFunc mocks the Recv method.
	RecvFunc func() (*tofnd.MessageOut, error)

	// RecvMsgFunc mocks the RecvMsg method.
	RecvMsgFunc func(m interface{}) error

	// SendFunc mocks the Send method.
	SendFunc func(messageIn *tofnd.MessageIn) error

	// SendMsgFunc mocks the SendMsg method.
	SendMsgFunc func(m interface{}) error

	// TrailerFunc mocks the Trailer method.
	TrailerFunc func() metadata.MD

	// calls tracks calls to the methods.
	calls struct {
		// CloseSend holds details about calls to the CloseSend method.
		CloseSend []struct {
		}
		// Context holds details about calls to the Context method.
		Context []struct {
		}
		// Header holds details about calls to the Header method.
		Header []struct {
		}
		// Recv holds details about calls to the Recv method.
		Recv []struct {
		}
		// RecvMsg holds details about calls to the RecvMsg method.
		RecvMsg []struct {
			// M is the m argument value.
			M interface{}
		}
		// Send holds details about calls to the Send method.
		Send []struct {
			// MessageIn is the messageIn argument value.
			MessageIn *tofnd.MessageIn
		}
		// SendMsg holds details about calls to the SendMsg method.
		SendMsg []struct {
			// M is the m argument value.
			M interface{}
		}
		// Trailer holds details about calls to the Trailer method.
		Trailer []struct {
		}
	}
	lockCloseSend sync.RWMutex
	lockContext   sync.RWMutex
	lockHeader    sync.RWMutex
	lockRecv      sync.RWMutex
	lockRecvMsg   sync.RWMutex
	lockSend      sync.RWMutex
	lockSendMsg   sync.RWMutex
	lockTrailer   sync.RWMutex
}

// CloseSend calls CloseSendFunc.
func (mock *TofndSignClientMock) CloseSend() error {
	if mock.CloseSendFunc == nil {
		panic("TofndSignClientMock.CloseSendFunc: method is nil but TofndSignClient.CloseSend was just called")
	}
	callInfo := struct {
	}{}
	mock.lockCloseSend.Lock()
	mock.calls.CloseSend = append(mock.calls.CloseSend, callInfo)
	mock.lockCloseSend.Unlock()
	return mock.CloseSendFunc()
}

// CloseSendCalls gets all the calls that were made to CloseSend.
// Check the length with:
//     len(mockedTofndSignClient.CloseSendCalls())
func (mock *TofndSignClientMock) CloseSendCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCloseSend.RLock()
	calls = mock.calls.CloseSend
	mock.lockCloseSend.RUnlock()
	return calls
}

// Context calls ContextFunc.
func (mock *TofndSignClientMock) Context() context.Context {
	if mock.ContextFunc == nil {
		panic("TofndSignClientMock.ContextFunc: method is nil but TofndSignClient.Context was just called")
	}
	callInfo := struct {
	}{}
	mock.lockContext.Lock()
	mock.calls.Context = append(mock.calls.Context, callInfo)
	mock.lockContext.Unlock()
	return mock.ContextFunc()
}

// ContextCalls gets all the calls that were made to Context.
// Check the length with:
//     len(mockedTofndSignClient.ContextCalls())
func (mock *TofndSignClientMock) ContextCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockContext.RLock()
	calls = mock.calls.Context
	mock.lockContext.RUnlock()
	return calls
}

// Header calls HeaderFunc.
func (mock *TofndSignClientMock) Header() (metadata.MD, error) {
	if mock.HeaderFunc == nil {
		panic("TofndSignClientMock.HeaderFunc: method is nil but TofndSignClient.Header was just called")
	}
	callInfo := struct {
	}{}
	mock.lockHeader.Lock()
	mock.calls.Header = append(mock.calls.Header, callInfo)
	mock.lockHeader.Unlock()
	return mock.HeaderFunc()
}

// HeaderCalls gets all the calls that were made to Header.
// Check the length with:
//     len(mockedTofndSignClient.HeaderCalls())
func (mock *TofndSignClientMock) HeaderCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHeader.RLock()
	calls = mock.calls.Header
	mock.lockHeader.RUnlock()
	return calls
}

// Recv calls RecvFunc.
func (mock *TofndSignClientMock) Recv() (*tofnd.MessageOut, error) {
	if mock.RecvFunc == nil {
		panic("TofndSignClientMock.RecvFunc: method is nil but TofndSignClient.Recv was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRecv.Lock()
	mock.calls.Recv = append(mock.calls.Recv, callInfo)
	mock.lockRecv.Unlock()
	return mock.RecvFunc()
}

// RecvCalls gets all the calls that were made to Recv.
// Check the length with:
//     len(mockedTofndSignClient.RecvCalls())
func (mock *TofndSignClientMock) RecvCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRecv.RLock()
	calls = mock.calls.Recv
	mock.lockRecv.RUnlock()
	return calls
}

// RecvMsg calls RecvMsgFunc.
func (mock *TofndSignClientMock) RecvMsg(m interface{}) error {
	if mock.RecvMsgFunc == nil {
		panic("TofndSignClientMock.RecvMsgFunc: method is nil but TofndSignClient.RecvMsg was just called")
	}
	callInfo := struct {
		M interface{}
	}{
		M: m,
	}
	mock.lockRecvMsg.Lock()
	mock.calls.RecvMsg = append(mock.calls.RecvMsg, callInfo)
	mock.lockRecvMsg.Unlock()
	return mock.RecvMsgFunc(m)
}

// RecvMsgCalls gets all the calls that were made to RecvMsg.
// Check the length with:
//     len(mockedTofndSignClient.RecvMsgCalls())
func (mock *TofndSignClientMock) RecvMsgCalls() []struct {
	M interface{}
} {
	var calls []struct {
		M interface{}
	}
	mock.lockRecvMsg.RLock()
	calls = mock.calls.RecvMsg
	mock.lockRecvMsg.RUnlock()
	return calls
}

// Send calls SendFunc.
func (mock *TofndSignClientMock) Send(messageIn *tofnd.MessageIn) error {
	if mock.SendFunc == nil {
		panic("TofndSignClientMock.SendFunc: method is nil but TofndSignClient.Send was just called")
	}
	callInfo := struct {
		MessageIn *tofnd.MessageIn
	}{
		MessageIn: messageIn,
	}
	mock.lockSend.Lock()
	mock.calls.Send = append(mock.calls.Send, callInfo)
	mock.lockSend.Unlock()
	return mock.SendFunc(messageIn)
}

// SendCalls gets all the calls that were made to Send.
// Check the length with:
//     len(mockedTofndSignClient.SendCalls())
func (mock *TofndSignClientMock) SendCalls() []struct {
	MessageIn *tofnd.MessageIn
} {
	var calls []struct {
		MessageIn *tofnd.MessageIn
	}
	mock.lockSend.RLock()
	calls = mock.calls.Send
	mock.lockSend.RUnlock()
	return calls
}

// SendMsg calls SendMsgFunc.
func (mock *TofndSignClientMock) SendMsg(m interface{}) error {
	if mock.SendMsgFunc == nil {
		panic("TofndSignClientMock.SendMsgFunc: method is nil but TofndSignClient.SendMsg was just called")
	}
	callInfo := struct {
		M interface{}
	}{
		M: m,
	}
	mock.lockSendMsg.Lock()
	mock.calls.SendMsg = append(mock.calls.SendMsg, callInfo)
	mock.lockSendMsg.Unlock()
	return mock.SendMsgFunc(m)
}

// SendMsgCalls gets all the calls that were made to SendMsg.
// Check the length with:
//     len(mockedTofndSignClient.SendMsgCalls())
func (mock *TofndSignClientMock) SendMsgCalls() []struct {
	M interface{}
} {
	var calls []struct {
		M interface{}
	}
	mock.lockSendMsg.RLock()
	calls = mock.calls.SendMsg
	mock.lockSendMsg.RUnlock()
	return calls
}

// Trailer calls TrailerFunc.
func (mock *TofndSignClientMock) Trailer() metadata.MD {
	if mock.TrailerFunc == nil {
		panic("TofndSignClientMock.TrailerFunc: method is nil but TofndSignClient.Trailer was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTrailer.Lock()
	mock.calls.Trailer = append(mock.calls.Trailer, callInfo)
	mock.lockTrailer.Unlock()
	return mock.TrailerFunc()
}

// TrailerCalls gets all the calls that were made to Trailer.
// Check the length with:
//     len(mockedTofndSignClient.TrailerCalls())
func (mock *TofndSignClientMock) TrailerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTrailer.RLock()
	calls = mock.calls.Trailer
	mock.lockTrailer.RUnlock()
	return calls
}

// Ensure, that VoterMock does implement tsstypes.Voter.
// If this is not the case, regenerate this file with moq.
var _ tsstypes.Voter = &VoterMock{}

// VoterMock is a mock implementation of tsstypes.Voter.
//
// 	func TestSomethingThatUsesVoter(t *testing.T) {
//
// 		// make and configure a mocked tsstypes.Voter
// 		mockedVoter := &VoterMock{
// 			DeletePollFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported1.PollMeta)  {
// 				panic("mock out the DeletePoll method")
// 			},
// 			GetPollFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, pollMeta exported1.PollMeta) *votetypes.Poll {
// 				panic("mock out the GetPoll method")
// 			},
// 			InitPollFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported1.PollMeta, snapshotCounter int64, expireAt int64) error {
// 				panic("mock out the InitPoll method")
// 			},
// 			ResultFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported1.PollMeta) exported1.VotingData {
// 				panic("mock out the Result method")
// 			},
// 			TallyVoteFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, sender github_com_cosmos_cosmos_sdk_types.AccAddress, pollMeta exported1.PollMeta, data exported1.VotingData) error {
// 				panic("mock out the TallyVote method")
// 			},
// 		}
//
// 		// use mockedVoter in code that requires tsstypes.Voter
// 		// and then make assertions.
//
// 	}
type VoterMock struct {
	// DeletePollFunc mocks the DeletePoll method.
	DeletePollFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported1.PollMeta)

	// GetPollFunc mocks the GetPoll method.
	GetPollFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, pollMeta exported1.PollMeta) *votetypes.Poll

	// InitPollFunc mocks the InitPoll method.
	InitPollFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported1.PollMeta, snapshotCounter int64, expireAt int64) error

	// ResultFunc mocks the Result method.
	ResultFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported1.PollMeta) exported1.VotingData

	// TallyVoteFunc mocks the TallyVote method.
	TallyVoteFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, sender github_com_cosmos_cosmos_sdk_types.AccAddress, pollMeta exported1.PollMeta, data exported1.VotingData) error

	// calls tracks calls to the methods.
	calls struct {
		// DeletePoll holds details about calls to the DeletePoll method.
		DeletePoll []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Poll is the poll argument value.
			Poll exported1.PollMeta
		}
		// GetPoll holds details about calls to the GetPoll method.
		GetPoll []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// PollMeta is the pollMeta argument value.
			PollMeta exported1.PollMeta
		}
		// InitPoll holds details about calls to the InitPoll method.
		InitPoll []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Poll is the poll argument value.
			Poll exported1.PollMeta
			// SnapshotCounter is the snapshotCounter argument value.
			SnapshotCounter int64
			// ExpireAt is the expireAt argument value.
			ExpireAt int64
		}
		// Result holds details about calls to the Result method.
		Result []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Poll is the poll argument value.
			Poll exported1.PollMeta
		}
		// TallyVote holds details about calls to the TallyVote method.
		TallyVote []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Sender is the sender argument value.
			Sender github_com_cosmos_cosmos_sdk_types.AccAddress
			// PollMeta is the pollMeta argument value.
			PollMeta exported1.PollMeta
			// Data is the data argument value.
			Data exported1.VotingData
		}
	}
	lockDeletePoll sync.RWMutex
	lockGetPoll    sync.RWMutex
	lockInitPoll   sync.RWMutex
	lockResult     sync.RWMutex
	lockTallyVote  sync.RWMutex
}

// DeletePoll calls DeletePollFunc.
func (mock *VoterMock) DeletePoll(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported1.PollMeta) {
	if mock.DeletePollFunc == nil {
		panic("VoterMock.DeletePollFunc: method is nil but Voter.DeletePoll was just called")
	}
	callInfo := struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Poll exported1.PollMeta
	}{
		Ctx:  ctx,
		Poll: poll,
	}
	mock.lockDeletePoll.Lock()
	mock.calls.DeletePoll = append(mock.calls.DeletePoll, callInfo)
	mock.lockDeletePoll.Unlock()
	mock.DeletePollFunc(ctx, poll)
}

// DeletePollCalls gets all the calls that were made to DeletePoll.
// Check the length with:
//     len(mockedVoter.DeletePollCalls())
func (mock *VoterMock) DeletePollCalls() []struct {
	Ctx  github_com_cosmos_cosmos_sdk_types.Context
	Poll exported1.PollMeta
} {
	var calls []struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Poll exported1.PollMeta
	}
	mock.lockDeletePoll.RLock()
	calls = mock.calls.DeletePoll
	mock.lockDeletePoll.RUnlock()
	return calls
}

// GetPoll calls GetPollFunc.
func (mock *VoterMock) GetPoll(ctx github_com_cosmos_cosmos_sdk_types.Context, pollMeta exported1.PollMeta) *votetypes.Poll {
	if mock.GetPollFunc == nil {
		panic("VoterMock.GetPollFunc: method is nil but Voter.GetPoll was just called")
	}
	callInfo := struct {
		Ctx      github_com_cosmos_cosmos_sdk_types.Context
		PollMeta exported1.PollMeta
	}{
		Ctx:      ctx,
		PollMeta: pollMeta,
	}
	mock.lockGetPoll.Lock()
	mock.calls.GetPoll = append(mock.calls.GetPoll, callInfo)
	mock.lockGetPoll.Unlock()
	return mock.GetPollFunc(ctx, pollMeta)
}

// GetPollCalls gets all the calls that were made to GetPoll.
// Check the length with:
//     len(mockedVoter.GetPollCalls())
func (mock *VoterMock) GetPollCalls() []struct {
	Ctx      github_com_cosmos_cosmos_sdk_types.Context
	PollMeta exported1.PollMeta
} {
	var calls []struct {
		Ctx      github_com_cosmos_cosmos_sdk_types.Context
		PollMeta exported1.PollMeta
	}
	mock.lockGetPoll.RLock()
	calls = mock.calls.GetPoll
	mock.lockGetPoll.RUnlock()
	return calls
}

// InitPoll calls InitPollFunc.
func (mock *VoterMock) InitPoll(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported1.PollMeta, snapshotCounter int64, expireAt int64) error {
	if mock.InitPollFunc == nil {
		panic("VoterMock.InitPollFunc: method is nil but Voter.InitPoll was just called")
	}
	callInfo := struct {
		Ctx             github_com_cosmos_cosmos_sdk_types.Context
		Poll            exported1.PollMeta
		SnapshotCounter int64
		ExpireAt        int64
	}{
		Ctx:             ctx,
		Poll:            poll,
		SnapshotCounter: snapshotCounter,
		ExpireAt:        expireAt,
	}
	mock.lockInitPoll.Lock()
	mock.calls.InitPoll = append(mock.calls.InitPoll, callInfo)
	mock.lockInitPoll.Unlock()
	return mock.InitPollFunc(ctx, poll, snapshotCounter, expireAt)
}

// InitPollCalls gets all the calls that were made to InitPoll.
// Check the length with:
//     len(mockedVoter.InitPollCalls())
func (mock *VoterMock) InitPollCalls() []struct {
	Ctx             github_com_cosmos_cosmos_sdk_types.Context
	Poll            exported1.PollMeta
	SnapshotCounter int64
	ExpireAt        int64
} {
	var calls []struct {
		Ctx             github_com_cosmos_cosmos_sdk_types.Context
		Poll            exported1.PollMeta
		SnapshotCounter int64
		ExpireAt        int64
	}
	mock.lockInitPoll.RLock()
	calls = mock.calls.InitPoll
	mock.lockInitPoll.RUnlock()
	return calls
}

// Result calls ResultFunc.
func (mock *VoterMock) Result(ctx github_com_cosmos_cosmos_sdk_types.Context, poll exported1.PollMeta) exported1.VotingData {
	if mock.ResultFunc == nil {
		panic("VoterMock.ResultFunc: method is nil but Voter.Result was just called")
	}
	callInfo := struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Poll exported1.PollMeta
	}{
		Ctx:  ctx,
		Poll: poll,
	}
	mock.lockResult.Lock()
	mock.calls.Result = append(mock.calls.Result, callInfo)
	mock.lockResult.Unlock()
	return mock.ResultFunc(ctx, poll)
}

// ResultCalls gets all the calls that were made to Result.
// Check the length with:
//     len(mockedVoter.ResultCalls())
func (mock *VoterMock) ResultCalls() []struct {
	Ctx  github_com_cosmos_cosmos_sdk_types.Context
	Poll exported1.PollMeta
} {
	var calls []struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Poll exported1.PollMeta
	}
	mock.lockResult.RLock()
	calls = mock.calls.Result
	mock.lockResult.RUnlock()
	return calls
}

// TallyVote calls TallyVoteFunc.
func (mock *VoterMock) TallyVote(ctx github_com_cosmos_cosmos_sdk_types.Context, sender github_com_cosmos_cosmos_sdk_types.AccAddress, pollMeta exported1.PollMeta, data exported1.VotingData) error {
	if mock.TallyVoteFunc == nil {
		panic("VoterMock.TallyVoteFunc: method is nil but Voter.TallyVote was just called")
	}
	callInfo := struct {
		Ctx      github_com_cosmos_cosmos_sdk_types.Context
		Sender   github_com_cosmos_cosmos_sdk_types.AccAddress
		PollMeta exported1.PollMeta
		Data     exported1.VotingData
	}{
		Ctx:      ctx,
		Sender:   sender,
		PollMeta: pollMeta,
		Data:     data,
	}
	mock.lockTallyVote.Lock()
	mock.calls.TallyVote = append(mock.calls.TallyVote, callInfo)
	mock.lockTallyVote.Unlock()
	return mock.TallyVoteFunc(ctx, sender, pollMeta, data)
}

// TallyVoteCalls gets all the calls that were made to TallyVote.
// Check the length with:
//     len(mockedVoter.TallyVoteCalls())
func (mock *VoterMock) TallyVoteCalls() []struct {
	Ctx      github_com_cosmos_cosmos_sdk_types.Context
	Sender   github_com_cosmos_cosmos_sdk_types.AccAddress
	PollMeta exported1.PollMeta
	Data     exported1.VotingData
} {
	var calls []struct {
		Ctx      github_com_cosmos_cosmos_sdk_types.Context
		Sender   github_com_cosmos_cosmos_sdk_types.AccAddress
		PollMeta exported1.PollMeta
		Data     exported1.VotingData
	}
	mock.lockTallyVote.RLock()
	calls = mock.calls.TallyVote
	mock.lockTallyVote.RUnlock()
	return calls
}

// Ensure, that StakingKeeperMock does implement tsstypes.StakingKeeper.
// If this is not the case, regenerate this file with moq.
var _ tsstypes.StakingKeeper = &StakingKeeperMock{}

// StakingKeeperMock is a mock implementation of tsstypes.StakingKeeper.
//
// 	func TestSomethingThatUsesStakingKeeper(t *testing.T) {
//
// 		// make and configure a mocked tsstypes.StakingKeeper
// 		mockedStakingKeeper := &StakingKeeperMock{
// 			GetLastTotalPowerFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context) github_com_cosmos_cosmos_sdk_types.Int {
// 				panic("mock out the GetLastTotalPower method")
// 			},
// 			GetValidatorFunc: func(ctx github_com_cosmos_cosmos_sdk_types.Context, addr github_com_cosmos_cosmos_sdk_types.ValAddress) (stakingtypes.Validator, bool) {
// 				panic("mock out the GetValidator method")
// 			},
// 		}
//
// 		// use mockedStakingKeeper in code that requires tsstypes.StakingKeeper
// 		// and then make assertions.
//
// 	}
type StakingKeeperMock struct {
	// GetLastTotalPowerFunc mocks the GetLastTotalPower method.
	GetLastTotalPowerFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context) github_com_cosmos_cosmos_sdk_types.Int

	// GetValidatorFunc mocks the GetValidator method.
	GetValidatorFunc func(ctx github_com_cosmos_cosmos_sdk_types.Context, addr github_com_cosmos_cosmos_sdk_types.ValAddress) (stakingtypes.Validator, bool)

	// calls tracks calls to the methods.
	calls struct {
		// GetLastTotalPower holds details about calls to the GetLastTotalPower method.
		GetLastTotalPower []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
		}
		// GetValidator holds details about calls to the GetValidator method.
		GetValidator []struct {
			// Ctx is the ctx argument value.
			Ctx github_com_cosmos_cosmos_sdk_types.Context
			// Addr is the addr argument value.
			Addr github_com_cosmos_cosmos_sdk_types.ValAddress
		}
	}
	lockGetLastTotalPower sync.RWMutex
	lockGetValidator      sync.RWMutex
}

// GetLastTotalPower calls GetLastTotalPowerFunc.
func (mock *StakingKeeperMock) GetLastTotalPower(ctx github_com_cosmos_cosmos_sdk_types.Context) github_com_cosmos_cosmos_sdk_types.Int {
	if mock.GetLastTotalPowerFunc == nil {
		panic("StakingKeeperMock.GetLastTotalPowerFunc: method is nil but StakingKeeper.GetLastTotalPower was just called")
	}
	callInfo := struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetLastTotalPower.Lock()
	mock.calls.GetLastTotalPower = append(mock.calls.GetLastTotalPower, callInfo)
	mock.lockGetLastTotalPower.Unlock()
	return mock.GetLastTotalPowerFunc(ctx)
}

// GetLastTotalPowerCalls gets all the calls that were made to GetLastTotalPower.
// Check the length with:
//     len(mockedStakingKeeper.GetLastTotalPowerCalls())
func (mock *StakingKeeperMock) GetLastTotalPowerCalls() []struct {
	Ctx github_com_cosmos_cosmos_sdk_types.Context
} {
	var calls []struct {
		Ctx github_com_cosmos_cosmos_sdk_types.Context
	}
	mock.lockGetLastTotalPower.RLock()
	calls = mock.calls.GetLastTotalPower
	mock.lockGetLastTotalPower.RUnlock()
	return calls
}

// GetValidator calls GetValidatorFunc.
func (mock *StakingKeeperMock) GetValidator(ctx github_com_cosmos_cosmos_sdk_types.Context, addr github_com_cosmos_cosmos_sdk_types.ValAddress) (stakingtypes.Validator, bool) {
	if mock.GetValidatorFunc == nil {
		panic("StakingKeeperMock.GetValidatorFunc: method is nil but StakingKeeper.GetValidator was just called")
	}
	callInfo := struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Addr github_com_cosmos_cosmos_sdk_types.ValAddress
	}{
		Ctx:  ctx,
		Addr: addr,
	}
	mock.lockGetValidator.Lock()
	mock.calls.GetValidator = append(mock.calls.GetValidator, callInfo)
	mock.lockGetValidator.Unlock()
	return mock.GetValidatorFunc(ctx, addr)
}

// GetValidatorCalls gets all the calls that were made to GetValidator.
// Check the length with:
//     len(mockedStakingKeeper.GetValidatorCalls())
func (mock *StakingKeeperMock) GetValidatorCalls() []struct {
	Ctx  github_com_cosmos_cosmos_sdk_types.Context
	Addr github_com_cosmos_cosmos_sdk_types.ValAddress
} {
	var calls []struct {
		Ctx  github_com_cosmos_cosmos_sdk_types.Context
		Addr github_com_cosmos_cosmos_sdk_types.ValAddress
	}
	mock.lockGetValidator.RLock()
	calls = mock.calls.GetValidator
	mock.lockGetValidator.RUnlock()
	return calls
}
