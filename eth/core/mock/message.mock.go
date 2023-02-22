// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"pkg.berachain.dev/stargazer/eth/core"
	"sync"
)

// Ensure, that MessageMock does implement core.Message.
// If this is not the case, regenerate this file with moq.
var _ core.Message = &MessageMock{}

// MessageMock is a mock implementation of core.Message.
//
//	func TestSomethingThatUsesMessage(t *testing.T) {
//
//		// make and configure a mocked core.Message
//		mockedMessage := &MessageMock{
//			AccessListFunc: func() types.AccessList {
//				panic("mock out the AccessList method")
//			},
//			DataFunc: func() []byte {
//				panic("mock out the Data method")
//			},
//			FromFunc: func() common.Address {
//				panic("mock out the From method")
//			},
//			GasFunc: func() uint64 {
//				panic("mock out the Gas method")
//			},
//			GasFeeCapFunc: func() *big.Int {
//				panic("mock out the GasFeeCap method")
//			},
//			GasPriceFunc: func() *big.Int {
//				panic("mock out the GasPrice method")
//			},
//			GasTipCapFunc: func() *big.Int {
//				panic("mock out the GasTipCap method")
//			},
//			IsFakeFunc: func() bool {
//				panic("mock out the IsFake method")
//			},
//			NonceFunc: func() uint64 {
//				panic("mock out the Nonce method")
//			},
//			ToFunc: func() *common.Address {
//				panic("mock out the To method")
//			},
//			ValueFunc: func() *big.Int {
//				panic("mock out the Value method")
//			},
//		}
//
//		// use mockedMessage in code that requires core.Message
//		// and then make assertions.
//
//	}
type MessageMock struct {
	// AccessListFunc mocks the AccessList method.
	AccessListFunc func() types.AccessList

	// DataFunc mocks the Data method.
	DataFunc func() []byte

	// FromFunc mocks the From method.
	FromFunc func() common.Address

	// GasFunc mocks the Gas method.
	GasFunc func() uint64

	// GasFeeCapFunc mocks the GasFeeCap method.
	GasFeeCapFunc func() *big.Int

	// GasPriceFunc mocks the GasPrice method.
	GasPriceFunc func() *big.Int

	// GasTipCapFunc mocks the GasTipCap method.
	GasTipCapFunc func() *big.Int

	// IsFakeFunc mocks the IsFake method.
	IsFakeFunc func() bool

	// NonceFunc mocks the Nonce method.
	NonceFunc func() uint64

	// ToFunc mocks the To method.
	ToFunc func() *common.Address

	// ValueFunc mocks the Value method.
	ValueFunc func() *big.Int

	// calls tracks calls to the methods.
	calls struct {
		// AccessList holds details about calls to the AccessList method.
		AccessList []struct {
		}
		// Data holds details about calls to the Data method.
		Data []struct {
		}
		// From holds details about calls to the From method.
		From []struct {
		}
		// Gas holds details about calls to the Gas method.
		Gas []struct {
		}
		// GasFeeCap holds details about calls to the GasFeeCap method.
		GasFeeCap []struct {
		}
		// GasPrice holds details about calls to the GasPrice method.
		GasPrice []struct {
		}
		// GasTipCap holds details about calls to the GasTipCap method.
		GasTipCap []struct {
		}
		// IsFake holds details about calls to the IsFake method.
		IsFake []struct {
		}
		// Nonce holds details about calls to the Nonce method.
		Nonce []struct {
		}
		// To holds details about calls to the To method.
		To []struct {
		}
		// Value holds details about calls to the Value method.
		Value []struct {
		}
	}
	lockAccessList sync.RWMutex
	lockData       sync.RWMutex
	lockFrom       sync.RWMutex
	lockGas        sync.RWMutex
	lockGasFeeCap  sync.RWMutex
	lockGasPrice   sync.RWMutex
	lockGasTipCap  sync.RWMutex
	lockIsFake     sync.RWMutex
	lockNonce      sync.RWMutex
	lockTo         sync.RWMutex
	lockValue      sync.RWMutex
}

// AccessList calls AccessListFunc.
func (mock *MessageMock) AccessList() types.AccessList {
	if mock.AccessListFunc == nil {
		panic("MessageMock.AccessListFunc: method is nil but Message.AccessList was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAccessList.Lock()
	mock.calls.AccessList = append(mock.calls.AccessList, callInfo)
	mock.lockAccessList.Unlock()
	return mock.AccessListFunc()
}

// AccessListCalls gets all the calls that were made to AccessList.
// Check the length with:
//
//	len(mockedMessage.AccessListCalls())
func (mock *MessageMock) AccessListCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAccessList.RLock()
	calls = mock.calls.AccessList
	mock.lockAccessList.RUnlock()
	return calls
}

// Data calls DataFunc.
func (mock *MessageMock) Data() []byte {
	if mock.DataFunc == nil {
		panic("MessageMock.DataFunc: method is nil but Message.Data was just called")
	}
	callInfo := struct {
	}{}
	mock.lockData.Lock()
	mock.calls.Data = append(mock.calls.Data, callInfo)
	mock.lockData.Unlock()
	return mock.DataFunc()
}

// DataCalls gets all the calls that were made to Data.
// Check the length with:
//
//	len(mockedMessage.DataCalls())
func (mock *MessageMock) DataCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockData.RLock()
	calls = mock.calls.Data
	mock.lockData.RUnlock()
	return calls
}

// From calls FromFunc.
func (mock *MessageMock) From() common.Address {
	if mock.FromFunc == nil {
		panic("MessageMock.FromFunc: method is nil but Message.From was just called")
	}
	callInfo := struct {
	}{}
	mock.lockFrom.Lock()
	mock.calls.From = append(mock.calls.From, callInfo)
	mock.lockFrom.Unlock()
	return mock.FromFunc()
}

// FromCalls gets all the calls that were made to From.
// Check the length with:
//
//	len(mockedMessage.FromCalls())
func (mock *MessageMock) FromCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockFrom.RLock()
	calls = mock.calls.From
	mock.lockFrom.RUnlock()
	return calls
}

// Gas calls GasFunc.
func (mock *MessageMock) Gas() uint64 {
	if mock.GasFunc == nil {
		panic("MessageMock.GasFunc: method is nil but Message.Gas was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGas.Lock()
	mock.calls.Gas = append(mock.calls.Gas, callInfo)
	mock.lockGas.Unlock()
	return mock.GasFunc()
}

// GasCalls gets all the calls that were made to Gas.
// Check the length with:
//
//	len(mockedMessage.GasCalls())
func (mock *MessageMock) GasCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGas.RLock()
	calls = mock.calls.Gas
	mock.lockGas.RUnlock()
	return calls
}

// GasFeeCap calls GasFeeCapFunc.
func (mock *MessageMock) GasFeeCap() *big.Int {
	if mock.GasFeeCapFunc == nil {
		panic("MessageMock.GasFeeCapFunc: method is nil but Message.GasFeeCap was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGasFeeCap.Lock()
	mock.calls.GasFeeCap = append(mock.calls.GasFeeCap, callInfo)
	mock.lockGasFeeCap.Unlock()
	return mock.GasFeeCapFunc()
}

// GasFeeCapCalls gets all the calls that were made to GasFeeCap.
// Check the length with:
//
//	len(mockedMessage.GasFeeCapCalls())
func (mock *MessageMock) GasFeeCapCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGasFeeCap.RLock()
	calls = mock.calls.GasFeeCap
	mock.lockGasFeeCap.RUnlock()
	return calls
}

// GasPrice calls GasPriceFunc.
func (mock *MessageMock) GasPrice() *big.Int {
	if mock.GasPriceFunc == nil {
		panic("MessageMock.GasPriceFunc: method is nil but Message.GasPrice was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGasPrice.Lock()
	mock.calls.GasPrice = append(mock.calls.GasPrice, callInfo)
	mock.lockGasPrice.Unlock()
	return mock.GasPriceFunc()
}

// GasPriceCalls gets all the calls that were made to GasPrice.
// Check the length with:
//
//	len(mockedMessage.GasPriceCalls())
func (mock *MessageMock) GasPriceCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGasPrice.RLock()
	calls = mock.calls.GasPrice
	mock.lockGasPrice.RUnlock()
	return calls
}

// GasTipCap calls GasTipCapFunc.
func (mock *MessageMock) GasTipCap() *big.Int {
	if mock.GasTipCapFunc == nil {
		panic("MessageMock.GasTipCapFunc: method is nil but Message.GasTipCap was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGasTipCap.Lock()
	mock.calls.GasTipCap = append(mock.calls.GasTipCap, callInfo)
	mock.lockGasTipCap.Unlock()
	return mock.GasTipCapFunc()
}

// GasTipCapCalls gets all the calls that were made to GasTipCap.
// Check the length with:
//
//	len(mockedMessage.GasTipCapCalls())
func (mock *MessageMock) GasTipCapCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGasTipCap.RLock()
	calls = mock.calls.GasTipCap
	mock.lockGasTipCap.RUnlock()
	return calls
}

// IsFake calls IsFakeFunc.
func (mock *MessageMock) IsFake() bool {
	if mock.IsFakeFunc == nil {
		panic("MessageMock.IsFakeFunc: method is nil but Message.IsFake was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsFake.Lock()
	mock.calls.IsFake = append(mock.calls.IsFake, callInfo)
	mock.lockIsFake.Unlock()
	return mock.IsFakeFunc()
}

// IsFakeCalls gets all the calls that were made to IsFake.
// Check the length with:
//
//	len(mockedMessage.IsFakeCalls())
func (mock *MessageMock) IsFakeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsFake.RLock()
	calls = mock.calls.IsFake
	mock.lockIsFake.RUnlock()
	return calls
}

// Nonce calls NonceFunc.
func (mock *MessageMock) Nonce() uint64 {
	if mock.NonceFunc == nil {
		panic("MessageMock.NonceFunc: method is nil but Message.Nonce was just called")
	}
	callInfo := struct {
	}{}
	mock.lockNonce.Lock()
	mock.calls.Nonce = append(mock.calls.Nonce, callInfo)
	mock.lockNonce.Unlock()
	return mock.NonceFunc()
}

// NonceCalls gets all the calls that were made to Nonce.
// Check the length with:
//
//	len(mockedMessage.NonceCalls())
func (mock *MessageMock) NonceCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockNonce.RLock()
	calls = mock.calls.Nonce
	mock.lockNonce.RUnlock()
	return calls
}

// To calls ToFunc.
func (mock *MessageMock) To() *common.Address {
	if mock.ToFunc == nil {
		panic("MessageMock.ToFunc: method is nil but Message.To was just called")
	}
	callInfo := struct {
	}{}
	mock.lockTo.Lock()
	mock.calls.To = append(mock.calls.To, callInfo)
	mock.lockTo.Unlock()
	return mock.ToFunc()
}

// ToCalls gets all the calls that were made to To.
// Check the length with:
//
//	len(mockedMessage.ToCalls())
func (mock *MessageMock) ToCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTo.RLock()
	calls = mock.calls.To
	mock.lockTo.RUnlock()
	return calls
}

// Value calls ValueFunc.
func (mock *MessageMock) Value() *big.Int {
	if mock.ValueFunc == nil {
		panic("MessageMock.ValueFunc: method is nil but Message.Value was just called")
	}
	callInfo := struct {
	}{}
	mock.lockValue.Lock()
	mock.calls.Value = append(mock.calls.Value, callInfo)
	mock.lockValue.Unlock()
	return mock.ValueFunc()
}

// ValueCalls gets all the calls that were made to Value.
// Check the length with:
//
//	len(mockedMessage.ValueCalls())
func (mock *MessageMock) ValueCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockValue.RLock()
	calls = mock.calls.Value
	mock.lockValue.RUnlock()
	return calls
}