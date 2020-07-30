package mapper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(splits{}, ModuleRoute+"/"+"splits", nil)
	codec.RegisterConcrete(split{}, ModuleRoute+"/"+"split", nil)
	codec.RegisterConcrete(splitID{}, ModuleRoute+"/"+"splitID", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	packageCodec.Seal()
}