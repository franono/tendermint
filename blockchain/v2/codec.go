package v2

import (
	amino "github.com/tendermint/go-amino"

	"github.com/franono/tendermint/types"
)

var cdc = amino.NewCodec()

func init() {
	RegisterBlockchainMessages(cdc)
	types.RegisterBlockAmino(cdc)
}
