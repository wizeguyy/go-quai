package node

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/dominant-strategies/go-quai/cmd/utils"
	"github.com/dominant-strategies/go-quai/core/types"
	"github.com/dominant-strategies/go-quai/log"
	"github.com/dominant-strategies/go-quai/p2p"
	quaiprotocol "github.com/dominant-strategies/go-quai/p2p/protocol"
	"github.com/dominant-strategies/go-quai/quai"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

// Starts the node and all of its services
func (p *P2PNode) Start() error {
	log.Infof("starting P2P node...")

	// Start any async processes belonging to this node
	log.Debugf("starting node processes...")
	go p.eventLoop()
	go p.statsLoop()

	// Is this node expected to have bootstrap peers to dial?
	if !viper.GetBool(utils.BootNodeFlag.Name) && !viper.GetBool(utils.SoloFlag.Name) && len(p.bootpeers) == 0 {
		err := errors.New("no bootpeers provided. Unable to join network")
		log.Errorf("%s", err)
		return err
	}

	// Register the Quai protocol handler
	p.SetStreamHandler(quaiprotocol.ProtocolVersion, func(s network.Stream) {
		quaiprotocol.QuaiProtocolHandler(s, p)
	})

	// If the node is a bootnode, start the bootnode service
	if viper.GetBool(utils.BootNodeFlag.Name) {
		log.Infof("starting node as a bootnode...")
		return nil
	}

	// Start the pubsub manager
	p.pubsub.Start(p.handleBroadcast)

	// Open data streams with connected Quai peers
	go quaiprotocol.OpenPeerStreams(p)

	return nil
}

func (p *P2PNode) Subscribe(location common.Location, data interface{}) error {
	return p.pubsub.Subscribe(location, data)
}

func (p *P2PNode) Broadcast(location common.Location, data interface{}) error {
	return p.pubsub.Broadcast(location, data)
}

func (p *P2PNode) SetConsensusBackend(be quai.ConsensusAPI) {
	p.consensus = be
}

type stopFunc func() error

// Function to gracefully shtudown all running services
func (p *P2PNode) Stop() error {
	// define a list of functions to stop the services the node is running
	stopFuncs := []stopFunc{
		p.Host.Close,
		p.dht.Close,
	}
	// create a channel to collect errors
	errs := make(chan error, len(stopFuncs))
	// run each stop function in a goroutine
	for _, fn := range stopFuncs {
		go func(fn stopFunc) {
			errs <- fn()
		}(fn)
	}

	var allErrors []error
	for i := 0; i < len(stopFuncs); i++ {
		select {
		case err := <-errs:
			if err != nil {
				log.Errorf("error during shutdown: %s", err)
				allErrors = append(allErrors, err)
			}
		case <-time.After(5 * time.Second):
			err := errors.New("timeout during shutdown")
			log.Warnf("error: %s", err)
			allErrors = append(allErrors, err)
		}
	}
	close(errs)
	if len(allErrors) > 0 {
		return errors.Errorf("errors during shutdown: %v", allErrors)
	} else {
		return nil
	}
}

// Request a block from the network for the specified location
func (p *P2PNode) RequestBlock(hash common.Hash, location common.Location) chan *types.Block {
	resultChan := make(chan *types.Block, 1)
	go func() {
		defer close(resultChan)
		var block *types.Block
		// 1. Check if the block is in the local cache
		block, ok := p.blockCache.Get(hash)
		if ok {
			log.Debugf("Block %s found in cache", hash)
			resultChan <- block
			return
		}
		// 2. If not, query the topic peers for the block
		peers, err := p.pubsub.PeersForTopic(location, types.Block{})
		if err != nil {
			log.Errorf("Error requesting block: ", err)
			return
		}
		for _, peerID := range peers {
			block, err := p.requestBlockFromPeer(hash, location, peerID)
			if err == nil {
				log.Debugf("Received block %s from peer %s", block.Hash, peerID)
				// add the block to the cache
				p.blockCache.Add(hash, block)
				// send the block to the result channel
				resultChan <- block
				// announce ourselves as a provider of the block topic
				err = p.announceToDHT(slice, types.Block{}, block)
				if err != nil {
					log.Errorf("Error announcing block: ", err)
				}
				return
			}
		}

		// 3. If block is not found, query the DHT for peers in the slice
		// TODO: evaluate making this configurable
		const (
			maxDHTQueryRetries    = 3  // Maximum number of retries for DHT queries
			peersPerDHTQuery      = 10 // Number of peers to query per DHT attempt
			dhtQueryRetryInterval = 5  // Time to wait between DHT query retries
		)
<<<<<<< HEAD
		// create a Cid from the slice location
		shardCid := locationToCid(location)
=======
		// create a Cid from the slice ID
		topic, err := p.pubsub.TopicName(slice, types.Block{})
		cid := topicToCid(topic)
		if err != nil {
			log.Errorf("Error creating Cid from slice ID: ", err)
			return
		}
>>>>>>> 1725839ac (announce to the DHT everytime we have a block)
		for retries := 0; retries < maxDHTQueryRetries; retries++ {
			log.Debugf("Querying DHT for slice Cid %s (retry %d)", cid, retries)
			// query the DHT for peers in the slice
			peerChan := p.dht.FindProvidersAsync(p.ctx, cid, peersPerDHTQuery)
			for peerInfo := range peerChan {
				block, err := p.requestBlockFromPeer(hash, location, peerInfo.ID)
				if err == nil {
					log.Debugf("Received block %s from peer %s", block.Hash, peerInfo.ID)
					p.blockCache.Add(hash, block)
					resultChan <- block
					// announce ourselves as a provider of the block topic
					err = p.announceToDHT(slice, types.Block{}, block)
					if err != nil {
						log.Errorf("Error announcing block: ", err)
					}
					return
				}
			}
			// if the block is not found, wait for a bit and try again
			log.Debugf("Block %s not found in slice %s. Retrying...", hash, location)
			time.Sleep(dhtQueryRetryInterval * time.Second)
		}
		log.Debugf("Block %s not found in slice %s", hash, location)
	}()
	return resultChan
}

func (p *P2PNode) RequestTransaction(hash common.Hash, loc common.Location) chan *types.Transaction {
	panic("todo")
}

func (p *P2PNode) ReportBadPeer(peer p2p.PeerID) {
	panic("todo")
}

// Returns the list of bootpeers
func (p *P2PNode) GetBootPeers() []peer.AddrInfo {
	return p.bootpeers
}

// Opens a new stream to the given peer using the given protocol ID
func (p *P2PNode) NewStream(peerID peer.ID, protocolID protocol.ID) (network.Stream, error) {
	return p.Host.NewStream(p.ctx, peerID, protocolID)
}

// Connects to the given peer
func (p *P2PNode) Connect(pi peer.AddrInfo) error {
	return p.Host.Connect(p.ctx, pi)
}

// Start gossipsub protocol
func (p *P2PNode) StartGossipSub(ctx context.Context) error {
	return nil
}

// Search for a block in the node's cache, or query the consensus backend if it's not found in cache.
// Returns nil if the block is not found.
func (p *P2PNode) GetBlock(hash common.Hash, location common.Location) *types.Block {
	block, ok := p.blockCache.Get(hash)
	if ok {
		return block
	}
	return p.consensus.LookupBlock(hash, location)
}

func (p *P2PNode) handleBroadcast(data interface{}) {
	switch v := data.(type) {
	case types.Block:
		p.blockCache.Add(v.Hash(), &v)
	// TODO: send it to consensus
	default:
		log.Debugf("received unsupported block broadcast")
		// TODO: ban the peer which sent it?
	}
}

