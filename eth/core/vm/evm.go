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

package vm

import (
	"pkg.berachain.dev/stargazer/eth/params"
	"pkg.berachain.dev/stargazer/lib/utils"
)

// `StargazerEVM` is the wrapper for the Go-Ethereum EVM.
type stargazerEVM struct {
	*GethEVM
}

// `NewStargazerEVM` creates and returns a new `StargazerEVM`.
func NewStargazerEVM(
	blockCtx BlockContext,
	txCtx TxContext,
	stateDB StargazerStateDB,
	chainConfig *params.ChainConfig,
	config Config,
	pcmgr PrecompileManager,
) StargazerEVM {
	return &stargazerEVM{
		GethEVM: NewGethEVMWithPrecompiles(
			blockCtx, txCtx, stateDB, chainConfig, config, pcmgr,
		),
	}
}

// `UnderlyingEVM` implements `StargazerEVM`.
func (evm *stargazerEVM) UnderlyingEVM() *GethEVM {
	return evm.GethEVM
}

// `SetTxContext` implements `StargazerEVM`.
func (evm *stargazerEVM) SetTxContext(txCtx TxContext) {
	evm.GethEVM.TxContext = txCtx
}

// `StateDB` implements `StargazerEVM`.
func (evm *stargazerEVM) StateDB() StargazerStateDB {
	return utils.MustGetAs[StargazerStateDB](evm.GethEVM.StateDB)
}