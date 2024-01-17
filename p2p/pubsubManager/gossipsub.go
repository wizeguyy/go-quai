package pubsubManager

import (
	"context"
	"errors"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/core/types"
	"github.com/dominant-strategies/go-quai/log"
	"github.com/dominant-strategies/go-quai/p2p/pb"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

const (
	// Data types for gossipsub topics
	C_blockType       = "blocks"
	C_transactionType = "transactions"
	C_headerType      = "headers"
)

var (
	ErrUnsupportedType = errors.New("data type not supported")
)

type PubsubManager struct {
	*pubsub.PubSub
	ctx           context.Context
	subscriptions map[string]*pubsub.Subscription
	topics        map[string]*pubsub.Topic

	// Callback function to handle received data
	onReceived func(interface{})
}

// gets the name of the topic for the given type of data
func TopicName(location common.Location, data interface{}) (string, error) {
	switch data.(type) {
<<<<<<< HEAD
	case types.Block:
		return location.Name() + "/blocks", nil
=======
	case *types.Block:
		return slice.String() + "/" + C_blockType, nil
	case *types.Transaction:
		return slice.String() + "/" + C_transactionType, nil
	case *types.Header:
		return slice.String() + "/" + C_headerType, nil
>>>>>>> 658627191 (update gossip sub to use new protobuf API)
	default:
		log.Debugf("unsupported data type: %T", data)
		return "", ErrUnsupportedType
	}
}

// creates a new gossipsub instance
// TODO: what options do we need for quai network? See:
// See https://pkg.go.dev/github.com/libp2p/go-libp2p-pubsub@v0.10.0#Option
func NewGossipSubManager(ctx context.Context, h host.Host) (*PubsubManager, error) {
	ps, err := pubsub.NewGossipSub(ctx, h)
	if err != nil {
		return nil, err
	}
	return &PubsubManager{
		ps,
		ctx,
		make(map[string]*pubsub.Subscription),
		make(map[string]*pubsub.Topic),
		nil,
	}, nil
}

func (g *PubsubManager) Start(receiveCb func(interface{})) {
	g.onReceived = receiveCb
	go g.handleSubscriptions()
}

// subscribe to broadcasts of the given type of data
func (g *PubsubManager) Subscribe(location common.Location, data interface{}) error {
	// build topic name
	topicName, err := TopicName(location, data)
	if err != nil {
		return err
	}

	// join the topic
	topic, err := g.Join(topicName)
	if err != nil {
		return err
	}
	g.topics[topicName] = topic

	// subscribe to the topic
	subscription, err := topic.Subscribe()
	if err != nil {
		return err
	}
	g.subscriptions[topicName] = subscription

	return nil
}

// broadcasts data to subscribing peers
func (g *PubsubManager) Broadcast(location common.Location, data interface{}) error {
	topicName, err := TopicName(location, data)
	if err != nil {
		return err
	}

	// verify we are subscribed to the topic
	if _, ok := g.subscriptions[topicName]; !ok {
		return errors.New("not subscribed to topic: " + topicName)
	}

	// marshal the data
	pbData, err := pb.ConvertAndMarshal(data)
	if err != nil {
		return err
	}
	log.Debugf("publishing data to topic: %s", topicName)
	return g.topics[topicName].Publish(g.ctx, pbData)
}

// lists our peers which provide the associated topic
func (g *PubsubManager) PeersForTopic(location common.Location, data interface{}) ([]peer.ID, error) {
	topicName, err := TopicName(location, data)
	if err != nil {
		return nil, err
	}
	return g.topics[topicName].ListPeers(), nil
}

// handles any data received on any of our subscribed topics
func (g *PubsubManager) handleSubscriptions() {
	for {
		//! TODO: consider using a context with a timeout here or goroutines with select
		for _, sub := range g.subscriptions {
			log.Debugf("waiting for next message on subscription: %s", sub.Topic())
			msg, err := sub.Next(g.ctx)
			if err != nil {
				// if context was cancelled, then we are shutting down
				if g.ctx.Err() != nil {
					return
				}
				log.Errorf("error getting next message from subscription: %s", err)
				continue
			}

			var data interface{}
			// unmarshal the received data depending on the topic's type
			topicType := getTopicType(*msg.Topic)
			log.Debugf("received message on topic: %s", *msg.Topic)
			switch topicType {
			case C_blockType:
				block := types.Block{}
				err = pb.UnmarshalAndConvert(msg.Data, &block)
				if err != nil {
					log.Errorf("error unmarshalling block: %s", err)
					continue
				}
				log.Tracef("received block: %+v", block)
				data = block
			case C_headerType:
				header := types.Header{}
				err = pb.UnmarshalAndConvert(msg.Data, &header)
				if err != nil {
					log.Errorf("error unmarshalling header: %s", err)
					continue
				}
				log.Tracef("received header: %+v", header)
				data = header
			default:
				log.Errorf("unknown topic type: %s", topicType)
				continue
			}

			// handle the received data
			if g.onReceived != nil {
				g.onReceived(data)
			}
		}
	}
}
