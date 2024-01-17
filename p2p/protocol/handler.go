package protocol

import (
	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/log"
	"github.com/dominant-strategies/go-quai/p2p/pb"
	"github.com/libp2p/go-libp2p/core/network"
	"google.golang.org/protobuf/proto"
)

func QuaiProtocolHandler(stream network.Stream, node QuaiP2PNode) {
	defer stream.Close()

	log.Debugf("Received a new stream from %s", stream.Conn().RemotePeer())

	// if there is a protocol mismatch, close the stream
	if stream.Protocol() != ProtocolVersion {
		log.Warnf("Invalid protocol: %s", stream.Protocol())
		// TODO: add logic to drop the peer
		return
	}

	// Enter the read loop for the stream and handle messages
	for {
		data, err := common.ReadMessageFromStream(stream)
		if err != nil {
			log.Errorf("error reading message from stream: %s", err)
			return
		}

		var protoMessage proto.Message
		err = pb.UnmarshalProtoMessage(data, protoMessage)
		if err != nil {
			log.Errorf("error unmarshalling message: %s", err)
			return
		}

		switch msg := protoMessage.(type) {
		case *pb.BlockRequest:
			// get the hash from the block request
<<<<<<< HEAD
			blockReq := msg
			hash := common.HexToHash(blockReq.Hash)
			// get the location from the block request
			location := blockReq.Location
=======
			hash := types.Hash{}
			pbHash := msg.GetHash()
			if pbHash == nil {
				log.Errorf("block request did not contain a hash")
				// handle error
				return
			}
			hash.FromProto(pbHash)

			// get the slice from the block request
			slice := types.SliceID{}
			pbSlice := msg.GetSliceId()
			if pbSlice == nil {
				log.Errorf("block request did not contain a slice")
				// handle error
				return
			}
			slice.FromProto(pbSlice)
>>>>>>> c684e256a (refactor QuaiProtocolHandler() to use protobuf generic functions)

			// check if we have the block in our cache
			block := node.GetBlock(hash, []byte(location))
			if block == nil {
				// TODO: handle block not found
				log.Warnf("block not found")
				return
			}
			// convert the block to a protocol buffer and send it back to the peer
			data, err := pb.ConvertAndMarshal(block)
			if err != nil {
				log.Errorf("error marshalling block: %s", err)
				// TODO: handle error
				return
			}
			err = common.WriteMessageToStream(stream, data)
			if err != nil {
				log.Errorf("error writing message to stream: %s", err)
				// TODO: handle error
				return
			}
			log.Debugf("Sent block %s to peer %s", block.Hash, stream.Conn().RemotePeer())

		case *pb.QuaiProtocolMessage:
			// TODO: handle quai protocol message
		default:
			log.Errorf("unknown message type received: %s", msg)
			// TODO: handle unknown message type
		}
	}
}
