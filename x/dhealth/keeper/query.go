package keeper

import (
	"github.com/dhealthproject/dhealth/x/dhealth/types"
)

var _ types.QueryServer = Keeper{}
