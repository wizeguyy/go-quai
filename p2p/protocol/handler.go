package protocol

import (
	"errors"
	"io"

	"github.com/dominant-strategies/go-quai/common"
	"github.com/dominant-strategies/go-quai/log"
	"github.com/dominant-strategies/go-quai/p2p/pb"
	"github.com/libp2p/go-libp2p/core/network"
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
			if errors.Is(err, io.EOF) {
				log.Debugf("stream closed by peer %s", stream.Conn().RemotePeer())
				break
			}

			log.Errorf("error reading message from stream: %s", err)
			// TODO: handle error
			break
		}
		action, slice, hash, err := pb.DecodeQuaiRequest(data)
		if err != nil {
			log.Errorf("error decoding quai request: %s", err)
			// TODO: handle error
			break
		}
		log.Tracef("Decoded Quai Request - action: %s, slice: %+v, hash: %s", action, slice, hash)

<<<<<<< HEAD
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
=======
		switch action {
		case pb.QuaiRequestMessage_REQUEST_BLOCK:
			log.Debugf("Received block request for slice %+v and hash %s", slice, hash)
			handleBlockRequest(slice, hash, stream, node)
		case pb.QuaiRequestMessage_REQUEST_HEADER:
			log.Debugf("Received header request for slice %+v and hash %s", slice, hash)
			handleHeaderRequest(slice, hash, stream, node)
		case pb.QuaiRequestMessage_REQUEST_TRANSACTION:
			handleTransactionRequest(slice, hash, stream, node)
>>>>>>> 10cdfc288 (update protocol handler to use new protobuf API)
		default:
			log.Errorf("invalid action type: %s", action)
			// TODO: handle error
			return
		}
	}
}

// Seeks the block in the cache or database and sends it to the peer in a pb.QuaiResponseMessage
func handleBlockRequest(slice *types.SliceID, hash *common.Hash, stream network.Stream, node QuaiP2PNode) {
	// check if we have the block in our cache or database
	block := node.GetBlock(*hash, *slice)
	if block == nil {
		// TODO: handle block not found
		log.Warnf("block not found")
		return
	}
	// create a Quai Message Response with the block
	action := pb.QuaiResponseMessage_RESPONSE_BLOCK
	data, err := pb.EncodeQuaiResponse(action, block)
	if err != nil {
		log.Errorf("error encoding quai response: %s", err)
		return
	}
	err = common.WriteMessageToStream(stream, data)
	if err != nil {
		log.Errorf("error writing message to stream: %s", err)
		// TODO: handle error
		return
	}
	log.Debugf("Sent block %s to peer %s", block.Hash, stream.Conn().RemotePeer())
}

// Seeks the header in the cache or database and sends it to the peer in a pb.QuaiResponseMessage
func handleHeaderRequest(slice *types.SliceID, hash *common.Hash, stream network.Stream, node QuaiP2PNode) {
	header := node.GetHeader(*hash, *slice)
	if header == nil {
		// TODO: handle header not found
		log.Warnf("header not found")
		return
	}
	log.Tracef("header found: %+v", header)
	// create a Quai Message Response with the header
	action := pb.QuaiResponseMessage_RESPONSE_HEADER
	data, err := pb.EncodeQuaiResponse(action, header)
	if err != nil {
		log.Errorf("error encoding quai response: %s", err)
		return
	}
	err = common.WriteMessageToStream(stream, data)
	if err != nil {
		log.Errorf("error writing message to stream: %s", err)
		// TODO: handle error
		return
	}
}

func handleTransactionRequest(slice *types.SliceID, hash *common.Hash, stream network.Stream, node QuaiP2PNode) {
	// TODO: implement
}
