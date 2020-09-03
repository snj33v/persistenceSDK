/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package simulator

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/genesis"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/parameters"
	"math/rand"
)

func (simulator) RandomizedGenesisState(simulationState *module.SimulationState) {

	var dummyParameter sdkTypes.Dec
	simulationState.AppParams.GetOrGenerate(
		simulationState.Cdc,
		DummyParameter,
		&dummyParameter,
		simulationState.Rand,
		func(r *rand.Rand) { dummyParameter = generateDummyParameter(r) },
	)

	Parameters := parameters.NewParameters(dummyParameter)

	// TODO add assetList
	genesisState := genesis.NewGenesisState(nil, Parameters)

	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", codec.MustMarshalJSONIndent(simulationState.Cdc, genesisState))
	simulationState.GenState[mapper.ModuleName] = simulationState.Cdc.MustMarshalJSON(genesisState)
}

func generateDummyParameter(r *rand.Rand) sdkTypes.Dec {
	return sdkTypes.NewDecWithPrec(int64(r.Intn(99)), 2)
}