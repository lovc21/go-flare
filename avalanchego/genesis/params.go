// Copyright (C) 2019-2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
)

type StakingConfig struct {
	// Staking uptime requirements
	UptimeRequirement float64 `json:"uptimeRequirement"`
	// Minimum stake, in nAVAX, required to validate the primary network
	MinValidatorStake uint64 `json:"minValidatorStake"`
	// Maximum stake, in nAVAX, allowed to be placed on a single validator in
	// the primary network
	MaxValidatorStake uint64 `json:"maxValidatorStake"`
	// Minimum stake, in nAVAX, that can be delegated on the primary network
	MinDelegatorStake uint64 `json:"minDelegatorStake"`
	// Minimum delegation fee, in the range [0, 1000000], that can be charged
	// for delegation on the primary network.
	MinDelegationFee uint32 `json:"minDelegationFee"`
	// MinStakeDuration is the minimum amount of time a validator can validate
	// for in a single period.
	MinStakeDuration time.Duration `json:"minStakeDuration"`
	// MaxStakeDuration is the maximum amount of time a validator can validate
	// for in a single period.
	MaxStakeDuration time.Duration `json:"maxStakeDuration"`
	// RewardConfig is the config for the reward function.
	RewardConfig reward.Config `json:"rewardConfig"`
}

type TxFeeConfig struct {
	// Transaction fee
	TxFee uint64 `json:"txFee"`
	// Transaction fee for create asset transactions
	CreateAssetTxFee uint64 `json:"createAssetTxFee"`
	// Transaction fee for create subnet transactions
	CreateSubnetTxFee uint64 `json:"createSubnetTxFee"`
	// Transaction fee for transform subnet transactions
	TransformSubnetTxFee uint64 `json:"transformSubnetTxFee"`
	// Transaction fee for create blockchain transactions
	CreateBlockchainTxFee uint64 `json:"createBlockchainTxFee"`
	// Transaction fee for adding a primary network validator
	AddPrimaryNetworkValidatorFee uint64 `json:"addPrimaryNetworkValidatorFee"`
	// Transaction fee for adding a primary network delegator
	AddPrimaryNetworkDelegatorFee uint64 `json:"addPrimaryNetworkDelegatorFee"`
	// Transaction fee for adding a subnet validator
	AddSubnetValidatorFee uint64 `json:"addSubnetValidatorFee"`
	// Transaction fee for adding a subnet delegator
	AddSubnetDelegatorFee uint64 `json:"addSubnetDelegatorFee"`
}

type Params struct {
	StakingConfig
	TxFeeConfig
}

func GetTxFeeConfig(networkID uint32) TxFeeConfig {
	switch networkID {
	case constants.MainnetID:
		return MainnetParams.TxFeeConfig
	case constants.LocalID:
		return LocalParams.TxFeeConfig
	case constants.FlareID:
		return FlareParams.TxFeeConfig
	case constants.CostwoID:
		return CostwoParams.TxFeeConfig
	case constants.LocalFlareID:
		return LocalFlareParams.TxFeeConfig
	case constants.SongbirdID:
		return SongbirdParams.TxFeeConfig
	case constants.CostonID:
		return CostonParams.TxFeeConfig
	default:
		return LocalParams.TxFeeConfig
	}
}

func GetStakingConfig(networkID uint32) StakingConfig {
	switch networkID {
	case constants.MainnetID:
		return MainnetParams.StakingConfig
	case constants.LocalID:
		return LocalParams.StakingConfig
	case constants.FlareID:
		return FlareParams.StakingConfig
	case constants.CostwoID:
		return CostwoParams.StakingConfig
	case constants.LocalFlareID:
		return LocalFlareParams.StakingConfig
	case constants.SongbirdID:
		return SongbirdParams.StakingConfig
	case constants.CostonID:
		return CostonParams.StakingConfig
	default:
		return LocalParams.StakingConfig
	}
}
