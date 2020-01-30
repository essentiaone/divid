// Copyright (c) 2017 The essentiaone developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/essentiaone/divid/blockchain"
	"github.com/essentiaone/divid/chaincfg"
	"github.com/essentiaone/divid/chaincfg/chainhash"
	"github.com/essentiaone/divid/mempool"
	"github.com/essentiaone/divid/peer"
	"github.com/essentiaone/divid/wire"
	"github.com/essentiaone/btcutil"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *btcutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockchain.BlockChain
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int

	FeeEstimator *mempool.FeeEstimator
}
