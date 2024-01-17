package quai

import (
	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/core/types"
	"github.com/dominant-strategies/go-quai/p2p"
)

// QuaiBackend implements the quai consensus protocol
type QuaiBackend struct {
	p2p NetworkingAPI
<<<<<<< HEAD
=======

	runningSlices map[types.SliceID]*types.SliceID
>>>>>>> 7ab2b819b (refactor slice usage to avoid build error)
}

// Create a new instance of the QuaiBackend consensus service
func NewQuaiBackend() (*QuaiBackend, error) {
	return &QuaiBackend{}, nil
}

// Assign the p2p client interface to use for interacting with the p2p network
func (qbe *QuaiBackend) SetP2PNode(api NetworkingAPI) {
	qbe.p2p = api
}

// Start the QuaiBackend consensus service
func (qbe *QuaiBackend) Start() error {
	return nil
}

// Handle blocks received from the P2P client
func (qbe *QuaiBackend) OnNewBlock(sourcePeer p2p.PeerID, block types.Block) bool {
	panic("todo")
}

// Handle transactions received from the P2P client
func (qbe *QuaiBackend) OnNewTransaction(sourcePeer p2p.PeerID, tx types.Transaction) bool {
	panic("todo")
}

// Returns the current block height for the given location
func (qbe *QuaiBackend) GetHeight(location common.Location) uint64 {
	// Example/mock implementation
	panic("todo")
}

<<<<<<< HEAD
func (qbe *QuaiBackend) LookupBlock(hash common.Hash, location common.Location) *types.Block {
=======
func (qbe *QuaiBackend) GetSlice(slice types.SliceID) *types.SliceID {
	return qbe.runningSlices[slice]
}

func (qbe *QuaiBackend) GetRunningSlices() map[types.SliceID]*types.SliceID {
	return qbe.runningSlices
}

func (qbe *QuaiBackend) SetRunningSlices(slices []types.SliceID) {
	qbe.runningSlices = make(map[types.SliceID]*types.SliceID)
	for _, slice := range slices {
		qbe.runningSlices[slice] = &slice
	}
}

func (qbe *QuaiBackend) LookupBlock(hash common.Hash, slice types.SliceID) *types.Block {
>>>>>>> 7ab2b819b (refactor slice usage to avoid build error)
	panic("todo")
}
