/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package make

import (
	"github.com/persistenceOne/persistenceSDK/constants/flags"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/mapper"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Transaction = base.NewTransaction(
	mapper.ModuleName,
	TransactionName,
	TransactionRoute,
	TransactionShort,
	TransactionLong,
	registerCodec,
	initializeTransactionKeeper,
	requestPrototype,
	[]helpers.CLIFlag{flags.ClassificationID, flags.FromID, flags.ToID, flags.MakerOwnableSplit, flags.MakerOwnableID, flags.ExchangeRate, flags.TakerOwnableID, flags.ExpiresIn, flags.ImmutableMetaProperties, flags.ImmutableProperties, flags.MutableMetaProperties, flags.MutableProperties},
)
