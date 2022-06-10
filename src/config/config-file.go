package config

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/latoken/bridge-balancer-service/src/models"
	"github.com/latoken/bridge-balancer-service/src/service/storage"

	"github.com/ethereum/go-ethereum/common"
)

// Reads Service params from config.json
func (v *viperConfig) ReadServiceConfig() string {
	return fmt.Sprintf("%s:%s", v.GetString("service.host"), v.GetString("service.port"))
}

// ReadLachainConfig reads lachain chain params from config.json
func (v *viperConfig) ReadLachainConfig() *models.WorkerConfig {
	return v.readWorkerConfig("LA")
}

// ReadEthWorkerConfig reads ethereum chain worker params from config.json
func (v *viperConfig) ReadWorkersConfig() []*models.WorkerConfig {
	chains := v.GetStringSlice("chains")
	chainCfgs := make([]*models.WorkerConfig, 0)

	for _, chain := range chains {
		chainCfgs = append(chainCfgs, v.readWorkerConfig(chain))
	}

	return chainCfgs
}

// readETHWorkerConfig reads ethereum chain worker params from config.json
func (v *viperConfig) readWorkerConfig(name string) *models.WorkerConfig {
	return &models.WorkerConfig{
		NetworkType:        v.GetString(fmt.Sprintf("workers.%s.type", name)),
		ChainName:          strings.ToUpper(name),
		User:               v.GetString(fmt.Sprintf("workers.%s.user", name)),
		Password:           v.GetString(fmt.Sprintf("workers.%s.password", name)),
		WorkerAddr:         common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.worker_addr", name))),
		PrivateKey:         v.GetString(fmt.Sprintf("workers.%s.private_key", name)),
		Provider:           v.GetString(fmt.Sprintf("workers.%s.provider", name)),
		ContractAddr:       common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.contract_addr", name))),
		TokenContractAddr:  common.HexToAddress(v.GetString(fmt.Sprintf("workers.%s.token_addr", name))),
		GasLimit:           v.GetInt64(fmt.Sprintf("workers.%s.gas_limit", name)),
		GasPrice:           big.NewInt(v.GetInt64(fmt.Sprintf("workers.%s.gas_price", name))),
		FetchInterval:      v.GetInt64(fmt.Sprintf("workers.%s.fetch_interval", name)),
		ConfirmNum:         v.GetInt64(fmt.Sprintf("workers.%s.confirm_num", name)),
		StartBlockHeight:   v.GetInt64(fmt.Sprintf("workers.%s.start_block_height", name)),
		DestinationChainID: v.GetString(fmt.Sprintf("workers.%s.dest_id", name)),
	}
}

// Reads storage params from config.json
func (v *viperConfig) ReadDBConfig() *models.StorageConfig {
	return &models.StorageConfig{
		URL:        v.GetString("storage.url"),
		DBDriver:   v.GetString("storage.driver"),
		DBHOST:     v.GetString("storage.host"),
		DBPORT:     v.GetInt64("storage.port"),
		DBSSL:      v.GetString("storage.ssl_mode"),
		DBName:     v.GetString("storage.db_name"),
		DBUser:     v.GetString("storage.user"),
		DBPassword: v.GetString("storage.password"),
	}
}

func (v *viperConfig) ReadFetcherConfig() *models.FetcherConfig {
	return &models.FetcherConfig{
		AllTokens: v.GetStringSlice("all_tokens"),
	}
}

func (v *viperConfig) ReadFarmerConfig() *models.FarmerConfig {
	return &models.FarmerConfig{
		ApyURL:    v.GetString("farmer.apy_url"),
		LpsURL:    v.GetString("farmer.lps_url"),
		PricesURL: v.GetString("farmer.prices_url"),
	}
}

func (v *viperConfig) ReadFarmsConfig() map[string]*models.FarmConfig {
	farms := v.GetStringSlice("farms.list")
	farmsCfgs := make(map[string]*models.FarmConfig)

	for _, farmId := range farms {
		farmsCfgs[farmId] = &models.FarmConfig{
			ID:                  farmId,
			ChainName:           v.GetString(fmt.Sprintf("farms.%s.chain_name", farmId)),
			Type:                v.GetString(fmt.Sprintf("farms.%s.type", farmId)),
			FarmId:              v.GetString(fmt.Sprintf("farms.%s.farm_id", farmId)),
			Oracle:              v.GetString(fmt.Sprintf("farms.%s.oracle", farmId)),
			VaultAddress:        common.HexToAddress(v.GetString(fmt.Sprintf("farms.%s.vault_address", farmId))),
			PairAddress:         common.HexToAddress(v.GetString(fmt.Sprintf("farms.%s.pair_address", farmId))),
			Token0:              v.GetString(fmt.Sprintf("farms.%s.token0", farmId)),
			Token1:              v.GetString(fmt.Sprintf("farms.%s.token1", farmId)),
			ChainId:             v.GetString(fmt.Sprintf("farms.%s.chain_id", farmId)),
			Name:                v.GetString(fmt.Sprintf("farms.%s.name", farmId)),
			Protocol:            v.GetString(fmt.Sprintf("farms.%s.protocol", farmId)),
			DepositToken:        v.GetString(fmt.Sprintf("farms.%s.deposit_token", farmId)),
			Logo0:               v.GetString(fmt.Sprintf("farms.%s.logo0", farmId)),
			Logo1:               v.GetString(fmt.Sprintf("farms.%s.logo1", farmId)),
			DepositResourceID:   v.GetString(fmt.Sprintf("farms.%s.deposit_resource_id", farmId)),
			WithdrawResourceID:  v.GetString(fmt.Sprintf("farms.%s.withdraw_resource_id", farmId)),
			Gas:                 v.GetInt64(fmt.Sprintf("farms.%s.withdraw_resource_id", farmId)),
			WrappedTokenAddress: v.GetString(fmt.Sprintf("farms.%s.wrapped_token_address", farmId)),
		}
	}

	return farmsCfgs
}

func (v *viperConfig) ReadResourceIDs(fetcher *models.FetcherConfig) []*storage.ResourceId {
	resouceIDs := make([]*storage.ResourceId, len(fetcher.AllTokens))
	for index, name := range fetcher.AllTokens {
		resouceIDs[index] = &storage.ResourceId{
			Name: name,
			ID:   v.GetString(fmt.Sprintf("resourceIDs.%s", name)),
		}
	}
	return resouceIDs
}

func (v *viperConfig) ReadChains() []string {
	return v.GetStringSlice("chains")
}
