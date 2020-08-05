package custody

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/exchanges/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/burn"
	"github.com/persistenceOne/persistenceSDK/modules/splits/auxiliaries/mint"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

type auxiliaryKeeper struct {
	mapper              helpers.Mapper
	splitsMintAuxiliary helpers.Auxiliary
	splitsBurnAuxiliary helpers.Auxiliary
}

var _ helpers.AuxiliaryKeeper = (*auxiliaryKeeper)(nil)

func (auxiliaryKeeper auxiliaryKeeper) Help(context sdkTypes.Context, AuxiliaryRequest helpers.AuxiliaryRequest) error {
	auxiliaryRequest := auxiliaryRequestFromInterface(AuxiliaryRequest)
	order := auxiliaryRequest.Order
	if auxiliaryRequest.Reverse {
		if Error := auxiliaryKeeper.splitsBurnAuxiliary.GetKeeper().Help(context,
			burn.NewAuxiliaryRequest(base.NewID(mapper.ModuleName), order.GetMakerAssetData(), order.GetMakerAssetAmount())); Error != nil {
			return Error
		}
		if Error := auxiliaryKeeper.splitsMintAuxiliary.GetKeeper().Help(context,
			mint.NewAuxiliaryRequest(order.GetMakerID(), order.GetMakerAssetData(), order.GetMakerAssetAmount())); Error != nil {
			return Error
		}
	} else {
		if Error := auxiliaryKeeper.splitsBurnAuxiliary.GetKeeper().Help(context,
			burn.NewAuxiliaryRequest(order.GetMakerID(), order.GetMakerAssetData(), order.GetMakerAssetAmount())); Error != nil {
			return Error
		}
		if Error := auxiliaryKeeper.splitsMintAuxiliary.GetKeeper().Help(context,
			mint.NewAuxiliaryRequest(base.NewID(mapper.ModuleName), order.GetMakerAssetData(), order.GetMakerAssetAmount())); Error != nil {
			return Error
		}

	}
	return nil
}

func initializeAuxiliaryKeeper(mapper helpers.Mapper, externalKeepers []interface{}) helpers.AuxiliaryKeeper {
	auxiliaryKeeper := auxiliaryKeeper{mapper: mapper}
	for _, externalKeeper := range externalKeepers {
		switch value := externalKeeper.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case mint.Auxiliary.GetName():
				auxiliaryKeeper.splitsMintAuxiliary = value
			case burn.Auxiliary.GetName():
				auxiliaryKeeper.splitsBurnAuxiliary = value
			}
		}
	}
	return auxiliaryKeeper
}
