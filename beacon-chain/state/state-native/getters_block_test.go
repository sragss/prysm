package state_native

import (
	"testing"

	"github.com/prysmaticlabs/prysm/beacon-chain/state"
	testtmpl "github.com/prysmaticlabs/prysm/beacon-chain/state/testing"
	ethpb "github.com/prysmaticlabs/prysm/proto/prysm/v1alpha1"
)

func TestBeaconState_LatestBlockHeader_Phase0(t *testing.T) {
	testtmpl.VerifyBeaconStateLatestBlockHeader(
		t,
		func() (state.BeaconState, error) {
			return InitializeFromProtoPhase0(&ethpb.BeaconState{})
		},
		func(BH *ethpb.BeaconBlockHeader) (state.BeaconState, error) {
			return InitializeFromProtoPhase0(&ethpb.BeaconState{LatestBlockHeader: BH})
		},
	)
}

func TestBeaconState_LatestBlockHeader_Altair(t *testing.T) {
	testtmpl.VerifyBeaconStateLatestBlockHeader(
		t,
		func() (state.BeaconState, error) {
			return InitializeFromProtoAltair(&ethpb.BeaconStateAltair{})
		},
		func(BH *ethpb.BeaconBlockHeader) (state.BeaconState, error) {
			return InitializeFromProtoAltair(&ethpb.BeaconStateAltair{LatestBlockHeader: BH})
		},
	)
}

func TestBeaconState_LatestBlockHeader_Bellatrix(t *testing.T) {
	testtmpl.VerifyBeaconStateLatestBlockHeader(
		t,
		func() (state.BeaconState, error) {
			return InitializeFromProtoBellatrix(&ethpb.BeaconStateBellatrix{})
		},
		func(BH *ethpb.BeaconBlockHeader) (state.BeaconState, error) {
			return InitializeFromProtoBellatrix(&ethpb.BeaconStateBellatrix{LatestBlockHeader: BH})
		},
	)
}

func TestBeaconState_BlockRoots_Phase0(t *testing.T) {
	testtmpl.VerifyBeaconStateBlockRootsNative(
		t,
		func() (state.BeaconState, error) {
			return InitializeFromProtoPhase0(&ethpb.BeaconState{})
		},
		func(BR [][]byte) (state.BeaconState, error) {
			return InitializeFromProtoPhase0(&ethpb.BeaconState{BlockRoots: BR})
		},
	)
}

func TestBeaconState_BlockRoots_Altair(t *testing.T) {
	testtmpl.VerifyBeaconStateBlockRootsNative(
		t,
		func() (state.BeaconState, error) {
			return InitializeFromProtoAltair(&ethpb.BeaconStateAltair{})
		},
		func(BR [][]byte) (state.BeaconState, error) {
			return InitializeFromProtoAltair(&ethpb.BeaconStateAltair{BlockRoots: BR})
		},
	)
}

func TestBeaconState_BlockRoots_Bellatrix(t *testing.T) {
	testtmpl.VerifyBeaconStateBlockRootsNative(
		t,
		func() (state.BeaconState, error) {
			return InitializeFromProtoBellatrix(&ethpb.BeaconStateBellatrix{})
		},
		func(BR [][]byte) (state.BeaconState, error) {
			return InitializeFromProtoBellatrix(&ethpb.BeaconStateBellatrix{BlockRoots: BR})
		},
	)
}

func TestBeaconState_BlockRootAtIndex_Phase0(t *testing.T) {
	testtmpl.VerifyBeaconStateBlockRootAtIndexNative(
		t,
		func() (state.BeaconState, error) {
			return InitializeFromProtoPhase0(&ethpb.BeaconState{})
		},
		func(BR [][]byte) (state.BeaconState, error) {
			return InitializeFromProtoPhase0(&ethpb.BeaconState{BlockRoots: BR})
		},
	)
}

func TestBeaconState_BlockRootAtIndex_Altair(t *testing.T) {
	testtmpl.VerifyBeaconStateBlockRootAtIndexNative(
		t,
		func() (state.BeaconState, error) {
			return InitializeFromProtoAltair(&ethpb.BeaconStateAltair{})
		},
		func(BR [][]byte) (state.BeaconState, error) {
			return InitializeFromProtoAltair(&ethpb.BeaconStateAltair{BlockRoots: BR})
		},
	)
}

func TestBeaconState_BlockRootAtIndex_Bellatrix(t *testing.T) {
	testtmpl.VerifyBeaconStateBlockRootAtIndexNative(
		t,
		func() (state.BeaconState, error) {
			return InitializeFromProtoBellatrix(&ethpb.BeaconStateBellatrix{})
		},
		func(BR [][]byte) (state.BeaconState, error) {
			return InitializeFromProtoBellatrix(&ethpb.BeaconStateBellatrix{BlockRoots: BR})
		},
	)
}
