package epoch_processing

import (
	"path"
	"testing"

	"github.com/prysmaticlabs/prysm/beacon-chain/core/epoch"
	"github.com/prysmaticlabs/prysm/beacon-chain/core/helpers"
	"github.com/prysmaticlabs/prysm/beacon-chain/state"
	"github.com/prysmaticlabs/prysm/config/params"
	"github.com/prysmaticlabs/prysm/testing/require"
	"github.com/prysmaticlabs/prysm/testing/spectest/utils"
)

// RunSlashingsTests executes "epoch_processing/slashings" tests.
func RunSlashingsTests(t *testing.T, config string) {
	require.NoError(t, utils.SetConfig(t, config))

	testFolders, testsFolderPath := utils.TestFolders(t, config, "bellatrix", "epoch_processing/slashings/pyspec_tests")
	for _, folder := range testFolders {
		t.Run(folder.Name(), func(t *testing.T) {
			folderPath := path.Join(testsFolderPath, folder.Name())
			helpers.ClearCache()
			RunEpochOperationTest(t, folderPath, processSlashingsWrapper)
		})
	}
}

func processSlashingsWrapper(t *testing.T, st state.BeaconState) (state.BeaconState, error) {
	st, err := epoch.ProcessSlashings(st, params.BeaconConfig().ProportionalSlashingMultiplierBellatrix)
	require.NoError(t, err, "Could not process slashings")
	return st, nil
}
