package assets

import (
	"github.com/persistenceOne/persistenceSDK/modules/assets/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/queries/asset"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assets/transactions/mutate"
	"github.com/persistenceOne/persistenceSDK/types"
)

var Module = types.NewModule(
	mapper.ModuleName,
	mapper.StoreKey,
	mapper.DefaultParamspace,
	mapper.QueryRoute,
	mapper.TransactionRoute,
	genesis.GenesisState,
	mapper.Mapper,
	[]types.Query{asset.Query},
	[]types.Transaction{burn.Transaction, mint.Transaction, mutate.Transaction},
)