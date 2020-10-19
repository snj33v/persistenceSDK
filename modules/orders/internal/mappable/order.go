/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/orders/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type order struct {
	ID         types.ID         `json:"key" valid:"required~required field key missing"`
	Immutables types.Immutables `json:"immutables" valid:"required field immutables missing"`
	Mutables   types.Mutables   `json:"mutables" valid:"required~required field mutables missing"`
}

var _ mappables.Order = (*order)(nil)

func (order order) GetClassificationID() types.ID {
	return key.ReadClassificationID(order.ID)
}
func (order order) GetMakerOwnableID() types.ID {
	return key.ReadMakerOwnableID(order.ID)
}
func (order order) GetTakerOwnableID() types.ID {
	return key.ReadTakerOwnableID(order.ID)
}
func (order order) GetMakerID() types.ID {
	return key.ReadMakerID(order.ID)
}
func (order order) GetTakerID() types.Property {
	if takerID := order.Immutables.Get().Get(base.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else if takerID := order.Mutables.Get().Get(base.NewID(properties.TakerID)); takerID != nil {
		return takerID
	} else {
		data, _ := base.ReadIDData("")
		return base.NewProperty(base.NewID(properties.TakerID), base.NewFact(data))
	}
}
func (order order) GetExchangeRate() types.Property {
	if exchangeRate := order.Immutables.Get().Get(base.NewID(properties.ExchangeRate)); exchangeRate != nil {
		return exchangeRate
	} else if exchangeRate := order.Mutables.Get().Get(base.NewID(properties.ExchangeRate)); exchangeRate != nil {
		return exchangeRate
	} else {
		data := base.NewDecData(sdkTypes.OneDec())
		return base.NewProperty(base.NewID(properties.ExchangeRate), base.NewFact(data))
	}
}
func (order order) GetCreation() types.Property {
	if creation := order.Immutables.Get().Get(base.NewID(properties.Creation)); creation != nil {
		return creation
	} else if creation := order.Mutables.Get().Get(base.NewID(properties.Creation)); creation != nil {
		return creation
	} else {
		data, _ := base.ReadHeightData("")
		return base.NewProperty(base.NewID(properties.Creation), base.NewFact(data))
	}
}
func (order order) GetExpiry() types.Property {
	if expiry := order.Immutables.Get().Get(base.NewID(properties.Expiry)); expiry != nil {
		return expiry
	} else if creation := order.Mutables.Get().Get(base.NewID(properties.Expiry)); creation != nil {
		return creation
	} else {
		data, _ := base.ReadHeightData("")
		return base.NewProperty(base.NewID(properties.Expiry), base.NewFact(data))
	}
}
func (order order) GetMakerOwnableSplit() types.Property {
	if split := order.Immutables.Get().Get(base.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else if split := order.Mutables.Get().Get(base.NewID(properties.MakerOwnableSplit)); split != nil {
		return split
	} else {
		data, _ := base.ReadDecData("")
		return base.NewProperty(base.NewID(properties.MakerOwnableSplit), base.NewFact(data))
	}
}
func (order order) GetImmutables() types.Immutables {
	return order.Immutables
}
func (order order) GetMutables() types.Mutables {
	return order.Mutables
}
func (order order) GetID() types.ID {
	return order.ID
}
func (order order) GetKey() helpers.Key {
	return key.New(order.ID)
}
func (order) RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(order{}, constants.ProjectRoute+"/"+"order", nil)
}

func NewOrder(orderID types.ID, immutables types.Immutables, mutables types.Mutables) mappables.Order {
	return order{
		ID:         orderID,
		Immutables: immutables,
		Mutables:   mutables,
	}
}