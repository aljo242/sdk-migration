package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/aljo242/sdk-migration/testutil/keeper"
	"github.com/aljo242/sdk-migration/x/sdkmigration/keeper"
	"github.com/aljo242/sdk-migration/x/sdkmigration/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SdkmigrationKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
