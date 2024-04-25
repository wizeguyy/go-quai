// Copyright 2015 The go-ethereum Authors
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
	"math/big"

	"github.com/dominant-strategies/go-quai/common"
)

const (
	GasLimitBoundDivisor            uint64 = 1024 // The bound divisor of the gas limit, used in update calculations.
	PercentGasUsedThreshold         uint64 = 50   // Percent Gas used threshold at which the gas limit adjusts
	GasLimitStepOneBlockThreshold   uint64 = 150000
	GasLimitStepTwoBlockThreshold   uint64 = 300000
	GasLimitStepThreeBlockThreshold uint64 = 450000
	MinGasLimit                     uint64 = 5000000 // Minimum the gas limit may ever be.
	GenesisGasLimit                 uint64 = 5000000 // Gas limit of the Genesis block.

	MaximumExtraDataSize  uint64 = 32    // Maximum size extra data may be after Genesis.
	ExpByteGas            uint64 = 10    // Times ceil(log256(exponent)) for the EXP instruction.
	CallValueTransferGas  uint64 = 9000  // Paid for CALL when the value transfer is non-zero.
	CallNewAccountGas     uint64 = 25000 // Paid for CALL when the destination address didn't exist prior.
	TxGas                 uint64 = 21000 // Per transaction not creating a contract. NOTE: Not payable on data of calls between transactions.
	TxGasContractCreation uint64 = 53000 // Per transaction that creates a contract. NOTE: Not payable on data of calls between transactions.
	TxDataZeroGas         uint64 = 4     // Per byte of data attached to a transaction that equals zero. NOTE: Not payable on data of calls between transactions.
	QuadCoeffDiv          uint64 = 512   // Divisor for the quadratic particle of the memory cost equation.
	LogDataGas            uint64 = 8     // Per byte in a LOG* operation's data.
	CallStipend           uint64 = 2300  // Free gas given at beginning of call.
	ETXGas                uint64 = 21000 // Per ETX generated by opETX or normal cross-chain transfer.
	//  The etx fractions  should be based on the current  expansion number
	ETXRegionMaxFraction int    = 1   // The maximum fraction of transactions for cross-region ETXs
	ETXPrimeMaxFraction  int    = 1   // The maximum fraction of transactions for cross-prime ETXs
	ETXRLimitMin         int    = 10  // Minimum possible cross-region ETX limit
	ETXPLimitMin         int    = 10  // Minimum possible cross-prime ETX limit
	EtxExpirationAge     uint64 = 100 // Number of blocks an ETX may wait for inclusion at the destination

	Sha3Gas     uint64 = 30 // Once per SHA3 operation.
	Sha3WordGas uint64 = 6  // Once per word of the SHA3 operation's data.

	SstoreClearGas  uint64 = 5000  // Once per SSTORE operation if the zeroness doesn't change.
	SstoreRefundGas uint64 = 15000 // Once per SSTORE operation if the zeroness changes to zero.

	NetSstoreNoopGas  uint64 = 200   // Once per SSTORE operation if the value doesn't change.
	NetSstoreInitGas  uint64 = 20000 // Once per SSTORE operation from clean zero.
	NetSstoreCleanGas uint64 = 5000  // Once per SSTORE operation from clean non-zero.
	NetSstoreDirtyGas uint64 = 200   // Once per SSTORE operation from dirty.

	NetSstoreClearRefund      uint64 = 15000 // Once per SSTORE operation for clearing an originally existing storage slot
	NetSstoreResetRefund      uint64 = 4800  // Once per SSTORE operation for resetting to the original non-zero value
	NetSstoreResetClearRefund uint64 = 19800 // Once per SSTORE operation for resetting to the original zero value

	SstoreSentryGas uint64 = 2300  // Minimum gas required to be present for an SSTORE call, not consumed
	SstoreSetGas    uint64 = 20000 // Once per SSTORE operation from clean zero to non-zero
	SstoreResetGas  uint64 = 5000  // Once per SSTORE operation from clean non-zero to something else

	ColdAccountAccessCost = uint64(2600) // COLD_ACCOUNT_ACCESS_COST
	ColdSloadCost         = uint64(2100) // COLD_SLOAD_COST
	WarmStorageReadCost   = uint64(100)  // WARM_STORAGE_READ_COST

	// SSTORE_CLEARS_SCHEDULE is defined as SSTORE_RESET_GAS + ACCESS_LIST_STORAGE_KEY_COST
	// Which becomes: 5000 - 2100 + 1900 = 4800
	SstoreClearsScheduleRefund uint64 = SstoreResetGas - ColdSloadCost + TxAccessListStorageKeyGas // Once per SSTORE operation for clearing an originally existing storage slot

	JumpdestGas   uint64 = 1     // Once per JUMPDEST operation.
	EpochDuration uint64 = 30000 // Duration between proof-of-work epochs.

	CreateDataGas         uint64 = 200   //
	CallCreateDepth       uint64 = 1024  // Maximum depth of call/create stack.
	ExpGas                uint64 = 10    // Once per EXP instruction
	LogGas                uint64 = 375   // Per LOG* operation.
	CopyGas               uint64 = 3     //
	StackLimit            uint64 = 1024  // Maximum size of VM stack allowed.
	TierStepGas           uint64 = 0     // Once per operation, for a selection of them.
	LogTopicGas           uint64 = 375   // Multiplied by the * of the LOG*, per LOG transaction. e.g. LOG0 incurs 0 * c_txLogTopicGas, LOG4 incurs 4 * c_txLogTopicGas.
	CreateGas             uint64 = 32000 // Once per CREATE operation & contract-creation transaction.
	Create2Gas            uint64 = 32000 // Once per CREATE2 operation
	SelfdestructRefundGas uint64 = 24000 // Refunded following a selfdestruct operation.
	MemoryGas             uint64 = 3     // Times the address of the (highest referenced byte in memory + 1). NOTE: referencing happens on read, write and in instructions such as RETURN and CALL.

	TxDataNonZeroGas          uint64 = 16   // Per byte of data attached to a transaction that is not equal to zero. NOTE: Not payable on data of calls between transactions.
	TxAccessListAddressGas    uint64 = 2400 // Per address specified in access list
	TxAccessListStorageKeyGas uint64 = 1900 // Per storage key specified in access list

	// These have been changed during the course of the chain
	CallGas         uint64 = 700 // Static portion of gas for CALL-derivates
	BalanceGas      uint64 = 700 // The cost of a BALANCE operation
	ExtcodeSizeGas  uint64 = 700 // Cost of EXTCODESIZE
	SloadGas        uint64 = 800
	ExtcodeHashGas  uint64 = 700  // Cost of EXTCODEHASH
	SelfdestructGas uint64 = 5000 // Cost of SELFDESTRUCT

	// EXP has a dynamic portion depending on the size of the exponent
	ExpByte uint64 = 50 // was raised to 50

	// Extcodecopy has a dynamic AND a static cost. This represents only the
	// static portion of the gas.
	ExtcodeCopyBase uint64 = 700

	// CreateBySelfdestructGas is used when the refunded account is one that does
	// not exist. This logic is similar to call.
	CreateBySelfdestructGas uint64 = 25000

	BaseFeeChangeDenominator = 8          // Bounds the amount the base fee can change between blocks.
	ElasticityMultiplier     = 2          // Bounds the maximum gas limit a block may have.
	InitialBaseFee           = 1 * GWei   // Initial base fee for blocks.
	MaxBaseFee               = 100 * GWei // Maximum base fee for blocks.

	MaxCodeSize = 24576 // Maximum bytecode to permit for a contract

	// Precompiled contract gas prices

	EcrecoverGas        uint64 = 3000 // Elliptic curve sender recovery gas price
	Sha256BaseGas       uint64 = 60   // Base price for a SHA256 operation
	Sha256PerWordGas    uint64 = 12   // Per-word price for a SHA256 operation
	Ripemd160BaseGas    uint64 = 600  // Base price for a RIPEMD160 operation
	Ripemd160PerWordGas uint64 = 120  // Per-word price for a RIPEMD160 operation
	IdentityBaseGas     uint64 = 15   // Base price for a data copy operation
	IdentityPerWordGas  uint64 = 3    // Per-work price for a data copy operation

	Bn256AddGas             uint64 = 150   // Gas needed for an elliptic curve addition
	Bn256ScalarMulGas       uint64 = 6000  // Gas needed for an elliptic curve scalar multiplication
	Bn256PairingBaseGas     uint64 = 45000 // Base price for an elliptic curve pairing check
	Bn256PairingPerPointGas uint64 = 34000 // Per-point price for an elliptic curve pairing check

	// The Refund Quotient is the cap on how much of the used gas can be refunded
	RefundQuotient uint64 = 5

	MaxAddressGrindAttempts int = 1000 // Maximum number of attempts to grind an address to a valid one
	MinimumEtxGasDivisor        = 5    // The divisor for the minimum gas for inbound ETXs (Block gas limit / MinimumEtxGasDivisor)
	MaximumEtxGasMultiplier     = 2    // Multiplied with the minimum ETX gas for inbound ETXs (Block gas limit / MinimumEtxGasDivisor) * MaximumEtxGasMultiplier

	// Dynamic Expansion parameters

	//  This is the threshold (range 0-100) above which the
	// score will begin the tree expansion decision process. This threshold should be
	// 	chosen high enough to not be easily triggered by minor changes in node
	// 	operating behavior, but not so high that the security efficiency becomes
	// 	unacceptably low.
	TREE_EXPANSION_THRESHOLD uint16 = 15

	// This is the smoothing factor (range 0-10) used by each zone in its low-pass
	// filter to gather a long running average of the zone's security efficiency
	// score. Choosing a larger will make the filter less responsive; the tree
	// expansion algorithm will be less susceptible to short term variations in the
	// efficiency score, but will take longer to decide to trigger an expansion when
	// one becomes necessary.
	TREE_EXPANSION_FILTER_ALPHA uint16 = 9

	//  Once all chains have confirmed above TREE_EXPANSION_THRESHOLD, this is
	//  the number of consecutive prime blocks that must remain above the
	//  threshold to confirm the decision to expand the tree.
	TREE_EXPANSION_TRIGGER_WINDOW uint16 = 144

	// Once the network has confirmed the decision to expand the tree, this is
	// the number of prime blocks to wait until the expansion is activated. This
	// should be chosen to give node operators some time to adjust their
	// infrastructure, if needed, to account for the upcoming network change.
	TREE_EXPANSION_WAIT_COUNT = 1024

	ConversionLockPeriod          int64 = 10 // The number of zone blocks that a conversion output is locked for
	MinQiConversionDenomination         = 1
	ConversionConfirmationContext       = common.PRIME_CTX // A conversion requires a single coincident Dom confirmation
)

var (
	GasCeil                    uint64 = 20000000
	ColosseumGasCeil           uint64 = 70000000
	GardenGasCeil              uint64 = 160000000
	OrchardGasCeil             uint64 = 50000000
	LighthouseGasCeil          uint64 = 160000000
	LocalGasCeil               uint64 = 20000000
	DifficultyBoundDivisor            = big.NewInt(2048)  // The bound divisor of the difficulty, used in the update calculations.
	ZoneMinDifficulty                 = big.NewInt(1000)  // The minimum difficulty in a zone. Prime & regions should be multiples of this value
	MinimumDifficulty                 = ZoneMinDifficulty // The minimum that the difficulty may ever be.
	GenesisDifficulty                 = ZoneMinDifficulty // Difficulty of the Genesis block.
	DurationLimit                     = big.NewInt(12)    // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	GardenDurationLimit               = big.NewInt(7)     // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	OrchardDurationLimit              = big.NewInt(12)    // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	LighthouseDurationLimit           = big.NewInt(7)     // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	LocalDurationLimit                = big.NewInt(2)     // The decision boundary on the blocktime duration used to determine whether difficulty should go up or not.
	TimeFactor                        = big.NewInt(7)
	TimeToStartTx              uint64 = 0 * BlocksPerDay
	BlocksPerDay               uint64 = new(big.Int).Div(big.NewInt(86400), DurationLimit).Uint64() // BlocksPerDay is the number of blocks per day assuming 12 second block time
	PrimeEntropyTarget                = big.NewInt(441)                                             // This is TimeFactor*TimeFactor*common.NumZonesInRegion*common.NumRegionsInPrime
	RegionEntropyTarget               = big.NewInt(21)                                              // This is TimeFactor*common.NumZonesInRegion
	DifficultyAdjustmentPeriod        = big.NewInt(360)                                             // This is the number of blocks over which the average has to be taken
	DifficultyAdjustmentFactor int64  = 40                                                          // This is the factor that divides the log of the change in the difficulty
	MinQuaiConversionAmount           = new(big.Int).Mul(big.NewInt(1), big.NewInt(GWei))           // 0.000000001 Quai
	MaxWorkShareCount                 = 16
	WorkSharesThresholdDiff           = 3 // Number of bits lower than the target that the default consensus engine uses
	WorkSharesInclusionDepth          = 7 // Number of blocks upto which the work shares can be referenced and this is protocol enforced
)
