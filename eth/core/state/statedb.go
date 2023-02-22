// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package state

import (
	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/state/journal"
	coretypes "pkg.berachain.dev/stargazer/eth/core/types"
	"pkg.berachain.dev/stargazer/eth/core/vm"
	"pkg.berachain.dev/stargazer/eth/crypto"
	"pkg.berachain.dev/stargazer/eth/params"
	"pkg.berachain.dev/stargazer/lib/snapshot"
	libtypes "pkg.berachain.dev/stargazer/lib/types"
)

var (
	// `emptyCodeHash` is the Keccak256 Hash of empty code
	// 0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470.
	emptyCodeHash = crypto.Keccak256Hash(nil)
)

// `stateDB` is a struct that holds the plugins and controller to manage Ethereum state.
type stateDB struct {
	// Plugin is injected by the chain running the Stargazer EVM.
	Plugin

	// Journals built internally and required for the stateDB.
	LogsJournal
	RefundJournal
	//TODO: Add a journal for the `accesslist journal` slice.

	// `ctrl` is used to manage snapshots and reverts across plugins and journals.
	ctrl libtypes.Controller[string, libtypes.Controllable[string]]

	// Dirty tracking of suicided accounts, we have to keep track of these manually, in order
	// for the code and state to still be accessible even after the account has been deleted.
	// We chose to keep track of them in a separate slice, rather than a map, because the
	// number of accounts that will be suicided in a single transaction is expected to be
	// very low.
	suicides []common.Address
}

// `NewStateDB` returns a `vm.StargazerStateDB` with the given `StatePlugin`.
func NewStateDB(sp Plugin) vm.StargazerStateDB {
	// Build the journals required for the stateDB
	lj := journal.NewLogs()
	rj := journal.NewRefund()

	// Build the controller and register the plugins and journals

	// TODO: journal registration
	ctrl := snapshot.NewController[string, libtypes.Controllable[string]]()
	_ = ctrl.Register(lj)
	_ = ctrl.Register(rj)
	_ = ctrl.Register(sp)

	return &stateDB{
		Plugin:        sp,
		LogsJournal:   lj,
		RefundJournal: rj,
		ctrl:          ctrl,
		suicides:      make([]common.Address, 1), // very rare to suicide, so we alloc 1 slot.
	}
}

// =============================================================================
// Suicide
// =============================================================================

// Suicide implements the StargazerStateDB interface by marking the given address as suicided.
// This clears the account balance, but the code and state of the address remains available
// until after Commit is called.
func (sdb *stateDB) Suicide(addr common.Address) bool {
	// only smart contracts can commit suicide
	ch := sdb.GetCodeHash(addr)
	if (ch == common.Hash{}) || ch == emptyCodeHash {
		return false
	}

	// Reduce it's balance to 0.
	sdb.SubBalance(addr, sdb.GetBalance(addr))

	// Mark the underlying account for deletion in `Commit()`.
	sdb.suicides = append(sdb.suicides, addr)
	return true
}

// `HasSuicided` implements the `StargazerStateDB` interface by returning if the contract was suicided
// in current transaction.
func (sdb *stateDB) HasSuicided(addr common.Address) bool {
	for _, suicide := range sdb.suicides {
		if addr == suicide {
			return true
		}
	}
	return false
}

// `Empty` implements the `StargazerStateDB` interface by returning whether the state object
// is either non-existent or empty according to the EIP161 epecification
// (balance = nonce = code = 0)
// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-161.md
func (sdb *stateDB) Empty(addr common.Address) bool {
	ch := sdb.GetCodeHash(addr)
	return sdb.GetNonce(addr) == 0 &&
		(ch == emptyCodeHash || ch == common.Hash{}) &&
		sdb.GetBalance(addr).Sign() == 0
}

// =============================================================================
// Snapshot
// =============================================================================

// `Snapshot` implements `stateDB`.
func (sdb *stateDB) Snapshot() int {
	return sdb.ctrl.Snapshot()
}

// `RevertToSnapshot` implements `stateDB`.
func (sdb *stateDB) RevertToSnapshot(id int) {
	sdb.ctrl.RevertToSnapshot(id)
}

// =============================================================================
// Finalize
// =============================================================================

// `Finalize` deletes the suicided accounts, clears the suicides list, and finalizes all plugins.
func (sdb *stateDB) Finalize() {
	sdb.DeleteSuicides(sdb.suicides)
	sdb.suicides = make([]common.Address, 1)
	sdb.ctrl.Finalize()
}

// =============================================================================
// AccessList
// =============================================================================

// TODO: implement `AddAddressToAccessList`.
func (sdb *stateDB) AddAddressToAccessList(addr common.Address) {
	panic("not supported by Stargazer")
}

// TODO: implement `AddSlotToAccessList`
func (sdb *stateDB) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	panic("not supported by Stargazer")
}

// TODO: implement `AddressInAccessList`
func (sdb *stateDB) AddressInAccessList(addr common.Address) bool {
	return false
}

// TODO: implement `SlotInAccessList`
func (sdb *stateDB) SlotInAccessList(addr common.Address, slot common.Hash) (bool, bool) {
	return false, false
}

// TODO: `GetTransientState` implements the `StargazerStateDB` interface by returning the transient state
func (sdb *stateDB) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	panic("not supported by Stargazer")
}

// TODO: `SetTransientState` implements the `StargazerStateDB` interface by setting the transient state
func (sdb *stateDB) SetTransientState(addr common.Address, key, value common.Hash) {
	panic("not supported by Stargazer")
}

// TODO: `Prepare` implements the `StargazerStateDB` interface by preparing the stateDB for a new transaction.
func (sdb *stateDB) Prepare(rules params.Rules, sender, coinbase common.Address,
	dest *common.Address, precompiles []common.Address, txAccesses coretypes.AccessList) {
	panic("not supported by Stargazer")
}

// =============================================================================
// PreImage
// =============================================================================

// AddPreimage implements the the `StateDB“ interface, but currently
// performs a no-op since the EnablePreimageRecording flag is disabled.
func (sdb *stateDB) AddPreimage(hash common.Hash, preimage []byte) {}