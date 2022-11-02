package keeper

import (
	"github.com/aljo242/sdk-migration/x/sdkmigration/types"
)

var _ types.QueryServer = Keeper{}
