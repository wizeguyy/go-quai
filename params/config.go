// Cojyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
	"math/big"

	"github.com/dominant-strategies/go-quai/common"
)

// Genesis hashes to enforce below configs on.
var (
	// Progpow GenesisHashes
	ProgpowColosseumGenesisHash = common.HexToHash("0x8e484a0047b01b33cfc6769f73c70e695ab27ec89b99c988b31641cb86e3bc3b")
	ProgpowGardenGenesisHash    = common.HexToHash("0xb0bd2ac6e04044cbe4e1e1113dc511d0a3d3d52c55bfd7240cffeb15f8b5ca1d")
	ProgpowOrchardGenesisHash   = common.HexToHash("0x7ea3f58395328b0ceec45b1254cd00c0d2200804c92e0de0cf0b7dbaec15beda")
	ProgpowLocalGenesisHash     = common.HexToHash("0xa03d19a0bf028c214def9621a8495c5182275be113e073767c8d2cbe33081315")
	ProgpowGalenaGenesisHash    = common.HexToHash("0x8e484a0047b01b33cfc6769f73c70e695ab27ec89b99c988b31641cb86e3bc3b")

	// Blake3GenesisHashes
	Blake3PowColosseumGenesisHash = common.HexToHash("0x5746089cbee3cde719c3d0599e31504793028c68e5df7acff956e917f72866f5")
	Blake3PowGardenGenesisHash    = common.HexToHash("0x9a6a442ca5fe302c4f6ab3cb5ed81d3cdbddd0bde28a55ad864ac84cda07fea9")
	Blake3PowOrchardGenesisHash   = common.HexToHash("0xddf6d6a29a05fc646df65a4965f6b6ebe40fa0c27189c5fac01e31b1f2763b37")
	Blake3PowLocalGenesisHash     = common.HexToHash("0x279e2d67eb5828010e67f752c7686f2c963df7cc41cc24f19e6610612fff42f1")
	Blake3PowGalenaGenesisHash    = common.HexToHash("0x9a6a442ca5fe302c4f6ab3cb5ed81d3cdbddd0bde28a55ad864ac84cda07fea9")
)

var (
	// ColosseumChainConfig is the chain parameters to run a node on the Colosseum network.
	ProgpowColosseumChainConfig = &ChainConfig{
		ChainID:     big.NewInt(9000),
		Progpow:     new(ProgpowConfig),
		GenesisHash: ProgpowColosseumGenesisHash,
	}

	Blake3PowColosseumChainConfig = &ChainConfig{
		ChainID:     big.NewInt(9000),
		Blake3Pow:   new(Blake3powConfig),
		GenesisHash: Blake3PowColosseumGenesisHash,
	}

	// GardenChainConfig contains the chain parameters to run a node on the Garden test network.
	ProgpowGardenChainConfig = &ChainConfig{
		ChainID:     big.NewInt(12000),
		Progpow:     new(ProgpowConfig),
		GenesisHash: ProgpowGardenGenesisHash,
	}

	Blake3PowGardenChainConfig = &ChainConfig{
		ChainID:     big.NewInt(12000),
		Blake3Pow:   new(Blake3powConfig),
		GenesisHash: Blake3PowGardenGenesisHash,
	}

	// OrchardChainConfig contains the chain parameters to run a node on the Orchard test network.
	ProgpowOrchardChainConfig = &ChainConfig{
		ChainID:     big.NewInt(15000),
		Progpow:     new(ProgpowConfig),
		GenesisHash: ProgpowOrchardGenesisHash,
	}

	Blake3PowOrchardChainConfig = &ChainConfig{
		ChainID:     big.NewInt(15000),
		Blake3Pow:   new(Blake3powConfig),
		GenesisHash: Blake3PowOrchardGenesisHash,
	}

	// GalenaChainConfig contains the chain parameters to run a node on the Galena test network.
	ProgpowGalenaChainConfig = &ChainConfig{
		ChainID:     big.NewInt(17000),
		Blake3Pow:   new(Blake3powConfig),
		Progpow:     new(ProgpowConfig),
		GenesisHash: ProgpowGalenaGenesisHash,
	}

	Blake3PowGalenaChainConfig = &ChainConfig{
		ChainID:     big.NewInt(17000),
		Blake3Pow:   new(Blake3powConfig),
		GenesisHash: Blake3PowGalenaGenesisHash,
	}

	// LocalChainConfig contains the chain parameters to run a node on the Local test network.
	ProgpowLocalChainConfig = &ChainConfig{
		ChainID:     big.NewInt(1337),
		Progpow:     new(ProgpowConfig),
		GenesisHash: ProgpowLocalGenesisHash,
	}

	Blake3PowLocalChainConfig = &ChainConfig{
		ChainID:     big.NewInt(1337),
		Blake3Pow:   new(Blake3powConfig),
		GenesisHash: Blake3PowLocalGenesisHash,
	}

	// AllProgpowProtocolChanges contains every protocol change introduced
	// and accepted by the Quai core developers into the Progpow consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllProgpowProtocolChanges = &ChainConfig{big.NewInt(1337), "progpow", new(Blake3powConfig), new(ProgpowConfig), common.Hash{}}

	TestChainConfig = &ChainConfig{big.NewInt(1), "progpow", new(Blake3powConfig), new(ProgpowConfig), common.Hash{}}
	TestRules       = TestChainConfig.Rules(new(big.Int))
)

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection
	// Various consensus engines
	ConsensusEngine string
	Blake3Pow       *Blake3powConfig `json:"blake3pow,omitempty"`
	Progpow         *ProgpowConfig   `json:"progpow,omitempty"`
	GenesisHash     common.Hash
}

// Blake3powConfig is the consensus engine configs for proof-of-work based sealing.
type Blake3powConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *Blake3powConfig) String() string {
	return "blake3pow"
}

// ProgpowConfig is the consensus engine configs for proof-of-work based sealing.
type ProgpowConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *ProgpowConfig) String() string {
	return "progpow"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Blake3Pow != nil:
		engine = c.Blake3Pow
	case c.Progpow != nil:
		engine = c.Progpow
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v, Engine: %v}",
		c.ChainID,
		engine,
	)
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID *big.Int
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID: new(big.Int).Set(chainID),
	}
}
