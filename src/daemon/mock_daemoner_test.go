// Code generated by mockery v1.0.0. DO NOT EDIT.

package daemon

import cipher "github.com/MDLlife/MDL/src/cipher"
import coin "github.com/MDLlife/MDL/src/coin"
import gnet "github.com/MDLlife/MDL/src/daemon/gnet"
import mock "github.com/stretchr/testify/mock"
import pex "github.com/MDLlife/MDL/src/daemon/pex"
import visor "github.com/MDLlife/MDL/src/visor"

// mockDaemoner is an autogenerated mock type for the daemoner type
type mockDaemoner struct {
	mock.Mock
}

// addPeers provides a mock function with given fields: addrs
func (_m *mockDaemoner) addPeers(addrs []string) int {
	ret := _m.Called(addrs)

	var r0 int
	if rf, ok := ret.Get(0).(func([]string) int); ok {
		r0 = rf(addrs)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// announceAllValidTxns provides a mock function with given fields:
func (_m *mockDaemoner) announceAllValidTxns() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// broadcastMessage provides a mock function with given fields: msg
func (_m *mockDaemoner) broadcastMessage(msg gnet.Message) ([]uint64, error) {
	ret := _m.Called(msg)

	var r0 []uint64
	if rf, ok := ret.Get(0).(func(gnet.Message) []uint64); ok {
		r0 = rf(msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gnet.Message) error); ok {
		r1 = rf(msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// connectionIntroduced provides a mock function with given fields: addr, gnetID, m
func (_m *mockDaemoner) connectionIntroduced(addr string, gnetID uint64, m *IntroductionMessage) (*connection, error) {
	ret := _m.Called(addr, gnetID, m)

	var r0 *connection
	if rf, ok := ret.Get(0).(func(string, uint64, *IntroductionMessage) *connection); ok {
		r0 = rf(addr, gnetID, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint64, *IntroductionMessage) error); ok {
		r1 = rf(addr, gnetID, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// daemonConfig provides a mock function with given fields:
func (_m *mockDaemoner) daemonConfig() DaemonConfig {
	ret := _m.Called()

	var r0 DaemonConfig
	if rf, ok := ret.Get(0).(func() DaemonConfig); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(DaemonConfig)
	}

	return r0
}

// disconnectNow provides a mock function with given fields: addr, r
func (_m *mockDaemoner) disconnectNow(addr string, r gnet.DisconnectReason) error {
	ret := _m.Called(addr, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, gnet.DisconnectReason) error); ok {
		r0 = rf(addr, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// executeSignedBlock provides a mock function with given fields: b
func (_m *mockDaemoner) executeSignedBlock(b coin.SignedBlock) error {
	ret := _m.Called(b)

	var r0 error
	if rf, ok := ret.Get(0).(func(coin.SignedBlock) error); ok {
		r0 = rf(b)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// filterKnownUnconfirmed provides a mock function with given fields: txns
func (_m *mockDaemoner) filterKnownUnconfirmed(txns []cipher.SHA256) ([]cipher.SHA256, error) {
	ret := _m.Called(txns)

	var r0 []cipher.SHA256
	if rf, ok := ret.Get(0).(func([]cipher.SHA256) []cipher.SHA256); ok {
		r0 = rf(txns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cipher.SHA256)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]cipher.SHA256) error); ok {
		r1 = rf(txns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getKnownUnconfirmed provides a mock function with given fields: txns
func (_m *mockDaemoner) getKnownUnconfirmed(txns []cipher.SHA256) (coin.Transactions, error) {
	ret := _m.Called(txns)

	var r0 coin.Transactions
	if rf, ok := ret.Get(0).(func([]cipher.SHA256) coin.Transactions); ok {
		r0 = rf(txns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(coin.Transactions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]cipher.SHA256) error); ok {
		r1 = rf(txns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// getSignedBlocksSince provides a mock function with given fields: seq, count
func (_m *mockDaemoner) getSignedBlocksSince(seq uint64, count uint64) ([]coin.SignedBlock, error) {
	ret := _m.Called(seq, count)

	var r0 []coin.SignedBlock
	if rf, ok := ret.Get(0).(func(uint64, uint64) []coin.SignedBlock); ok {
		r0 = rf(seq, count)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]coin.SignedBlock)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(seq, count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// headBkSeq provides a mock function with given fields:
func (_m *mockDaemoner) headBkSeq() (uint64, bool, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func() bool); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// injectTransaction provides a mock function with given fields: txn
func (_m *mockDaemoner) injectTransaction(txn coin.Transaction) (bool, *visor.ErrTxnViolatesSoftConstraint, error) {
	ret := _m.Called(txn)

	var r0 bool
	if rf, ok := ret.Get(0).(func(coin.Transaction) bool); ok {
		r0 = rf(txn)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *visor.ErrTxnViolatesSoftConstraint
	if rf, ok := ret.Get(1).(func(coin.Transaction) *visor.ErrTxnViolatesSoftConstraint); ok {
		r1 = rf(txn)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*visor.ErrTxnViolatesSoftConstraint)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(coin.Transaction) error); ok {
		r2 = rf(txn)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// pexConfig provides a mock function with given fields:
func (_m *mockDaemoner) pexConfig() pex.Config {
	ret := _m.Called()

	var r0 pex.Config
	if rf, ok := ret.Get(0).(func() pex.Config); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pex.Config)
	}

	return r0
}

// recordMessageEvent provides a mock function with given fields: m, c
func (_m *mockDaemoner) recordMessageEvent(m asyncMessage, c *gnet.MessageContext) error {
	ret := _m.Called(m, c)

	var r0 error
	if rf, ok := ret.Get(0).(func(asyncMessage, *gnet.MessageContext) error); ok {
		r0 = rf(m, c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// recordPeerHeight provides a mock function with given fields: addr, gnetID, height
func (_m *mockDaemoner) recordPeerHeight(addr string, gnetID uint64, height uint64) {
	_m.Called(addr, gnetID, height)
}

// requestBlocksFromAddr provides a mock function with given fields: addr
func (_m *mockDaemoner) requestBlocksFromAddr(addr string) error {
	ret := _m.Called(addr)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// sendMessage provides a mock function with given fields: addr, msg
func (_m *mockDaemoner) sendMessage(addr string, msg gnet.Message) error {
	ret := _m.Called(addr, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, gnet.Message) error); ok {
		r0 = rf(addr, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// sendRandomPeers provides a mock function with given fields: addr
func (_m *mockDaemoner) sendRandomPeers(addr string) error {
	ret := _m.Called(addr)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(addr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disconnect provides a mock function with given fields: addr, r
func (_m *mockDaemoner) Disconnect(addr string, r gnet.DisconnectReason) error {
	ret := _m.Called(addr, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, gnet.DisconnectReason) error); ok {
		r0 = rf(addr, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
