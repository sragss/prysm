package params

import (
	eth1Params "github.com/ethereum/go-ethereum/params"
)

// UseSepoliaNetworkConfig uses the Sepolia beacon chain specific network config.
func UseSepoliaNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 1273020
	cfg.BootstrapNodes = []string{
		// EF boot node
		"enr:-Iq4QMCTfIMXnow27baRUb35Q8iiFHSIDBJh6hQM5Axohhf4b6Kr_cOCu0htQ5WvVqKvFgY28893DHAg8gnBAXsAVqmGAX53x8JggmlkgnY0gmlwhLKAlv6Jc2VjcDI1NmsxoQK6S-Cii_KmfFdUJL2TANL3ksaKUnNXvTCv1tLwXs0QgIN1ZHCCIyk",
	}
	OverrideBeaconNetworkConfig(cfg)
}

// SepoliaConfig defines the config for the Sepolia beacon chain testnet.
func SepoliaConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()
	cfg.MinGenesisTime = 1655647200
	cfg.GenesisDelay = 86400
	cfg.MinGenesisActiveValidatorCount = 1300
	cfg.ConfigName = SepoliaName
	cfg.GenesisForkVersion = []byte{0x90, 0x00, 0x00, 0x69}
	cfg.SecondsPerETH1Block = 14
	cfg.DepositChainID = eth1Params.SepoliaChainConfig.ChainID.Uint64()
	cfg.DepositNetworkID = eth1Params.SepoliaChainConfig.ChainID.Uint64()
	cfg.AltairForkEpoch = 50
	cfg.AltairForkVersion = []byte{0x90, 0x00, 0x00, 0x70}
	cfg.BellatrixForkEpoch = 100
	cfg.BellatrixForkVersion = []byte{0x90, 0x00, 0x00, 0x71}
	cfg.TerminalTotalDifficulty = "17000000000000000"
	cfg.DepositContractAddress = "0x7f02C3E3c98b133055B8B348B2Ac625669Ed295D"
	cfg.InitializeForkSchedule()
	return cfg
}
